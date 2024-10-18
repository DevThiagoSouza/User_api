package routes

import (
	controller "modulo/src/controller/rest"

	"github.com/gin-gonic/gin"
)

func Routes(rg *gin.RouterGroup) {
	rg.GET("/userId:userId", controller.UserId)
	rg.GET("/userEmail:userEmail", controller.UserEmail)
	rg.POST("/createUser", controller.CreateUser)
	rg.PUT("/updateUser:userId", controller.UpdateUser)
	rg.DELETE("/:userId", controller.DeleteUser)
}
