package routes

import (
	controllers "github.com/Desmond51/go-jwt-project/controllers"
	"github.com/Desmond51/go-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUser())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())

}