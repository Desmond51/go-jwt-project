package helpers

import (
	controller "go-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)


func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}