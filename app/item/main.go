package main

import (
	"log"
	item "saas_basics/kitex_gen/cwg/item/itemservice"
)

func main() {
	svr := item.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
