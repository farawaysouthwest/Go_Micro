package service

import (
	model "go_micro/model"
	"log"
)

// Service build assets
type Service interface {
	FetchUser(id int64) model.User
}

type service struct {
	model model.Model
}

func NewService(model model.Model) Service {
	return &service{
		model: model,
	}
}


// Example service functions
// Don't forget to add any functions to the interface!

func (r service) FetchUser(id int64) model.User{

	user, err := r.model.GetUser(id)
	if err != nil {
		log.Panicln("problem getting user")
	}

	formattedUser, err := r.formatUser(user)
	if err != nil {
		log.Panicln("problem formatting user")
	}

	return formattedUser
}

// example of a service layer data manipulation func.
func (r service) formatUser(user model.User) (model.User, error) {
	return model.User{
		Id: user.Id,
		Name: user.Name,
		Role: user.Role,
	}, nil
}