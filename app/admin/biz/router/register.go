// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	admin_admin "saas/app/admin/biz/router/admin/admin"
	admin_dictionary "saas/app/admin/biz/router/admin/dictionary"
	admin_logs "saas/app/admin/biz/router/admin/logs"
	admin_member "saas/app/admin/biz/router/admin/member"
	admin_menu "saas/app/admin/biz/router/admin/menu"
	admin_order "saas/app/admin/biz/router/admin/order"
	admin_product "saas/app/admin/biz/router/admin/product"
	admin_pub "saas/app/admin/biz/router/admin/pub"
	admin_role "saas/app/admin/biz/router/admin/role"
	admin_sys "saas/app/admin/biz/router/admin/sys"
	admin_token "saas/app/admin/biz/router/admin/token"
	admin_user "saas/app/admin/biz/router/admin/user"
	admin_venue "saas/app/admin/biz/router/admin/venue"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	admin_sys.Register(r)

	admin_pub.Register(r)

	admin_member.Register(r)

	admin_order.Register(r)

	admin_venue.Register(r)

	admin_product.Register(r)

	admin_role.Register(r)

	admin_menu.Register(r)

	admin_logs.Register(r)

	admin_dictionary.Register(r)

	admin_admin.Register(r)

	admin_user.Register(r)

	admin_token.Register(r)
}
