package service

import (
	"fmt"

	"github.com/heisenxxd/crud-go/crud-go/src/configuration/logger"
	resterr "github.com/heisenxxd/crud-go/crud-go/src/configuration/rest_err"
	"github.com/heisenxxd/crud-go/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomain) *resterr.Resterr {
    logger.Info("Init CreateUser Model", zap.String("Journey", "CreateUser"))

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}