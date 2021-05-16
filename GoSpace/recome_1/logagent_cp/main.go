package main

import (
	"fmt"
	"log"
	"logagent/conf"
	"logagent/etcd"
	"logagent/kafka"
	"logagent/taillog"
	"logagent/utils"
	"os"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

// logAgent入口程序

var (
	cfg = new(conf.AppConf)
)

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
	fmt.Fprintln(os.Stdout, "init kafka success.")

	// 2. 初始化 etcd连接
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		log.Fatalf("init etcd failed, err: %v\n", err)
	}
	fmt.Println("init etcd success.")

	// 2.1 From etcd Get 日志收集项的配置信息
	// Get local ip
	ipStr, err := utils.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	
	fmt.Println(etcdConfKey)

	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		log.Fatalf("etcd.GetConf failed, err: %v\n", err)
	}

	// 2.2 派一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现热加载配置）
	// cli.Get
	fmt.Printf("%#v\n", logEntryConf)
	for idx, val := range logEntryConf {
		fmt.Printf("idx: %v, val: %v\n", idx, val)
	}

	// 3. 收集日志发往Kafka
	taillog.Init(logEntryConf)
	// 因为NewConfChan访问了tskMgr的newConfChan, 这个channel是在taillog.Init(logEntryConf) 执行的初始化
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) // 哨兵发现最新的配置信息会通知上面的那个通道
	wg.Wait()
}
