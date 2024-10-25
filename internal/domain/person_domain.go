package domain

type User struct {
	ID   string    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Hobbies []string `json:"hobbies"`
}


type UserUseCase interface{
	GetAll() interface {}
	CreateUser(user User) interface {}
	UpdateUser(id string, user User) interface {}
	DeleteUser(id string) interface {}
	GetUserById(id string) interface {}
}

type UserRepository interface{
	GetAll() ([]User, error)
	CreateUser(user User) (error)
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string) error
	GetUserById(id string) (User, error)
}