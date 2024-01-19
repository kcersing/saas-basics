package main

import (
	"log"
	order "saas_basics/kitex_gen/cwg/order/orderservice"
)

func main() {
	svr := order.NewServer(new(OrderServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
