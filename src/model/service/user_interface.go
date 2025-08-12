package service

import (
	resterr "github.com/heisenxxd/crud-go/crud-go/src/configuration/rest_err"
	"github.com/heisenxxd/crud-go/crud-go/src/model"
)
func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
	model.UserDomainInterface
}


type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *resterr.Resterr
	DeleteUser(string) *resterr.Resterr
	UpdateUser(string, model.UserDomainInterface) *resterr.Resterr
	FindUser(string) (*model.UserDomainInterface, *resterr.Resterr)
}