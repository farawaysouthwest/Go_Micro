package service

import model "go_mvc/model"

// Service build assets
type Service interface {
	FetchUser(id int64) model.User
}

type service struct {
	model model.Model
}

func NewService(model model.Model) Service {
	return service{
		model: model,
	}
}


// Example functions
// Replace with real data sources, such as a database connection.

	// static data.
var users = []model.DataBaseUser{
	{
		Id: 123,
		Name: "test 1",
		Role: "tester",
		LastUpdate: 123,
	},
	{
		Id: 456,
		Name: "test 2",
		Role: "admin",
		LastUpdate: 23455,
	},
}

func (r service) FetchUser(id int64) model.User{



	var user model.DataBaseUser

	// fetch static data for this example, replace with live db connection here.
	for _, v := range users {
		if id == v.Id {
			user = v
			break
		}
	}

	formattedUser := r.model.ModelUser(user)
	return formattedUser
}