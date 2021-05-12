# 其他数据类型渲染

## 一、结构体

```
后端：
type User struct {
    Id int
    Name string

}

func Hello(context *gin.Context)  {
    user := User{Id:1,Name:"hallen"}
    context.HTML(http.StatusOK,"user/index.html",user)
}

前端：
{{.Id}}
{{.Name}}
```

## 二、数组

```
后端：
func Hello(context *gin.Context)  {
    arr := [5]int{1,2,3,4,5}
    context.HTML(http.StatusOK,"user/index.html",arr)
}

前端：
{{range .}}
    {{.}}
{{end}}

或者:

{{range $i,$v := .}}
    {{$i}}
    {{$v}}
{{end}}

注意：range后面有两个变量，那就是角标和对应的元素值
    如果只有一个值，就是数组的元素值
```

## 三、结构体数组

```
后端
type User struct {
    Id int
    Name string

}
func Hello(context *gin.Context)  {
    arr_struct := [2]User{{Id:1,Name:"hallen1"},{Id:2,Name:"hallen2"}}
    context.HTML(http.StatusOK,"user/index.html",arr_struct)
}

前端：
{{range $v := .}}
    {{$v.Id}}
    {{$v.Name}}
    <br>
{{end}}
```

## 四、map

```
后端：

func Hello(context *gin.Context)  {
    map_data := map[string]string{
        "name":"hallen",
        "age":"18",
    }
    context.HTML(http.StatusOK,"user/index.html",map_data)
}
前端：
{{.name}}
{{.age}}
```

## 五、结构体+map

```
后端：
type User struct {
    Id int
    Name string

}

func Hello(context *gin.Context)  {

    map_struct_data := map[string]User{
        "user":User{Id:1,Name:"hallen"},
    }
    context.HTML(http.StatusOK,"user/index.html",map_struct_data )
}

前端：
{{.user.Id}}
{{.user.Name}}
```

## 六、切片

和数组类似，唯一的区别是不用指定长度了，长度是可变的