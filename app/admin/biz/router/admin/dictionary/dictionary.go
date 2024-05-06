// Code generated by hertz generator. DO NOT EDIT.

package dictionary

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	dictionary "saas/app/admin/biz/handler/admin/dictionary"
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
			_admin.POST("/dict", append(_deletedictionaryMw(), dictionary.DeleteDictionary)...)
			_dict := _admin.Group("/dict", _dictMw()...)
			_dict.GET("/detail", append(_deletedictionarydetailMw(), dictionary.DeleteDictionaryDetail)...)
			_detail := _dict.Group("/detail", _detailMw()...)
			_detail.POST("/list", append(_detailbydictionarylistMw(), dictionary.DetailByDictionaryList)...)
			_dict.GET("/list", append(_dictionarylistMw(), dictionary.DictionaryList)...)
			{
				_detail0 := _dict.Group("/detail", _detail0Mw()...)
				_detail0.POST("/create", append(_createdictionarydetailMw(), dictionary.CreateDictionaryDetail)...)
				_detail0.POST("/update", append(_updatedictionarydetailMw(), dictionary.UpdateDictionaryDetail)...)
			}
			{
				_dict0 := _admin.Group("/dict", _dict0Mw()...)
				_dict0.POST("/create", append(_createdictionaryMw(), dictionary.CreateDictionary)...)
				_dict0.POST("/update", append(_updatedictionaryMw(), dictionary.UpdateDictionary)...)
			}
		}
	}
}
