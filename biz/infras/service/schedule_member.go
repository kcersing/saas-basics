package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/memberproduct"
	schedule2 "saas/biz/dal/db/ent/schedule"
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
	m, err := s.db.Member.Query().
		Where(member.ID(req.MemberId)).
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
	//TODO implement me
	panic("implement me")
}

func (s Schedule) UpdateScheduleMemberStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (s Schedule) ScheduleMemberInfo(ID int64) (roleInfo *schedule.ScheduleMemberInfo, err error) {
	//TODO implement me
	panic("implement me")
}
