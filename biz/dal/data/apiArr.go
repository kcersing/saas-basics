package data

type api struct {
	Method      string
	Path        string
	Description string
	Group       string
	Title       string
}

var ApiArr = []api{

	{Method: "POST", Path: "/api/admin/entry/list", Description: "member.EntryList", Title: "进馆记录", Group: "entry"},

	{Method: "POST", Path: "/api/admin/contract", Description: "member.ContractByID", Title: "合同详情", Group: "contract"},
	{Method: "POST", Path: "/api/admin/contract/create", Description: "member.CreateMember", Title: "创建合同", Group: "contract"},
	{Method: "POST", Path: "/api/admin/contract/list", Description: "member.CreateMember", Title: "合同列表", Group: "contract"},
	{Method: "POST", Path: "/api/admin/contract/status", Description: "member.CreateMember", Title: "更新合同状态", Group: "contract"},
	{Method: "POST", Path: "/api/admin/contract/update", Description: "member.CreateMember", Title: "更新合同", Group: "contract"},

	{Method: "POST", Path: "/api/admin/member/create", Description: "member.CreateMember", Title: "创建会员", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/info", Description: "member.MemberInfo", Title: "会员详情", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/list", Description: "member.MemberList", Title: "会员列表", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/status", Description: "member.UpdateMemberStatus", Title: "更新会员状态", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/update", Description: "member.UpdateMember", Title: "更新会员", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/search", Description: "member.MemberSearch", Title: "查询会员", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/contract-list", Description: "member.MemberContractList", Title: "会员合同列表", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/product-detail", Description: "member.MemberProductDetail", Title: "会员产品详情", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/product-list", Description: "member.MemberProductList", Title: "会员产品列表", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/property-detail", Description: "member.MemberPropertyDetail", Title: "会员产品属性详情", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/property-list", Description: "member.MemberPropertyList", Title: "会员产品属性列表", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/property-update", Description: "member.MemberPropertyUpdate", Title: "更新会员产品属性", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/search-product", Description: "member.MemberProductSearch", Title: "搜索会员产品", Group: "member"},
	{Method: "POST", Path: "/api/admin/member/search-property", Description: "member.MemberPropertySearch", Title: "搜索会员产品属性", Group: "member"},

	{Method: "POST", Path: "/api/admin/order/create", Description: "order.CreateOrder", Title: "创建订单", Group: "order"},
	{Method: "GET", Path: "/api/admin/order/info", Description: "order.GetOrderById", Title: "订单详情", Group: "order"},
	{Method: "POST", Path: "/api/admin/order/list", Description: "order.ListOrder", Title: "订单列表", Group: "order"},
	{Method: "POST", Path: "/api/admin/order/status", Description: "order.UpdateStatus", Title: "更新订单状态", Group: "order"},
	{Method: "POST", Path: "/api/admin/order/update", Description: "order.UpdateOrder", Title: "更新订单", Group: "order"},

	{Method: "POST", Path: "/api/admin/order/unifyPay", Description: "order.UnifyPay", Title: "订单支付", Group: "order"},
	{Method: "POST", Path: "/api/admin/order/QRPay", Description: "order.QRPay", Title: "订单支付二维码", Group: "order"},

	{Method: "POST", Path: "/api/admin/place/create", Description: "venue.CreatePlace", Title: "创建场地", Group: "venue"},
	{Method: "POST", Path: "/api/admin/place/list", Description: "venue.PlaceList", Title: "场地列表", Group: "venue"},
	{Method: "POST", Path: "/api/admin/place/status", Description: "venue.PlaceUpdateStatus", Title: "更新场地状态", Group: "venue"},
	{Method: "POST", Path: "/api/admin/place/update", Description: "venue.UpdatePlace", Title: "更新场地", Group: "venue"},
	{Method: "POST", Path: "/api/admin/venue/create", Description: "venue.Create", Title: "创建场馆", Group: "venue"},
	{Method: "POST", Path: "/api/admin/venue/list", Description: "venue.List", Title: "场馆列表", Group: "venue"},
	{Method: "POST", Path: "/api/admin/venue/status", Description: "venue.UpdateStatus", Title: "更新场馆状态", Group: "venue"},
	{Method: "POST", Path: "/api/admin/venue/update", Description: "venue.Update", Title: "更新场馆", Group: "venue"},

	{Method: "POST", Path: "/api/admin/product/create", Description: "product.Create", Title: "创建产品", Group: "product"},
	//{Method: "POST", Path: "/api/admin/product/del", Description: "product.Delete", Title: "删除产品",Group: "product"},
	{Method: "POST", Path: "/api/admin/product/info", Description: "product.InfoByID", Title: "产品详情", Group: "product"},
	{Method: "POST", Path: "/api/admin/product/list", Description: "product.List", Title: "产品列表", Group: "product"},
	{Method: "POST", Path: "/api/admin/product/status", Description: "product.UpdateStatus", Title: "更新产品状态", Group: "product"},
	{Method: "POST", Path: "/api/admin/product/update", Description: "product.Update", Title: "更新产品", Group: "product"},
	{Method: "POST", Path: "/api/admin/property/create", Description: "product.CreateProperty", Title: "创建属性", Group: "product"},
	//{Method: "POST", Path: "/api/admin/property/del", Description: "product.DeleteProperty",Title: "删除属性", Group: "product"},
	{Method: "POST", Path: "/api/admin/property/list", Description: "product.ListProperty", Title: "属性列表", Group: "product"},
	{Method: "POST", Path: "/api/admin/property/update", Description: "product.UpdateProperty", Title: "更新属性", Group: "product"},

	//{Method: "POST", Path: "/api/admin/api", Description: "role.DeleteApi",Title: "删除", Group: "role"},
	{Method: "POST", Path: "/api/admin/api/list", Description: "role.ApiList", Title: "API列表", Group: "role"},
	{Method: "POST", Path: "/api/admin/api/tree", Description: "role.ApiTree", Title: "API树状列表", Group: "role"},

	//{Method: "POST", Path: "/api/admin/api/create", Description: "role.CreateApi", Title: "创建", Group: "role"},
	//{Method: "POST", Path: "/api/admin/api/update", Description: "role.UpdateApi", Title: "更新", Group: "role"},

	//{Method: "POST", Path: "/api/admin/authority/api/create", Description: "role.CreateAuthority", Title: "创建", Group: "role"},
	{Method: "POST", Path: "/api/admin/authority/api/role", Description: "role.ApiAuthority", Title: "角色API详情", Group: "role"},
	{Method: "POST", Path: "/api/admin/authority/api/update", Description: "role.UpdateApiAuthority", Title: "更新角色API", Group: "role"},

	//{Method: "POST", Path: "/api/admin/authority/menu/create", Description: "role.CreateMenuAuthority", Title: "创建", Group: "role"},
	{Method: "POST", Path: "/api/admin/authority/menu/role", Description: "role.MenuAuthority", Title: "角色菜单列表", Group: "role"},
	{Method: "POST", Path: "/api/admin/authority/menu/update", Description: "role.UpdateMenuAuthority", Title: "更新角色菜单", Group: "role"},

	{Method: "GET", Path: "/api/admin/role/del", Description: "role.RoleByID", Title: "角色详情", Group: "role"},
	{Method: "POST", Path: "/api/admin/role", Description: "role.DeleteRole", Title: "删除角色", Group: "role"},
	{Method: "GET", Path: "/api/admin/role/list", Description: "role.RoleList", Title: "角色列表", Group: "role"},
	{Method: "POST", Path: "/api/admin/role/status", Description: "role.UpdateRoleStatus", Title: "更新角色状态", Group: "role"},
	{Method: "POST", Path: "/api/admin/role/create", Description: "role.CreateRole", Title: "创建角色", Group: "role"},
	{Method: "POST", Path: "/api/admin/role/update", Description: "role.UpdateRole", Title: "更新角色", Group: "role"},

	//{Method: "POST", Path: "/api/admin/menu", Description: "menu.DeleteMenu", Title: "删除菜单", Group: "menu"},
	{Method: "GET", Path: "/api/admin/menu/list", Description: "menu.MenuLists", Title: "菜单列表", Group: "menu"},
	//{Method: "POST", Path: "/api/admin/menu/param", Description: "menu.MenuParam", Group: "menu"},
	//{Method: "POST", Path: "/api/admin/menu/param/list", Description: "menu.MenuParamListByMenuID", Title: "列表", Group: "menu"},
	{Method: "POST", Path: "/api/admin/menu/role", Description: "menu.MenuByRole", Title: "角色菜单列表", Group: "menu"},
	{Method: "POST", Path: "/api/admin/menu/tree", Description: "menu.ApiTree", Title: "菜单树状列表", Group: "menu"},
	//{Method: "POST", Path: "/api/admin/menu/param/create", Description: "menu.CreateMenuParam", Title: "创建菜单", Group: "menu"},
	//{Method: "POST", Path: "/api/admin/menu/param/update", Description: "menu.UpdateMenuParam", Title: "更新菜单", Group: "menu"},
	//{Method: "POST", Path: "/api/admin/menu/create", Description: "menu.CreateMenu", Title: "创建菜单", Group: "menu"},
	//{Method: "POST", Path: "/api/admin/menu/update", Description: "menu.UpdateMenu", Title: "更新菜单", Group: "menu"},

	//{Method: "POST", Path: "/api/admin/logs/deleteAll", Description: "logs.DeleteLogs",Title: "删除全部日志列表", Group: "logs"},
	{Method: "POST", Path: "/api/admin/logs/list", Description: "logs.GetLogsList", Title: "日志列表", Group: "logs"},

	//{Method: "POST", Path: "/api/admin/dict", Description: "dictionary.DeleteDictionary", Title: "删除", Group: "dictionary"},
	{Method: "GET", Path: "/api/admin/dict/detail", Description: "dictionary.DictionaryDetail", Title: "字典详情", Group: "dictionary"},
	{Method: "POST", Path: "/api/admin/dict/detail/list", Description: "dictionary.DetailByDictionaryList", Title: "字典详情列表", Group: "dictionary"},
	{Method: "GET", Path: "/api/admin/dict/list", Description: "dictionary.DictionaryList", Title: "字典列表", Group: "dictionary"},
	{Method: "POST", Path: "/api/admin/dict/detail/create", Description: "dictionary.CreateDictionaryDetail", Title: "创建字典详情", Group: "dictionary"},
	{Method: "POST", Path: "/api/admin/dict/detail/update", Description: "dictionary.UpdateDictionaryDetail", Title: "更新字典详情", Group: "dictionary"},
	{Method: "POST", Path: "/api/admin/dict/create", Description: "dictionary.CreateDictionary", Title: "创建字典", Group: "dictionary"},
	{Method: "POST", Path: "/api/admin/dict/update", Description: "dictionary.UpdateDictionary", Title: "更新字典", Group: "dictionary"},

	//{Method: "POST", Path: "/api/admin/user", Description: "user.DeleteUser", Group: "user"},
	{Method: "GET", Path: "/api/admin/user/profile", Description: "user.UserProfile", Title: "员工详情", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/profile-update", Description: "user.UpdateProfile", Title: "更新员工", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/status", Description: "user.UpdateUserStatus", Title: "更新员工状态", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/change-password", Description: "user.ChangePassword", Title: "修改密码", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/create", Description: "user.CreateUser", Title: "创建员工", Group: "user"},
	{Method: "GET", Path: "/api/admin/user/info", Description: "user.UserInfo", Title: "员工详情", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/list", Description: "user.UserList", Title: "员工列表", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/perm", Description: "user.UserPermCode", Title: "员工权限", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/update", Description: "user.UpdateUser", Title: "更新员工", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/set-role", Description: "user.SetUserRole", Title: "更新用户角色", Group: "user"},
	{Method: "POST", Path: "/api/admin/user/set-default-venue", Description: "user.SetDefaultVenue", Title: "更新用户登陆后默认场馆", Group: "user"},
	//{Method: "POST", Path: "/api/admin/token", Description: "token.DeleteToken", Group: "token"},
	//{Method: "POST", Path: "/api/admin/token/list", Description: "token.TokenList", Title: "列表", Group: "token"},
	//{Method: "POST", Path: "/api/admin/token/update", Description: "token.UpdateToken", Title: "更新", Group: "token"},

	{Method: "POST", Path: "/api/admin/schedule/create", Description: "schedule.CreateSchedule", Title: "创建课程", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/info", Description: "schedule.GetScheduleById", Title: "课程详情", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/list", Description: "schedule.ListSchedule", Title: "课程列表", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/status", Description: "schedule.UpdateStatus", Title: "更新课程状态", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/update", Description: "schedule.UpdateSchedule", Title: "更新课程", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/date-list", Description: "schedule.DateListSchedule", Title: "课程日期列表", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/member-subscribe", Description: "schedule.MemberSubscribe", Title: "会员课程列表", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/schedule-member-list", Description: "schedule.GetScheduleMemberList", Title: "课程会员列表", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/schedule-member-status", Description: "schedule.UpdateMemberStatus", Title: "更新会员课程状态", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/search-subscribe-by-member", Description: "schedule.SearchSubscribeByMember", Title: "查询可上课程会员", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/schedule-coach-list", Description: "schedule.GetScheduleCoachList", Title: "课程教练列表", Group: "schedule"},
	{Method: "POST", Path: "/api/admin/schedule/schedule-coach-status", Description: "schedule.UpdateCoachStatus", Title: "更新教练课程状态", Group: "schedule"},
}
