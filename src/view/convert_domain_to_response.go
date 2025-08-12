package view

import (
	"github.com/heisenxxd/crud-go/crud-go/src/model"
	"github.com/heisenxxd/crud-go/crud-go/src/model/response"
)

func ConvertDomainToResponse(UserDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{ 
		ID: "",
		Email: UserDomain.GetEmail(),
		Name: UserDomain.GetName(),
		Age: UserDomain.GetAge(),
	}
}