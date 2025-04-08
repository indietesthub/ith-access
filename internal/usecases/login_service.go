package usecases

import (
	"errors"

	"github.com/indietesthub/ith-access/internal/auth"
	"github.com/indietesthub/ith-access/internal/domain"
	"github.com/indietesthub/ith-access/internal/repository"
)

type loginService struct {
	userRepo repository.UserRepositoryInterface
}

func NewLoginService(userRepo repository.UserRepositoryInterface) domain.LoginService {
	return &loginService{userRepo: userRepo}
}

func (s *loginService) Login(request *domain.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if !auth.CheckPassword(request.Password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
