package service

import (
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	"time"

	"saas/idl_gen/model/schedule"
)

func (s Schedule) ScheduleInfo(ID int64) (roleInfo *schedule.ScheduleInfo, err error) {
	//TODO implement me
	panic("implement me")
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
		StartTime:       req.SignStartTime.Format(time.DateTime),
		EndTime:         req.EndTime.Format(time.DateTime),
		SignStartTime:   req.SignStartTime.Format(time.DateTime),
		SignEndTime:     req.SignEndTime.Format(time.DateTime),
		Status:          req.Status,
		MemberProductId: req.MemberProductID,

		//VenueName:         req2.VenueName,
		//PlaceName:         req2.PlaceName,
		MemberName:        req.MemberName,
		MemberProductName: req.MemberProductName,

		Mobile: member.Mobile,
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

		StartTime:     req.StartTime.Format(time.DateTime),
		EndTime:       req.EndTime.Format(time.DateTime),
		SignStartTime: req.SignStartTime.Format(time.DateTime),
		SignEndTime:   req.SignEndTime.Format(time.DateTime),
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
	//coach, _ := req.QueryCoachs().First(s.ctx)
	return &schedule.ScheduleInfo{
		ID:         req.ID,
		Type:       req.Type,
		VenueId:    req.VenueID,
		PlaceId:    req.PlaceID,
		Num:        req.Num,
		NumSurplus: req.NumSurplus,
		Date:       req.Date,
		StartTime:  req.StartTime.Format(time.DateTime),
		EndTime:    req.EndTime.Format(time.DateTime),

		Name: req.Name,
		//CoachId: coach.CoachID,
		//Price:              0,
		//MemberId:           0,
		//MemberProductId:    0,
		//MemberName:         "",
		//MemberProductName:  "",
		//CoachName: coach.CoachName,

		Status:             req.Status,
		VenueName:          req.VenueName,
		PlaceName:          req.PlaceName,
		MemberCourseRecord: nil,
		CoachCourseRecord:  nil,
		CreatedAt:          req.CreatedAt.Format(time.DateTime),
		UpdatedAt:          req.UpdatedAt.Format(time.DateTime),
	}
}
