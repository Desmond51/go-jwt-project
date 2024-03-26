package routes

import (
	"github.com/Desmond51/go-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("user/signup", controller.Signup())
	incomingRoutes.POST("user/login", controller.Login())
}