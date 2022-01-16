package model

import "log"

// Model build assets.
type Model interface {
// User model formatter.
GetUser(id int64) (User, error)
}



type model struct {}


// add database connection logic to this function. pass the connection creds as args.
func NewModel() Model {
	return &model{}
}

// Example model functions
// Replace with real model.
// Don't forget to add any functions to the interface!

// this example fetches a user from a database,
// but this structure should work for almost any microservice use case.

type User struct{
	Id int64
	Name string
	Role string
}

func (r model) GetUser(id int64) (User, error){

	user, err := r.mockDB(id)
	if err != nil {
		log.Panic("error querying mock database")
	}

	return user, nil
}


// mock database connection, emulates a id query. Replace with your favorite SQL or NoSQL driver.
func (r model) mockDB(id int64) (User, error){

	// static data to emulate database.
var users = []User{
	{
		Id: 123,
		Name: "test 1",
		Role: "tester",
	},
	{
		Id: 456,
		Name: "test 2",
		Role: "admin",
	},
}

	var user = User{}

	// fetch static data for this example, replace with live db connection here.
	for _, v := range users {
		if id == v.Id {
			user = v
			break
		}
	}

	return user, nil

}