// Code generated by hertz generator.

package contest

import (
	"github.com/cloudwego/hertz/pkg/app"
	"saas/biz/dal/casbin"
	"saas/biz/handler/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _contestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createcontestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createparticipantMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _contestinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _contestparticipantinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _contestlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _participantlistlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatememberstatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateparticipantstatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatecontestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateparticipantMw() []app.HandlerFunc {
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

func _participantMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _promotionallinksMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _resultsuploadMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatecontestshowMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateconteststatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _delcontestMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _participantlistlistexportMw() []app.HandlerFunc {
	// your code...
	return nil
}
