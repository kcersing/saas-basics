package rpc

import (
	"common/clientsuite"
	"facade/conf"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"rpc_gen/kitex_gen/member/memberservice"
	"rpc_gen/kitex_gen/orders/order/orderservice"
	"sync"
)

var (
	OrderClient  orderservice.Client
	MemberClient memberservice.Client

	once         sync.Once
	err          error
	registryAddr string
	commonSuite  client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: "facade",
		})

		initMemberClient()
	})
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	if err != nil {
		hlog.Fatal(err)
	}
}
func initMemberClient() {
	MemberClient, err = memberservice.NewClient("member", commonSuite)
	if err != nil {
		hlog.Fatal(err)
	}
}
