package route

import (
	"github.com/gin-gonic/gin"
	"github.com/yordil/mereb-tech-challenge/internal/controllers"
	"github.com/yordil/mereb-tech-challenge/internal/repository"
	"github.com/yordil/mereb-tech-challenge/internal/usecase"
)

func NewPersonRoute(r *gin.Engine) {

	repo := repository.NewUserRepository()
	usecase := usecase.NewUserUsecase(repo)
	
	personcontroller := controllers.UserController{
		UserUsecase: usecase,
	}
	
	userRoute := r.Group("/person")
	{
		userRoute.POST("/" , personcontroller.CreateUser)
		userRoute.GET("/" , personcontroller.GetAllUser)
		userRoute.GET("/:personId" , personcontroller.GetUserById)
		userRoute.PUT("/:personId"  , personcontroller.UpdateUser)
		userRoute.DELETE("/:personId" , personcontroller.DeleteUser)
	}
	// handle all route except the above should return error

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "The page You are looking for Not Found"})
	} , 

	)





}