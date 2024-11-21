package service

import (
	"fmt"
	"github.com/pkg/errors"
	"rpc_gen/kitex_gen/schedule"

	"time"
	"user/biz/dal/mysql/ent/user"
)

func (s Schedule) CreateSchedule(req schedule.CreateOrUpdateScheduleReq) error {

	// 解析字符串到 time.Time 类型
	startTime, _ := time.Parse(time.DateTime, *req.StartTime)
	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	if *req.Type == "course" {
		mpp, err := tx.MemberProductProperty.Query().
			Where(memberproductproperty.ID(req.MemberProductPropertyId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员产品属性")
			return rollback(tx, err)
		}
		req.PropertyId = mpp.PropertyID
	}

	property, err := s.db.ProductProperty.Query().
		Where(productproperty.ID(req.PropertyId)).
		First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到改课程属性")
		return err
	}

	venueName := ""
	if req.VenueId != 0 {
		ven, err := s.db.Venue.Query().
			Where(venue.ID(req.PropertyId)).
			First(s.ctx)
		if err == nil {
			venueName = ven.Name
		}
	}
	placeName := ""
	if req.PlaceID != 0 {
		place, err := s.db.VenuePlace.Query().
			Where(venueplace.ID(req.PropertyId)).
			First(s.ctx)
		if err == nil {
			placeName = place.Name
		}
	}

	one, err := tx.Schedule.Create().
		SetType(req.Type).
		SetStatus(1).
		SetVenueID(req.VenueId).
		SetPlaceID(req.PlaceID).
		SetPropertyID(req.PropertyId).
		SetNum(req.Num).
		SetNumSurplus(req.Num).
		SetDate(startTime.Format(time.DateOnly)).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(property.Length) * time.Minute)).
		SetPrice(req.Price).
		SetRemark(req.Remark).
		SetLength(property.Length).
		SetName(property.Name).
		SetVenueName(venueName).
		SetPlaceName(placeName).
		Save(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Course Record Schedule failed")
		return rollback(tx, err)
	}

	if req.CoachID != 0 {

		u, err := tx.User.Query().
			Where(user.ID(req.CoachID)).
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
			SetCoachID(req.CoachID).
			SetStartTime(startTime).
			SetEndTime(startTime.Add(time.Duration(property.Length) * time.Minute)).
			SetStatus(1).
			SetCoachName(u.Nickname).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return rollback(tx, err)
		}

	}

	if req.MemberID != 0 {
		m, err := tx.Member.Query().
			Where(member.ID(req.MemberID)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员")
			return rollback(tx, err)
		}
		mp, err := tx.MemberProduct.Query().
			Where(memberproduct.ID(req.MemberProductID)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员产品")
			return rollback(tx, err)
		}
		mpp, err := tx.MemberProductProperty.Query().
			Where(memberproductproperty.ID(req.MemberProductPropertyId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员产品属性")
			return rollback(tx, err)
		}
		if (mpp.CountSurplus - 1) < 0 {
			err = errors.Wrap(err, "该会员课程不足")
			return rollback(tx, err)
		}

		_, err = tx.ScheduleMember.Create().
			SetSchedule(one).
			SetType(req.Type).
			SetMemberID(req.MemberID).
			SetVenueID(req.VenueId).
			SetStartTime(startTime).
			SetEndTime(startTime.Add(time.Duration(property.Length) * time.Minute)).
			SetMemberProductID(req.MemberProductID).
			SetMemberProductPropertyID(req.MemberProductPropertyId).
			SetStatus(0).
			SetMemberName(m.Name).
			SetMemberProductName(mp.Name).
			SetMemberProductPropertyName(mpp.Name).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return rollback(tx, err)
		}

		_, err = tx.MemberProductProperty.
			Update().
			Where(memberproductproperty.IDEQ(req.MemberProductPropertyId)).
			AddCountSurplus(-1).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "Member Product Item failed")
			return rollback(tx, err)
		}

	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
