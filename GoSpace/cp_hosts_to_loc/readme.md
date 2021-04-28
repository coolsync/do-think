
## golang string 去最后一个字符
```go
func main() {
	fmt.Println("Hello, 世界")
	var s string
	s = "222,"
	strings.TrimRight(s, ",")
	fmt.Println(s)
	s = strings.TrimRight(s, ",")
	fmt.Println(s)
}
```


## golang 切片合并

```go
func main() {
    a := []int{1, 2, 3}
    b := []int{2, 3, 4, 5, 6}
    a = append(a, b...)
    fmt.Println(a)
}
```
