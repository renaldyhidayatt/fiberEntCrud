package services

import (
	"github.com/renaldyhidayatt/fiberEntCrud/ent"
	"github.com/renaldyhidayatt/fiberEntCrud/interfaces"
	"github.com/renaldyhidayatt/fiberEntCrud/repository"
	"github.com/renaldyhidayatt/fiberEntCrud/schemas"
)

type AuthService = interfaces.IAuthService

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository: repository}
}

func (s *authService) Register(input schemas.SchemaUsers) (*ent.Users, error) {
	var userSchema schemas.SchemaUsers

	userSchema.FirstName = input.FirstName
	userSchema.LastName = input.LastName
	userSchema.Email = input.Email
	userSchema.Password = input.Password

	row, err := s.repository.Register(userSchema)

	return row, err

}

func (s *authService) Login(input schemas.SchemaUsers) (*ent.Users, error) {
	var loginSchema schemas.SchemaUsers

	loginSchema.Email = input.Email
	loginSchema.Password = input.Password

	row, err := s.repository.Login(loginSchema)

	return row, err
}
