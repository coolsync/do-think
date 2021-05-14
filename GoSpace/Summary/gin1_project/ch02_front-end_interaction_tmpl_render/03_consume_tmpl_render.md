## Custom template renderer

## 一、使用"html/template"

要指定所有的html路径，不推荐 Not recommended

```go
router := gin.Default()
html := template.Must(template.ParseFiles("test1.html", "test2.html"))
router.SetHTMLTemplate(html)
router.Run(":8080")
```