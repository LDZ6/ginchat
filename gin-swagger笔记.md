
---

# 使用 `gin-swagger` 生成 API 文档

`gin-swagger` 是一个用于 Go 语言的 `gin` 框架生成 Swagger 文档的库。Swagger 是一种用于描述和测试 RESTful APIs 的规范。使用 `gin-swagger` 可以自动生成 API 文档，方便开发者和用户查看和测试 API。

## 1. 安装依赖

首先，需要安装 `gin` 和 `gin-swagger` 及其依赖：

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/swaggo/swag/cmd/swag
```

## 2. 初始化 Gin 项目

创建一个新的 Go 项目并初始化 Gin：

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run() // listen and serve on 0.0.0.0:8080
}
```

## 3. 使用 Swagger 注解

在你的 handler 函数上方添加 Swagger 注解，以便生成 API 文档。例如：

```go
package main

import (
    "github.com/gin-gonic/gin"
    _ "your_project/docs" // 如果你的项目目录是 `your_project`
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @Summary Ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Swagger 页面
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run() // listen and serve on 0.0.0.0:8080
}
```

## 4. 生成 Swagger 文档

在你的项目根目录运行以下命令生成 Swagger 文档：

```bash
swag init
```

这会在你的项目目录下生成 `docs` 文件夹，里面包含生成的 Swagger 文档。

## 5. 运行项目并查看 Swagger 文档

启动你的项目后，打开浏览器访问 `http://localhost:8080/swagger/index.html`，你将看到自动生成的 Swagger API 文档。

## 注解详解

### 文件级注解

这些注解通常放在文件的顶部，用于描述整个 API 文档。

- `@title`：API 文档的标题。
- `@version`：API 的版本。
- `@description`：对 API 的简要描述。
- `@termsOfService`：服务条款的 URL。
- `@contact.name`：联系人姓名。
- `@contact.url`：联系人的 URL。
- `@contact.email`：联系人的电子邮件。
- `@host`：API 服务的主机名。
- `@BasePath`：API 的基础路径。

示例：

```go
// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /
```

### 函数级注解

这些注解用于描述具体的 API 端点。

- `@Summary`：简要说明这个端点的作用。
- `@Schemes`：支持的传输协议（http、https、ws、wss）。
- `@Description`：对这个端点的详细描述。
- `@Tags`：给端点打标签，便于分类。
- `@Accept`：接受的请求类型（如 json、xml）。
- `@Produce`：返回的响应类型（如 json、xml）。
- `@Success`：成功响应的描述，包括 HTTP 状态码、返回的数据类型和描述。
- `@Failure`：失败响应的描述，包括 HTTP 状态码、返回的数据类型和描述。
- `@Param`：请求参数的详细描述，包括名称、位置（query、path、body、header）、数据类型和是否必须。
- `@Router`：定义路由和 HTTP 方法。

示例：

```go
// @Summary Ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func pingHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}
```

### `@Param` 注解详解

`@Param` 注解用于描述请求参数的详细信息。其格式如下：

```
@Param name location dataType isRequired description
```

- `name`：参数名称。
- `location`：参数的位置，可以是 query（查询参数）、path（路径参数）、body（请求体）、header（请求头）。
- `dataType`：参数的数据类型，如 string、integer、boolean 等。
- `isRequired`：参数是否必须，值为 true 或 false。
- `description`：对参数的详细描述。

示例：

```go
// @Param id path int true "Account ID"
```

### `@Success` 和 `@Failure` 注解详解

这些注解用于描述 API 响应的详细信息。

- `@Success` 和 `@Failure` 的格式为：
  ```
  @Success httpStatusCode {responseType} responseModel "description"
  @Failure httpStatusCode {responseType} responseModel "description"
  ```

- `httpStatusCode`：HTTP 状态码，如 200、400、404 等。
- `responseType`：响应的数据类型，如 string、object 等。
- `responseModel`：响应的数据模型，可以是自定义的结构体。
- `description`：对响应的详细描述。

示例：

```go
// @Success 200 {object} Account "Account data"
// @Failure 400 {string} string "Bad request"
```

## 完整示例代码

以下是一个完整的示例代码，展示如何使用 `gin-swagger` 生成和展示 API 文档：

```go
package main

import (
    "github.com/gin-gonic/gin"
    _ "your_project/docs" // 这里需要替换成你的项目目录
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /

// PingExample godoc
// @Summary Ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Swagger 页面
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run() // listen and serve on 0.0.0.0:8080
}
```

---