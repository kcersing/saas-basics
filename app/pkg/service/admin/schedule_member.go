package admin

import (
	"saas/app/pkg/do"
)

func (c Schedule) MemberCreate(req do.ScheduleMemberInfo) error {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (c Schedule) UpdateMemberStatus(ID int64, status int64) error {
	//TODO implement me
	panic("implement me")
}

func (c Schedule) MemberInfo(ID int64) (roleInfo *do.ScheduleMemberInfo, err error) {
	//TODO implement me
	panic("implement me")
}
