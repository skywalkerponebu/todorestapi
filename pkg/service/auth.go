package service

import (
	"crypto/sha1"
	"fmt"

	todo "example.com/m/v2"
	"example.com/m/v2/pkg/repository"
)

const salt = "fdfsqe432hkapfap"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = geteranePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func geteranePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
