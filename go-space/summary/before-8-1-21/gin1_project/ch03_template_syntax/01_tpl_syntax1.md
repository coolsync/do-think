# Template Syntax One



Source code path: ch03/01_tpl_syntax.go 	

Template path:  templates/ch03/01.tpl_syntax.html



## 1、统一使用 `{{` 和 `}}` 作为左右标签

Unified use of `{{` and `}}` as the left and right labels





## 2、Context

- .      Access the context of the current location  访问当前位置 上下文
- $     Refer to the context of the root level of the current template 引用当前模板语句根级 上下文
- $.    Reference to the root-level context in the template 引用模板中的根级 上下文





## 3、支持go语言的符号，这里只是符号的支持

Support the symbols of the go language, here is only the symbol support



1. String：	{ { "mark" } }

2. Original string: 	{{ ‘mark‘}} will not be escaped  (原始字符串：{{ 'mark' }} 不会转义)

3. Byte type: 	{ { ' a' } } -->97 ascll码对应表： http://ascii.911cha.com/

4. Nil type:	{{ print nil } } { {nil } }只有nil会报错：nil is not a command





## 4、Consume Define variables

1. Define：

```html
{{$username := "xxxx"}}
```

2. use：

```html
{{$username }}
```

注意：只能在当前模板中使用





## 5、pipline

1.可以是上下文的变量输出，也可以是函数通过管道传递的返回值

e.g.

- { {.Name} }   是上下文的变量输出，是个pipline
- { { "mark" | len } } 是函数通过管道传递的返回值，是个pipline





## 6、if

1.if...else

```html
{{if .name}}
    有姓名
{{else}}
    没有姓名
{{end}}
```

2.if嵌套

```
成年人而且带了身份证的准进：
{{if .A}}
    {{if .C }}
        可以进
    {{else}}
        不能进
    {{end}}
{{else}}
    未成年不能进
{{end}}
```





## 7、range

```html
第一种：
{{range $v := .arr_struct}}
    {{$v.Name}}
    {{$v.Age}}
    {{$v.Gender}}
{{end}}

第二种：
{{ range .arr_struct }}
    {{.Name }}
    {{.Age}}
    {{ $.total}} // 使用 $. 引用模板中的根级上下文
{{end}}


<p>
    {{range .arr}}
    {{.}}
    {{end}}
</p>

<p>
    {{range $v := .arr}}
    {{/*$v*/}}
    {{.}}
    {{end}}
</p>
```



range也支持else，当长度为0时，执行else

```html
{{range .total}}
    {{.}}
{{else}}
    {{ 0 }}                 {{/* 当 .total 为空 或者 长度为 0 时会执行这里 */}}
{{end}}

<p>
    {{range .arr1}}
    {{.}}
    {{else}}
    No Data
    {{end}}
</p>
```