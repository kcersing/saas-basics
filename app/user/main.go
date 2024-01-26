package main

import (
	"context"
	"saas/kitex_gen/cwg/user/userservice"

	"log"
	"net"
	"saas/app/user/config"
	"saas/app/user/infras"
	"saas/app/user/mysql"
	"saas/pkg/consts"
	db2 "saas/pkg/db"
	"saas/pkg/md5"
	"saas/pkg/paseto"
	"saas/pkg/wechat"
	"strconv"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
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
	tg, err := paseto.NewTokenGenerator(
		config.GlobalServerConfig.PasetoInfo.SecretKey,
		[]byte(config.GlobalServerConfig.PasetoInfo.Implicit),
	)

	// Create new server.
	svr := userservice.NewServer(
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
		server.WithLimit(&limit.Option{
			MaxConnections: 2000,
			MaxQPS:         500,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.GlobalServerConfig.Name,
		}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
