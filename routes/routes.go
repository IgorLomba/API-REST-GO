package routes

import (
	"github.com/IgorLomba/API-REST-GO/API-REST-GO/controllers"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api")
	{
		main.GET("/", controllers.Home)
		person := main.Group("person")
		{
			person.GET("/", controllers.GetAllPerson)
			person.GET("/:id", controllers.GetPersonId)
			person.GET("/name/:id", controllers.GetPersonName)
			person.GET("/address/:id", controllers.GetPersonAddress)
			person.POST("/", controllers.CreatePerson)
			person.PUT("/", controllers.UpdatePerson)
			person.DELETE("/:id", controllers.DeletePersonID)

		}
	}
	return router

}
