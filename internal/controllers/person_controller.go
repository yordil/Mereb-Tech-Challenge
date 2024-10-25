package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yordil/mereb-tech-challenge/internal/domain"
)

// UserController handles user-related requests
type UserController struct {
	UserUsecase domain.UserUseCase
}

// NewUserController creates a new UserController
func NewUserController(userUsecase domain.UserUseCase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Creates a new user in the system
// @Tags Users
// @Accept json
// @Produce json
// @Param user body domain.User true "User data"
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /person [post]
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

// GetAllUser gets all users
// @Summary Get all users
// @Description Returns a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} dtos.AllUserResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /person [get]
func (uc *UserController) GetAllUser(c *gin.Context) {
	result := uc.UserUsecase.GetAll()
	HandleResponse(c, result)
}

// GetUserById gets a user by ID
// @Summary Get user by ID
// @Description Retrieves a user based on ID
// @Tags Users
// @Accept json
// @Produce json
// @Param personId path string true "User ID"
// @Success 200 {object} dtos.UserResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /person/{personId} [get]
func (uc *UserController) GetUserById(c *gin.Context) {
	personID := c.Param("personId")
	result := uc.UserUsecase.GetUserById(personID)
	HandleResponse(c, result)
}

// UpdateUser updates a user by ID
// @Summary Update a user by ID
// @Description Updates the details of an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param personId path string true "User ID"
// @Param user body domain.User true "Updated user data"
// @Success 200 {object} dtos.SuccessResponse
// @Failure 400 {object} dtos.ErrorResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /person/{personId} [put]
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

// DeleteUser deletes a user by ID
// @Summary Delete a user by ID
// @Description Deletes a user from the system based on ID
// @Tags Users
// @Param personId path string true "User ID"
// @Success 200 {object} dtos.SuccessResponse
// @Failure 404 {object} dtos.ErrorResponse
// @Failure 500 {object} dtos.ErrorResponse
// @Router /person/{personId} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	personID := c.Param("personId")
	response := uc.UserUsecase.DeleteUser(personID)
	HandleResponse(c, response)
}
