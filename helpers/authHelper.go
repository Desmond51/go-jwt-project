package helpers

import (
	controller "github.com/Desmond51/go-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)


func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}