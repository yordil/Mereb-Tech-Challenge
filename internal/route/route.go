package route

import "github.com/gin-gonic/gin"

func SetupAllRoutes(r *gin.Engine) {
	
	NewPersonRoute(r)
}