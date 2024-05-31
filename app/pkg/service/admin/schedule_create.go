package admin

import (
	"github.com/pkg/errors"
	"saas/app/pkg/do"
	"saas/pkg/db/ent/memberproductproperty"
	"saas/pkg/db/ent/productproperty"
	"time"
)

func (c Schedule) ScheduleCreate(req do.CreateOrUpdateScheduleReq) error {

	// 解析字符串到 time.Time 类型
	startTime, _ := time.Parse(time.DateTime, req.StartTime)
	data, _ := time.Parse(time.DateOnly, req.StartTime)
	property, err := c.db.ProductProperty.Query().
		Where(productproperty.ID(req.PropertyId)).
		First(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "为查询到改课程属性")
		return err
	}

	one, err := c.db.Schedule.Create().
		SetType(req.Type).
		SetStatus(0).
		SetVenueID(req.VenueId).
		SetPlaceID(req.PlaceID).
		SetPropertyID(req.PropertyId).
		SetNum(req.Num).
		SetDate(data).
		SetStartTime(startTime).
		SetEndTime(startTime.Add(time.Duration(property.Length) * time.Minute)).
		SetPrice(req.Price).
		SetRemark(req.Remark).
		SetName(property.Name).
		Save(c.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Course Record Schedule failed")
		return err
	}

	if req.CoachID != 0 {
		_, err = c.db.ScheduleCoach.Create().
			SetSchedule(one).
			SetType(req.Type).
			SetVenueID(req.VenueId).
			SetCoachID(req.CoachID).
			SetStartTime(startTime).
			SetEndTime(startTime.Add(time.Duration(property.Length) * time.Minute)).
			SetStatus(0).
			Save(c.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return err
		}

	}

	if req.MemberID != 0 {
		_, err = c.db.ScheduleMember.Create().
			SetSchedule(one).
			SetType(req.Type).
			SetMemberID(req.MemberID).
			SetVenueID(req.VenueId).
			SetStartTime(startTime).
			SetEndTime(startTime.Add(time.Duration(property.Length) * time.Minute)).
			SetMemberProductID(req.MemberProductID).
			SetMemberProductPropertyID(req.MemberProductPropertyId).
			SetStatus(0).
			Save(c.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return err
		}

		if req.MemberProductPropertyId != 0 {
			_, err = c.db.MemberProductProperty.
				Update().
				Where(memberproductproperty.IDEQ(req.MemberProductPropertyId)).
				AddCountSurplus(-1).
				Save(c.ctx)
			if err != nil {
				err = errors.Wrap(err, "Member Product Item failed")
				return err
			}
		}
	}

	return nil
}
