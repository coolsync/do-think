# 模板函数二

## 一、Format

1.实现了时间的格式化，返回字符串，也可以在后端转好，前端直接使用

2.使用方法

- 使用方法 { {.time_data.Format "2006/01/02 15:04:05"} }

- 设置时间格式比较特殊，需要按如下方式，一定不能变

  - "2006/01/02 15:04:05"
  - 貌似是GO的诞生的时间

- 和go语言的使用方法类似

- ```
  now := time.Now().Format("2006/01/02 15:04:05")
  ```

## 二、html

1.作用：转义文本中的html标签

如将“<”转义为“<”，“>”转义为“>”等

## 三、js

1.作用：返回用JavaScript的escape处理后的文本

escape函数可对字符串进行编码，这样就可以在所有的计算机上读取该字符串

可以使用unescape解吗

```
{{js "<script>xx</script>"}}

结果：\x3Cscript\x3Exx\x3C/script\x3E
```

模板函数整理

```
var builtins = FuncMap{
"and":      and,
"call":     call,
"html":     HTMLEscaper,
"index":    index,
"js":       JSEscaper,
"len":      length,
"not":      not,
"or":       or,
"print":    fmt.Sprint,
"printf":   fmt.Sprintf,
"println":  fmt.Sprintln,
"urlquery": URLQueryEscaper,


// Comparisons
"eq": eq, // ==
"ge": ge, // >=
"gt": gt, // >
"le": le, // <=
"lt": lt, // <
"ne": ne, // !=
}
```