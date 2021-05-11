### Transactions交易次数



## Disable Default Transaction禁用默认交易

GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency, you can disable it during initialization if it is not required, you will gain about 30%+ performance improvement after thatGORM执行写入（创建/更新/删除）操作在事务内运行以确保数据一致性，如果不需要，可以在初始化期间将其禁用，此后性能将提高约30％以上

```
// Globally disable
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  SkipDefaultTransaction: true,
})

// Continuous session mode
tx := db.Session(&Session{SkipDefaultTransaction: true})
tx.First(&user, 1)
tx.Find(&users)
tx.Model(&user).Update("Age", 18)
```

## Transaction交易

To perform a set of operations within a transaction, the general flow is as below.要在一个事务中执行一组操作，一般流程如下。

```
db.Transaction(func(tx *gorm.DB) error {
  // do some database operations in the transaction (use 'tx' from this point, not 'db')
  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
    // return any error will rollback
    return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
    return err
  }

  // return nil will commit the whole transaction
  return nil
})
```

### Nested Transactions嵌套交易

GORM supports nested transactions, you can rollback a subset of operations performed within the scope of a larger transaction, for example:GORM支持嵌套事务，您可以回滚在较大事务范围内执行的操作的子集，例如：

```
db.Transaction(func(tx *gorm.DB) error {
  tx.Create(&user1)

  tx.Transaction(func(tx2 *gorm.DB) error {
    tx2.Create(&user2)
    return errors.New("rollback user2") // Rollback user2
  })

  tx.Transaction(func(tx2 *gorm.DB) error {
    tx2.Create(&user3)
    return nil
  })

  return nil
})

// Commit user1, user3
```

## Transactions by manual手动交易

```
// begin a transaction
tx := db.Begin()

// do some database operations in the transaction (use 'tx' from this point, not 'db')
tx.Create(...)

// ...

// rollback the transaction in case of error
tx.Rollback()

// Or commit the transaction
tx.Commit()
```

### A Specific Example一个具体的例子

```go
{
	if err := CreateUsers(db); err != nil {
		p("add users failed.")
	}
}

func CreateUsers(db *gorm.DB) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
	  if r := recover(); r != nil {
		tx.Rollback()
	  }
	}()
  
	if err := tx.Error; err != nil {
	  return err
	}
	
	// add users failed. 已有 id 11， 后面不执行， 直接撤销
	if err := tx.Create(&relate_tables.User{ID: 11, Name: "mark222", Age: 40, Addr: "xx"}).Error; err != nil {
	   tx.Rollback()
	   return err
	}
  
	if err := tx.Create(&relate_tables.User{ID: 15, Name: "mark222", Age: 40, Addr: "xx"}).Error; err != nil {
	   tx.Rollback()
	   return err
	}
  
	return tx.Commit().Error
  }

```



```
func CreateAnimals(db *gorm.DB) error {
  // Note the use of tx as the database handle once you are within a transaction
  tx := db.Begin()
  defer func() {
    if r := recover(); r != nil {
      tx.Rollback()
    }
  }()

  if err := tx.Error; err != nil {
    return err
  }

  if err := tx.Create(&Animal{Name: "Giraffe"}).Error; err != nil {
     tx.Rollback()
     return err
  }

  if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
     tx.Rollback()
     return err
  }

  return tx.Commit().Error
}
```

## SavePoint, RollbackTo保存点，回滚到

GORM provides `SavePoint`, `RollbackTo` to save points and roll back to a savepoint, for example:GORM提供SavePoint，RollbackTo以保存点并回滚到保存点，例如：

```
tx := db.Begin()
tx.Create(&user1)

tx.SavePoint("sp1")
tx.Create(&user2)
tx.RollbackTo("sp1") // Rollback user2

tx.Commit() // Commit user1
```