

# 日志介绍

## 一、日志的作用

1. 记录用户操作的审计日志，甚至有的时候就是管理部门的要求。
2. 快速定位问题的根源
3. 追踪程序执行的过程。
4. 追踪数据的变化
5. 数据统计和性能分析
6. 采集运行环境数据

日志是程序的重要组成部分

## 二、日志模板

1.什么是日志模板？

一种统一的格式，一种规范

2.日志模板的作用？

- 可读性
- 数据分析，二次挖掘





# 日志的使用

基于gin的日志中间件

## 一、使用日志文件

```go
// 1.创建日志文件
f, _ := os.Create("gin.log")

// 2.重新赋值DefaultWriter
gin.DefaultWriter = io.MultiWriter(f)

// 同时在控制台打印信息
gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
```

## 二、自定义日志格式

LoggerWithFormatter中间件指定日志格式

```go
router := gin.New()
router.Use(gin.Recovery())

router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

    // 你的自定义格式
    return fmt.Sprintf("%s\t|\t%s|\t%s|\t%s|\t%s|\t%d|\t%s|\t%s|\t%s \n",
        //客户端IP
        param.ClientIP,
        //时间格式
        param.TimeStamp.Format("2006-01-02 15:04:05"),
        //http请求方式 get post等
        param.Method,
        //客户端请求的路径
        param.Path,
        //http请求协议版本
        param.Request.Proto,
        //http请求状态码
        param.StatusCode,
        //耗时
        param.Latency,
        //http请求代理头
        param.Request.UserAgent(),
        //处理请求错误时设置错误消息
        param.ErrorMessage,
        )
    }))
```