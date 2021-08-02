# Template Syntax Two

## 一、with：伴随

1.作用：用于重定向 pipeline

2.使用

```
{{with .user}}
    {{.Id}}
    {{.Name}}
{{end}}
```

不用每个字段前面都加user.了

3.支持else，当长度为0时，显示else中的数据

```
{{with .user}}
    {{.Id}}
    {{.Name}}
{{else}}
    暂无数据
{{end}}
```

场景：登录的用户显示用户名，没有登录的用户显示"游客"

## 二、template

1.作用：引入另一个模板文件，对于模板的分模块处理很有用处，

2.使用：{ {template "模板名" pipeline} }

e.g.

```html
{{template "user/test.html" .}}



<p>
    {{/*template "ch03/01.tpl_syntax.html" . */}}
</p>

<p>
    {{template "ch03/02.base.html" . }}         
</p>
```



注意：

- 引入的模板文件中也要用{ {define "user/test.html"} } { {end} }包含
- 如果想在引入的模板中也需要获取动态数据，必须使用.访问当前位置的上下文



occur err:

```shell
executing "ch03/02.tpl_syntax.html" at <{{template "ch03/02.tpl_syntax.html" .}}>: exceeded maximum template depth (100000), 超过最大模板深度

self call self err
```



## 三、Annotation

允许多行文本注释，不允许嵌套

```

{{/* comment content
support new line */}}


{{/*
这是模板块注释
这是模板块注释
这是模板块注释
这是模板块注释
*/}}
```