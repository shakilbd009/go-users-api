package app

import (
	"github.com/shakilbd009/go-users-api/controllers/ping"
	"github.com/shakilbd009/go-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search",.ser)
	router.POST("/users", users.CreateUser)

}
