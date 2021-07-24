package taillog

import "logagent/etcd"

// 具体管理每一个 tail read log, write to kafka 进程处理

// global task manager
var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntryConf
	//tskMap map[string]*TailTask
}

func Init(logEntry []*etcd.LogEntryConf) {
	tskMgr = &tailLogMgr{
		logEntry: logEntry, // 把当前的日志收集项配置信息保存起来
	}

	for _, logEntryObj := range logEntry {
		NewTailTask(logEntryObj.Path, logEntryObj.Topic)
	}
}
