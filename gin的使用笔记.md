## 创建一个基本的 Gin 应用

```go
package main

import "github.com/gin-gonic/gin"

func main() {
r := gin.Default() // 创建一个默认的 Gin 路由器
r.GET("/ping", func(c *gin.Context) {
c.JSON(200, gin.H{"message": "pong"}) // 定义一个 GET 路由，返回 JSON 响应
})
r.Run() // 启动 HTTP 服务器，默认监听在 0.0.0.0:8080
}
```