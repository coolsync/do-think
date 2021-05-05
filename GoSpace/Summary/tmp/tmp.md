# proto install



解压后

mv bin/proto 到 /usr/local/bin/



测试

$ proto



proto-gen-go



# config parse



kafka

https://kafka.apache.org/quickstart



#### 1. [TERMINATE THE KAFKA ENVIRONMENT](https://kafka.apache.org/quickstart#quickstart_kafkaterminate)

终止 kafka 运行环境

If you also want to delete any data of your local Kafka environment including any events you have created along the way, run the command:

```bash
$ rm -rf /tmp/kafka-logs /tmp/zookeeper
```



#### 2. [START THE KAFKA ENVIRONMENT](https://kafka.apache.org/quickstart#quickstart_startserver):

首先

```bash
# Start the ZooKeeper service
# Note: Soon, ZooKeeper will no longer be required by Apache Kafka.
$ bin/zookeeper-server-start.sh config/zookeeper.properties
```

然后

```bash
# Start the Kafka broker service
$ bin/kafka-server-start.sh config/server.properties
```

Once all services have successfully launched, you will have a basic Kafka environment running and ready to use.



#### 3 [WRITE SOME EVENTS INTO THE TOPIC](https://kafka.apache.org/quickstart#quickstart_send) Proudcer





```bash
$ bin/kafka-console-producer.sh --topic quickstart-events --bootstrap-server localhost:9092
This is my first event
This is my second event
```



#### 4 [READ THE EVENTS](https://kafka.apache.org/quickstart#quickstart_consume) Consumer



```bash
$ bin/kafka-console-consumer.sh --topic quickstart-events --from-beginning --bootstrap-server localhost:9092
This is my first event
This is my second event
```





go-ini package:	parse conf file

https://ini.unknwon.io/







# etcd

etcd

https://etcd.io/

https://www.zhihu.com/column/c_1248405562469597184



### install

script .sh run:

https://linux.cn/article-13106-1.html



https://github.com/etcd-io/etcd/releases

```shell
ETCD_VER=v3.4.15

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GITHUB_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
```





### simple operation

PUT：

```bash
etcdctl --endpoints=http://127.0.0.1:2379 put key value
```

GET：

```bash
etcdctl --endpoints=http://127.0.0.1:2379 GET key
```

DEL:

```bash
etcdctl --endpoints=http://127.0.0.1:2379 DEL key
```





### 解决go mod拉取etcd依赖包报错的问题



https://blog.csdn.net/skh2015java/article/details/111060465

https://www.imooc.com/article/315263?block_id=tuijian_wz



解决方法
删除原来已生成得go.mod和go.sum

go mod init

go mod edit -replace github.com/coreos/bbolt@v1.3.4=go.etcd.io/bbolt@v1.3.4

go mod edit -replace google.golang.org/grpc@v1.37.0=google.golang.org/grpc@v1.26.0

go mod tidy



### Libraries and tools

https://etcd.io/docs/v3.4/integrations/



### go connect etcd

https://github.com/etcd-io/etcd/tree/master/client/v3



注意包名：

"github.com/coreos/etcd/clientv3"

"go.etcd.io/etcd/clientv3"



put和get操作
put 命令用来设置键值对数据， get 命令用来根据key获取值。



```go
package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

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
	_, err = cli.Put(ctx, "bob", "quick")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "bob")
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



watch操作
watch 用来获取未来更改的通知。

```go
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
```



补充：

```go
type Watcher interface {
    
	// (see https://github.com/etcd-io/etcd/issues/8980)
	Watch(ctx context.Context, key string, opts ...OpOption) WatchChan

	// RequestProgress requests a progress notify response be sent in all watch channels.
	RequestProgress(ctx context.Context) error

	// Close closes the watcher and cancels all watch requests.
	Close() error
}

func (clientv3.Watcher).Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan



type WatchChan <-chan WatchResponse

type WatchResponse struct {
	Header pb.ResponseHeader
	Events []*Event

	// CompactRevision is the minimum revision the watcher may receive.
	CompactRevision int64

	// Canceled is used to indicate watch failure.
	// If the watch failed and the stream was about to close, before the channel is closed,
	// the channel sends a final response that has Canceled set to true with a non-nil Err().
	Canceled bool

	// Created is used to indicate the creation of the watcher.
	Created bool

	closeErr error

	// cancelReason is a reason of canceling watch
	cancelReason string
}


type Event mvccpb.Event

type Event struct {

	Type Event_EventType `protobuf:"varint,1,opt,name=type,proto3,enum=mvccpb.Event_EventType" json:"type,omitempty"`

	Kv *KeyValue `protobuf:"bytes,2,opt,name=kv,proto3" json:"kv,omitempty"`
	
    // prev_kv holds the key-value pair before the event happens.
    
	PrevKv               *KeyValue `protobuf:"bytes,3,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

```





# logagent 整合 ectd



1. 在 conf.ini 配置 ectd 条目
2. 在 ectd.go 创建 Init function, 初始化 ectd client 连接
3. 在 ectd_put 向 ectd server put json data
4. 创建 GetConf function, 用于get ectd 中 json data
5. 在 main.go , 引入 Init and GetConf, 打印 json 解析出来的 data 



```go
需要收集日志入口
存放日志的path
日志要发往 kafka 哪个topic
```





设置 ectd key 配置到 conf.ini

把 GetConf param 更改为 cfg.Ectd.Key 映射













按照

将从taillog获取 topic, line data, 发送到当下kafka内部channel

从通道中读取数据

具体管理每一个 tail read log, write to kafka 进程处理， 





taillog.go

type TailTask struct



func NewTailTask(path, topic string) (tailObj *TailTask)



// init read log file conf

func (t *TailTask) init() (err error) 

   

go t.run() // 直接去采集日志发送到kafka

func (t *TailTask) run()





