package admin

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"saas/app/pkg/do"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/schedule"
	"saas/pkg/db/ent/schedulemember"
)

func (c Schedule) MemberCreate(req do.ScheduleMemberCreate) error {

	sc, err := c.db.Schedule.Query().
		Where(schedule.ID(req.Schedule)).
		First(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到该课程")
		return err
	}
	for _, v := range req.Member {

		m, err := c.db.Member.Query().
			Where(member.ID(v)).
			First(c.ctx)
		if err == nil {
			c.db.ScheduleMember.Create().
				SetSchedule(sc).
				SetMemberName(m.Name).
				SetStatus(1).
				SetStartTime(sc.StartTime).
				SetEndTime(sc.EndTime).
				SetMemberID(v).
				SetType(sc.Type).
				SetVenueID(sc.VenueID).
				SetMemberProductID().
				SetMemberProductName().
				SetMemberProductPropertyID().
				SetMemberProductPropertyName().
				Save(c.ctx)
		}

	}

	return nil

}

func (c Schedule) MemberUpdate(req do.ScheduleMemberInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) MemberDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) MemberList(req do.ScheduleMemberListReq) (resp []*do.ScheduleMemberInfo, total int, err error) {

	var predicates []predicate.ScheduleMember
	if req.Member > 0 {
		predicates = append(predicates, schedulemember.MemberID(req.Member))
	}
	if req.Schedule > 0 {
		predicates = append(predicates, schedulemember.ScheduleID(req.Member))
	}

	lists, err := c.db.ScheduleMember.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Member list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Schedule Member info failed")
		return resp, 0, err
	}
	total, _ = c.db.ScheduleMember.Query().Where(predicates...).Count(c.ctx)
	return

}

func (c Schedule) UpdateMemberStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) MemberInfo(ID int64) (roleInfo *do.ScheduleMemberInfo, err error) {
	//TODO implement me
	panic("implement me")
}
