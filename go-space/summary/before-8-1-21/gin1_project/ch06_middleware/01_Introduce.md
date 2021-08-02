# Middleware



# introduce

中间件是介于应用系统和[系统软件](https://baike.baidu.com/item/系统软件/215962)之间的一类软件，它使用系统软件所提供的基础服务（功能），衔接网络上应用系统的各个部分或不同的应用，能够达到资源共享、功能共享的目的。目前，它并没有很严格的定义，但是普遍接受IDC的定义：中间件是一种独立的系统软件服务程序，分布式应用软件借助这种软件在不同的技术之间共享资源，中间件位于客户机服务器的操作系统之上，管理计算资源和网络通信。从这个意义上可以用一个等式来表示中间件：中间件=平台+通信，这也就限定了只有用于分布式系统中才能叫中间件，同时也把它与支撑软件和实用软件区分开来。



使用 BasicAuth 中间件：https://gin-gonic.com/zh-cn/docs/examples/using-basicauth-middleware/

在中间件中使用 Goroutine：https://gin-gonic.com/zh-cn/docs/examples/goroutines-inside-a-middleware/



## 一、什么是中间件

- 开发者自定义的钩子（Hook） function 
- 类似python中的装饰器

## 二、中间件的作用

- 中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等
- 需要对某一类 function 进行通用的前置或者后置处理

## 三、使用场景

1. 逻辑执行之前
   - token等认证
   - 权限校验
   - 限流
   - 数据过滤
   - 白名单   refused  specify ip visit
   - ...
2. 逻辑执行之后
   - 数据过滤，比如敏感词等
   - 统一的响应头等
   - ...