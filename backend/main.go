package main

import (
    "crud_go/controller"
    "crud_go/util"

    "github.com/gin-gonic/gin"
)

func main() {
    server := controller.Server{
  		DB: util.InitDB(),
  	}
    router := gin.Default()
    router.GET("/", server.GetStorePage)

  	router.Run(":8080")
}
