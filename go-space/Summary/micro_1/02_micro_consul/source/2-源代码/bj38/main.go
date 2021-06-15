package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"bj38/handler"

	bj38 "bj38/proto/bj38"
)

func main() {
	// New Service   -- 初始化服务器对象.
	service := micro.NewService(
		micro.Name("go.micro.srv.bj38"),   // 服务器名
		micro.Version("latest"),			// 版本
	)

	// Initialise service 与newService作用一致,但优先级高.后续代码运行期,初始化才有使用的必要.
	//service.Init()


	// Register Handler --- 注册服务
	bj38.RegisterBj38Handler(service.Server(), new(handler.Bj38))

	// Register Struct as Subscriber -- redis 发布订阅.
	//micro.RegisterSubscriber("go.micro.srv.bj38", service.Server(), new(subscriber.Bj38))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.bj38", service.Server(), subscriber.Handler)

	// Run service  --- 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
