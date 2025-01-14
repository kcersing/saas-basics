package service

import (
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/memberproduct"
	schedule2 "saas/biz/dal/db/ent/schedule"
	"saas/biz/dal/db/ent/venueplace"
	"time"

	"saas/idl_gen/model/schedule"
)

func (s Schedule) ScheduleInfo(ID int64) (resp *schedule.ScheduleInfo, err error) {
	one, err := s.db.Schedule.Query().Where(schedule2.IDEQ(ID)).First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get venue failed")
		return resp, err
	}
	resp = s.entScheduleInfo(one)

	coachAll, _ := one.QueryCoachs().All(s.ctx)
	if coachAll != nil {
		for _, v := range coachAll {
			resp.CoachCourseRecord = append(resp.CoachCourseRecord, s.entScheduleCoachInfo(v, one))
		}
	}

	memberAll, _ := one.QueryMembers().All(s.ctx)
	if memberAll != nil {
		for _, v := range memberAll {
			resp.MemberCourseRecord = append(resp.MemberCourseRecord, s.entScheduleMemberInfo(v, one))
		}
	}

	return
}

func (s Schedule) entScheduleMemberInfo(req *ent.ScheduleMember, sch *ent.Schedule) (info *schedule.ScheduleMemberInfo) {

	info = &schedule.ScheduleMemberInfo{
		ID:         req.ID,
		MemberId:   req.MemberID,
		VenueId:    req.VenueID,
		PlaceId:    req.PlaceID,
		ScheduleId: req.ScheduleID,

		Type:              req.Type,
		CreatedAt:         req.CreatedAt.Format(time.DateTime),
		UpdatedAt:         req.UpdatedAt.Format(time.DateTime),
		StartTime:         req.SignStartTime.Format("15:04"),
		EndTime:           req.EndTime.Format("15:04"),
		SignStartTime:     req.SignStartTime.Format("15:04"),
		SignEndTime:       req.SignEndTime.Format("15:04"),
		Status:            req.Status,
		MemberProductId:   req.MemberProductID,
		MemberName:        req.MemberName,
		MemberProductName: req.MemberProductName,

		Date: req.SignStartTime.Format(time.DateOnly),
		Seat: &req.Seat,
	}
	member, _ := s.db.Member.Query().Where(member2.ID(req.MemberID)).First(s.ctx)
	if member != nil {
		info.Mobile = member.Mobile
	}

	if sch != nil {
		sch, _ = req.QuerySchedule().First(s.ctx)
	}
	if sch != nil {
		info.ScheduleName = sch.Name
		info.PlaceName = sch.PlaceName
		info.VenueName = sch.VenueName
		info.ProductId = sch.ProductID
		coach, _ := sch.QueryCoachs().First(s.ctx)
		if coach != nil {
			info.CoachId = coach.CoachID
			info.CoachName = coach.CoachName
		}
	}

	mp, _ := s.db.MemberProduct.Query().Where(memberproduct.ID(req.MemberProductID)).First(s.ctx)
	if mp != nil {
		info.MemberProductSn = mp.Sn
	}
	return
}
func (s Schedule) entScheduleCoachInfo(req *ent.ScheduleCoach, sch *ent.Schedule) (info *schedule.ScheduleCoachInfo) {

	info = &schedule.ScheduleCoachInfo{
		ID:      req.ID,
		CoachId: req.CoachID,
		VenueId: req.VenueID,
		PlaceId: req.PlaceID,

		ScheduleId: req.ScheduleID,
		Type:       req.Type,
		CreatedAt:  req.CreatedAt.Format(time.DateTime),
		UpdatedAt:  req.UpdatedAt.Format(time.DateTime),

		StartTime:     req.StartTime.Format("15:04"),
		EndTime:       req.EndTime.Format("15:04"),
		SignStartTime: req.SignStartTime.Format("15:04"),
		SignEndTime:   req.SignEndTime.Format("15:04"),
		Status:        req.Status,
		Date:          req.SignStartTime.Format(time.DateOnly),

		CoachName: req.CoachName,
		ProductId: req.ProductID,
	}
	if sch == nil {
		sch, _ = req.QuerySchedule().First(s.ctx)
	}

	if sch != nil {
		info.ScheduleName = sch.Name
		info.VenueName = sch.VenueName
		info.PlaceName = sch.PlaceName
		m, _ := sch.QueryMembers().First(s.ctx)
		if m != nil {
			member, _ := s.db.Member.Query().Where(member2.ID(m.MemberID)).First(s.ctx)
			info.MemberId = m.MemberID
			info.MemberMobile = member.Mobile
			info.MemberName = m.MemberName
			info.MemberProductName = m.MemberProductName
			info.MemberProductId = m.MemberProductID
		}
	}

	return
}

func (s Schedule) entScheduleInfo(req *ent.Schedule) (info *schedule.ScheduleInfo) {

	info = &schedule.ScheduleInfo{
		ID:         req.ID,
		Type:       req.Type,
		VenueId:    req.VenueID,
		PlaceId:    req.PlaceID,
		Num:        req.Num,
		NumSurplus: req.NumSurplus,
		Date:       req.Date.Format(time.DateOnly),
		StartTime:  req.StartTime.Format("15:04"),
		EndTime:    req.EndTime.Format("15:04"),

		Name:      req.Name,
		ProductId: req.ProductID,
		Status:    req.Status,
		VenueName: req.VenueName,
		PlaceName: req.PlaceName,
		CreatedAt: req.CreatedAt.Format(time.DateTime),
		UpdatedAt: req.UpdatedAt.Format(time.DateTime),

		MemberCourseRecord: nil,
		CoachCourseRecord:  nil,
	}

	seats, _ := s.db.VenuePlace.Query().Where(venueplace.IDEQ(req.PlaceID)).First(s.ctx)
	if seats != nil {
		info.Seats = seats.Seat
	}
	coach, _ := req.QueryCoachs().First(s.ctx)
	if coach != nil {
		info.CoachId = coach.CoachID
		info.CoachName = coach.CoachName
	}

	return
}
