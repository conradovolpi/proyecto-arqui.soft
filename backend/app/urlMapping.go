package app

import (
	user "backend/controllers/user"
)

func mapUrls() {
	// User routes
	router.POST("/signup", user.Signup)
	router.POST("/login", user.Login)
}
