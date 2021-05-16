package main

import (
	"fmt"
	"log"
	"logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

// logAgent入口程序

var (
	cfg = new(conf.AppConf)
)

// func run() {
// 	// 1. 读取日志
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			// 2. 发送到kafka
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 			fmt.Println(line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

func main() {
	// 0. Get init conf file
	err := ini.MapTo(cfg, "conf/conf.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		log.Fatalf("init Kafka failed,err: %v\n", err)
	}
	// fmt.Println("init kafka success.")
	fmt.Fprintln(os.Stdout, "init kafka success.")

	// 2. 初始化 etcd连接
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		log.Fatalf("init etcd failed, err: %v\n", err)
	}
	fmt.Println("init etcd success.")

	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		log.Fatalf("etcd.GetConf failed, err: %v\n", err)
	}

	// 2.2 拍一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现热加载配置）
	// cli.Get
	fmt.Printf("%#v\n", logEntryConf)
	for idx, val := range logEntryConf {
		fmt.Printf("idx: %v, val: %v\n", idx, val)
	}

	// 3. 收集日志发往Kafka
	taillog.Init(logEntryConf)



	// // 3. 具体的业务
	// run()
}
