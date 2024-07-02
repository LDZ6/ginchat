## `AutoMigrate` 

1. **自动创建表**: 如果数据库中没有与结构体对应的表，`AutoMigrate` 会自动创建该表。
2. **自动更新表结构**: 如果表已经存在，但与结构体定义不完全匹配（例如，新增了字段），`AutoMigrate` 会自动更新表结构以匹配结构体。
3. **简化开发流程**: 在开发过程中频繁修改数据结构时，使用 `AutoMigrate` 可以减少手动更新数据库表结构的工作量。

### 示例代码

假设你有一个名为 `UserBasic` 的结构体，如下所示：

```go
package models

type UserBasic struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"unique;not null"`
    Email     string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

你可以通过以下代码将 `UserBasic` 结构体对应的表结构自动迁移到数据库中：

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "myapp/models" // 假设你的模型在这个路径下
)

func main() {
    // 初始化数据库连接
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }

    // 自动迁移
    db.AutoMigrate(&models.UserBasic{})
}
```

运行上述代码后，GORM 会自动检查 `UserBasic` 结构体，并在数据库中创建或更新相应的表结构。

---

## `db.Create()`

### 作用
`db.Create()` 方法用于将一个新的记录插入到数据库中。它会根据传入的结构体创建一条新记录，并将生成的主键值（如果有）回填到结构体的相应字段中。

### 示例代码

假设你有一个名为 `UserBasic` 的结构体，如下所示：

```go
package models

type UserBasic struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"unique;not null"`
    Email     string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

你可以通过以下代码将一个新的用户记录插入到数据库中：

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "myapp/models" // 假设你的模型在这个路径下
)

func main() {
    // 初始化数据库连接
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }

    // 创建新用户记录
    user := models.UserBasic{
        Username: "johndoe",
        Email:    "johndoe@example.com",
        Password: "securepassword",
    }
    result := db.Create(&user)
    if result.Error != nil {
        log.Fatal("failed to create user: ", result.Error)
    }

    log.Println("User ID:", user.ID)
}
```

### 注意事项

- `db.Create()` 会立即插入记录并保存更改。
- 如果插入记录时发生错误，例如违反唯一约束，`result.Error` 将包含相关错误信息。

----

## `db.Model()`

### 作用
`db.Model()` 方法用于指定查询或更新操作的模型类型。在进行链式调用时，它通常用于设置操作的目标模型。

### 示例代码

以下示例展示了如何使用 `db.Model()` 来执行更新操作：

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "myapp/models" // 假设你的模型在这个路径下
)

func main() {
    // 初始化数据库连接
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }

    // 更新用户记录
    result := db.Model(&models.UserBasic{}).Where("username = ?", "johndoe").Update("email", "newemail@example.com")
    if result.Error != nil {
        log.Fatal("failed to update user: ", result.Error)
    }

    log.Println("Rows affected:", result.RowsAffected)
}
```

### 注意事项

- `db.Model()` 通常与查询、更新、删除等操作链式调用，以指定操作的目标模型。
- 它不会立即执行操作，直到链式调用中的查询或更新方法被调用。

---

## `db.Delete()`

### 作用
`db.Delete()` 方法用于删除数据库中的记录。可以根据结构体实例、主键值或指定条件删除记录。

### 示例代码

假设你有一个名为 `UserBasic` 的结构体，如下所示：

```go
package models

type UserBasic struct {
    ID        uint   `gorm:"primaryKey"`
    Username  string `gorm:"unique;not null"`
    Email     string `gorm:"unique;not null"`
    Password  string `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

你可以通过以下代码删除一个用户记录：

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "myapp/models" // 假设你的模型在这个路径下
)

func main() {
    // 初始化数据库连接
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }

    // 删除用户记录，假设用户 ID 为 1
    result := db.Delete(&models.UserBasic{}, 1)
    if result.Error != nil {
        log.Fatal("failed to delete user: ", result.Error)
    }

    log.Println("Rows affected:", result.RowsAffected)
}
```

### 注意事项

- `db.Delete()` 会删除匹配条件的记录。
- 可以通过传入结构体实例、主键值或条件来删除记录。

---

## `db.First()`

### 作用
`db.First()` 方法用于从数据库中查询第一条匹配的记录。通常根据主键进行排序并返回第一条记录。

### 示例代码

以下示例展示了如何使用 `db.First()` 来查询第一条用户记录：

```go
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "myapp/models" // 假设你的模型在这个路径下
)

func main() {
    // 初始化数据库连接
    dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to database: ", err)
    }

    // 查询第一条用户记录
    var user models.UserBasic
    result := db.First(&user)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            log.Println("No user found")
        } else {
            log.Fatal("failed to query user: ", result.Error)
        }
    }

    log.Println("User:", user)
}
```

### 注意事项

- `db.First()` 会根据主键排序并返回第一条记录。
- 如果没有找到记录，`result.Error` 将返回 `gorm.ErrRecordNotFound`。
