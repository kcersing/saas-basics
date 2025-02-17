package do

import (
	"saas/idl_gen/model/member"
	"saas/idl_gen/model/wx"
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

	ContestList(req member.MemberContestListReq) (resp []*member.MemberContestInfo, total int, err error)
	BootcampList(req member.MemberBootcampListReq) (resp []*member.MemberBootcampInfo, total int, err error)
	CommunityList(req member.MemberCommunityListReq) (resp []*member.MemberCommunityInfo, total int, err error)

	// wx
	Login(res *wx.MemberLoginReq) (info *member.MemberInfo, err error)
	Captcha(res *wx.MemberCaptchaReq) (code string, err error)
	Registe(res *wx.MemberRegisterReq) (info *member.MemberInfo, err error)
	Logout(id int64) (ok bool, err error)
}
