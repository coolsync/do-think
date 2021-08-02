# 错误处理

如果在执行SQL查询的时候，出现错误，GORM 会将错误信息保存到 *gorm.DB 的Error字段，我们只要检测Error字段就可以知道是否存在错误。

## 一、处理单个错误

```
err := db.Where("name = ?", "tizi365").First(&user).Error
if err != nil {
  // 错误处理
}

// 或者

result := db.Where("name = ?", "jinzhu").First(&user)

if result.Error != nil {
  // 错误处理
}
```

## 二、处理多个错误

通过GetErrors获取错误列表

```
errors := db.First(&user).Limit(10).Find(&users).GetErrors()
fmt.Println(len(errors)) // 打印错误数量

// 遍历错误内容
for _, err := range errors {
  fmt.Println(err)
}
```

## 三、错误种类

gorm.错误类型进行判断，如gorm.ErrRecordNotFound

1.RecordNotFound：查询不到数据，不适用于切片

```
gorm.IsRecordNotFoundError(err) {
    // 没有查询到数据
}
```

2.ErrInvalidSQL：无效sql

3.ErrInvalidTransaction：事务有错

4.ErrCantStartTransaction：无法开启事务，出现在使用Begin的情况下

5.ErrUnaddressable：使用不可寻址的值，传递的指针值不对