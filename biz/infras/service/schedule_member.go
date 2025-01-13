package service

import (
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/predicate"
	schedule2 "saas/biz/dal/db/ent/schedule"
	"saas/biz/dal/db/ent/schedulemember"
	"saas/idl_gen/model/schedule"
	"time"
)

func (s Schedule) ScheduleMemberList(req schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int, err error) {
	var predicates []predicate.ScheduleMember

	if req.StartTime != "" && req.EndTime != "" {
		startTime, _ := time.Parse(time.DateOnly, req.StartTime)
		endTime, _ := time.Parse(time.DateOnly, req.EndTime)

		predicates = append(predicates, schedulemember.DateGTE(startTime))
		predicates = append(predicates, schedulemember.DateLTE(endTime))
	}
	if len(req.Status) > 0 {
		predicates = append(predicates, schedulemember.StatusIn(req.Status...))
	}
	if req.MemberId > 0 {
		predicates = append(predicates, schedulemember.MemberID(req.MemberId))
	}
	if req.ScheduleId > 0 {
		predicates = append(predicates, schedulemember.ScheduleID(req.ScheduleId))
	}
	if req.Type != "" {
		predicates = append(predicates, schedulemember.Type(req.Type))

	}
	if req.Name != "" {
		predicates = append(predicates, schedulemember.HasScheduleWith(schedule2.NameEQ(req.Name)))
	}
	if req.VenueId == 0 {
		return resp, total, errors.New("场馆ID不能为空")
	}
	predicates = append(predicates, schedulemember.VenueIDEQ(req.VenueId))

	predicates = append(predicates, schedulemember.HasScheduleWith(schedule2.StatusNotIn(0, 5)))

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
