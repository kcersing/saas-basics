package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	paseto1 "github.com/hertz-contrib/paseto"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"net"
	"saas/app/user/config"
	"saas/app/user/infras"
	"saas/app/user/mysql"
	"saas/kitex_gen/cwg/user/userservice"
	"saas/pkg/consts"
	db2 "saas/pkg/db"
	"saas/pkg/md5"
	"saas/pkg/paseto"
	"saas/pkg/wechat"
	"strconv"
)

func main() {

	infras.InitLogger()
	infras.InitConfig()

	IP, Port := infras.InitFlag()

	r, info := infras.InitRegistry(Port)

	db := db2.InitDB(config.GlobalServerConfig.MysqlInfo.Source)

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.GlobalServerConfig.Name),
		provider.WithExportEndpoint(config.GlobalServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	//tg, err := paseto.NewTokenGenerator(
	//	config.GlobalServerConfig.PasetoInfo.SecretKey,
	//	[]byte(config.GlobalServerConfig.PasetoInfo.Implicit),
	//)

	tg, err := paseto.NewTokenGenerator(
		paseto1.DefaultPrivateKey,
		[]byte(paseto1.DefaultImplicit),
	)

	if err != nil {
		klog.Fatal(err)
	}
	// Create new server.

	srv := userservice.NewServer(
		&UserServiceImpl{
			OpenIDResolver: &wechat.AuthServiceImpl{
				AppID:     config.GlobalServerConfig.WXInfo.AppId,
				AppSecret: config.GlobalServerConfig.WXInfo.AppSecret,
			},
			EncryptManager:    &md5.EncryptManager{Salt: config.GlobalServerConfig.MysqlInfo.Salt},
			AdminMysqlManager: mysql.NewAdminManager(db, config.GlobalServerConfig.MysqlInfo.Salt),
			UserMysqlManager:  mysql.NewUserManager(db, config.GlobalServerConfig.MysqlInfo.Salt),
			TokenGenerator:    tg,
		},

		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		//server.WithReadWriteTimeout(time.Minute*5),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}),
	)

	err = srv.Run()
	if err != nil {
		klog.Fatal(err)
	}

}
