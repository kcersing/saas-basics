// Code generated by hertz generator. DO NOT EDIT.

package role

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	role "saas/app/admin/biz/handler/admin/role"
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
			_admin.POST("/api", append(_deleteapiMw(), role.DeleteApi)...)
			_api0 := _admin.Group("/api", _api0Mw()...)
			_api0.POST("/list", append(_apilistMw(), role.ApiList)...)
			_api0.POST("/tree", append(_apitreeMw(), role.ApiTree)...)
			_admin.GET("/role", append(_rolebyidMw(), role.RoleByID)...)
			{
				_api1 := _admin.Group("/api", _api1Mw()...)
				_api1.POST("/create", append(_createapiMw(), role.CreateApi)...)
				_api1.POST("/update", append(_updateapiMw(), role.UpdateApi)...)
			}
			{
				_authority := _admin.Group("/authority", _authorityMw()...)
				{
					_api2 := _authority.Group("/api", _api2Mw()...)
					_api2.POST("/create", append(_createauthorityMw(), role.CreateAuthority)...)
					_api2.POST("/role", append(_apiauthorityMw(), role.ApiAuthority)...)
					_api2.POST("/update", append(_updateapiauthorityMw(), role.UpdateApiAuthority)...)
				}
				{
					_menu := _authority.Group("/menu", _menuMw()...)
					_menu.POST("/create", append(_createmenuauthorityMw(), role.CreateMenuAuthority)...)
					_menu.POST("/role", append(_menuauthorityMw(), role.MenuAuthority)...)
					_menu.POST("/update", append(_updatemenuauthorityMw(), role.UpdateMenuAuthority)...)
				}
			}
			_admin.POST("/role", append(_deleteroleMw(), role.DeleteRole)...)
			_role := _admin.Group("/role", _roleMw()...)
			_role.GET("/list", append(_rolelistMw(), role.RoleList)...)
			_role.POST("/status", append(_updaterolestatusMw(), role.UpdateRoleStatus)...)
			{
				_role0 := _admin.Group("/role", _role0Mw()...)
				_role0.POST("/create", append(_createroleMw(), role.CreateRole)...)
				_role0.POST("/update", append(_updateroleMw(), role.UpdateRole)...)
			}
		}
	}
}
