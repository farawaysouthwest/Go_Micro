package main


func main () {

	r := InitializeRouter()

	router := r.GetRouter()
	controller := r.GetController()

	// Routes
	router.GET("/user/:userid", controller.GetUser)

	router.Run(":8080")
}