// Code generated by hertz generator.

package dictionary

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

func _adminMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		mw.GetJWTMw(infras.CasbinEnforcer()).MiddlewareFunc(),
		mw.LogMw(),
	}
}

func _dictMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletedictionaryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _detailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletedictionarydetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _detailbydictionarynameMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _dictionarylistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _detail0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createdictionarydetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatedictionarydetailMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _dict0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createdictionaryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatedictionaryMw() []app.HandlerFunc {
	// your code...
	return nil
}
