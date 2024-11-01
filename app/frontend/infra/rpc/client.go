package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"rpc_gen/kitex_gen/orders/order/orderservice"
	"sync"
)

var (
	OrderClient orderservice.Client
	once        sync.Once
	err         error
	commonSuite client.Option
)

func InitClient() {
	once.Do(func() {
		initOrderClient()
	})
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	if err != nil {
		hlog.Fatal(err)
	}
}
