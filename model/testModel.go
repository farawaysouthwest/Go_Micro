package model

// Model build assets.
type Model interface {
// User model formatter.
ModelUser(dbRecord DataBaseUser) User
}

type model struct {}

func NewModel() Model {
	return model{}
}

// Example model

type User struct{
	Id int64
	Name string
	Role string
}

type DataBaseUser struct {
	Id int64
	Name string
	Role string
	LastUpdate int64
}

func (r model) ModelUser(dbUser DataBaseUser) User{
	return User{
		Id: dbUser.Id,
		Name: dbUser.Name,
		Role: dbUser.Role,

	}
}