package logs_source

import (
	"os"

	"github.com/sirupsen/logrus"
)

// create 一个实例
var Log = logrus.New()

func init() {
	// var log_conf LogConf
	log_conf := LoadLogConf()

	// 设置 log 输出文件
	// f, err := os.OpenFile(log_conf.LogDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// f, err := os.OpenFile(log_conf.LogDir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	f, err := os.OpenFile(log_conf.LogDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
	// defer f.Close()
	Log.Out = f

	// var AllLevels = []Level{
	// 	PanicLevel,
	// 	FatalLevel,/
	// 	ErrorLevel,
	// 	WarnLevel,
	// 	InfoLevel,
	// 	DebugLevel,
	// 	TraceLevel,
	// }

	// 设置 log 级别
	level_mapping := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
	}

	Log.SetLevel(level_mapping[log_conf.LogLevel])

	// 格式化 log
	Log.SetFormatter(&logrus.TextFormatter{})
}
