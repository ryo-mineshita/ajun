package main
import (
    "crud_go/controller"
    "crud_go/util"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)
func main() {
    server := controller.Server{
      DB: util.InitDB(),
    }
    router := gin.Default()
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:3000"}
    router.Use(cors.New(config))
    router.GET("/", server.GetStorePage)
    router.POST("/", server.CreateStoreHandler)

    router.Run(":8080")
}
