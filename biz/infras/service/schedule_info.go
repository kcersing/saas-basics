package service

import (
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	schedule2 "saas/biz/dal/db/ent/schedule"
	"saas/biz/dal/db/ent/venueplace"
	"saas/idl_gen/model/base"
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
			resp.CoachCourseRecord = append(resp.CoachCourseRecord, s.entScheduleCoachInfo(v))
		}
	}

	memberAll, _ := one.QueryMembers().All(s.ctx)
	if memberAll != nil {
		for _, v := range memberAll {
			resp.MemberCourseRecord = append(resp.MemberCourseRecord, s.entScheduleMemberInfo(v))
		}
	}

	return
}

func (s Schedule) entScheduleMemberInfo(req *ent.ScheduleMember) (info *schedule.ScheduleMemberInfo) {
	member, _ := s.db.Member.Query().Where(member2.ID(req.MemberID)).First(s.ctx)

	return &schedule.ScheduleMemberInfo{
		ID:              req.ID,
		MemberId:        req.MemberID,
		VenueId:         req.VenueID,
		PlaceId:         req.PlaceID,
		ScheduleId:      req.ScheduleID,
		ScheduleName:    req.ScheduleName,
		Type:            req.Type,
		CreatedAt:       req.CreatedAt.Format(time.DateTime),
		UpdatedAt:       req.UpdatedAt.Format(time.DateTime),
		StartTime:       req.SignStartTime.Format("15:04"),
		EndTime:         req.EndTime.Format("15:04"),
		SignStartTime:   req.SignStartTime.Format("15:04"),
		SignEndTime:     req.SignEndTime.Format("15:04"),
		Status:          req.Status,
		MemberProductId: req.MemberProductID,

		//VenueName:         req2.VenueName,
		//PlaceName:         req2.PlaceName,
		MemberName:        req.MemberName,
		MemberProductName: req.MemberProductName,

		Mobile: member.Mobile,

		Date:            req.SignStartTime.Format(time.DateOnly),
		MemberProductSn: "",
		CoachId:         0,
		CoachName:       "",
		Seat:            &req.Seat,
	}
}
func (s Schedule) entScheduleCoachInfo(req *ent.ScheduleCoach) (info *schedule.ScheduleCoachInfo) {
	return &schedule.ScheduleCoachInfo{
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
		ScheduleName:  req.ScheduleName,
		CoachName:     req.CoachName,

		ProductId: req.ProductID,
		//Date:      req.Date,
		//ProductName:   req.ProductName,
		//VenueName:     r,
		//PlaceName:     "",
		//CoachAvatar:   req.CoachAvatar,
	}
}

func (s Schedule) entScheduleInfo(req *ent.Schedule) (info *schedule.ScheduleInfo) {
	var seat [][]*base.Seat
	seats, _ := s.db.VenuePlace.Query().Where(venueplace.IDEQ(req.PlaceID)).First(s.ctx)
	if seats != nil {
		seat = seats.Seat
	}
	coach, _ := req.QueryCoachs().First(s.ctx)
	return &schedule.ScheduleInfo{
		ID:         req.ID,
		Type:       req.Type,
		VenueId:    req.VenueID,
		PlaceId:    req.PlaceID,
		Num:        req.Num,
		NumSurplus: req.NumSurplus,
		Date:       req.Date,
		StartTime:  req.StartTime.Format("15:04"),
		EndTime:    req.EndTime.Format("15:04"),
		Seats:      seat,
		Name:       req.Name,
		CoachId:    coach.CoachID,
		//Price:              0,
		//MemberId:           0,
		//MemberProductId:    0,
		//MemberName:         "",
		//MemberProductName:  "",
		CoachName:          coach.CoachName,
		ProductId:          req.ProductID,
		Status:             req.Status,
		VenueName:          req.VenueName,
		PlaceName:          req.PlaceName,
		MemberCourseRecord: nil,
		CoachCourseRecord:  nil,
		CreatedAt:          req.CreatedAt.Format(time.DateTime),
		UpdatedAt:          req.UpdatedAt.Format(time.DateTime),
	}
}
