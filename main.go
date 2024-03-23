package main

import (
	"os"

	"github.com/Desmond51/go-jwt-project/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	port = os.Getenv("PORT")

	if port ==""{
		port = "8000"
	}
	router := gin. New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H("success":"Access granted ap-1"))
	})

	routes.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H("success":"Access granted ap-2"))
	}) 

	router.Run(":" + port)
}	