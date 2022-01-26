package service

import model "go_micro/model"

type MockDB interface {
	
	User(id int64) (model.User, error)
}

type mockDB struct {}


func NewMockDB() MockDB {
	return &mockDB{}
}

// mock database connection, emulates a id query. Replace with your favorite SQL or NoSQL driver.
func (r mockDB) User(id int64) (model.User, error){

	// static data to emulate database.
var users = []model.User{
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

	var user = model.User{}

	// fetch static data for this example, replace with live db connection here.
	for _, v := range users {
		if id == v.Id {
			user = v
			break
		}
	}

	return user, nil

}
