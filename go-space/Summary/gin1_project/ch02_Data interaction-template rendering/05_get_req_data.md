# Get request parameters

## 一、Route with parameters: add parameter values directly to the path

带参数的路由：路径中直接加上参数值

e.g.      http://127.0.0.1:8080/user/hallen

1 .	The first case: use a placeholder **`:`** , you must specify this path

- 路由：engine.GET("/user/:name",Index)
- 如：[http://127.0.0.1:8080/user/hallen](http://127.0.0.1:8080/user/zhiliao)，这里必须指定name这个路径，不然会找不到
- 获取方式：context.Param("name")

2 .   第二种情况：使用占位符 *，可以不用匹配这个路径

- 路由：engine.GET("/user/*name",Index)
- 这里可以指定name这个路径，也可以不用指定
- 如：下面两种都可以访问
  - [http://127.0.0.1:8080/user/bob](http://127.0.0.1:8080/user/bob)，	---> /bob
  - [http://127.0.0.1:8080/user](http://127.0.0.1:8080/user),   	---> /
- 获取方式：ctx.Param("name")

区别：参数前面是使用冒号还是使用通配符，冒号的比如指定路径，通配符的可以不用

代码示例：

```go
func Index(c *gin.Context)  {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)

}

// engine.GET("/user/*name",Index)
engine.GET("/user/:name",Index)
```

## 二、Route with parameters: use the parameter name in the path

带参数的路由：路径中使用参数名

1.context.Query

- 传参：[http://127.0.0.1:8080/user/?name=hallen](http://127.0.0.1:8080/user?name=hallen)
- 获取：context.Query("name")

2.context.DefaultQuery

- 传参：[http://127.0.0.1:8080/user/?name=hallen](http://127.0.0.1:8080/user?name=hallen)
- 获取：context.DefaultQuery("name","hallen")

区别：DefaultQuery比Query多了个默认值，如果没有获取到会使用默认值

3.context.QueryArray

1. 传参：http://127.0.0.1:8080/user?name=1,2,3,4,5
2. 获取：names := context.QueryArray("name")
3. 数据结构：[1,2,3,4,5]

4.context.QueryMap

- 传参：

  [http://127.0.0.1:8080/user?name[1\]=hallen1&name[2]=hallen2](http://127.0.0.1:8080/user?name[1]=xx1&name[2]=xx2)

  http://localhost:8090/query_map?user[name]=haha&user[age]=18

- 获取：name_map := context.QueryMap("name")

- 数据结构：

  map[1:hallen1 2:hallen2]

  map[age:18 name:haha]







# Get POST request data

注意：是post请求

## 一、Get the data submitted by the form

1.context.PostForm("username")  获取表单中的name属性对应的值

示例代码：

```go
前端：submit提交
<form action="/hello_add" method="post">
    <input type="text" name="username"><br>
    <input type="text" name="age"><br>
    <input type="submit" value="提交">
</form>

后端：

func IndexAdd(context *gin.Context)  {

    name := context.PostForm("username")
    age := context.PostForm("age")
    context.String(200,"hello,%s,年龄为:%s",name,age)

}

func main() {
    engine := gin.Default()
    engine.LoadHTMLGlob("templates/**/*")
    engine.Static("/static","static")

    engine.POST("/hello_add",IndexAdd)

    engine.Run()

}
```

2.context.DefaultPostForm("username", "hallen")   如果没有获取到则使用指定的默认值

3.context.PostFormArray("love")          如果提交的数据有多个相同的name，获取数组

```go
前端：
<form action="/hello_add" method="post">
    <input type="text" name="username"><br>
    <input type="text" name="age"><br>
    ck1:<input type="checkbox" name="ck" value="1">
    ck2:<input type="checkbox" name="ck" value="2">
    ck3:<input type="checkbox" name="ck" value="3">
    <input type="submit" value="提交">
</form>


后端：
arr_ck := context.PostFormArray("ck")
```

1. context.PostFormMap("username")

```go
前端代码：
<form action="/hello_add" method="post">
    <input type="text" name="username[1]"><br>
    <input type="text" name="username[2]"><br>
    <input type="submit" value="提交">
</form>


后端代码：
map_name := context.PostFormMap("username")

数据结构：map[1:xx1 2:xx2]

注意：name要以map的格式定义，指定key，用户输入value，
```



## 二、Ajax interaction

前端使用ajax提交，后端和form表单的获取方式一样，唯一的区别就是返回的是json

```go
前端：

<script src="/static/js/jquery.min.js"></script>
<form>
    姓名:<input type="text" id="name">
    年龄:<input type="text" id="age">
    <input type="button" value="提交" id="btn_add">
</form>

<script>
    var btn_add = document.getElementById("btn_add");
    btn_add.onclick = function (ev) {
        var name = document.getElementById("name").value;
        var age = document.getElementById("age").value;

        $.ajax({
            url:"/hello3_add",
            type:"POST",
            data:{
                "name":name,
                "age":age
            },
            success:function (data) {
                alert(data["code"]);
                alert(data["msg"]);
            },
            fail:function (data) {

            }
        })

    }

</script>


注意：引入jquery.min.js：


后端：
name := context.PostForm("name")
age := context.PostForm("age")
fmt.Println(name)
fmt.Println(age)
messgae_map := map[string]interface{}{
    "code":200,
    "msg":"提交成功",
}
context.JSON(http.StatusOK,messgae_map)

//context.JSON(http.StatusOK,gin.H{
//    "code":200,
//    "msg":"提交成功",
//})
```