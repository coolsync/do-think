

# 复合主键

联合主键、组合主键

```
type Product struct {
    ID           int `gorm:"primary_key"`
    ERPID        int `gorm:"primary_key"`
}
```



# 原生sql

## 一、查询用Raw

```
var users []relate_tables.User
db.Raw("select * from users").Find(&users)
```

## 二、增改删用 Exec

```
db.Exec("insert into users (name,age) values(?,?)","hallen222",111)
db.Exec("update users set name = ? where id = ?","hallen111",1)
db.Exec("delete from  users where id = ?",1)
```

## 三、返回单条

```
row,_ := db.Raw("select * from users").Row()
```

## 四、返回多条

```
row,_ := db.Raw("select * from users").Rows()
```