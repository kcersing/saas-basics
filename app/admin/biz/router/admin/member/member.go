// Code generated by hertz generator. DO NOT EDIT.

package member

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	member "saas/app/admin/biz/handler/admin/member"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_admin := _api.Group("/admin", _adminMw()...)
			{
				_member := _admin.Group("/member", _memberMw()...)
				_member.POST("/contract-list", append(_membercontractlistMw(), member.MemberContractList)...)
				_member.POST("/create", append(_creatememberMw(), member.CreateMember)...)
				_member.GET("/info", append(_memberinfoMw(), member.MemberInfo)...)
				_member.POST("/list", append(_memberlistMw(), member.MemberList)...)
				_member.GET("/product-detail", append(_memberproductdetailMw(), member.MemberProductDetail)...)
				_member.POST("/product-list", append(_memberproductlistMw(), member.MemberProductList)...)
				_member.GET("/property-detail", append(_memberpropertydetailMw(), member.MemberPropertyDetail)...)
				_member.POST("/property-list", append(_memberpropertylistMw(), member.MemberPropertyList)...)
				_member.POST("/property-update", append(_memberpropertyupdateMw(), member.MemberPropertyUpdate)...)
				_member.POST("/search", append(_membersearchMw(), member.MemberSearch)...)
				_member.POST("/status", append(_updatememberstatusMw(), member.UpdateMemberStatus)...)
				_member.POST("/update", append(_updatememberMw(), member.UpdateMember)...)
			}
		}
	}
}
