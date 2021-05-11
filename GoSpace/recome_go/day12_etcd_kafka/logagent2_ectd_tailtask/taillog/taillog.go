package taillog

import (
	"fmt"
	"logagent/kafka"

	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

// TailTask： 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

// New TailTask obj
func NewTailTask(path, topic string) *TailTask {
	tailObj := &TailTask{
		path:  path,
		topic: topic,
	}
	err := tailObj.init() // 根据路径去打开对应的日志
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v\n", err)
		return nil
	}
	return tailObj
}

// init read log file conf
func (t *TailTask) init() (err error) {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}

	// according to path and config, instance *Tail obj
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Printf("tail file failed, err: %v\n", err)
		return
	}

	go t.run() // 直接去采集日志发送到kafka

	return nil
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines: // 从tailObj的通道中一行一行的读取日志数据
			// 3.2 发往Kafka
			// kafka.SendToKafka(t.topic, line.Text) // 函数调用函数
			// 先把日志数据发到一个通道中
			kafka.SendToChan(t.topic, line.Text)
			// kafka那个包中有单独的goroutine去取日志数据发到kafka
		}

	}
}