package usecase

import (
	"github.com/google/uuid"
	"github.com/yordil/mereb-tech-challenge/internal/domain"
	"github.com/yordil/mereb-tech-challenge/internal/dtos"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

// NewPersonRepository is a function that returns a new instance of PersonRepository
func NewUserUsecase(useRepo domain.UserRepository) domain.UserUseCase {
	return &UserUseCase{
		userRepo: useRepo,
	}
}

// // SavePerson is a function that saves a person to the repository
func (uc *UserUseCase) CreateUser(user domain.User) interface {}{	
	
	if user.Name == "" {
		return dtos.ErrorResponse{
			Status:  400,
			Message: "Name is required",
		}
	}

	if user.Age == 0 {
	return dtos.ErrorResponse{
		Status:  400,
		Message: "Age is required",
	}
}


	// attach random id for the user generate froom uuid googgle
	user.ID = uuid.New().String()

	err := uc.userRepo.CreateUser(user)
	
	if err != nil {
		return dtos.ErrorResponse{
			Status:  500,
			Message: "Internal server error",
		}
	}

	return dtos.UserResponse{
		Status: 200,
		Message: "User created successfully",
		User: user,
	}
}



func (uc *UserUseCase) GetAll() interface {} {

	users , err := uc.userRepo.GetAll()

	if err != nil {
		return dtos.ErrorResponse{
			Status:  500,
			Message: "Internal server error",
		}
	}

	if len(users) == 0 {
		return dtos.ErrorResponse{
			Status:  404,
			Message: "No user found on the system",
		}
	}

	return dtos.AllUserResponse{
		Status: 200,
		Message: "All users",
		Users: users,

	}

}


func (uc *UserUseCase) GetUserById(id string) interface {} {

	user , err := uc.userRepo.GetUserById(id)

	if err != nil {
		return dtos.ErrorResponse{
			Status:  404,
			Message: "User not found",
	}

}

	return dtos.UserResponse{
		Status: 200,
		Message: "User found",
		User: user,
	}

}

func (uc *UserUseCase) UpdateUser(id string, user domain.User) interface {} {
	
	// check if the user exist
	existingUser , err := uc.userRepo.GetUserById(id)

	if err != nil {
		return dtos.ErrorResponse{
			Status:  404,
			Message: "User not found",
		}
	}

	// update the user
	user.ID = existingUser.ID

	// check empty fields and if it is leave as it is else updatge it

	if user.Name == "" {
		user.Name = existingUser.Name
	}

	if user.Age == 0 {
		user.Age = existingUser.Age
	}

	if user.Hobbies == nil {
		user.Hobbies = existingUser.Hobbies
	}

	updatedUser  , err := uc.userRepo.UpdateUser(existingUser.ID , user)

	if err != nil {
		return dtos.ErrorResponse{
			Status:  500,
			Message: "Internal server error",
		}
	}

	return dtos.UserResponse{
		Status: 200,
		Message: "User updated successfully",
		User: updatedUser,

	}
}

func (uc *UserUseCase) DeleteUser(id string) interface{} {

	// check if the user exists

	_ , err := uc.userRepo.GetUserById(id)

	if err != nil {
		return dtos.ErrorResponse{
			Status:  404,
			Message: "User not found",
		}
	}

	err = uc.userRepo.DeleteUser(id)
	
	if err != nil {
		return dtos.ErrorResponse{
			Status:  500,
			Message: "Internal server error",
		}
	}

	return dtos.SuccessResponse{
		Status: 200,	
		Message: "User deleted successfully",
	}

	
}