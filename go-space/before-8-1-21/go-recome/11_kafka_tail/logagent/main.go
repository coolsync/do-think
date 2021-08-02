package main

import (
	"fmt"
	"log"
	"logagent/conf"
	"logagent/kafka"
	"logagent/taillog"
	"time"

	"gopkg.in/ini.v1"
)

// logAgent入口程序

var (
	cfg = new(conf.AppConf)
)

func run() {
	// 1. 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
			fmt.Println(line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0. Get init conf file
	// cfg, err := ini.Load("conf/conf.ini")
	// if err != nil {
	// 	log.Fatal("Fail to read file: ", err)
	// }
	// fmt.Println(cfg.Section("kafka").Key("address"))
	// fmt.Println(cfg.Section("kafka").Key("topic"))
	// fmt.Println(cfg.Section("taillog").Key("filename"))

	err := ini.MapTo(cfg, "conf/conf.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success.")

	// 2. 打开日志文件准备收集日志
	err = taillog.Init(cfg.TaillogConf.FileName)
	// fmt.Println(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success.")
	run()
}
