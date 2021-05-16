package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

// 需要收集日志 config 入口
type LogEntryConf struct {
	Path  string `json:"path"`  // 存放日志的path
	Topic string `json:"topic"` // 日志要发往 kafka 哪个topic
}

func Init(addr string, timeout time.Duration) (err error) {
	// new cli obj
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})

	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}

	fmt.Fprintln(os.Stdout, "connect to etcd success")
	return
}

// 从ETCD中根据key获取配置项
func GetConf(key string) (logEntryConf []*LogEntryConf, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("etcd cli.Get err: %v", err)
		return
	}

	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%v\n", err)
			return
		}
	}
	return
}

// etcd watch
func WatchConf(key string, newConfCh chan<- []*LogEntryConf) {
	ch := cli.Watch(context.Background(), key)
	// 从通道尝试取值(监视的信息)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			// 通知taillog.tskMgr
			// 1. 先判断操作的类型
			var newConf []*LogEntryConf
			if evt.Type != clientv3.EventTypeDelete {
				// 如果是删除操作，手动传递一个空的配置项
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed, err:%v\n", err)
					continue
				}
			}
			fmt.Printf(" get new conf:%v\n", newConf)
			newConfCh <- newConf
		}
	}
}

// // 从ETCD中根据key获取配置项
// func GetConf1(key string) (logEntryConf []*LogEntryConf, err error) {
// 	// get
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	resp, err := cli.Get(ctx, key)
// 	cancel()
// 	if err != nil {
// 		fmt.Printf("get from etcd failed, err:%v\n", err)
// 		return
// 	}
// 	for _, ev := range resp.Kvs {
// 		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
// 		err = json.Unmarshal(ev.Value, &logEntryConf)
// 		if err != nil {
// 			fmt.Printf("unmarshal etcd value failed,err:%v\n", err)
// 			return
// 		}
// 	}
// 	return
// }
