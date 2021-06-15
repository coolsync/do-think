# BasicAuth MiddleWare



## 1、BasicAuth Middleware Use





```
// 局部使用中间价
chap05.GET("/basic",gin.BasicAuth(gin.Accounts{
        "zs":"123456",
        "ls":"123",
        "ww":"1234",
    }),BasicAuthTest)


// 私有数据
var map_data map[string]interface{} = map[string]interface{}{
    "zs":gin.H{"age":18,"addr":"zs-xx"},
    "ls":gin.H{"age":19,"addr":"ls-xx"},
    "ww":gin.H{"age":20,"addr":"ww-xx"},
}


// 获取私有数据。如果没有权限则获取不到
func BasicAuthTest(ctx *gin.Context)  {

    user_name := ctx.Query("user_name")

    data ,ok := map_data[user_name]

    if ok{
        ctx.JSON(http.StatusOK,gin.H{"user_name":user_name,"data":data})
    }else {
        ctx.JSON(http.StatusOK,gin.H{"user_name":user_name,"data":"没有权限"})
    }
}
```

一文读懂HTTP Basic身份认证:https://juejin.im/entry/6844903586405564430

执行逻辑：

登录页面(没有中间件) -- 会设置session -- 其他路由回去session的key -- 获取对应的数据



controllers/ch05/ch05_routers.go:

```go
func Routers(ch05_routers *gin.RouterGroup) {
	// Use of BasicAuth middleware
	ch05_routers.GET("/auth", gin.BasicAuth(gin.Accounts{
		"bob":   "123456",
		"paul":  "123",
		"jerry": "1234",
	}), AuthBasicHandler)
}
```



controllers/ch05/02auth_basic.go:

```go
package ch05

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// set private data
var map_data map[string]interface{} = map[string]interface{}{
	"bob":   gin.H{"age": 18, "addr": "bob--xxx"},
	"paul":  gin.H{"age": 19, "addr": "paul--xxx"},
	"jerry": gin.H{"age": 20, "addr": "jerry--xxx"},
}

func AuthBasicHandler(ctx *gin.Context) {
	// get query params val
	user_name := ctx.Query("user_name")

	// query user is not exists
	user_info, ok := map_data[user_name]

	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"user": user_name,
			"data": user_info,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"user": user_name,
			"data": "user is not admin",
		})
	}
}
```





```go

1.
// BasicAuth returns a Basic HTTP Authorization middleware. It takes as argument a map[string]string where
// the key is the user name and the value is the password.
func BasicAuth(accounts Accounts) HandlerFunc {
	return BasicAuthForRealm(accounts, "")
}

2.
// BasicAuthForRealm returns a Basic HTTP Authorization middleware. It takes as arguments a map[string]string where
// the key is the user name and the value is the password, as well as the name of the Realm.
// If the realm is empty, "Authorization Required" will be used by default.
// (see http://tools.ietf.org/html/rfc2617#section-1.2)
func BasicAuthForRealm(accounts Accounts, realm string) HandlerFunc {
	if realm == "" {
		realm = "Authorization Required"
	}
	realm = "Basic realm=" + strconv.Quote(realm)
	pairs := processAccounts(accounts)
	return func(c *Context) {
		// Search user in the slice of allowed credentials
		user, found := pairs.searchCredential(c.requestHeader("Authorization"))
		if !found {
			// Credentials doesn't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
		c.Set(AuthUserKey, user)
	}
}

3.
func processAccounts(accounts Accounts) authPairs {
	length := len(accounts)
	assert1(length > 0, "Empty list of authorized credentials")
	pairs := make(authPairs, 0, length)
	for user, password := range accounts {
		assert1(user != "", "User can not be empty")
		value := authorizationHeader(user, password)
		pairs = append(pairs, authPair{
			value: value,
			user:  user,
		})
	}
	return pairs
}

4.
// Quote returns a double-quoted Go string literal representing s. The
// returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for
// control characters and non-printable characters as defined by
// IsPrint.
func Quote(s string) string {
	return quoteWith(s, '"', false, false)
}
```







## 2、WrapF

```go
gin.WrapF(IndexHandler)


func IndexHandler(w http.ResponseWriter, r *http.Request)  {
    ...
}
```

```go
func WrapFHandler(w http.ResponseWriter, r *http.Request) {

}

func Routers(ch05_routers *gin.RouterGroup) {
	// Use of BasicAuth middleware
	ch05_routers.GET("/auth", gin.BasicAuth(gin.Accounts{
		"bob":   "123456",
		"paul":  "123",
		"jerry": "1234",
	}), gin.WrapF(WrapFHandler), AuthBasicHandler)
}
```





## 3、The difference between *WrapH* and *WrapF*

```
需要自己去定义struct实现这个Handler接口

type TestStruct struct {}

func (test *TestStruct) TestH(w http.ResponseWriter, r *http.Request) {
    ...
}
```





```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// WrapF is a helper function for wrapping http.HandlerFunc and returns a Gin middleware.
func WrapF(f http.HandlerFunc) HandlerFunc {
	return func(c *Context) {
		f(c.Writer, c.Request)
	}
}

------------------------------

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// WrapH is a helper function for wrapping http.Handler and returns a Gin middleware.
func WrapH(h http.Handler) HandlerFunc {
	return func(c *Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
```

