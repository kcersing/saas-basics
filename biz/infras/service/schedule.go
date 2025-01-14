package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"saas/biz/dal/cache"
	"saas/biz/dal/db"
	"saas/biz/dal/db/ent"
	member2 "saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/predicate"
	schedule2 "saas/biz/dal/db/ent/schedule"
	"saas/biz/dal/db/ent/schedulecoach"
	"saas/biz/dal/db/ent/schedulemember"
	"saas/pkg/enums"
	"time"

	"saas/biz/infras/do"
	"saas/config"

	"saas/idl_gen/model/schedule"
)

type Schedule struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (s Schedule) ScheduleLessonsPublish(ids []int64) error {
	_, err := s.db.Schedule.Update().
		Where(schedule2.IDIn(ids...)).
		SetStatus(2).Save(s.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s Schedule) DeleteSchedule(id int64) error {
	_, err := s.db.Schedule.Update().Where(schedule2.IDEQ(id)).SetDelete(1).Save(s.ctx)
	return err
}

func (s Schedule) ScheduleList(req schedule.ScheduleListReq, isSubList bool) (resp []*schedule.ScheduleInfo, total int, err error) {
	var predicates []predicate.Schedule

	if req.StartTime != "" && req.EndTime != "" {
		startTime, _ := time.Parse(time.DateOnly, req.StartTime)
		endTime, _ := time.Parse(time.DateOnly, req.EndTime)

		predicates = append(predicates, schedule2.DateGTE(startTime))
		predicates = append(predicates, schedule2.DateLTE(endTime))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, schedule2.VenueID(req.VenueId))
	}
	if len(req.CoachId) > 0 {
		predicates = append(predicates, schedule2.HasCoachsWith(schedulecoach.CoachIDIn(req.CoachId...)))
	}

	if len(req.MemberId) > 0 {
		predicates = append(predicates, schedule2.HasMembersWith(schedulemember.ScheduleIDIn(req.MemberId...)))
	}
	if len(req.ProductId) > 0 {
		predicates = append(predicates, schedule2.ProductIDIn(req.ProductId...))
	}

	if req.Type != "" {
		if req.Type == enums.Lessons {
			predicates = append(predicates, schedule2.TypeEQ(req.Type))
		}
		if req.Type == enums.Course {
			predicates = append(predicates, schedule2.TypeIn(enums.CourseOne, enums.CourseMore))
		}
	}
	if len(req.SubType) > 0 {
		predicates = append(predicates, schedule2.TypeIn(req.SubType...))
	}

	if len(req.Status) > 0 {
		predicates = append(predicates, schedule2.StatusIn(req.Status...))
	}

	if req.Name != "" {
		predicates = append(predicates, schedule2.NameEQ(req.Name))
	}
	if req.VenueId == 0 {
		return resp, total, errors.New("场馆ID不能为空")
	}

	if req.MemberName != "" {
		predicates = append(predicates, schedule2.HasMembersWith(schedulemember.MemberNameEQ(req.MemberName)))
	}
	if req.MemberMobile != "" {
		first, _ := s.db.Member.Query().Where(
			member2.MobileEQ(req.MemberMobile),
			member2.Delete(0),
		).First(s.ctx)
		if first == nil {
			return nil, 0, errors.New("未找到符合条件的会员")
		}
		predicates = append(predicates, schedule2.HasMembersWith(schedulemember.MemberIDEQ(first.ID)))
	}

	predicates = append(predicates, schedule2.VenueIDEQ(req.VenueId))

	predicates = append(predicates, schedule2.StatusNotIn(0, 5))

	lists, err := s.db.Schedule.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule list failed")
		return resp, total, err
	}
	for _, v := range lists {

		sc := s.entScheduleInfo(v)

		if isSubList {
			coachList, _, _ := s.ScheduleCoachList(schedule.ScheduleCoachListReq{
				Page:       1,
				PageSize:   999,
				ScheduleId: sc.ID,
			})

			list, _, _ := s.ScheduleMemberList(schedule.ScheduleMemberListReq{
				Page:       1,
				PageSize:   999,
				ScheduleId: sc.ID,
			})

			sc.CoachCourseRecord = coachList
			sc.MemberCourseRecord = list
		}

		resp = append(resp, sc)
	}

	total, _ = s.db.Schedule.Query().Where(predicates...).Count(s.ctx)
	return
}

func (s Schedule) ScheduleDateList(req schedule.ScheduleListReq) (map[string][]*schedule.ScheduleInfo, int, error) {
	req.Page = 1
	req.PageSize = 1000
	req.Status = []int64{1, 2, 3, 4}
	req.Type = enums.Lessons
	lists, total, err := s.ScheduleList(req, false)
	m := make(map[string][]*schedule.ScheduleInfo)
	for _, v := range lists {
		m[v.Date] = append(m[v.Date], v)
	}

	return m, total, err
}

func (s Schedule) UpdateScheduleStatus(ID int64, status int64) error {
	_, err := s.db.Schedule.Update().Where(schedule2.IDEQ(ID)).SetStatus(status).Save(s.ctx)
	return err
}

func NewSchedule(ctx context.Context, c *app.RequestContext) do.Schedule {
	return &Schedule{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
