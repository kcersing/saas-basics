// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	admin "saas/biz/router/admin"
	auth "saas/biz/router/auth"
	banner "saas/biz/router/banner"
	bootcamp "saas/biz/router/bootcamp"
	community "saas/biz/router/community"
	order "saas/biz/router/order"
	payment "saas/biz/router/payment"
	product "saas/biz/router/product"
	schedule "saas/biz/router/schedule"
	sms "saas/biz/router/sms"
	sys "saas/biz/router/sys"

	contest "saas/biz/router/contest"
	contract "saas/biz/router/contract"
	dictionary "saas/biz/router/dictionary"
	entry "saas/biz/router/entry"
	member "saas/biz/router/member"
	menu "saas/biz/router/menu"
	pub "saas/biz/router/pub"
	token "saas/biz/router/token"
	user "saas/biz/router/user"
	venue "saas/biz/router/venue"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	sys.Register(r)

	schedule.Register(r)

	sms.Register(r)

	product.Register(r)

	community.Register(r)

	bootcamp.Register(r)

	order.Register(r)

	payment.Register(r)

	banner.Register(r)

	entry.Register(r)

	contract.Register(r)

	member.Register(r)

	venue.Register(r)

	token.Register(r)

	user.Register(r)

	pub.Register(r)

	admin.Register(r)

	dictionary.Register(r)

	menu.Register(r)

	auth.Register(r)

	contest.Register(r)
}
