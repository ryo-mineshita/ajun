package main
import (
    "reactgo/controller"
    "reactgo/util"
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

    router.GET("/w", server.GetFoodPage)
    router.POST("/w", server.CreateFoodHandler)

    router.POST("/user", server.CreateUserHandler)

    router.POST("/usershop", server.UserShopHandler)
    router.Run(":8080")
}
