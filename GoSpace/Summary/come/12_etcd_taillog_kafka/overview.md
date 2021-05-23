# day12



## Etcd doc

https://etcd.io/docs/v3.4/

https://www.liwenzhou.com/posts/Go/go_etcd/

## Go operate etcd



### 1 put, get

run etcd server:

```shell
$ etcd
```



```go
package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd client put/get demo
// use etcd/clientv3

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
    fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "q1mi", "dsb")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "q1mi")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
```



### 2 watch

`watch`用来获取未来更改的通知。	



```go
package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	// create etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()

    // 派一个哨兵 watch "topic" 这个key的变化
	wch := cli.Watch(context.Background(), "topic")	// context.Background() 可以一直for wch
	
    // 从 channel 尝试取值 (监视的信息)
	for wresp := range wch { // <-chan WatchResponse
		for _, evt := range wresp.Events {
			fmt.Printf("Type: %s, Key: %s, Value: %s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
		}
	}
}
```



启动 etcd 客户端

```go
etcdctl --endpoints=127.0.0.1:2379 put topic 300
etcdctl --endpoints=127.0.0.1:2379 del topic
```



result:

```shell
connect to etcd success
Type: PUT, Key: topic, Value: 300
Type: DELETE, Key: topic, Value:
```





occur probrem:

```shell
go: found github.com/coreos/bbolt in github.com/coreos/bbolt v1.3.5
go: day12/etcd_ex imports
        go.etcd.io/etcd/clientv3 tested by
        go.etcd.io/etcd/clientv3.test imports
        github.com/coreos/etcd/auth imports
        github.com/coreos/etcd/mvcc/backend imports
        github.com/coreos/bbolt: github.com/coreos/bbolt@v1.3.5: parsing go.mod:
        module declares its path as: go.etcd.io/bbolt
                but was required as: github.com/coreos/bbolt
```



```shell
go: finding module for package google.golang.org/grpc/naming
day12/etcd_ex imports
        go.etcd.io/etcd/clientv3 tested by
        go.etcd.io/etcd/clientv3.test imports
        github.com/coreos/etcd/integration imports
        github.com/coreos/etcd/proxy/grpcproxy imports
        google.golang.org/grpc/naming: module google.golang.org/grpc@latest found (v1.37.1), but does not contain package google.golang.org/grpc/naming
```



solution:

in  go.mod:

```go
module day12

go 1.16

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc v1.37.1 => google.golang.org/grpc v1.26.0
```





## 日志收集项目 etcd



启动kafka

ectd_put:

写入 key and value to ectd



logagent/main.go

从 etcd 获取日志收集配置项的信息

打印 多个配置项

 交给

logagent/etcd/etcd.go:

type LogEntry struct 

GetConf(key string)





日志 要发往 哪个 kafka topic





occur error info:

```shell
{"level":"warn","ts":"2021-05-16T09:49:11.238+0800","caller":"clientv3/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"endpoint://client-5af3bfd3-293e-4a5a-94c8-c012cd952cf4/","attempt":0,"error":"rpc error: code = DeadlineExceeded desc = latest balancer error: all SubConns are in TransientFailure, latest connection error: connection error: desc = \"transport: Error while dialing dial tcp: missing address\""}
get from etcd failed, err:context deadline exceeded
```



solution:

无法配置 ini, 从 windows 上写的文件 复制到 linux 上 要小心

clientv3.New, 空地址 也可初始化





## taillog and kafka intract

taillog.go

single read log file

send to kafka topic and line data



taillog_mgr.go

manage multiple taillog



kafka.go

SendToChan() 暴露一个channel, 接受来自 taillog info

sendTokafka()  将得到的data 交给 kafka produce msg





## 