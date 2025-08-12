package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/heisenxxd/crud-go/crud-go/src/model/service"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService,) UserControllerInterface {
       return &userControllerInterface{
		service: serviceInterface,
	   }
}

type UserControllerInterface interface {
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindUserById(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}