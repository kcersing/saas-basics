package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/memberprofile"
	"saas/biz/dal/db/ent/product"
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

	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}
	first, err := tx.MemberProduct.Query().
		Where(memberproduct.ID(req.MemberProductId)).
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
		SetStatus(1).
		SetVenueID(req.VenueId).
		SetProductID(req.ProductId).
		SetLength(first.Length).
		//SetPlaceID(req.PlaceId)
		//SetNum(req.Num).
		SetDate(startTime.Format(time.DateOnly)).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
		SetVenueName(venueName).
		Save(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Course Record Schedule failed")
		return rollback(tx, err)
	}

	if req.CoachId != 0 {

		u, err := tx.User.Query().
			Where(user.ID(req.CoachId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该员工")
			return rollback(tx, err)
		}

		_, err = tx.ScheduleCoach.Create().
			SetSchedule(one).
			SetScheduleName(one.Name).
			SetType(req.Type).
			SetVenueID(req.VenueId).
			SetCoachID(req.CoachId).
			SetStartTime(startTime).
			SetProductID(req.ProductId).
			SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
			SetStatus(1).
			SetCoachName(u.Name).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return rollback(tx, err)
		}

	}

	if req.MemberId != 0 {
		err = NewSchedule(s.ctx, s.c).CreateScheduleMemberCourse(do.CreateScheduleMemberCourse{
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
		if len(req.MemberIds) > 0 {

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
		err = errors.Wrap(err, "未查询到该会员")
		return err
	}

	var memberProduct *ent.MemberProduct
	if req.MemberProductId != 0 {
		memberProduct, err = s.db.MemberProduct.Query().
			Where(memberproduct.ID(req.MemberProductId)).
			First(s.ctx)
		if err != nil {
			return err
		}
	} else {
		memberProduct, err = s.db.MemberProduct.Query().
			Where(memberproduct.ProductID(req.ProductId), memberproduct.CountSurplusGT(0)).
			First(s.ctx)
		if err != nil {
			return err
		}
	}
	if memberProduct == nil {
		hlog.Error("未查询到该会员课程:", req)
		err = errors.Wrap(err, "未查询到该会员课程")
		return err
	}

	if (memberProduct.CountSurplus - 1) < 0 {
		hlog.Error("该会员课程不足:", req)
		err = errors.Wrap(err, "该会员课程不足")
		return err
	}

	_, err = s.db.ScheduleMember.Create().
		SetSchedule(req.One).
		SetType(req.Type).
		SetMemberID(req.MemberId).
		SetPlaceID(req.PlaceId).
		SetVenueID(req.VenueId).
		SetStartTime(req.StartTime).
		SetEndTime(req.StartTime.Add(time.Duration(memberProduct.Length) * time.Minute)).
		SetMemberProductID(memberProduct.ID).
		SetStatus(0).
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
		AddCountSurplus(-1).
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
		SetDate(startTime.Format(time.DateOnly)).
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
		SetScheduleName(one.Name).
		SetType("lessons").
		SetVenueID(req.VenueId).
		SetCoachID(req.CoachId).
		SetPlaceID(req.PlaceId).
		SetProductID(req.ProductId).
		SetStartTime(startTime).
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
