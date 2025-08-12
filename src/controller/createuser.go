package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/heisenxxd/crud-go/crud-go/src/configuration/logger"
	"github.com/heisenxxd/crud-go/crud-go/src/configuration/validation"
	"github.com/heisenxxd/crud-go/crud-go/src/model"
	"github.com/heisenxxd/crud-go/crud-go/src/model/request"
	"github.com/heisenxxd/crud-go/crud-go/src/view"

	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}


	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
	i := int8(domain.GetAge())
	ageString := strconv.Itoa(int(i))
	logger.Info("User Created Sucessfully",
    zap.String("journey", "CreateUser"),
    zap.String("name", domain.GetName()),
    zap.String("email", domain.GetEmail()),
    zap.String("age", ageString),
)
}
