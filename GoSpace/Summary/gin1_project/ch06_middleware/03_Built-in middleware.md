# Built-in middleware



## 一、gin内置中间件

- func BasicAuth(accounts Accounts) HandlerFunc
- func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc    realm：认证分组
- func Bind(val interface{}) HandlerFunc //拦截请求参数并进行绑定
- func ErrorLogger() HandlerFunc       //错误日志处理
- func ErrorLoggerT(typ ErrorType) HandlerFunc //自定义类型的错误日志处理
- func Logger() HandlerFunc //日志记录
- func LoggerWithConfig(conf LoggerConfig) HandlerFunc
- func LoggerWithFormatter(f LogFormatter) HandlerFunc
- func LoggerWithWriter(out io.Writer, notlogged ...string) HandlerFunc
- func Recovery() HandlerFunc
- func RecoveryWithWriter(out io.Writer) HandlerFunc
- func WrapF(f http.HandlerFunc) HandlerFunc //将http.HandlerFunc包装成中间件
- func WrapH(h http.Handler) HandlerFunc       //将http.Handler包装成中间件