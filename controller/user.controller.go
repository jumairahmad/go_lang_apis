package controller

import (
	"github.com/gin-gonic/gin"
	"restapi.com/my-apis/services"
)

type UserController struct {
	UserServices services.UserServices
}

func New(userservice services.UserServices) UserController {
	return UserController{
		UserServices: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {

	ctx.JSON(200, "")
}

func (uc *UserController) GetUsers(ctx *gin.Context) {

	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")

	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUsers)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete", uc.DeleteUser)

}
