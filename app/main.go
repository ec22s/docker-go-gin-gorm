package main

import (
  "go_gin_gorm/controllers"
	"go_gin_gorm/middlewares"
  "go_gin_gorm/models"

  "github.com/gin-gonic/gin"
)

func main() {
  models.ConnectDataBase()
  router := gin.Default()
  public := router.Group("/api")
  public.POST("/register", controllers.Register)
  public.POST("/login", controllers.Login)
  protected := router.Group("/api/admin")
  protected.Use(middlewares.JwtAuthMiddleware())
  protected.GET("/user", controllers.CurrentUser)
  router.Run(":80")
}
