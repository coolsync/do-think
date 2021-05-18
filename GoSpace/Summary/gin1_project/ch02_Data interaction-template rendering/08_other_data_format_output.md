# Other data format output

## JSON

```
context.JSON(http.StatusOK,gin.H{
        "code":200,
        "tag":"<br>",
        "msg":"提交成功",
        "html":"<b>Hello, world!</b>",
    })

结果：{"code":200,"html":"\u003cb\u003eHello, world!\u003c/b\u003e","msg":"提交成功","tag":"\u003cbr\u003e"}
```

## AsciiJSON

生成具有转义的非 ASCII 字符的 ASCII-only JSON

```
context.AsciiJSON(http.StatusOK,gin.H{
        "code":200,
        "tag":"<br>",
        "msg":"提交成功",
        "html":"<b>Hello, world!</b>",
    })

结果：{"code":200,"html":"\u003cb\u003eHello, world!\u003c/b\u003e","msg":"\u63d0\u4ea4\u6210\u529f","tag":"\u003cbr\u003e"}
```

## JSONP

使用 JSONP 向不同域的服务器请求数据。如果查询参数存在回调，则将回调添加到响应体中

```
context.JSONP(http.StatusOK,gin.H{
        "code":200,
        "tag":"<br>",
        "msg":"提交成功",
        "html":"<b>Hello, world!</b>",
    })

结果：{"code":200,"html":"\u003cb\u003eHello, world!\u003c/b\u003e","msg":"提交成功","tag":"\u003cbr\u003e"}
```

如果传输的数据在两个不同的域，由于在javascript里无法跨域获取数据，所以一般采取script标签的方式获取数据，传入一些callback来获取最终的数据，这就有可能造成敏感信息被劫持

## PureJSON

```
context.PureJSON(http.StatusOK,gin.H{
        "code":200,
        "tag":"<br>",
        "msg":"提交成功",
    })

结果：{"code":200,"html":"<b>Hello, world!</b>","msg":"提交成功","tag":"<br>"}
```

## SecureJSON

使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 "while(1)," 到响应体。

```
names := []string{"lena", "austin", "foo"}

// 将输出：while(1);["lena","austin","foo"]
context.SecureJSON(http.StatusOK, names)
```

json劫持：利用网站的cookie未过期，然后访问了攻击者的虚假页面，那么该页面就可以拿到json形式的用户敏感信息

## XML

```
context.XML(http.StatusOK,gin.H{
        "code":200,
        "tag":"<br>",
        "msg":"提交成功",
        "html":"<b>Hello, world!</b>",
 })

结果：
<map>
<code>200</code>
<tag><br></tag>
<msg>提交成功</msg>
<html><b>Hello, world!</b></html>
</map>
```

## YAML

```
context.YAML(http.StatusOK,gin.H{
        "code":200,
        "tag":"<br>",
        "user":gin.H{"name":"zhiliao","age":18},
        "html":"<b>Hello, world!</b>",
    })

结果：
code: 200
html: <b>Hello, world!</b>
tag: <br>
user:
  age: 18
  name: zhiliao
```

## ProtoBuf

定义proto文件

user.proto

```
syntax = 'proto3';

package user;

message User {
    string name = 1;
    int32 age = 2;
}

messages类型：message，server,enum
```

导出go文件

```
protoc --go_out=. user.proto
```

使用

```
user_data := &user.User{
        Name:"zhiliao",
        Age:18,
    }
context.ProtoBuf(200,user_data)

注意：是指针
```