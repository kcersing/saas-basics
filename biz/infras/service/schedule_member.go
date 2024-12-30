package service

import "saas/idl_gen/model/schedule"

func (s Schedule) CreateMemberSubscribeLessons(req schedule.MemberSubscribeReq) error {
	//TODO implement me
	panic("implement me")
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
