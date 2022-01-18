package main

func main () {

	r := InitializeRouter()

	// Routes
	r.Router.GET("/user/:userid", r.Controller.GetUser) 

	r.Router.Run(":8080")
}