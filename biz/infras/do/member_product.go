package do

import "saas/biz/dal/db/ent"

type MemberProduct interface {
	CreateMemberProduct(req CreateMemberProductReq) error
	//UpdateMemberProduct(req member.CreateOrUpdateMemberReq) error
	//MemberProductInfo(id int64) (info *member.MemberInfo, err error)
	//MemberProductList(req member.MemberListReq) (resp []*member.MemberInfo, total int, err error)
}

type CreateMemberProductReq struct {
	MemberId    int64
	ProductId   int64
	VenueId     int64
	Order       *ent.Order
	OrderAmount *ent.OrderAmount
	OrderItem   *ent.OrderItem
}
