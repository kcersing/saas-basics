package do

import (
	"saas/idl_gen/model/member"
)

type Member interface {
	CreateMember(req member.CreateOrUpdateMemberReq) error
	UpdateMember(req member.CreateOrUpdateMemberReq) error
	DeleteMember(id []int64) error
	MemberInfo(id int64) (info *member.MemberInfo, err error)

	MemberFullList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error)
	MemberPotentialList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error)

	MemberFullListExport(req member.MemberListReq) (string, error)
	MemberPotentialListExport(req member.MemberListReq) (string, error)

	ChangePassword(ID int64, oldPassword, newPassword string) error

	UpdateMemberStatus(ID int64, status int64) error

	UpdateMemberFollow(req *member.UpdateMemberFollowReq) (err error)

	//ProductSearch(members []int64) (info *ProductInfo, err error)
	//PropertySearch(memberProducts []int64) (info *PropertyInfo, err error)
	ContractList(req member.MemberContractListReq) (resp []*member.MemberContractInfo, total int, err error)
}
