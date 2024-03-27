// Code generated by hertz generator.

package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"saas/app/admin/biz/handler/mw"
	"saas/app/admin/infras"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminMw() []app.HandlerFunc {
	// your code...

	return []app.HandlerFunc{
		mw.GetJWTMw(infras.CasbinEnforcer()).MiddlewareFunc(),
		mw.LogMw(),
	}

}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userprofileMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateprofileMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserstatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _user0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _changepasswordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userpermcodeMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _memberMw() []app.HandlerFunc {
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

func _memberprofileMw() []app.HandlerFunc {
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
