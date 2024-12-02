package do

import "saas/idl_gen/model/member"

type Member interface {
	CreateMember(req member.CreateOrUpdateMemberReq) error
	UpdateMember(req member.CreateOrUpdateMemberReq) error
	DeleteMember(id int64) error
	MemberList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error)
	MemberInfo(id int64) (info *member.MemberInfo, err error)
	ChangePassword(ID int64, oldPassword, newPassword string) error
	UpdateMemberStatus(ID int64, status int64) error
	MemberSearch(option string, value string) (memberInfo *member.MemberInfo, err error)

	MemberPrivacy(ID int64) (info *member.MemberPrivacy, err error)
	MemberNode(ID int64) (info *member.MemberNode, err error)

	//ProductSearch(members []int64) (info *ProductInfo, err error)
	//PropertySearch(memberProducts []int64) (info *PropertyInfo, err error)
	ContractList(req member.MemberContractListReq) (resp []*member.MemberContractInfo, total int, err error)
}

//base.NilResponse MemberProductList(1: MemberProductListReq req) (api.post = "/service/member/product-list")
//
//base.NilResponse MemberPropertyList(1: MemberPropertyListReq req) (api.post = "/service/member/property-list")
//
//base.NilResponse MemberProductDetail(1: base.IDReq req) (api.post = "/service/member/product-detail")
//
//base.NilResponse MemberPropertyDetail(1: base.IDReq req) (api.post = "/service/member/property-detail")
//
//base.NilResponse MemberPropertyUpdate(1: MemberPropertyListReq req) (api.post = "/service/member/property-update")
//
//base.NilResponse MemberProductSearch(1: MemberProductSearchReq req) (api.post = "/service/member/search-product")
//
//base.NilResponse MemberPropertySearch(1: MemberPropertySearchReq req) (api.post = "/service/member/search-property")
