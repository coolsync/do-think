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

// 需要收集日志 config入口
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

// Form etcd 根据key 获取配置项
func GetConf(key string) (logEntryConf []*LogEntryConf, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("etcd cli.Get err: %v", err)
		return
	}

	//
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%v\n", err)
			return
		}
	}
	return
}
