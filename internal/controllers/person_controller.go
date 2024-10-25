package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yordil/mereb-tech-challenge/internal/domain"
)

type UserController struct {
	UserUsecase domain.UserUseCase
}

func NewUserController(userUsecase domain.UserUseCase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}


func (uc *UserController) CreateUser(c *gin.Context) {
	var user domain.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := uc.UserUsecase.CreateUser(user)
	
	HandleResponse(c, result)
}

// get all user

func (uc *UserController)GetAllUser(c *gin.Context) {

	personID := c.Param("personId")

	if personID != "" {
		fmt.Println("personID is not empty")
	}
	
	result  := uc.UserUsecase.GetAll()

	
	HandleResponse(c, result)
}

func (uc *UserController) GetUserById(c *gin.Context) {
	
	personID := c.Param("personId")

	result := uc.UserUsecase.GetUserById(personID)

	
	HandleResponse(c, result)

}

func (uc *UserController) UpdateUser(c *gin.Context) {
	
	personID := c.Param("personId")	

	var user domain.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := uc.UserUsecase.UpdateUser(personID, user)

	HandleResponse(c, result)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	
	personID := c.Param("personId")
	fmt.Println(personID)

	response := uc.UserUsecase.DeleteUser(personID)

	
	HandleResponse(c, response)
}