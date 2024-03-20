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

	"github.com/casbin/casbin/v2"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cockroachdb/errors"
)

type InitDatabase struct {
	DB  *ent.Client
	Csb *casbin.Enforcer
	Mu  *sync.Mutex
}

var (
	mutex              = new(sync.Mutex)
	InitDatabaseStatus bool
)

func NewInitDatabase() *InitDatabase {
	return &InitDatabase{
		DB:  DB,
		Csb: CasbinEnforcer(),
		Mu:  mutex,
	}
}

func (I *InitDatabase) InitDatabase() error {
	// add lock to avoid duplicate initialization
	I.Mu.Lock()
	defer I.Mu.Unlock()
	ctx := context.Background()
	// judge if the initialization had been done
	check, err := I.DB.API.Query().Count(ctx)
	if InitDatabaseStatus || check > 0 {
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

	err = I.insertMenuData(ctx)
	if err != nil {
		hlog.Error("insert menu data failed", err)
		err = errors.Wrap(err, "insert menu data failed")
		return err
	}

	err = I.insertApiData(ctx)
	if err != nil {
		hlog.Error("insert api data failed", err)
		err = errors.Wrap(err, "insert api data failed")
		return err
	}
	err = I.insertRoleMenuAuthorityData(ctx)
	if err != nil {
		hlog.Error("insert role menu authority data failed", err)
		err = errors.Wrap(err, "insert role menu authority data failed")
		return err
	}
	err = I.insertCasbinPoliciesData(ctx)
	if err != nil {
		hlog.Error("insert casbin policies data failed", err)
		err = errors.Wrap(err, "insert casbin policies data failed")
		return err
	}
	err = I.insertDictionariesData(ctx)
	if err != nil {
		hlog.Error("insert dictionaries data failed", err)
		err = errors.Wrap(err, "insert dictionaries data failed")
		return err
	}

	// set init status
	InitDatabaseStatus = true
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
	roles = make([]*ent.RoleCreate, 3)
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

	roles[2] = I.DB.Role.Create().
		SetName("role.member").
		SetValue("member").
		SetRemark("注册会员").
		SetOrderNo(3)

	err := I.DB.Role.CreateBulk(roles...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// insert init API data
func (I *InitDatabase) insertApiData(ctx context.Context) error {
	var apis []*ent.APICreate
	apis = make([]*ent.APICreate, 59)
	// USER
	apis[0] = I.DB.API.Create().
		SetPath("/api/admin/user/login").
		SetDescription("apiDesc.userLogin").
		SetAPIGroup("user").
		SetMethod("POST")

	apis[1] = I.DB.API.Create().
		SetPath("/api/admin/user/register").
		SetDescription("apiDesc.userRegister").
		SetAPIGroup("user").
		SetMethod("POST")

	apis[2] = I.DB.API.Create().
		SetPath("/api/admin/user/create").
		SetDescription("apiDesc.createUser").
		SetAPIGroup("user").
		SetMethod("POST")

	apis[3] = I.DB.API.Create().
		SetPath("/api/admin/user/update").
		SetDescription("apiDesc.updateUser").
		SetAPIGroup("user").
		SetMethod("POST")

	apis[4] = I.DB.API.Create().
		SetPath("/api/admin/user/change-password").
		SetDescription("apiDesc.userChangePassword").
		SetAPIGroup("user").
		SetMethod("POST")

	apis[5] = I.DB.API.Create().
		SetPath("/api/admin/user/info").
		SetDescription("apiDesc.OauthUserInfo").
		SetAPIGroup("user").
		SetMethod("GET")

	apis[6] = I.DB.API.Create().
		SetPath("/api/admin/user/list").
		SetDescription("apiDesc.userList").
		SetAPIGroup("user").
		SetMethod("GET")

	apis[7] = I.DB.API.Create().
		SetPath("/api/admin/user").
		SetDescription("apiDesc.deleteUser").
		SetAPIGroup("user").
		SetMethod("DELETE")

	apis[8] = I.DB.API.Create().
		SetPath("/api/admin/user/perm").
		SetDescription("apiDesc.userPermissions").
		SetAPIGroup("user").
		SetMethod("GET")

	apis[9] = I.DB.API.Create().
		SetPath("/api/admin/user/profile").
		SetDescription("apiDesc.userProfile").
		SetAPIGroup("user").
		SetMethod("GET")

	apis[10] = I.DB.API.Create().
		SetPath("/api/admin/user/profile").
		SetDescription("apiDesc.updateProfile").
		SetAPIGroup("user").
		SetMethod("POST")

	apis[11] = I.DB.API.Create().
		SetPath("/api/admin/user/logout").
		SetDescription("apiDesc.logout").
		SetAPIGroup("user").
		SetMethod("GET")

	apis[12] = I.DB.API.Create().
		SetPath("/api/admin/user/status").
		SetDescription("apiDesc.updateUserStatus").
		SetAPIGroup("user").
		SetMethod("POST")

	// ROLE
	apis[13] = I.DB.API.Create().
		SetPath("/api/admin/role/create").
		SetDescription("apiDesc.createRole").
		SetAPIGroup("role").
		SetMethod("POST")

	apis[14] = I.DB.API.Create().
		SetPath("/api/admin/role/update").
		SetDescription("apiDesc.updateRole").
		SetAPIGroup("role").
		SetMethod("POST")

	apis[15] = I.DB.API.Create().
		SetPath("/api/admin/role").
		SetDescription("apiDesc.deleteRole").
		SetAPIGroup("role").
		SetMethod("DELETE")

	apis[16] = I.DB.API.Create().
		SetPath("/api/admin/role/list").
		SetDescription("apiDesc.roleList").
		SetAPIGroup("role").
		SetMethod("GET")

	apis[17] = I.DB.API.Create().
		SetPath("/api/admin/role/status").
		SetDescription("apiDesc.setRoleStatus").
		SetAPIGroup("role").
		SetMethod("POST")

	// MENU
	apis[18] = I.DB.API.Create().
		SetPath("/api/admin/menu/create").
		SetDescription("apiDesc.createMenu").
		SetAPIGroup("menu").
		SetMethod("POST")

	apis[19] = I.DB.API.Create().
		SetPath("/api/admin/menu/update").
		SetDescription("apiDesc.updateMenu").
		SetAPIGroup("menu").
		SetMethod("POST")

	apis[20] = I.DB.API.Create().
		SetPath("/api/admin/menu").
		SetDescription("apiDesc.deleteMenu").
		SetAPIGroup("menu").
		SetMethod("DELETE")

	apis[21] = I.DB.API.Create().
		SetPath("/api/admin/menu/list").
		SetDescription("apiDesc.menuList").
		SetAPIGroup("menu").
		SetMethod("GET")

	apis[22] = I.DB.API.Create().
		SetPath("/api/admin/menu/role").
		SetDescription("apiDesc.roleMenu").
		SetAPIGroup("menu").
		SetMethod("GET")

	apis[23] = I.DB.API.Create().
		SetPath("/api/admin/menu/param/create").
		SetDescription("apiDesc.createMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST")

	apis[24] = I.DB.API.Create().
		SetPath("/api/admin/menu/param/update").
		SetDescription("apiDesc.updateMenuParam").
		SetAPIGroup("menu").
		SetMethod("POST")

	apis[25] = I.DB.API.Create().
		SetPath("/api/admin/menu/param/list").
		SetDescription("apiDesc.menuParamListByMenuID").
		SetAPIGroup("menu").
		SetMethod("GET")

	apis[26] = I.DB.API.Create().
		SetPath("/api/admin/menu/param").
		SetDescription("apiDesc.deleteMenuParam").
		SetAPIGroup("menu").
		SetMethod("DELETE")

	// CAPTCHA
	apis[27] = I.DB.API.Create().
		SetPath("/api/admin/captcha").
		SetDescription("apiDesc.captcha").
		SetAPIGroup("captcha").
		SetMethod("GET")

	// AUTHORIZATION
	apis[28] = I.DB.API.Create().
		SetPath("/api/admin/authority/api/create").
		SetDescription("apiDesc.createApiAuthority").
		SetAPIGroup("authority").
		SetMethod("POST")

	apis[29] = I.DB.API.Create().
		SetPath("/api/admin/authority/api/update").
		SetDescription("apiDesc.updateApiAuthority").
		SetAPIGroup("authority").
		SetMethod("POST")

	apis[30] = I.DB.API.Create().
		SetPath("/api/admin/authority/api/role").
		SetDescription("apiDesc.APIAuthorityOfRole").
		SetAPIGroup("authority").
		SetMethod("POST")

	apis[31] = I.DB.API.Create().
		SetPath("/api/admin/authority/menu/create").
		SetDescription("apiDesc.createMenuAuthority").
		SetAPIGroup("authority").
		SetMethod("POST")

	apis[32] = I.DB.API.Create().
		SetPath("/api/admin/authority/menu/update").
		SetDescription("apiDesc.updateMenuAuthority").
		SetAPIGroup("authority").
		SetMethod("POST")

	apis[33] = I.DB.API.Create().
		SetPath("/api/admin/authority/menu/role").
		SetDescription("apiDesc.menuAuthorityOfRole").
		SetAPIGroup("authority").
		SetMethod("POST")

	// API
	apis[34] = I.DB.API.Create().
		SetPath("/api/admin/api/create").
		SetDescription("apiDesc.createApi").
		SetAPIGroup("api").
		SetMethod("POST")

	apis[35] = I.DB.API.Create().
		SetPath("/api/admin/api/update").
		SetDescription("apiDesc.updateApi").
		SetAPIGroup("api").
		SetMethod("POST")

	apis[36] = I.DB.API.Create().
		SetPath("/api/admin/api").
		SetDescription("apiDesc.deleteAPI").
		SetAPIGroup("api").
		SetMethod("DELETE")

	apis[37] = I.DB.API.Create().
		SetPath("/api/admin/api/list").
		SetDescription("apiDesc.APIList").
		SetAPIGroup("api").
		SetMethod("GET")

	// DICTIONARY
	apis[38] = I.DB.API.Create().
		SetPath("/api/admin/dict/create").
		SetDescription("apiDesc.createDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST")

	apis[39] = I.DB.API.Create().
		SetPath("/api/admin/dict/update").
		SetDescription("apiDesc.updateDictionary").
		SetAPIGroup("dictionary").
		SetMethod("POST")

	apis[40] = I.DB.API.Create().
		SetPath("/api/admin/dict").
		SetDescription("apiDesc.deleteDictionary").
		SetAPIGroup("dictionary").
		SetMethod("DELETE")

	apis[41] = I.DB.API.Create().
		SetPath("/api/admin/dict/detail").
		SetDescription("apiDesc.deleteDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("DELETE")

	apis[42] = I.DB.API.Create().
		SetPath("/api/admin/dict/detail/create").
		SetDescription("apiDesc.createDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST")

	apis[43] = I.DB.API.Create().
		SetPath("/api/admin/dict/detail/update").
		SetDescription("apiDesc.updateDictionaryDetail").
		SetAPIGroup("dictionary").
		SetMethod("POST")

	apis[44] = I.DB.API.Create().
		SetPath("/api/admin/dict/detail/list").
		SetDescription("apiDesc.getDictionaryListDetail").
		SetAPIGroup("dictionary").
		SetMethod("GET")

	apis[45] = I.DB.API.Create().
		SetPath("/api/admin/dict/list").
		SetDescription("apiDesc.getDictionaryList").
		SetAPIGroup("dictionary").
		SetMethod("GET")

	// OAUTH
	apis[46] = I.DB.API.Create().
		SetPath("/api/admin/oauth/provider/create").
		SetDescription("apiDesc.createProvider").
		SetAPIGroup("oauth").
		SetMethod("POST")

	apis[47] = I.DB.API.Create().
		SetPath("/api/admin/oauth/provider/update").
		SetDescription("apiDesc.updateProvider").
		SetAPIGroup("oauth").
		SetMethod("POST")

	apis[48] = I.DB.API.Create().
		SetPath("/api/admin/oauth/provider").
		SetDescription("apiDesc.deleteProvider").
		SetAPIGroup("oauth").
		SetMethod("DELETE")

	apis[49] = I.DB.API.Create().
		SetPath("/api/admin/oauth/provider/list").
		SetDescription("apiDesc.geProviderList").
		SetAPIGroup("oauth").
		SetMethod("GET")

	apis[50] = I.DB.API.Create().
		SetPath("/api/admin/oauth/login").
		SetDescription("apiDesc.oauthLogin").
		SetAPIGroup("oauth").
		SetMethod("POST")

	// TOKEN
	apis[51] = I.DB.API.Create().
		SetPath("/api/admin/token/create").
		SetDescription("apiDesc.createToken").
		SetAPIGroup("token").
		SetMethod("POST")

	apis[52] = I.DB.API.Create().
		SetPath("/api/admin/token/update").
		SetDescription("apiDesc.updateToken").
		SetAPIGroup("token").
		SetMethod("POST")

	apis[53] = I.DB.API.Create().
		SetPath("/api/admin/token").
		SetDescription("apiDesc.deleteToken").
		SetAPIGroup("token").
		SetMethod("DELETE")

	apis[54] = I.DB.API.Create().
		SetPath("/api/admin/token/list").
		SetDescription("apiDesc.getTokenList").
		SetAPIGroup("token").
		SetMethod("GET")

	apis[55] = I.DB.API.Create().
		SetPath("/api/admin/token/status").
		SetDescription("apiDesc.setTokenStatus").
		SetAPIGroup("token").
		SetMethod("POST")

	apis[56] = I.DB.API.Create().
		SetPath("/api/admin/token/logout").
		SetDescription("user.forceLoggingOut").
		SetAPIGroup("token").
		SetMethod("POST")

	apis[57] = I.DB.API.Create().
		SetPath("/api/admin/logs/list").
		SetDescription("apiDesc.getLogsList").
		SetAPIGroup("logs").
		SetMethod("GET")

	apis[58] = I.DB.API.Create().
		SetPath("/api/admin/logs/deleteAll").
		SetDescription("apiDesc.deleteLogs").
		SetAPIGroup("logs").
		SetMethod("DELETE")

	err := I.DB.API.CreateBulk(apis...).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}

// init menu data
func (I *InitDatabase) insertMenuData(ctx context.Context) error {
	var menus []*ent.MenuCreate
	menus = make([]*ent.MenuCreate, 29)
	menus[0] = I.DB.Menu.Create().
		SetMenuLevel(0).
		SetMenuType(0).
		SetParentID(1).
		SetPath("/").
		SetName("root").
		SetComponent("").
		SetOrderNo(0).
		SetTitle("").
		SetIcon("").
		SetHideMenu(false)

	menus[1] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(1).
		SetPath("/dashboard/workbench/index").
		SetName("Dashboard").
		SetComponent("/dashboard/workbench/index").
		SetOrderNo(0).
		SetTitle("控制台").
		SetIcon("ant-design:home-outlined").
		SetHideMenu(false)

	menus[2] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("/").
		SetName("SystemManagement").
		SetComponent("LAYOUT").
		SetOrderNo(19).
		SetTitle("系统管理").
		SetIcon("ant-design:tool-outlined").
		SetHideMenu(false)

	menus[3] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("/").
		SetName("Maintenance").
		SetComponent("LAYOUT").
		SetOrderNo(21).
		SetTitle("运维管理").
		SetIcon("ant-design:bar-chart-outlined").
		SetHideMenu(false)

	menus[4] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(4).
		SetPath("/sys/menu/index").
		SetName("MenuManagement").
		SetComponent("/sys/menu/index").
		SetOrderNo(1).
		SetTitle("菜单管理").
		SetIcon("").
		SetHideMenu(false)

	menus[5] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/role/index").
		SetName("RoleManagement").
		SetComponent("/sys/role/index").
		SetOrderNo(1).
		SetTitle("角色管理").
		SetIcon("").
		SetHideMenu(false)

	menus[6] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(4).
		SetPath("/sys/api/index").
		SetName("APIManagement").
		SetComponent("/sys/api/index").
		SetOrderNo(3).
		SetTitle("API管理").
		SetIcon("").
		SetHideMenu(false)

	menus[7] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/user/index").
		SetName("UserManagement").
		SetComponent("/sys/user/index").
		SetOrderNo(2).
		SetTitle("用户管理").
		SetIcon("").
		SetHideMenu(false)

	menus[8] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/dictionary/index").
		SetName("DictionaryManagement").
		SetComponent("/sys/dictionary/index").
		SetOrderNo(3).
		SetTitle("字典管理").
		SetIcon("").
		SetHideMenu(false)

	menus[9] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(2).
		SetParentID(3).
		SetPath("/sys/dictionary/detail").
		SetName("DictionaryDetail").
		SetComponent("/sys/dictionary/detail").
		SetOrderNo(4).
		SetTitle("字典明细").
		SetIcon("").
		SetHideMenu(false)

	menus[10] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("/").
		SetName("Order").
		SetComponent("LAYOUT").
		SetOrderNo(4).
		SetTitle("订单管理").
		SetIcon("ant-design:property-safety-outlined").
		SetHideMenu(false)

	menus[11] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(4).
		SetPath("/sys/logs/index").
		SetName("LogsManagement").
		SetComponent("/sys/logs/index").
		SetOrderNo(9).
		SetTitle("日志管理").
		SetIcon("").
		SetHideMenu(false)

	menus[12] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(4).
		SetPath("/").
		SetName("OtherPages").
		SetComponent("LAYOUT").
		SetOrderNo(2).
		SetTitle("其他页面").
		SetIcon("").
		SetHideMenu(false)

	menus[13] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(13).
		SetPath("/sys/oauth/callback").
		SetName("OauthCallbackPage").
		SetComponent("/sys/oauth/callback").
		SetOrderNo(1).
		SetTitle("回调页面").
		SetIcon("").
		SetHideMenu(true)

	menus[14] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(1).
		SetParentID(13).
		SetPath("/sys/profile/index").
		SetName("Profile").
		SetComponent("/sys/profile/index").
		SetOrderNo(2).
		SetTitle("用户信息").
		SetIcon("").
		SetHideMenu(false)

	menus[15] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("/").
		SetName("Product").
		SetComponent("LAYOUT").
		SetOrderNo(3).
		SetTitle("产品管理").
		SetIcon("ant-design:inbox-outlined").
		SetHideMenu(false)

	menus[16] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(16).
		SetPath("/sys/product/index").
		SetName("ProductList").
		SetComponent("/sys/product/index").
		SetOrderNo(1).
		SetTitle("产品列表").
		SetIcon("").
		SetHideMenu(false)

	menus[17] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(16).
		SetPath("/sys/property/index").
		SetName("PropertyList").
		SetComponent("/sys/property/index").
		SetOrderNo(2).
		SetTitle("属性列表").
		SetIcon("").
		SetHideMenu(false)

	menus[18] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(11).
		SetPath("/sys/order/index").
		SetName("OrderList").
		SetComponent("/sys/order/index").
		SetOrderNo(1).
		SetTitle("订单列表").
		SetIcon("").
		SetHideMenu(false)

	menus[19] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("/").
		SetName("Member").
		SetComponent("LAYOUT").
		SetOrderNo(3).
		SetTitle("会员管理").
		SetIcon("ant-design:team-outlined").
		SetHideMenu(false)

	menus[20] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(20).
		SetPath("/sys/member/index").
		SetName("MemberList").
		SetComponent("/sys/member/index").
		SetOrderNo(1).
		SetTitle("会员列表").
		SetIcon("").
		SetHideMenu(false)

	menus[21] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("/").
		SetName("Schedule").
		SetComponent("LAYOUT").
		SetOrderNo(3).
		SetTitle("预约排期").
		SetIcon("ant-design:schedule-outlined").
		SetHideMenu(false)

	menus[22] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(22).
		SetPath("/sys/schedule/index").
		SetName("ScheduleReservation").
		SetComponent("/sys/schedule/index").
		SetOrderNo(8).
		SetTitle("课程预约").
		SetIcon("").
		SetHideMenu(false)

	menus[23] = I.DB.Menu.Create().
		SetMenuLevel(1).
		SetMenuType(0).
		SetParentID(1).
		SetPath("/").
		SetName("Statistical").
		SetComponent("LAYOUT").
		SetOrderNo(18).
		SetTitle("数据分析").
		SetIcon("ant-design:bar-chart-outlined").
		SetHideMenu(false)

	menus[24] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(24).
		SetPath("/sys/statistical/index").
		SetName("StatisticalAll").
		SetComponent("/sys/statistical/index").
		SetOrderNo(8).
		SetTitle("数据汇总").
		SetIcon("").
		SetHideMenu(false)

	menus[25] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(4).
		SetPath("/sys/token/index").
		SetName("TokenManagement").
		SetComponent("/sys/token/index").
		SetOrderNo(8).
		SetTitle("Token管理").
		SetIcon("").
		SetHideMenu(false)

	menus[26] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/venue/index").
		SetName("Venue").
		SetComponent("/sys/venue/index").
		SetOrderNo(1).
		SetTitle("场馆管理").
		SetIcon("").
		SetHideMenu(false)

	menus[27] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(2).
		SetParentID(5).
		SetPath("/sys/venue/place").
		SetName("VenuePlace").
		SetComponent("/sys/venue/place").
		SetOrderNo(1).
		SetTitle("场地管理").
		SetIcon("").
		SetHideMenu(true)

	menus[28] = I.DB.Menu.Create().
		SetMenuLevel(2).
		SetMenuType(1).
		SetParentID(3).
		SetPath("/sys/contract/index").
		SetName("Contract").
		SetComponent("/sys/contract/index").
		SetOrderNo(1).
		SetTitle("合同管理").
		SetIcon("").
		SetHideMenu(false)

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

// init menu data
func (I *InitDatabase) insertDictionariesData(ctx context.Context) error {

	propertyType, err := I.DB.Dictionary.Create().
		SetTitle("property_type").
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
		SetDictionary(propertyType).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("私教课").
		SetValue("私教课").
		SetKey("course").
		SetDictionary(propertyType).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	_, err = I.DB.DictionaryDetail.Create().
		SetStatus(1).
		SetTitle("团课").
		SetValue("团课").
		SetKey("class").
		SetDictionary(propertyType).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "db failed")
	}
	return nil
}
