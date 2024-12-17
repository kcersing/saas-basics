// Code generated by hertz generator.

package member

import (
	"github.com/cloudwego/hertz/pkg/app"
	"saas/biz/dal/casbin"
	"saas/biz/handler/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _serviceMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.GetJWTMw(casbin.CasbinEnforcer()).MiddlewareFunc(),
		mw.LogMw(),
	}
}

func _memberMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _membercontractlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _creatememberMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _membernodeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberprivacyMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberproductdetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberproductlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberpropertydetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberpropertylistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberpropertyupdateMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _membersearchMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberproductsearchMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberpropertysearchMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatememberstatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatememberMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberfulllistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberpotentiallistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatememberfollowMw() []app.HandlerFunc {
	// your code...
	return nil
}
