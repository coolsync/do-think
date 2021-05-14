# File Upload

## 一、Form upload file

1.Single file upload

```go
前端：
<form action="/upload2" method="post" enctype="multipart/form-data">
    <input type="file" name="file">
    <input type="submit" value="提交">
</form>

注意：设置enctype参数

后端：
func Upload2(context *gin.Context)  {
    fmt.Println("+++++++++++++++")
    file,_ := context.FormFile("file")  // 获取文件
    fmt.Println(file.Filename)

    file_path := "upload/" + file.Filename  // 设置保存文件的路径，不要忘了后面的文件名

    context.SaveUploadedFile(file, file_path)  // 保存文件

    context.String(http.StatusOK,"上传成功")
}


防止文件名冲突，使用时间戳命名：

unix_int := time.Now().Unix()   // 时间戳，int类型
time_unix_str := strconv.FormatInt(unix_int,10)   // 讲int类型转为string类型，方便拼接，使用sprinf也可以

file_path := "upload/" + time_unix_str + file.Filename   // 设置保存文件的路径，不要忘了后面的文件名
context.SaveUploadedFile(file, file_path)  // 保存文件
```

2.Multiple file upload

```go
前端：
<form action="/upload2" method="post" enctype="multipart/form-data">
    <input type="file" name="file">
    <input type="file" name="file">
    <input type="submit" value="提交">
</form>

注意：不要忘了enctype参数

后端：
func Upload2(context *gin.Context)  {
    form,_ := context.MultipartForm()
    files := form.File["file"]

    for _,file := range files {    // 循环
        fmt.Println(file.Filename)
        unix_int := time.Now().Unix() // 时间戳，int类型
        time_unix_str := strconv.FormatInt(unix_int,10) // 讲int类型转为string类型，方便拼接，使用sprinf也可以

        file_path := "upload/" + time_unix_str + file.Filename // 设置保存文件的路径，不要忘了后面的文件名
        context.SaveUploadedFile(file, file_path) // 保存文件
    }

    context.String(http.StatusOK,"上传成功")
}

注意：form.File["file"]   这里是中括号，不是小括号
```

二、Ajax upload file

后端代码和form表单方式一样的

1.Single file upload

```js
前端：
<script src="/static/js/jquery.min.js"></script>
<form>
    {{/*<input type="file" name="file">*/}}
    用户名:<input type="test" id="name"><br>
    <input type="file" id="file">
    <input type="button" value="提交" id="btn_add">
</form>


<script>
    var btn_add = document.getElementById("btn_add");
    btn_add.onclick = function (ev) {
        var name = document.getElementById("name").value;
        var file = $("#file")[0].files[0];

        var form_data = new FormData();
        form_data.append("name",name);
        form_data.append("file",file);



        $.ajax({
            url:"/upload2",
            type:"POST",
            data:form_data,
            contentType:false,
            processData:false,
            success:function (data) {
                console.log(data);
            },
            fail:function (data) {
                console.log(data);
            }
        })



    }
</script>

注意：
    1.引入juery.min.js文件
    2.ajax中需要加两个参数：
        contentType:false,
        processData:false,
```

processData:false   默认为true，当设置为true的时候,jquery ajax 提交的时候不会序列化 data，而是直接使用data

contentType: false 不使用默认的application/x-www-form-urlencoded这种contentType

- 分界符：目的是防止上传文件中出现分界符导致服务器无法正确识别文件起始位置
- ajax 中 contentType 设置为 false 是为了避免 JQuery 对其操作，从而失去分界符



2.Multiple file upload

name名称不相同就是个单文件上传

name名称相同

```js
<script>
    var btn_add = document.getElementById("btn_add");
    btn_add.onclick = function (ev) {
        var name = document.getElementById("name").value;
        console.log($(".file"));
        var files_tag = $(".file");
        var form_data = new FormData();

        for (var i=0;i<files_tag.length;i++){
            var file = files_tag[i].files[0];
            form_data.append("file",file);

        }

        console.log(files);
        form_data.append("name",name);

        $.ajax({
            url:"/upload2",
            type:"POST",
            data:form_data,
            contentType:false,
            processData:false,
            success:function (data) {
                console.log(data);
            },
            fail:function (data) {
                console.log(data);
            }
        })



    }
</script>
```