package service

import (
	"be/src/domain/repository"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func Login() {

}

func Register() {

}

func Logout() {

}
