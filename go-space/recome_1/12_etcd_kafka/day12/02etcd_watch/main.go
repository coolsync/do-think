package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	// new cli obj
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		log.Fatalf("connect to etcd failed, err:%v\n", err)
	}
	fmt.Fprintln(os.Stdout, "connect to etcd success")
	defer cli.Close()

	// watch
	// 派一个哨兵 一直 listen key的变化（新增、修改、删除）
	ch := cli.Watch(context.Background(), "hello")

	// 从通道尝试 get key-val info
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Fprintf(os.Stdout, "type: %s, key: %s, value: %s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}
}
