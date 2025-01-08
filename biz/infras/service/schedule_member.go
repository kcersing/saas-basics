package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/memberprofile"
	"saas/biz/dal/db/ent/predicate"
	schedule2 "saas/biz/dal/db/ent/schedule"
	"saas/biz/dal/db/ent/schedulemember"
	"saas/idl_gen/model/schedule"
)

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
		SetVenueID(one.VenueID).
		SetStartTime(one.StartTime).
		SetEndTime(one.StartTime).
		SetMemberProductID(req.MemberProductId).
		SetStatus(0).
		SetMemberName(m.Name).
		SetMemberProductName(mp.Name).
		SetSeat(*req.Seat).
		SetPlaceID(one.PlaceID).
		Save(s.ctx)
	if err != nil {
		hlog.Error("无法创建会员课程:", req)
		err = errors.Wrap(err, "无法创建会员课程")
		return err
	}
	return nil

}

func (s Schedule) ScheduleMemberList(req schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int, err error) {
	var predicates []predicate.ScheduleMember

	//if req.StartTime != "" {
	//	startTime, _ := time.Parse(time.DateOnly, req.StartTime)
	//	//大于
	//	predicates = append(predicates, schedulemember.StartTimeLTE(startTime))
	//	//小于
	//	predicates = append(predicates, schedulemember.EndTimeGTE(startTime.Add(7*24*time.Hour)))
	//}
	if req.MemberId > 0 {
		predicates = append(predicates, schedulemember.MemberID(req.MemberId))
	}
	if req.ScheduleId > 0 {
		predicates = append(predicates, schedulemember.ScheduleID(req.ScheduleId))
	}
	if req.Type != "" {
		predicates = append(predicates, schedulemember.Type(req.Type))

	}
	lists, err := s.db.ScheduleMember.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule member list failed")
		return resp, total, err
	}
	for _, v := range lists {
		resp = append(resp, s.entScheduleMemberInfo(v))

	}

	total, _ = s.db.ScheduleMember.Query().Where(predicates...).Count(s.ctx)
	return
}

func (s Schedule) UpdateScheduleMemberStatus(ID int64, status int64) error {
	_, err := s.db.ScheduleMember.Update().Where(schedulemember.ID(ID)).SetStatus(status).Save(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s Schedule) ScheduleMemberInfo(ID int64) (roleInfo *schedule.ScheduleMemberInfo, err error) {
	first, err := s.db.ScheduleMember.Query().Where(schedulemember.ID(ID)).First(s.ctx)
	if err != nil {
		return nil, err
	}
	roleInfo = s.entScheduleMemberInfo(first)
	return
}
