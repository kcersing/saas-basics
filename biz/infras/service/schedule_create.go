package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/memberprofile"
	"saas/biz/dal/db/ent/product"
	schedule2 "saas/biz/dal/db/ent/schedule"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/db/ent/venue"
	"saas/biz/dal/db/ent/venueplace"
	"saas/biz/infras/do"
	"saas/idl_gen/model/schedule"
	"time"

	"github.com/pkg/errors"
)

func (s Schedule) CreateScheduleCourse(req schedule.CreateOrUpdateScheduleCourseReq) error {
	startTime, err := time.Parse(time.DateTime, req.StartTime)
	if err != nil {
		return errors.New("日期类型传值错误")
	}

	if req.CoachId == 0 {
		return errors.New("教练ID不能为空")
	}
	if req.MemberId == 0 {
		return errors.New("会员ID不能为空")
	}

	date := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())

	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	first, err := tx.Product.Query().
		Where(product.ID(req.ProductId)).
		First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到该会员产品")
		return rollback(tx, err)
	}
	venueName := ""
	if req.VenueId != 0 {
		ven, err := s.db.Venue.Query().
			Where(venue.ID(req.VenueId)).
			First(s.ctx)
		if err == nil {
			venueName = ven.Name
		}
	}
	one, err := tx.Schedule.Create().
		SetType(req.Type).
		SetName(first.Name).
		SetStatus(2).
		SetVenueID(req.VenueId).
		SetProductID(req.ProductId).
		SetLength(first.Length).
		//SetPlaceID(req.PlaceId)
		//SetNum(req.Num).
		SetDate(date).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
		SetVenueName(venueName).
		Save(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Course Record Schedule failed")
		return rollback(tx, err)
	}

	u, err := tx.User.Query().
		Where(user.ID(req.CoachId)).
		First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到该员工")
		return rollback(tx, err)
	}

	_, err = tx.ScheduleCoach.Create().
		SetSchedule(one).
		SetType(req.Type).
		SetVenueID(req.VenueId).
		SetCoachID(req.CoachId).
		SetProductID(req.ProductId).
		SetDate(date).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
		SetStatus(1).
		SetCoachName(u.Name).
		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Course Record Coach failed")
		return rollback(tx, err)
	}
	err = s.CreateScheduleMemberCourse(do.CreateScheduleMemberCourse{
		One:       one,
		Type:      req.Type,
		VenueId:   req.VenueId,
		StartTime: startTime,
		MemberId:  req.MemberId,
		ProductId: req.MemberProductId,
	})
	if err != nil {
		return err
	}
	if len(req.MpIds) > 0 {
		for _, v := range req.MpIds {
			err = s.CreateScheduleMemberCourse(do.CreateScheduleMemberCourse{
				One:       one,
				Type:      req.Type,
				VenueId:   req.VenueId,
				StartTime: startTime,
				MemberId:  v.MemberId,
				ProductId: v.MemberProductId,
			})
			if err != nil {
				return err
			}
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (s Schedule) CreateScheduleMemberCourse(req do.CreateScheduleMemberCourse) error {
	m, err := s.db.MemberProfile.Query().
		Where(memberprofile.MemberID(req.MemberId)).
		First(s.ctx)
	if err != nil {
		hlog.Error("未查询到该会员:", req)

		return errors.Wrap(err, "未查询到该会员")
	}

	memberProduct, err := s.db.MemberProduct.Query().
		Where(
			memberproduct.ID(req.MemberProductId)).
		First(s.ctx)
	if err != nil {
		return errors.Wrap(err, "会员课程未找到")
	}

	if memberProduct.NumberSurplus < 1 {
		hlog.Error("该会员课程不足:", req)
		return errors.Wrap(err, "该会员课程不足")
	}

	_, err = s.db.ScheduleMember.Create().
		SetSchedule(req.One).
		SetType(req.Type).
		SetMemberID(req.MemberId).
		SetVenueID(req.VenueId).
		SetProductID(req.ProductId).
		SetStartTime(req.StartTime).
		SetDate(req.One.Date).
		SetEndTime(req.StartTime.Add(time.Duration(memberProduct.Length) * time.Minute)).
		SetMemberProductID(memberProduct.ID).
		SetStatus(1).
		SetMemberName(m.Name).
		SetMemberProductName(memberProduct.Name).
		Save(s.ctx)
	if err != nil {
		hlog.Error("无法创建会员课程:", req)
		err = errors.Wrap(err, "无法创建会员课程")
		return err
	}
	_, err = s.db.MemberProduct.
		Update().
		Where(memberproduct.IDEQ(memberProduct.ID)).
		SetNumberSurplus(-1).
		Save(s.ctx)
	if err != nil {
		hlog.Error("无法会员课程节数异常:", req)
		err = errors.Wrap(err, "无法会员课程节数异常")
		return err
	}
	return nil
}

func (s Schedule) CreateScheduleLessons(req schedule.CreateOrUpdateScheduleLessonsReq) error {

	// 解析字符串到 time.Time 类型
	startTime, err := time.Parse(time.DateTime, req.StartTime)

	if err != nil {
		return errors.New("日期类型传值错误")
	}
	date := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())

	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	first, err := tx.Product.Query().
		Where(product.ID(req.ProductId)).
		First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到该产品")
		return rollback(tx, err)
	}

	var venueName, placeName string

	ven, err := s.db.Venue.Query().
		Where(venue.ID(req.VenueId)).
		First(s.ctx)
	if err != nil {
		return rollback(tx, err)
	}
	venueName = ven.Name
	place, err := s.db.VenuePlace.Query().
		Where(venueplace.ID(req.PlaceId)).
		First(s.ctx)
	if err != nil {
		return rollback(tx, err)
	}
	placeName = place.Name

	one, err := tx.Schedule.Create().
		SetType("lessons").
		SetName(first.Name).
		SetStatus(1).
		SetVenueID(req.VenueId).
		SetProductID(req.ProductId).
		SetPlaceID(req.PlaceId).
		SetNum(place.Number).
		SetNumSurplus(place.Number).
		SetDate(date).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
		//SetPrice(req.Price).
		//SetRemark(req.Remark).
		SetLength(first.Length).
		SetVenueName(venueName).
		SetPlaceName(placeName).
		Save(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Course Record Schedule failed")
		return rollback(tx, err)
	}

	u, err := tx.User.Query().
		Where(user.ID(req.CoachId)).
		First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到该员工")
		return rollback(tx, err)
	}

	_, err = tx.ScheduleCoach.Create().
		SetSchedule(one).
		SetType("lessons").
		SetVenueID(req.VenueId).
		SetCoachID(req.CoachId).
		SetPlaceID(req.PlaceId).
		SetProductID(req.ProductId).
		SetStartTime(startTime).
		SetDate(date).
		SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
		SetStatus(1).
		SetCoachName(u.Name).
		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Course Record Coach failed")
		return rollback(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
func (s Schedule) CreateMemberSubscribeLessons(req schedule.MemberSubscribeReq) error {

	one, err := s.db.Schedule.Query().Where(schedule2.IDEQ(req.ScheduleId)).First(s.ctx)

	if err != nil {
		hlog.Error("未查询到课程:", req)
		err = errors.Wrap(err, "未查询到课程")
		return err
	}
	mp, err := s.db.MemberProduct.Query().
		Where(memberproduct.ID(req.MemberProductId)).
		First(s.ctx)
	if err != nil {
		hlog.Error("未查询到该会员:", req)
		err = errors.Wrap(err, "未查询到该会员")
		return err
	}
	m, err := s.db.MemberProfile.Query().
		Where(memberprofile.MemberIDEQ(req.MemberId)).
		First(s.ctx)
	if err != nil {
		hlog.Error("未查询到该会员:", req)
		err = errors.Wrap(err, "未查询到该会员")
		return err
	}
	_, err = s.db.ScheduleMember.Create().
		SetSchedule(one).
		SetType(one.Type).
		SetMemberID(req.MemberId).
		SetProductID(one.ProductID).
		SetVenueID(one.VenueID).
		SetPlaceID(one.PlaceID).
		SetMemberProductID(req.MemberProductId).
		SetDate(one.Date).
		SetStartTime(one.StartTime).
		SetEndTime(one.StartTime).
		SetStatus(0).
		SetMemberName(m.Name).
		SetMemberProductName(mp.Name).
		SetSeat(*req.Seat).
		Save(s.ctx)
	if err != nil {
		hlog.Error("无法创建会员课程:", req)
		err = errors.Wrap(err, "无法创建会员课程")
		return err
	}
	return nil

}
