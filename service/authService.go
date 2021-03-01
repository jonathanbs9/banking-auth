package service

import (
	"banking-auth/domain"
	"banking-auth/dto"
)

type AuthService interface {
	Login(dto.LoginRequest) (*string, error)
	Verify(urlParams map[string]string) (bool, error)
}

type DefaultService struct {
	repo domain.AuthRepository
	rolePermissions domain.RolePersmissions
}

func (s DefaultService) Login(req dto.LoginRequest) (*string, error) {
	login, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil{
		return nil, err
	}
	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s DefaultService) Verify(urlParams map[string]string) (bool, error){

}
