# logrus第三方库的使用

## 一、安装

```
github.com/sirupsen/logrus
```

## 二、使用logrus自定义日志中间件

```go
func LoggerFile(c *gin.Context)  {
    file_dir := "logs/" + "gin_project.log"

    //写入文件
    src, err := os.OpenFile(file_dir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

    if err != nil {
        fmt.Println("err", err)
    }


    //实例化
    logger := logrus.New()

    //设置输出
    logger.Out = src

    //设置日志级别
    logger.SetLevel(logrus.DebugLevel)

    //设置日志格式，格式化时间
    logger.SetFormatter(&logrus.TextFormatter{TimestampFormat:"2006-01-02 15:04:05"})


    // 开始时间
    startTime := time.Now()

    // 处理请求
    c.Next()

    // 结束时间
    endTime := time.Now()

    // 执行时间
    latencyTime := endTime.Sub(startTime)

    // 请求方式
    reqMethod := c.Request.Method

    // 请求路由
    reqUri := c.Request.RequestURI

    // 状态码
    statusCode := c.Writer.Status()

    // 请求IP
    clientIP := c.ClientIP()

    // 日志格式
    logger.Infof("| %3d | %13v | %15s | %s | %s |",
        statusCode,
        latencyTime,
        clientIP,
        reqMethod,
        reqUri,
    )

}
```





## 三、使用自定义日志中间件

```
router.Use(LoggerFile)
```

拆分：https://blog.csdn.net/u010918487/article/details/86146691

logrus 不支持输出文件名和行号

# logrus使用配置文件

Path:	`/home/dart/DoThinking/GoSpace/GoGin_2021_3_11/gin_project`

`main.go`

```go
package main

import (
	_ "ginproject/logs_source"
)
```



## 一、配置文件

`log_conf.json`:

```json
{
    "log_dir": "logs/gin_project.log",
    "log_level": "info"
}
```



## 二、加载配置

```go
package logs_source

import (
	"encoding/json"
	"os"
)

type LogConf struct {
	LogDir   string `json:"log_dir"`	// 相对路径 + 文件名：e.g: logs/gin_project.log
	LogLevel string `json:"log_level"`	// 日志級別
}

func LoadLogConf() *LogConf {
	var log_conf LogConf
	// get bs
	bs, err := os.ReadFile("./conf/log_conf.json")
	if err != nil {
		panic(err)
	}

	// json
	if err := json.Unmarshal(bs, &log_conf); err != nil {
		panic(err)
	}

	return &log_conf
}

```



## 三、初始化

```go
package logs

import (
    "os"
    "github.com/sirupsen/logrus"
)

var Log = logrus.New() // 创建一个log示例

func init() {

    // 日志級別映射
    log_level_mapping := map[string]logrus.Level{
        "trace":logrus.TraceLevel,
        "debug":logrus.DebugLevel,
        "info":logrus.InfoLevel,
        "warn":logrus.WarnLevel,
        "error":logrus.ErrorLevel,
        "fata":logrus.FatalLevel,
        "panic":logrus.PanicLevel,

    }

    // 初始化配置
    log_conf := LoadLogConf()

    //设置输出
    dir, err := os.OpenFile(log_conf.LogDir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        fmt.Println("err", err)
    }
    Log.Out = dir

    // 设置日志级别
    Log.Level = log_level_mapping[log_conf.LogLevel]

    //设置日志格式，格式化时间
    Log.SetFormatter(&logrus.TextFormatter{TimestampFormat:"2006-01-02 15:04:05"})
}


TextFormatter:文本格式
JSONFormatter：json格式
```

```go
package logs_source

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create 一个实例
var Log = logrus.New()

func init() {
	// var log_conf LogConf
	log_conf := LoadLogConf()

	// 设置 log 输出文件
	// f, err := os.OpenFile(log_conf.LogDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// f, err := os.OpenFile(log_conf.LogDir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	f, err := os.OpenFile(log_conf.LogDir, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// defer f.Close()
	Log.Out = f

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

	// 格式化 log, 设置 
	Log.SetFormatter(&logrus.TextFormatter{TimestampFormat:"2006-01-02 15:04:05"})
}
```



## 四、使用



`controller/chapter07/router.go`

```go
package chapter07

import "github.com/gin-gonic/gin"

func Router(ch07 *gin.RouterGroup) {
	ch07.GET("/log_data", LogTest)
}

```



`controller/chapter07/log_data.go`

```go
package chapter07

import (
	"ginproject/logs_source"

	"github.com/gin-gonic/gin"
)

func LogTest(ctx *gin.Context) {
	logs_source.Log.Info("this is info")
	logs_source.Log.Debug("this is debug")
	logs_source.Log.Warn("this is warn")

	logs_source.Log.WithField("name", "xx").Info("info")

	fileds_mapping := map[string]interface{}{
		"id":   18,
		"name": "haha",
	}

	logs_source.Log.WithFields(fileds_mapping).Info("info")
}
```



```go
logs.Log.Warn("这是一个warnning级别的日志")
logs.Log.WithFields(logrus.Fields{
            "msg": "测试的错误",
        }).Warn("这是一个warnning级别的日志")
```



ret:

```go
time="2021-04-09T15:50:05+08:00" level=info msg="this is info"
time="2021-04-09T15:50:05+08:00" level=warning msg="this is warn"
time="2021-04-09T15:50:05+08:00" level=info msg=info name=xx
time="2021-04-09T15:50:05+08:00" level=info msg=info id=18 name=haha
time="2021-04-09 19:15:47" level=info msg="this is info"
time="2021-04-09 19:15:47" level=warning msg="this is warn"
time="2021-04-09 19:15:47" level=info msg=info name=xx
time="2021-04-09 19:15:47" level=info msg=info id=18 name=haha
```

