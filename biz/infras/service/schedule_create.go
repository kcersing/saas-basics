package service

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/biz/dal/db/ent"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/memberproduct"
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
	startTime, _ := time.Parse(time.DateTime, req.StartTime)
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
		SetStatus(1).
		SetVenueID(req.VenueId).
		SetDate(startTime.Format(time.DateOnly)).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
		SetLength(first.Length).
		SetName(first.Name).
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
	m, err := s.db.Member.Query().
		Where(member.ID(req.MemberId)).
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
	startTime, _ := time.Parse(time.DateTime, req.StartTime)
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

	venueName := ""
	if req.VenueId != 0 {
		ven, err := s.db.Venue.Query().
			Where(venue.ID(req.VenueId)).
			First(s.ctx)
		if err == nil {
			venueName = ven.Name
		}
	}
	placeName := ""
	if req.PlaceId != 0 {
		place, err := s.db.VenuePlace.Query().
			Where(venueplace.ID(req.PlaceId)).
			First(s.ctx)
		if err == nil {
			placeName = place.Name
		}
	}

	one, err := tx.Schedule.Create().
		SetType("lessons").
		SetStatus(1).
		SetVenueID(req.VenueId).
		SetPlaceID(req.PlaceId).
		//SetNum(req.Num).
		//SetNumSurplus(req.Num).
		SetDate(startTime.Format(time.DateOnly)).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
		//SetPrice(req.Price).
		//SetRemark(req.Remark).
		SetLength(first.Length).
		SetName(first.Name).
		SetVenueName(venueName).
		SetPlaceName(placeName).
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
			SetType("lessons").
			SetVenueID(req.VenueId).
			SetCoachID(req.CoachId).
			SetStartTime(startTime).
			SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
			SetStatus(1).
			SetCoachName(u.Name).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return rollback(tx, err)
		}

	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
