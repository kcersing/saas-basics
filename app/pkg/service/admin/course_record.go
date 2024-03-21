package admin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"saas/app/admin/config"
	"saas/app/admin/infras"
	"saas/app/pkg/do"
	"saas/pkg/db/ent"
	"saas/pkg/db/ent/memberproductproperty"

	"github.com/pkg/errors"
	"saas/pkg/db/ent/courserecordschedule"
)

type CourseRecord struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	Cbs   *casbin.Enforcer
	db    *ent.Client
	cache *ristretto.Cache
}

func NewCourseRecord(ctx context.Context, c *app.RequestContext) do.CourseRecord {
	return &CourseRecord{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MysqlInfo.Salt,
		db:    infras.DB,
		cache: infras.Cache,
		Cbs:   infras.CasbinEnforcer(),
	}
}

func (c CourseRecord) ScheduleCreate(req do.CreateOrUpdateScheduleReq) error {
	one, err := c.db.CourseRecordSchedule.Create().
		SetType(req.Type).
		SetVenueID(req.VenueId).
		SetPlaceID(req.PlaceID).
		SetNum(req.Num).
		SetStartTime(req.StartTime).
		SetEndTime(req.EndTime).
		SetPrice(req.Price).
		SetStatus(0).
		Save(c.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Course Record Schedule failed")
		return err
	}

	if req.CoachID != 0 {
		_, err = c.db.CourseRecordCoach.Create().
			SetSchedule(one).
			SetType(req.Type).
			SetVenueID(req.VenueId).
			SetCoachID(req.CoachID).
			SetStartTime(req.StartTime).
			SetEndTime(req.EndTime).
			SetStatus(0).
			Save(c.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return err
		}

	}

	if req.MemberID != 0 {
		_, err = c.db.CourseRecordMember.Create().
			SetSchedule(one).
			SetType(req.Type).
			SetMemberID(req.MemberID).
			SetVenueID(req.VenueId).
			SetStartTime(req.StartTime).
			SetEndTime(req.EndTime).
			SetMemberProductID(req.MemberProductID).
			SetMemberProductItemID(req.MemberProductItemID).
			SetCoachID(req.CoachID).
			SetStatus(0).
			Save(c.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return err
		}

		if req.MemberProductItemID != 0 {
			_, err = c.db.MemberProductProperty.
				Update().
				Where(memberproductproperty.IDEQ(req.MemberProductItemID)).
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

func (c CourseRecord) ScheduleUpdate(req do.CreateOrUpdateScheduleReq) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) ScheduleDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) ScheduleList(req do.ScheduleListReq) (resp []*do.ScheduleInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) ScheduleUpdateStatus(ID int64, status int64) error {
	_, err := c.db.CourseRecordSchedule.Update().Where(courserecordschedule.IDEQ(ID)).SetStatus(int64(status)).Save(c.ctx)
	return err
}

func (c CourseRecord) ScheduleInfo(ID int64) (roleInfo *do.ScheduleInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) CoachCreate(req do.CourseRecordCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) CoachUpdate(req do.CourseRecordCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) CoachDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) CoachList(req do.CourseRecordCoachListReq) (resp []*do.CourseRecordCoachInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) CoachUpdateStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) CoachInfo(ID int64) (roleInfo *do.CourseRecordCoachInfo, err error) {
	//TODO implement me
	panic("implement me")
}
func (c CourseRecord) MemberCreate(req do.CourseRecordMemberInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) MemberUpdate(req do.CourseRecordMemberInfo) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) MemberDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) MemberList(req do.CourseRecordMemberListReq) (resp []*do.CourseRecordMemberInfo, total int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) UpdateMemberStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c CourseRecord) MemberInfo(ID int64) (roleInfo *do.CourseRecordMemberInfo, err error) {
	//TODO implement me
	panic("implement me")
}
