# 使用模板

## 一、engine.LoadHTMLGlob：推荐

只有一个参数，通配符，如：template/*    意思是找当前项目路径下template文件夹下所有的html文件

e.g.：engine.LoadHTMLGlob("templates/*")

## 二、engine.LoadHTMLFiles：不推荐

不定长参数，可以传多个字符串，使用这个方法需要指定所有要使用的html文件路径

e.g.：engine.LoadHTMLFiles("templates/index.html","template/user.html")

## 三.指定模板路径

```go
// 使用*gin.Context下的HTML方法

func Hello(context *gin.Context)  {
    name := "zhiliao"
    context.HTML(http.StatusOK,"index.html",name)
}
```

***注意***：不要使用goland里面run，否则会报错

```
panic: html/template: pattern matches no files: `templates/*`
```

在cmd运行即可

## 四、多级目录的模板指定

如果有多级目录，比如templates下有user和article两个目录，如果要使用里面的html文件，必须得在Load的时候指定多级才可以，比如：engine.LoadHTMLGlob("templates/**/*")

1.有几级目录，得在通配符上指明

```
两级：engine.LoadHTMLGlob("templates/**/*")
三级：engine.LoadHTMLGlob("templates/**/**/*")
```

2.Specify html file， 并在project下创建 templates/user/index.html 

```
// 除了第一级的templates路径不需要指定，后面的路径都要指定
e.g.：context.HTML(http.StatusOK,"user/index.html","mark")
```

3.In html file,

```html
必须使用
{{ define "user/index.html" }}

html内容

{{ end }}

包起来
```