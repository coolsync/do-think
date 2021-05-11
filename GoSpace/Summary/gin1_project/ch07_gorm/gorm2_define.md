# 模型定义一

## 一、模型定义的作用

用作数据库数据转换和自动建表

## 二、模型名和表名的映射关系

1. 规则
   - 第一个大写字母变为小写，
   - 遇到其他大写字母变为小写并且在前面加下划线，
   - 连着的几个大写字母，只有第一个遵循上面的两条规则，其他的大写字母转为小写，不加下划线，遇到小写，前面的第一个大写字母变小写并加下划线
   - 复数形式
2. 举例
   - User --> users       首字母小写，复数
   - UserInfo --> user_infos
   - DBUserInfo --> db_user_infos     
   - DBXXXXUserInfo --> dbxxxx_user_infos

## 三、在默认表名上加其他规则

```go
// 在默认表名前加sys_前缀

gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
    return "sys_" + defaultTableName;
}


自定义表名：
func (模型) TableName() string{
    return "新的表名"
}
```

## 四、结构体字段名和列名的对应规则

1. 规则 * 列名是字段名的蛇形小写
2. 举例
   - Name --> name
   - CreatedTime --> create_time
3. 可以通过gorm标签指定列名，AnimalId    int64     `gorm:"column:beast_id"`  

## 五、gorm.Model

基本模型定义gorm.Model，包括字段ID，CreatedAt，UpdatedAt，DeletedAt

只需要在自己的模型中指定gorm.Model匿名字段，即可使用上面的四个字段

```go
// 添加字段 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type User struct {
    gorm.Model
    Name string
}
```

ID：主键自增长

CreatedAt：用于存储记录的创建时间

UpdatedAt：用于存储记录的修改时间

DeletedAt：用于存储记录的删除时间



# 模型定义二

## 一、结构体标签gorm的使用

```
type UserInfo struct {
    Id int `gorm:"primary_key"`
    Name string `gorm:"index"`
    Age int 
}
```

## 二、gorm标签属性值

- -： 忽略，不映射这个字段 `gorm:"-"`
- primary_key：主键 `gorm:"primary_key"`
- AUTO_INCREMENT：自增 `gorm:"AUTO_INCREMENT"`
- not null：不为空，默认为空 `gorm:"not null"`
- index：索引， `gorm:"index"`
  - 创建索引并命名： `gorm:"index:idx_name_code"`
- - 优化查询，相当于图书的目录
- unique_index:唯一索引 `gorm:"unique_index"`
- unique：唯一 `gorm:"unique"`
- column：指定列名 `gorm:"column:user_name"`
- size：字符串长度,默认为255 `gorm:"size:64"`
- type：设置sql类型 `gorm:"type:varchar(100)"` // 不推荐直接改类型
- default `default:'galeone'` 默认值

多个属性值之间用分号分隔(英文的;):`gorm:"size:64;not null"`







# 模型定义三

## 一、一对一

1.属于

```go
     // UserProfile属于User，外键是在UserProfile模型，外键字段为：UId


     type User struct {
        Id int
        Name string
        Age int
        Addr string

    }


    type UserProfile struct {
        Id int
        Pic string
        CPic string
        Phone string
        User User `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`  // 关联关系
        //UserID int  // 默认关联字段为Id
        UId int // uid
    }

    注：

    // 外键默认使用UserID，如果不指定外键，则使用默认的外键字段,
    // 默认关联ID，通过AssociationForeignKey指定关联字段
```

2.包含

```go
     // UserProfile 包含一个 User, 外键在User模型，外键字段为:PId

    type User struct {
        Id int
        Name string
        Age int
        Addr string
        PId int

    }


    type UserProfile struct {
        Id int
        Pic string
        CPic string
        Phone string
        User User `gorm:"foreignKey:PId;AssociationForeignKey:Id"`  // 关联关系

    }
```

属于：关系和外键在同一方，有关系的那一方属于另外一个模型

包含：关系和外键不在同一方，有关系的那一方包含另外一个有外键的模型

## 二、一对多

```go
    type User2 struct {
        Id int
        Name string
        Age int
        Addr string
        Articles []Article `gorm:"foreignKey:UId;references:Id"`
    }

    type Article struct {
        Id int
        Title string
        Content string
        Desc string
        // 外键
        UId int
    }
```

## 三、多对多

```go
    type Article2 struct {
        AId int `gorm:"primary_key:true"`
        Title string
        Content string
        Desc string
        Tags []Tag `gorm:"many2many:Article2s2Tags"`  // ;foreignKey:AId;references:TId

    }

    type Tag struct {
        TId int `gorm:"primary_key:true"`
        Name string
        Desc string
    }
```

创建/更新时不会保存关联：`gorm:"save_associations:false"`