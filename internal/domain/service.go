package domain

type UserService interface {
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id string) error
}

type LoginService interface {
	Login(request *LoginRequest) (string, error)
}
