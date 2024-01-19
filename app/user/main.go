package main

import (
	"log"
	user "saas_basics/kitex_gen/cwg/user/userservice"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
