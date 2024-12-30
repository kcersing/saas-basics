package service

import (
	"fmt"
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/db/ent/venue"
	"saas/biz/dal/db/ent/venueplace"
	"saas/idl_gen/model/schedule"
	"time"
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
		m, err := tx.Member.Query().
			Where(member.ID(req.MemberId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员")
			return rollback(tx, err)
		}

		if (first.CountSurplus - 1) < 0 {
			err = errors.Wrap(err, "该会员课程不足")
			return rollback(tx, err)
		}

		_, err = tx.ScheduleMember.Create().
			SetSchedule(one).
			SetType(req.Type).
			SetMemberID(req.MemberId).
			SetVenueID(req.VenueId).
			SetStartTime(startTime).
			SetEndTime(startTime.Add(time.Duration(first.Length) * time.Minute)).
			SetMemberProductID(req.MemberProductId).
			SetStatus(0).
			SetMemberName(m.Name).
			SetMemberProductName(first.Name).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return rollback(tx, err)
		}

		_, err = tx.MemberProduct.
			Update().
			Where(memberproduct.IDEQ(req.MemberProductId)).
			AddCountSurplus(-1).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "Member Product Item failed")
			return rollback(tx, err)
		}

	}
	if len(req.MemberIds) > 0 {

	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s Schedule) CreateScheduleMemberCourse(req schedule.CreateOrUpdateScheduleLessonsReq) {
	m, err := s.db.Member.Query().
		Where(member.ID(req.MemberId)).
		First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到该会员")

	}
	memberProduct, err := s.db.MemberProduct.Query().
		Where(memberproduct.ProductID(req.ProductId), memberproduct.CountSurplusGT(0)).
		First(s.ctx)

	if (memberProduct.CountSurplus - 1) < 0 {
		err = errors.Wrap(err, "该会员课程不足")

	}

	_, err = s.db.ScheduleMember.Create().
		SetSchedule(one).
		SetType(req.Type).
		SetMemberID(req.MemberId).
		SetVenueID(req.VenueId).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(memberProduct.Length) * time.Minute)).
		SetMemberProductID(memberProduct.ID).
		SetStatus(0).
		SetMemberName(m.Name).
		SetMemberProductName(memberProduct.Name).
		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Course Record Coach failed")

	}

	_, err = s.db.MemberProduct.
		Update().
		Where(memberproduct.IDEQ(memberProduct.ID)).
		AddCountSurplus(-1).
		Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "Member Product Item failed")

	}
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
