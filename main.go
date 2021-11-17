package main

import (
	go_thrift_service "github.com/go/go-thrift-service/kitex_gen/github/go/go_thrift_service/helloservice"
	"log"
)

func main() {
	svr := go_thrift_service.NewServer(new(HelloServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
