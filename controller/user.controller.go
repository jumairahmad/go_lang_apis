package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"restapi.com/my-apis/models"
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

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserServices.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserServices.GetUsers(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserServices.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)

}

func (uc *UserController) UpdateUser(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserServices.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})

}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserServices.DeleteUser(&username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")

	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUsers)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:name", uc.DeleteUser)

}
