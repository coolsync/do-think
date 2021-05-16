# Template function 1

Source code path: 

Template path:  



## 一、print

```
print:对应 fmt.Sprint
printf:对应 fmt.Sprintf

println:对应 fmt.Sprintln

{{print "hallen"}}
```

###### 格式化输出:

- %c：字符型，可以把输入的数字按照ASCII码相应转换为对应的字符
- %d:一个十进制数值，基数为10
- %f：以标准计数法表示的浮点数或者复数值
- %s:字符串。输出字符串中的字符直至字符串中的空字符（字符串以'\0'结尾，这个'\0'即空字符）
- %t：以true或者false输出的布尔值
- %T：查看类型
- %v：自动匹配类型输出，适用于大多数类型

fmt.Sprint和fmt.Print的区别：

- Print有打印，Sprint没有打印，但是可以返回结果

## 二、管道符

变量可以使用符号 | 在函数间传递

```
{{.Name | printf "%s"}}
```

## 三、括号   优先级

```
{{printf "nums is %s %d" (printf "%d %d" 1 2) 3}}

{{printf "name:%s,addr:%s" "hallen" (printf "%s-%s" "北京市" "西城区")}}
```

## 四、and

1.作用：只要有一个为空，则整体为空，如果都不为空，则返回最后一个

2.使用：

```
{{and .X .Y .Z}}

{{and .name .age}}
// .name .age 都有值， 显示 .age val, 
// .name .age 其中之一有值， 都不显示

{{and .age .name}}
// .name .age 都有值， 显示 .name val, 
// .name .age 其中之一有值， 都不显示
```



## 五、or

1.作用：只要有一个不为空，则返回第一个不为空的，否则返回空



2.使用：

```
{{or .X .Y .Z}}
```

## 六、call(看自定义模板函数)

1.作用：可以调用函数，并传入参数

2.使用：

```
{{call .Field .Arg1 .Arg2}}
```

## 七、index

1.作用：读取指定类型对应下标的值

2.支持 map, slice, array, string

e.g.

```
数据准备：
arr := []int{1,2,3,4,5}
message := map[string]interface{}{
    "arr":arr,
    "name":map[string]interface{}{"name":"xx111","age":18},
}


使用：
{{index .name "name"}}
{{index .arr 1}}

注意：数组的角标从0开始的
```

## 八、len

1.作用：返回对应类型的长度

2.支持类型：map, slice, array, string, chan

e.g.

```
{{.arr | len}}
```

## 九、not

1.作用：返回输入参数的否定值

2.使用

```
{{not .arr1}}
```

## 十、urlquery

1.作用：有些符号在URL中是不能直接传递的，如果要在URL中传递这些特殊符号，那么就要使用他们的编码了。

2.使用

```
{{urlquery "http://www.baidu.com"}}

结果：http%3A%2F%2Fwww.baidu.com

% 后面的就是字符的16进制的字符码
```

## 十一、eq / ne / lt / le / gt / ge

1.eq：等于

2.ne：不等于

3.lt：小于    less

4.le：小于等于

5.gt：大于 greater

6.ge：大于等于

```
{{eq .num 18}}

{{ne .num 17}}

{{lt .num 16}}

{{le .num 18}}

{{gt .num 18}}

{{ge .num 18}}
```

eq 和其他函数不一样的地方是，支持多个参数

```
{{eq .num 10 11 17 18}}
```