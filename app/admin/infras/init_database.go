/*
 * Copyright 2023 FormulaGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 */

package infras

import (
	"context"
	"saas/pkg/db/ent"
	"saas/pkg/encrypt"

	"sync"

	"github.com/casbin/casbin/v2"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cockroachdb/errors"
)

type InitDatabase struct {
	DB  *ent.Client
	Csb *casbin.Enforcer
	Mu  *sync.Mutex
}

var (
	mutex = new(sync.Mutex)
)

func NewInitDatabase() *InitDatabase {
	return &InitDatabase{
		DB:  DB,
		Csb: CasbinEnforcer(),
		Mu:  mutex,
	}
}

func (I *InitDatabase) InitDatabaseUser() error {
	// add lock to avoid duplicate initialization
	I.Mu.Lock()
	defer I.Mu.Unlock()
	ctx := context.Background()
	// judge if the initialization had been done
	check, err := I.DB.User.Query().Count(ctx)
	if check > 0 {
		return errors.New("Database had been initialized")
	}

	// insert init data
	err = I.insertUserData(ctx)
	if err != nil {
		hlog.Error("insert user data failed", err)
		err = errors.Wrap(err, "insert user data failed")
		return err
	}
	err = I.insertRoleData(ctx)
	if err != nil {
		hlog.Error("insert role data failed", err)
		err = errors.Wrap(err, "insert role data failed")
		return err
	}
	// set init status
	return nil
}

func (I *InitDatabase) InitDatabaseDict() error {
	// add lock to avoid duplicate initialization
	I.Mu.Lock()
	defer I.Mu.Unlock()
	ctx := context.Background()
	// judge if the initialization had been done
	check, err := I.DB.Dictionary.Query().Count(ctx)
	if check > 0 {
		return errors.New("Database had been initialized")
	}

	err = I.insertDictionariesPropertyTypeData(ctx)
	if err != nil {
		hlog.Error("insert dictionaries data failed", err)
		err = errors.Wrap(err, "insert dictionaries data failed")
		return err
	}
	err = I.insertDictionariesOrganizationData(ctx)
	if err != nil {
		hlog.Error("insert dictionaries data failed", err)
		err = errors.Wrap(err, "insert dictionaries data failed")
		return err
	}
	err = I.insertDictionariesJobData(ctx)
	if err != nil {
		hlog.Error("insert dictionaries data failed", err)
		err = errors.Wrap(err, "insert dictionaries data failed")
		return err
	}

	err = I.insertDictionariesNatureData(ctx)
	if err != nil {
		hlog.Error("insert dictionaries data failed", err)
		err = errors.Wrap(err, "insert dictionaries data failed")
		return err
	}

	return nil
}
func (I *InitDatabase) InsertDatabaseMenuData() error {
	// add lock to avoid duplicate initialization
	I.Mu.Lock()
	defer I.Mu.Unlock()
	ctx := context.Background()
	// judge if the initialization had been done
	check, err := I.DB.Menu.Query().Count(ctx)
	if check > 0 {
		return errors.New("Database had been initialized")
	}

	err = I.insertMenuData(ctx)
	if err != nil {
		hlog.Error("insert menu data failed", err)
		err = errors.Wrap(err, "insert menu data failed")
		return err
	}
	err = I.insertRoleMenuAuthorityData(ctx)
	if err != nil {
		hlog.Error("insert role menu authority data failed", err)
		err = errors.Wrap(err, "insert role menu authority data failed")
		return err
	}
	return nil
}
func (I *InitDatabase) InitDatabaseApi() error {
	// add lock to avoid duplicate initialization
	I.Mu.Lock()
	defer I.Mu.Unlock()
	ctx := context.Background()
	// judge if the initialization had been done
	check, err := I.DB.API.Query().Count(ctx)
	if check > 0 {
		return errors.New("Database had been initialized")
	}

	err = I.insertApiData(ctx)
	if err != nil {
		hlog.Error("insert api data failed", err)
		err = errors.Wrap(err, "insert api data failed")
		return err
	}

	err = I.insertCasbinPoliciesData(ctx)
	if err != nil {
		hlog.Error("insert casbin policies data failed", err)
		err = errors.Wrap(err, "insert casbin policies data failed")
		return err
	}

	return nil
}

// insert init user data
func (I *InitDatabase) insertUserData(ctx context.Context) error {
	var users []*ent.UserCreate
	password, _ := encrypt.Crypt("123456")
	users = append(users, I.DB.User.Create().
		SetUsername("admin").
		SetNickname("admin").
		SetPassword(password).
		SetEmail("admin@gmail.com").
		SetMobile("12345678901").
		SetRoleID(1).
		SetWecom("admin"),
	)

	err := I.DB.User.CreateBulk(users...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// insert init apis data
func (I *InitDatabase) insertRoleData(ctx context.Context) error {
	var roles []*ent.RoleCreate
	roles = make([]*ent.RoleCreate, 2)
	roles[0] = I.DB.Role.Create().
		SetName("role.admin").
		SetValue("admin").
		SetRemark("超级管理员").
		SetOrderNo(1)

	roles[1] = I.DB.Role.Create().
		SetName("role.stuff").
		SetValue("stuff").
		SetRemark("普通员工").
		SetOrderNo(2)

	err := I.DB.Role.CreateBulk(roles...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// insert init API data
func (I *InitDatabase) insertApiData(ctx context.Context) error {

	type api struct {
		Method      string
		Path        string
		Description string
		Group       string
		Title       string
	}
	apiArr := []api{

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
	//
	var apis []*ent.APICreate
	apis = make([]*ent.APICreate, len(apiArr))

	for i, v := range apiArr {
		apis[i] = I.DB.API.Create().
			SetPath(v.Path).
			SetDescription(v.Description).
			SetAPIGroup(v.Group).
			SetTitle(v.Title).
			SetMethod(v.Method)
	}
	err := I.DB.API.CreateBulk(apis...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// init menu data
func (I *InitDatabase) insertMenuData(ctx context.Context) error {
	var menus []*ent.MenuCreate
	menus = make([]*ent.MenuCreate, 25)
	menus[0] = I.DB.Menu.Create().
		//SetMenuLevel(0).
		//SetMenuType(0).
		//SetParentID(1).
		SetParentID(1).
		SetPath("").
		SetName("root")
	//SetComponent("").
	//SetOrderNo(0).
	//SetTitle("").
	//SetIcon("").
	//SetHideMenu(false)

	menus[1] = I.DB.Menu.Create().
		SetPath("dashboard").
		SetParentID(1).
		SetName("仪表盘").SetOrderNo(0)
	menus[2] = I.DB.Menu.Create().
		SetPath("sys").
		SetParentID(1).
		SetName("系统管理").SetOrderNo(7)
	menus[3] = I.DB.Menu.Create().
		SetPath("devops").
		SetParentID(1).
		SetName("运维管理").SetOrderNo(8)
	menus[4] = I.DB.Menu.Create().
		SetPath("orders").
		SetParentID(1).
		SetName("订单管理").SetOrderNo(3)
	menus[5] = I.DB.Menu.Create().
		SetPath("products").
		SetParentID(1).
		SetName("产品管理").SetOrderNo(2)
	menus[6] = I.DB.Menu.Create().
		SetPath("members").
		SetParentID(1).
		SetName("会员管理").SetOrderNo(1)
	menus[7] = I.DB.Menu.Create().
		SetPath("schedules").
		SetParentID(1).
		SetName("预约排期").SetOrderNo(4)
	menus[8] = I.DB.Menu.Create().
		SetPath("statisticals").
		SetParentID(1).
		SetName("数据分析").SetOrderNo(6)

	menus[9] = I.DB.Menu.Create().
		SetPath("dashboard/workplace").
		SetParentID(2).
		SetName("控制台").SetOrderNo(0)

	menus[10] = I.DB.Menu.Create().
		SetPath("sys/role").
		SetParentID(3).
		SetName("角色管理").SetOrderNo(1)
	menus[11] = I.DB.Menu.Create().
		SetPath("sys/user").
		SetParentID(3).
		SetName("用户管理").SetOrderNo(2)
	menus[12] = I.DB.Menu.Create().
		SetPath("sys/dictionary").
		SetParentID(3).
		SetName("字典管理").SetOrderNo(3)
	menus[13] = I.DB.Menu.Create().
		SetPath("sys/venue").
		SetParentID(3).
		SetName("场馆管理").SetOrderNo(4)
	menus[14] = I.DB.Menu.Create().
		SetPath("sys/place").
		SetParentID(3).
		SetName("场地管理").SetOrderNo(5)
	menus[15] = I.DB.Menu.Create().
		SetPath("sys/contract").
		SetParentID(3).
		SetName("合同管理").SetOrderNo(6)

	menus[16] = I.DB.Menu.Create().
		SetPath("devops/logs").
		SetParentID(4).
		SetName("日志管理").SetOrderNo(0)

	menus[17] = I.DB.Menu.Create().
		SetPath("orders/add").
		SetParentID(5).
		SetName("创建订单").SetOrderNo(0)
	menus[18] = I.DB.Menu.Create().
		SetPath("orders/list").
		SetParentID(5).
		SetName("订单列表").SetOrderNo(1)

	menus[19] = I.DB.Menu.Create().
		SetPath("products/product").
		SetParentID(6).
		SetName("产品列表").SetOrderNo(0)
	menus[20] = I.DB.Menu.Create().
		SetPath("products/property").
		SetParentID(6).
		SetName("属性列表").SetOrderNo(1)

	menus[21] = I.DB.Menu.Create().
		SetPath("members/list").
		SetParentID(7).
		SetName("会员列表").SetOrderNo(0)
	menus[22] = I.DB.Menu.Create().
		SetPath("schedules/schedule").
		SetParentID(8).
		SetName("团课课程").SetOrderNo(0)
	menus[23] = I.DB.Menu.Create().
		SetPath("schedules/course").
		SetParentID(8).
		SetName("私教课程").SetOrderNo(0)
	menus[24] = I.DB.Menu.Create().
		SetPath("statisticals/all").
		SetParentID(9).
		SetName("数据汇总").SetOrderNo(0)

	//
	//menus[2] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(0).
	//	SetParentID(1).
	//	SetPath("sys/").
	//	SetName("SystemManagement").
	//	SetComponent("LAYOUT").
	//	SetOrderNo(19).
	//	SetTitle("系统管理").
	//	SetIcon("ant-design:tool-outlined").
	//	SetHideMenu(false)
	//
	//menus[3] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(0).
	//	SetParentID(1).
	//	SetPath("sys/").
	//	SetName("Maintenance").
	//	SetComponent("LAYOUT").
	//	SetOrderNo(21).
	//	SetTitle("运维管理").
	//	SetIcon("ant-design:bar-chart-outlined").
	//	SetHideMenu(false)
	//
	//menus[4] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(4).
	//	SetPath("sys/menu").
	//	SetName("MenuManagement").
	//	SetComponent("sys/menu/index").
	//	SetOrderNo(1).
	//	SetTitle("菜单管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[5] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(3).
	//	SetPath("sys/role").
	//	SetName("RoleManagement").
	//	SetComponent("sys/role/index").
	//	SetOrderNo(1).
	//	SetTitle("角色管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[6] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(4).
	//	SetPath("sys/api").
	//	SetName("APIManagement").
	//	SetComponent("sys/api/index").
	//	SetOrderNo(3).
	//	SetTitle("API管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[7] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(3).
	//	SetPath("sys/user").
	//	SetName("UserManagement").
	//	SetComponent("sys/user/index").
	//	SetOrderNo(2).
	//	SetTitle("用户管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[8] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(3).
	//	SetPath("sys/dictionary").
	//	SetName("DictionaryManagement").
	//	SetComponent("sys/dictionary/index").
	//	SetOrderNo(3).
	//	SetTitle("字典管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[9] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(2).
	//	SetParentID(3).
	//	SetPath("sys/dictionary/detail").
	//	SetName("DictionaryDetail").
	//	SetComponent("sys/dictionary/detail").
	//	SetOrderNo(4).
	//	SetTitle("字典明细").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[10] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(0).
	//	SetParentID(1).
	//	SetPath("sys/order").
	//	SetName("Order").
	//	SetComponent("LAYOUT").
	//	SetOrderNo(4).
	//	SetTitle("订单管理").
	//	SetIcon("ant-design:property-safety-outlined").
	//	SetHideMenu(false)
	//
	//menus[11] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(4).
	//	SetPath("sys/logs").
	//	SetName("LogsManagement").
	//	SetComponent("sys/logs/index").
	//	SetOrderNo(9).
	//	SetTitle("日志管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[12] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(0).
	//	SetParentID(4).
	//	SetPath("sys/").
	//	SetName("OtherPages").
	//	SetComponent("LAYOUT").
	//	SetOrderNo(2).
	//	SetTitle("其他页面").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[13] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(13).
	//	SetPath("sys/oauth/callback").
	//	SetName("OauthCallbackPage").
	//	SetComponent("sys/oauth/callback").
	//	SetOrderNo(1).
	//	SetTitle("回调页面").
	//	SetIcon("").
	//	SetHideMenu(true)
	//
	//menus[14] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(1).
	//	SetParentID(13).
	//	SetPath("sys/profile").
	//	SetName("Profile").
	//	SetComponent("sys/profile/index").
	//	SetOrderNo(2).
	//	SetTitle("用户信息").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[15] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(0).
	//	SetParentID(1).
	//	SetPath("sys/product").
	//	SetName("Product").
	//	SetComponent("LAYOUT").
	//	SetOrderNo(3).
	//	SetTitle("产品管理").
	//	SetIcon("ant-design:inbox-outlined").
	//	SetHideMenu(false)
	//
	//menus[16] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(16).
	//	SetPath("sys/product").
	//	SetName("ProductList").
	//	SetComponent("sys/product/index").
	//	SetOrderNo(1).
	//	SetTitle("产品列表").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[17] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(16).
	//	SetPath("sys/property").
	//	SetName("PropertyList").
	//	SetComponent("sys/property/index").
	//	SetOrderNo(2).
	//	SetTitle("属性列表").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[18] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(11).
	//	SetPath("sys/order").
	//	SetName("OrderList").
	//	SetComponent("sys/order/index").
	//	SetOrderNo(1).
	//	SetTitle("订单列表").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[19] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(0).
	//	SetParentID(1).
	//	SetPath("/").
	//	SetName("Member").
	//	SetComponent("LAYOUT").
	//	SetOrderNo(3).
	//	SetTitle("会员管理").
	//	SetIcon("ant-design:team-outlined").
	//	SetHideMenu(false)
	//
	//menus[20] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(20).
	//	SetPath("sys/member").
	//	SetName("MemberList").
	//	SetComponent("sys/member/index").
	//	SetOrderNo(1).
	//	SetTitle("会员列表").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[21] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(0).
	//	SetParentID(1).
	//	SetPath("/").
	//	SetName("Schedule").
	//	SetComponent("LAYOUT").
	//	SetOrderNo(3).
	//	SetTitle("预约排期").
	//	SetIcon("ant-design:schedule-outlined").
	//	SetHideMenu(false)
	//
	//menus[22] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(22).
	//	SetPath("sys/schedule").
	//	SetName("ScheduleReservation").
	//	SetComponent("sys/schedule/index").
	//	SetOrderNo(8).
	//	SetTitle("课程预约").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[23] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(0).
	//	SetParentID(1).
	//	SetPath("/").
	//	SetName("Statistical").
	//	SetComponent("LAYOUT").
	//	SetOrderNo(18).
	//	SetTitle("数据分析").
	//	SetIcon("ant-design:bar-chart-outlined").
	//	SetHideMenu(false)
	//
	//menus[24] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(24).
	//	SetPath("sys/statistical").
	//	SetName("StatisticalAll").
	//	SetComponent("sys/statistical/index").
	//	SetOrderNo(8).
	//	SetTitle("数据汇总").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[25] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(4).
	//	SetPath("sys/token").
	//	SetName("TokenManagement").
	//	SetComponent("sys/token/index").
	//	SetOrderNo(8).
	//	SetTitle("Token管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[26] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(3).
	//	SetPath("sys/venue").
	//	SetName("Venue").
	//	SetComponent("sys/venue/index").
	//	SetOrderNo(1).
	//	SetTitle("场馆管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[27] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(2).
	//	SetParentID(5).
	//	SetPath("sys/venue/place").
	//	SetName("VenuePlace").
	//	SetComponent("sys/venue/place").
	//	SetOrderNo(1).
	//	SetTitle("场地管理").
	//	SetIcon("").
	//	SetHideMenu(true)
	//
	//menus[28] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(1).
	//	SetParentID(3).
	//	SetPath("sys/contract").
	//	SetName("Contract").
	//	SetComponent("sys/contract/index").
	//	SetOrderNo(1).
	//	SetTitle("合同管理").
	//	SetIcon("").
	//	SetHideMenu(false)
	//
	//menus[29] = I.DB.Menu.Create().
	//	SetMenuLevel(2).
	//	SetMenuType(2).
	//	SetParentID(11).
	//	SetPath("sys/order/add").
	//	SetName("AddOrder").
	//	SetComponent("sys/order/add").
	//	SetOrderNo(1).
	//	SetTitle("添加订单").
	//	SetIcon("").
	//	SetHideMenu(true)
	//
	//menus[30] = I.DB.Menu.Create().
	//	SetMenuLevel(1).
	//	SetMenuType(1).
	//	SetParentID(2).
	//	SetPath("dashboard/workplace").
	//	SetName("Dashboard").
	//	SetComponent("dashboard/workplace/index").
	//	SetOrderNo(0).
	//	SetTitle("控制台").
	//	SetIcon("ant-design:home-outlined").
	//	SetHideMenu(false)

	err := I.DB.Menu.CreateBulk(menus...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// insert admin menu authority

func (I *InitDatabase) insertRoleMenuAuthorityData(ctx context.Context) error {
	count, err := I.DB.Menu.Query().Count(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	var menuIDs []int64
	menuIDs = make([]int64, count)

	for i := range menuIDs {
		menuIDs[i] = int64(i + 1)
	}
	err = I.DB.Role.Update().AddMenuIDs(menuIDs...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// init casbin policies

func (I *InitDatabase) insertCasbinPoliciesData(ctx context.Context) error {
	apis, err := I.DB.API.Query().All(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	var policies [][]string
	for _, v := range apis {
		policies = append(policies, []string{"1", v.Path, v.Method})
	}

	addResult, err := I.Csb.AddPolicies(policies)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	if !addResult {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// init propertyType data
func (I *InitDatabase) insertDictionariesPropertyTypeData(ctx context.Context) error {
	data, err := I.DB.Dictionary.Create().
		SetTitle("propertys").
		SetName("属性类型").
		SetStatus(1).
		SetDescription("属性类型").
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("卡").
		SetValue("卡").
		SetKey("card").
		SetDictionary(data).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("私教课").
		SetValue("私教课").
		SetKey("course").
		SetDictionary(data).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("团课").
		SetValue("团课").
		SetKey("class").
		SetDictionary(data).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// init job data
func (I *InitDatabase) insertDictionariesJobData(ctx context.Context) error {
	data, err := I.DB.Dictionary.Create().
		SetTitle("jobs").
		SetName("职业").
		SetStatus(1).
		SetDescription("职业").
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("销售").
		SetValue("销售").
		SetKey("sale").
		SetDictionary(data).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("教练").
		SetValue("教练").
		SetKey("coach").
		SetDictionary(data).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	return nil
}

// init organization data
func (I *InitDatabase) insertDictionariesOrganizationData(ctx context.Context) error {
	data, err := I.DB.Dictionary.Create().
		SetTitle("organizations").
		SetName("部门").
		SetStatus(1).
		SetDescription("部门").
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("销售").
		SetValue("销售").
		SetKey("sales").
		SetDictionary(data).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("教练").
		SetValue("教练").
		SetKey("coach").
		SetDictionary(data).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	return nil
}

// init nature data
func (I *InitDatabase) insertDictionariesNatureData(ctx context.Context) error {
	data, err := I.DB.Dictionary.Create().
		SetTitle("nature").
		SetName("业务类型").
		SetStatus(1).
		SetDescription("业务类型").
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("新单").
		SetValue("新单").
		SetKey("1").
		SetDictionary(data).
		Save(ctx)

	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("续约").
		SetValue("续约").
		SetKey("2").
		SetDictionary(data).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}

	return nil
}

//
