package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entAdapter "github.com/casbin/ent-adapter"
	"github.com/cloudwego/kitex/pkg/klog"
	"user/conf"
)

var casbinEnforcer *casbin.Enforcer

func Init() {
	var err error
	casbinEnforcer, err = newCasbin()
	if err != nil {
		klog.Fatal(err)
	}
}

func CasbinEnforcer() *casbin.Enforcer {
	return casbinEnforcer
}

func newCasbin() (enforcer *casbin.Enforcer, err error) {
	//adapter, err := entAdapter.NewAdapter("mysql", config.GlobalServerConfig.MySQLInfo.Host)

	adapter, err := entAdapter.NewAdapter("mysql", conf.GetConf().MySQL.DSN)

	if err != nil {
		klog.Error(err)
		return
	}

	var text string
	if conf.GetConf().Casbin.ModelText == "" {
		text = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	} else {
		text = conf.GetConf().Casbin.ModelText
	}

	m, err := model.NewModelFromString(text)
	if err != nil {
		klog.Error(err)
		return
	}
	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		klog.Error(err)
		return
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		klog.Error(err)
		return
	}
	return
}
