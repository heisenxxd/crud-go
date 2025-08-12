package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heisenxxd/crud-go/crud-go/src/configuration/logger"
	"github.com/heisenxxd/crud-go/crud-go/src/configuration/validation"
	"github.com/heisenxxd/crud-go/crud-go/src/model"
	"github.com/heisenxxd/crud-go/crud-go/src/model/request"
	"github.com/heisenxxd/crud-go/crud-go/src/model/service"

	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomain
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "CreateUser"),)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info, error:", err,
			zap.String("journey", "CreateUser"),)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Name,
		userRequest.Password,
		userRequest.Age,
	)
    service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User Created successfully",
	zap.String("journey", "CreateUser"))

	c.String(http.StatusOK, "")
}
