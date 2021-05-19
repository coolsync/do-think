# day13课上笔记



# 内容回顾

## LogAgent的实现

1. 配置文件版LogAnent实现
   1. kafka的使用
   2. tailf第三方包的使用
   3. ini配置文件解析
2. etcd版本的LogAgent实现
   1. 项目启动时从etcd拉取收集日志项信息
   2. 利用watch实时监听etcd中配置的变化
   3. 利用IP每个LogAgent分别从etcd拉取自己的配置

## 留的两个思考题(自主学习能力)

1. Raft协议

2. watch底层实现的原理

# 今日内容

## 1. LogTransfer

从kafka里面把日志取出来,写入ES,使用Kibana做可视化的展示

## 2. Elasticsearch

[ES介绍博客地址](https://www.liwenzhou.com/posts/Go/go_elasticsearch/)

[ES搭建指南](https://docs.qq.com/doc/DTmZxQUdHeFRXU2dP)



### Modify elastics host ?

1. network.host: 192.168.0.107
2. cluster.initial_master_nodes: ["node-1"]

### Run Er

bootstrap check failure [1] of [1]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]





### Kibana

详见[ES搭建指南](https://docs.qq.com/doc/DTmZxQUdHeFRXU2dP)

Elasticsearch.yml:

```yml
#node.name: node-1
node.name: node-1

#
#network.host: 192.168.0.1
network.host: 192.168.0.107

#cluster.initial_master_nodes: ["node-1", "node-2"]
cluster.initial_master_nodes: ["node-1"]
```



kibana.yml

```yml
# The URLs of the Elasticsearch instances to use for all your queries.
#elasticsearch.hosts: ["http://localhost:9200"]
elasticsearch.hosts: ["http://192.168.0.107:9200"]


# Specifies locale to be used for all localizable strings, dates and number formats.
# Supported languages are the following: English - en , by default , Chinese - zh-CN .
#i18n.locale: "en"
i18n.locale: "zh-CN"
```



### kafka消费

根据topic找所有的分区

每一个分区去消费数据

```go
package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			}
		}(pc)
	}
}
```

### LogTransfer实现

### 加载配置文件

```go
// 0. 加载配置文件
var cfg = new(conf.LogTransferCfg)
err := ini.MapTo(cfg, "./conf/cfg.ini")
if err != nil {
	fmt.Printf("init config failed, err:%v\n", err)
	return
}
fmt.Printf("cfg:%v\n", cfg)
```

两个坑:

1. 在一个函数中修改变量一定要传指针
2. 在配置文件对应的结构体中一定要设置tag(特别是嵌套的结构体)





## 系统监控

gopsutil做系统监控信息的采集,写入influxDB,使用grafana作展示

prometheus监控:采集性能指标数据,保存起来,使用grafana作展示

![1569664389751](/home/dart/DoThinking/GoSpace/Summary/come/day13/day13.assets/1569664389751.png)



![architecture](/home/dart/DoThinking/GoSpace/Summary/come/day13/day13.assets/architecture.png)

## 项目总结

1. 项目的架构(图)
2. 为什么不用ELK
3. logAgent里面如何保证日志不丢/重启之后继续收集日志(记录读取文件的offset)
4. kafka课上整理的那一些
5. etcd的watch的原理
6. es 相关知识点



找工作:

1. 找开发的话还是算法和数据结构(刷leetcode)
2. 找运维开发的话前端自己会一点会加分,时下热点的技术栈
3. 学历(尽快自己想办法)
4. 简历好好写
5. Boss直聘等该花钱花钱
6. 项目一定要自己写一遍



国庆后:

web框架:gin

微服务

Docker和K8s



# Ohter

https://github.com/valeriansaliou/sonic

1. Sonic是一个开源搜索索引服务器, 构建简单，高性能且轻量级。
2. Sonic接受用户查询，返回标识符（例如 消息，文章，CRM联系人等）
3.  Sonic不存储文档, 获取搜索结果的应用 必须从另一个数据库（例如，MongoDB，MySQL等）提取结果数据，因为搜索结果返回的是ID
4. 使用传统的开源搜索索引软件（例如Elasticsearch等）需要巨大的服务器CPU和RAM。
5. Sonic：简单的功能，简单的网络协议， “可搜索的Redis”。

# 今日分享

怎么玩 很重要.