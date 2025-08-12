package service

import (
	resterr "github.com/heisenxxd/crud-go/crud-go/src/configuration/rest_err"
	"github.com/heisenxxd/crud-go/crud-go/src/model"
)
func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
	model.UserDomain
}


type UserDomainService interface {
	CreateUser(model.UserDomain) *resterr.Resterr
	DeleteUser(string) *resterr.Resterr
	UpdateUser(string, model.UserDomain) *resterr.Resterr
	FindUser(string) (*model.UserDomain, *resterr.Resterr)
}