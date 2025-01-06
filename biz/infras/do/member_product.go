package do

import (
	"saas/biz/dal/db/ent"
	memberProduct "saas/idl_gen/model/member/memberProduct"
)

type MemberProduct interface {
	CreateMemberProduct(req CreateMemberProductReq) error
	//UpdateMemberProduct(req member.CreateOrUpdateMemberReq) error
	MemberProductInfo(id int64) (info *memberProduct.MemberProductInfo, err error)
	MemberProductList(req memberProduct.MemberProductListReq) (resp []*memberProduct.MemberProductInfo, total int, err error)
}

type CreateMemberProductReq struct {
	MemberId    int64
	ProductId   int64
	VenueId     int64
	Order       *ent.Order
	OrderAmount *ent.OrderAmount
	OrderItem   *ent.OrderItem
}
