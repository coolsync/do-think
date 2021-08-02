# Use of static files

## 一、指定静态文件路径

engine.Static("/static", "static")

第一个参数是url，第二个参数是url对应的文件夹

engine.StaticFS("/static", http.Dir("static"))



## 二、前端引入静态文件

```html
<link rel="stylesheet" href="/static/user/index.css">
```



## backup:

```go
http.FileServer()	?
http.Dir()
```

```go
// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//     router.Static("/static", "/var/www")
func (group *RouterGroup) Static(relativePath, root string) IRoutes {
	return group.StaticFS(relativePath, Dir(root, false))
}


// Dir returns a http.Filesystem that can be used by http.FileServer(). It is used internally
// in router.Static().
// if listDirectory == true, then it works the same as http.Dir() otherwise it returns
// a filesystem that prevents http.FileServer() to list the directory files.
func Dir(root string, listDirectory bool) http.FileSystem {
	fs := http.Dir(root)
	if listDirectory {
		return fs
	}
	return &onlyFilesFS{fs}
}
```

