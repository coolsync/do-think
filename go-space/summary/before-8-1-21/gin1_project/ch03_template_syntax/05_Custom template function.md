# Custom template function

## 一、Define function

```
func SubStr(str string,l int) string {

    ret := str[0:l]
    return (ret + "...")
}
```

## 二、SetFuncMap

```
engine.SetFuncMap(template.FuncMap{
        "SubStr": SubStr,      // 字符串名称是前端使用的名称
    })
```

## 三、Front end use (前端使用)

```
{{SubStr "qwertyuu" 3}}
```

注意：是左闭右开区间





```go

main.go:

// Create router
router := gin.Default()
// Set consume tpl func
router.SetFuncMap(template.FuncMap{
    "add_num":   ch03.AddNum,
    "str_len":   ch03.SubStr,
    "safe_html": ch03.SafeHTML,
})



ch03/02_tpl_func.go:

// define func
func AddNum(n1, n2 int) int {
	return n1 + n2
}

func SubStr(str string, n int) string {
	str_len := len(str)

	if str_len <= n {
		return str
	}

	return str[:n] + "..."
}

func SafeHTML(s string) template.HTML {
	return template.HTML(s)
}
```

