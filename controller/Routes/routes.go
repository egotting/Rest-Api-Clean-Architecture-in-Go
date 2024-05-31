package routes

import (
	"github.com/egotting/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/:userid", controller.FindUserById)
	r.DELETE("/user/:userid", controller.DeleteUser)
	r.PUT("/user/:userid", controller.UpdateUser)
	r.POST("/CreateUser", controller.CreateUser)
}
