package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yordil/mereb-tech-challenge/internal/dtos"
)

// HandleResponse handles the API response based on the type of the response object.
func HandleResponse(c *gin.Context, response interface{}) {

	// print the type of response
	fmt.Printf("Type of response: %T\n", response)
	switch res := response.(type) {
	case dtos.AllUserResponse:
		c.JSON(http.StatusOK, res)
	// case dtos.:
		c.JSON(http.StatusOK, res)
	case dtos.ErrorResponse:
		c.JSON(res.Status, res)
	case dtos.SuccessResponse:
		c.JSON(http.StatusOK, res)
	case dtos.UserResponse:
		c.JSON(http.StatusOK, res)
	default:
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{Message: "Internal Server Error", Status: 500})
	}
}
