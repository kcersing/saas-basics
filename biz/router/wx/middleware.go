// Code generated by hertz generator.

package wx

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
		mw.LogMw(),
	}
}

func _wxMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _activationMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _bootcampinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberbootcamplistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _bootcamplistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _buyMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _coachinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _coachlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _communityinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _communitylistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _membercommunitylistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _contestinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _contestlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _membercontestlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _membercontractlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _creatememberschedulecourseMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _creatememberschedulecourselessonsMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createplacescheduleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _joinbootcampMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _joincommunityMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _joincontestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberlogoutMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberorderlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _placeinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _placelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberproductinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _productinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberproductlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _productlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _scanqrMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _schedulememberlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _venueinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _venuelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _staffMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.GetJWTMw(casbin.CasbinEnforcer()).MiddlewareFunc(),
	}
}

func _mymemberMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _coachschedulelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _signmemberscheduleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _signstaffscheduleMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _membercaptchaMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberregisterMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _staffcaptchaMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _staffloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _stafflogoutMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _payMw() []app.HandlerFunc {
	// your code...
	return nil
}
