package service

import (
	model "go_micro/model"
	"log"
)

// Service build assets
type Service interface {
	GetUser(id int64) (model.User, error)
}

type service struct {
	database MockDB
}

func NewService(db MockDB) Service {
	return &service{
		database: db,
	}
}


// Example service functions
// Don't forget to add any functions to the interface!

func (r service) GetUser(id int64) (model.User, error){

	user, err := r.database.User(id)
	if err != nil {
		log.Panic("error querying mock database")
	}

		formattedUser, err := r.formatUser(user)
	if err != nil {
		log.Panicln("problem formatting user")
	}

	return formattedUser, nil
}




// example of a service layer data manipulation func.
func (r service) formatUser(user model.User) (model.User, error) {
	return model.User{
		Id: user.Id,
		Name: user.Name,
		Role: user.Role,
	}, nil
}