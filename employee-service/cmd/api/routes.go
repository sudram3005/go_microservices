package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() *gin.Engine {
	//automatically attaches logger and recovery middleware
	//router := gin.Default()

	//Blank gin without any middleware
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"PUT", "DELETE", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	//router.POST("/authenticate", app.Authenticate)
	router.GET("/emps", app.GetAllUser)
	router.POST("/emp", app.AddUser)
	router.DELETE("/emp/{id}", app.DeleteUser)

	return router
}
