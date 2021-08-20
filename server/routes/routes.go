package routes

import (
	"github.com/IgorLomba/API-REST-GO/controllers"
	"github.com/IgorLomba/API-REST-GO/server/middlewares"
	"github.com/gin-gonic/gin"
)

// By: DARTHxIKE

func LoadRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		main.GET("/", controllers.Home)
		person := main.Group("person", middlewares.Auth())
		{
			person.GET("/", controllers.GetAllPerson)
			person.GET("/:id", controllers.GetPersonId)
			person.GET("/name/:id", controllers.GetPersonName)
			person.GET("/address/:id", controllers.GetPersonAddress)
			person.DELETE("/:id", controllers.DeletePersonID)
		}
		create := main.Group("create")
		{
			create.POST("/", controllers.CreatePerson)
			create.PUT("/", controllers.UpdatePerson)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
	}
	return router
}
