package main

import (
	"fmt"
	"log"
	"logtrans/conf"
	"logtrans/es"
	"logtrans/kafka"
	"os"

	"gopkg.in/ini.v1"
)

// From kafka get data, send to es

func main() {
	// var cfg = new(conf.LogTransConf)
	var cfg conf.LogTransConf
	// 1 load conf file
	err := ini.MapTo(&cfg, "conf/conf.ini")
	if err != nil {
		log.Fatal("ini config failed, err: ", err)
	}
	fmt.Printf("cfg: %v\n", cfg)

	// 2 init es
	// get es connect client
	// 提供一个对外 write to Es function
	err = es.Init(cfg.EsConf.Address, cfg.EsConf.ChanMaxSize, cfg.Nums)
	if err != nil {
		log.Fatal("init es failed, err: ", err)
	}
	fmt.Fprintln(os.Stdout, "init es ok!")

	// 3 init kafka
	// connect kafka, create consumer side client
	// Get everyone partitions 对应 kafka consumer data, send to ES
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.Topic)
	if err != nil {
		log.Fatal("init kafka failed, err: ", err)
	}
	fmt.Fprintln(os.Stdout, "init kafka ok!")

	select {}
	// 3 from kafka get data
	// 4 send to es
}
