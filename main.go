package main

import "narwhal/views"
import "narwhal/database"

func main() {
	database.InitDB()

	// database.CreateUser("admin", "admin")
	// database.CreateUser("test", "test")
	// admin := database.GetUser("name", "admin")
	// admin.CreateProject("Whaler")
	// whaler := database.GetProject(admin, "Whaler")

	views.StartServer()
}
