package service

import (
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/schedulemember"
	"saas/idl_gen/model/schedule"
)

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
		k := s.entScheduleMemberInfo(v)
		f, _ := v.QuerySchedule().First(s.ctx)
		coach, _ := f.QueryCoachs().First(s.ctx)
		mp, _ := s.db.MemberProduct.Query().Where(memberproduct.ID(v.MemberProductID)).First(s.ctx)
		resp = append(resp, k)
		k.ScheduleName = f.Name
		k.PlaceName = f.PlaceName
		k.VenueName = f.VenueName
		k.MemberProductSn = mp.Sn
		k.CoachId = coach.CoachID
		k.CoachName = coach.CoachName
		k.ProductId = f.ProductID
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
