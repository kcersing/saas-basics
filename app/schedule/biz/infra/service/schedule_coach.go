package service

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"rpc_gen/kitex_gen/schedule"
	"schedule/biz/dal/mysql/ent/predicate"
	schedule2 "schedule/biz/dal/mysql/ent/schedule"
	"schedule/biz/dal/mysql/ent/schedulecoach"
	"time"
)

func (c Schedule) CreateCoach(req schedule.ScheduleCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) UpdateCoach(req schedule.ScheduleCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) DeleteCoach(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) CoachList(req schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int64, err error) {

	var predicates []predicate.ScheduleCoach
	if *req.Coach > 0 {
		predicates = append(predicates, schedulecoach.CoachIDEQ(*req.Coach))
	}
	if *req.Schedule > 0 {
		predicates = append(predicates, schedulecoach.ScheduleID(*req.Schedule))
	}
	if *req.Type != "" {
		predicates = append(predicates, schedulecoach.Type(*req.Type))
	}
	lists, err := c.db.ScheduleCoach.Query().Where(predicates...).
		Offset(int(*req.Page-1) * int(*req.PageSize)).
		Limit(int(*req.PageSize)).All(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Coach list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Schedule Coach info failed")
		return resp, 0, err
	}
	for i, v := range lists {
		vInfo, err := NewVenue(c.ctx, c.c).VenueInfo(v.VenueID)
		if err == nil {
			resp[i].VenueName = vInfo.Name
		}
		coach, err := NewUser(c.ctx, c.c).Info(v.CoachID)
		if err == nil {
			resp[i].CoachAvatar = coach.Avatar
		}

		s, err := v.QuerySchedule().First(c.ctx)
		if err == nil {
			resp[i].ScheduleName = s.Name
			resp[i].VenueName = s.VenueName
			resp[i].PlaceName = s.PlaceName
			resp[i].Remark = s.Remark
		}
		sm, err := s.QueryMembers().First(c.ctx)
		if err == nil {
			resp[i].MRemark = sm.Remark
			resp[i].MemberName = sm.MemberName
			resp[i].MemberProductName = sm.MemberProductName
			resp[i].MemberProductPropertyName = sm.MemberProductPropertyName

			m, err := NewMember(c.ctx, c.c).Info(sm.MemberID)
			if err == nil {
				resp[i].Mobile = m.Mobile
				resp[i].MemberAvatar = m.Avatar
			}
		}

		resp[i].CreatedAt = v.CreatedAt.Format(time.DateTime)
		resp[i].StartTime = v.StartTime.Format(time.TimeOnly)
		resp[i].EndTime = v.EndTime.Format(time.TimeOnly)
		resp[i].SignStartTime = v.SignStartTime.Format(time.TimeOnly)
		resp[i].SignEndTime = v.SignEndTime.Format(time.TimeOnly)
		resp[i].Date = v.StartTime.Format(time.DateOnly)

	}

	total, _ = c.db.ScheduleCoach.Query().Where(predicates...).Count(c.ctx)
	return

}

func (c Schedule) UpdateCoachStatus(ID int64, status int64) error {
	sc, err := c.db.ScheduleCoach.Query().Where().First(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Coach info failed")
		return err
	}
	_, err = c.db.ScheduleCoach.Update().Where(schedulecoach.ID(ID)).SetStatus(status).Save(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Schedule Coach status failed")
		return err
	}
	_, err = c.db.Schedule.Update().Where(schedule2.ID(sc.ScheduleID)).SetStatus(status).Save(c.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Schedule status failed")
		return err
	}

	return nil
}

func (c Schedule) CoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error) {
	//TODO implement me
	panic("implement me")
}
