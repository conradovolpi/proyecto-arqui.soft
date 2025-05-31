// para el login 
r.POST("/login", func(c *gin.Context) {
	controllers.LoginHandler(c, db)
})
