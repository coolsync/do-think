# Custom Validators 



https://gin-gonic.com/zh-cn/docs/examples/custom-validators/

## 1、安装包

go get [github.com/go-playground/validator](https://github.com/go-playground/validator)



## 2、使用

1.定义验证器

```go
var Len6Valid validator.Func = func(fl validator.FieldLevel) bool {
    data := fl.Field().Interface().(string)
    if len(data) > 6 {
        fmt.Println("false")
        return false
    }else {
        fmt.Println("true")
        return true
    }
}


注意：必须为validator.Func类型
```

2.注册验证器

```go
if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("len_valid", valid.Len6Valid)
    }

在路由匹配前，main中即可
```

3.结构体中使用

```go
type Article struct {
    Id int `form:"id"`
    Title string `form:"title" binding:"required,len_valid"`
    Desc string `form:"desc" binding:"required,len_valid"`
}

这里必须binding中，而且名称为前面注册的字符串名称
```



# beego中的验证器

#### 用于data valid ：

-  valid data是否合法

#### **安装**：

```
go get github.com/astaxie/beego/validation
```

#### 验证方法：

- Required 不为空，即各个类型要求不为其零值
- Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证
- Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
- Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
- MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
- MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
- Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
- Alpha alpha字符，有效类型：string，其他类型都将不能通过验证
- Numeric 数字，有效类型：string，其他类型都将不能通过验证
- AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
- Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf(“%v”, obj).Match)
- AlphaDash alpha 字符或数字或横杠 -_，有效类型：string，其他类型都将不能通过验证
- Email 邮箱格式，有效类型：string，其他类型都将不能通过验证
- IP IP 格式，目前只支持 IPv4 格式验证，有效类型：string，其他类型都将不能通过验证
- Base64 base64 编码，有效类型：string，其他类型都将不能通过验证
- Mobile 手机号，有效类型：string，其他类型都将不能通过验证
- Tel 固定电话号，有效类型：string，其他类型都将不能通过验证
- Phone 手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
- ZipCode 邮政编码，有效类型：string，其他类型都将不能通过验证

简单例子：

```go
     

type LoginParams struct {
      Name string valid:"Required"
      Age int    valid:"Required;MinSize(2)"
      Addr string    valid:"Required"
 }

func (l *LoginController) Post()  {
        valid := validation.Validation{}
     // 解析到结构体
      params := LoginParams{}
      if err := l.ParseForm(&params); err != nil {
          //handle error
          return
      }

      //重写错误信息：validation.SetDefaultMessage(map)
      var messages = map[string]string{
        "Required": "不能为空",
        "MinSize":  "最短长度为 %d",
        "Length":   "长度必须为 %d",
        "Numeric":  "必须是有效的数字",
        "Email":    "必须是有效的电子邮件地址",
        "Mobile":   "必须是有效的手机号码",
      }
      validation.SetDefaultMessage(messages)

      // 校验
      b, err := valid.Valid(&params)

      // 验证StructTag 是否正确
      if err != nil {

          fmt.Println(err)
      }

      if !b {   
          // 验证没通过，则b为false
          for _, err := range valid.Errors {
              fmt.Println(err.Key, err.Message)
              message := err.Key + err.Message
              l.Ctx.WriteString(message)
          }
      }
}
```



##### 1 Init beego valid

`chapter04/bind_valid_data.go`

```go
// Valid for Beego
type Article struct {
	Id      int    `form:"-"`
	Title   string `form:"title" valid:"Required"` // custom validator
	Content string `form:"content"`
	Desc    string `form:"desc"`
}

func ToBindValid(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_valid.html", nil)
}

func DoBindValid(ctx *gin.Context) {
	var article Article

	err := ctx.ShouldBind(&article)
	fmt.Println(err)
	fmt.Println(article)

	// init beego valid
	valid := validation.Validation{}

	fmt.Println(valid)
	bl, err1 := valid.Valid(&article)
	fmt.Println(err1)

	if !bl { // valid err
		for _, err := range valid.Errors {
			fmt.Println(err.Key)
			fmt.Println(err.Message)
		}
	}

	ctx.String(http.StatusOK, "OK")
}
```



##### 2 Rewrite error info：validation.SetDefaultMessage(map)	

```go
func DoBindValid(ctx *gin.Context) {
	var article Article

	err := ctx.ShouldBind(&article)
	fmt.Println(err)
	fmt.Println(article)

	// init beego valid
	valid := validation.Validation{}
	fmt.Println("valid: ", valid)
    
	// MessageTmpls store commond validate template
	var MessageTmpls = map[string]string{
		// "Required":     "Can not be empty",
		"Required":     "不能为空", // modify "Required"
		"Min":          "Minimum is %d",
		"Max":          "Maximum is %d",
		"Range":        "Range is %d to %d",
		"MinSize":      "Minimum size is %d",
		"MaxSize":      "Maximum size is %d",
		"Length":       "Required length is %d",
		"Alpha":        "Must be valid alpha characters",
		"Numeric":      "Must be valid numeric characters",
		"AlphaNumeric": "Must be valid alpha or numeric characters",
		"Match":        "Must match %s",
		"NoMatch":      "Must not match %s",
		"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
		"Email":        "Must be a valid email address",
		"IP":           "Must be a valid ip address",
		"Base64":       "Must be valid base64 characters",
		"Mobile":       "Must be valid mobile number",
		"Tel":          "Must be valid telephone number",
		"Phone":        "Must be valid telephone or mobile phone number",
		"ZipCode":      "Must be valid zipcode",
	}
	validation.SetDefaultMessage(MessageTmpls)

	bl, err1 := valid.Valid(&article)
	fmt.Println(err1)
	
    // 提示 map
    key_mapping := map[string]string{
		"Title.Required.": "标题",
	}
    
	if !bl { // valid err
		for _, err := range valid.Errors {
			fmt.Println(err.Key)
			fmt.Println(err.Message)
			// ctx.String(http.StatusOK, err.Message)
			ctx.String(http.StatusOK, key_mapping[err.Key]+ " " +err.Message)
        }
	}
}
```

##### 3  Multiple field valid

```go
func DoBindValid(ctx *gin.Context) {
	var article Article

	err := ctx.ShouldBind(&article)
	fmt.Println(err)
	fmt.Println(article)

	// init beego valid
	valid := validation.Validation{}
	fmt.Println("valid: ", valid)

	//重写错误信息：validation.SetDefaultMessage(map)
	// MessageTmpls store commond validate template
	var MessageTmpls = map[string]string{
		// "Required":     "Can not be empty",
		"Required": "不能为空", // modify "Required"
		"Min":      "Minimum is %d",
		"Max":      "Maximum is %d",
		"Range":    "Range is %d to %d",
		// "MinSize":      "Minimum size is %d",
		"MinSize": "字符串最小长度为 %d",
		"MaxSize": "Maximum size is %d",
		// "Length":       "Required length is %d",
		"Length":       "字符串固定长度为 %d",
		"Alpha":        "Must be valid alpha characters",
		"Numeric":      "Must be valid numeric characters",
		"AlphaNumeric": "Must be valid alpha or numeric characters",
		"Match":        "Must match %s",
		"NoMatch":      "Must not match %s",
		"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
		// "Email":        "Must be a valid email address",
		"Email":   "必须是有效邮箱的",
		"IP":      "Must be a valid ip address",
		"Base64":  "Must be valid base64 characters",
		"Mobile":  "Must be valid mobile number",
		"Tel":     "Must be valid telephone number",
		"Phone":   "Must be valid telephone or mobile phone number",
		"ZipCode": "Must be valid zipcode",
	}

	validation.SetDefaultMessage(MessageTmpls)

	key_mapping := map[string]string{
		"Title.Required.": "标题",
		"Content.Length.": "内容",
		"Desc.MinSize.":   "描述",
		"Email.Email.":    "邮箱",
	}

	bl, err1 := valid.Valid(&article)
	fmt.Println("err1: ", err1)

	if !bl { // valid err
		for _, err := range valid.Errors {
			fmt.Println(err.Key)
			fmt.Println(err.Message)
			ctx.String(http.StatusOK, key_mapping[err.Key]+" "+err.Message+"\n")
		}
	}

	ctx.String(http.StatusOK, "OK!")
}
```





#### **通过 StructTag valid data：**

- 验证function写在 "valid" tag 的标签里
- 各个验证规则之间用分号 ";" 分隔，分号后面可以有空格
- 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
- 正则function(Match)的匹配模式用两斜杠 "/" 括起来
- 各个function的结果的 key 值为字段名.验证function名

```

```

#### 多个StructTag[之间用空格隔开]()