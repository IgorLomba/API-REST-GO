package controllers

// By: DARTHxIKE

import (
	"fmt"
	"net/http"

	"github.com/IgorLomba/API-REST-GO/API-REST-GO/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, "API DARTHxIKE")
}

// list all people and addresses
func GetAllPerson(c *gin.Context) {
	person, _ := models.LoadPeople()
	c.JSON(http.StatusOK, person)

}

// get person by a ID
func GetPersonId(c *gin.Context) {
	id := c.Param("id")
	// log.Println("IDEZERAA", id)
	person, err := models.LoadPersonByID(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find person: " + id,
		})
		return
	}
	if person.ID == 0 {
		explain := "person " + id + " does not exist."
		c.JSON(400, gin.H{
			"error": explain,
		})
		return
	}

	c.JSON(200, person)

}

// get person by name
func GetPersonName(c *gin.Context) {
	id := c.Param("id")
	// log.Println("IDEZERAA", id)
	person, err := models.LoadPersonByName(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find person: " + id,
		})
		return
	}
	if len(person) == 0 {
		explain := "person " + id + " does not exist."
		c.JSON(400, gin.H{
			"error": explain,
		})
		return
	}
	c.JSON(200, person)
}

// get person by address
func GetPersonAddress(c *gin.Context) {
	id := c.Param("id")
	person, err := models.LoadPersonByAddress(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find person: " + id,
		})
		return
	}
	if len(person) == 0 {
		explain := "person " + id + " does not exist."
		c.JSON(400, gin.H{
			"error": explain,
		})
		return
	}
	c.JSON(200, person)
}

func CreatePerson(c *gin.Context) {
	var person models.Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	newPerson, err := models.CreatePerson(person)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind create: " + err.Error(),
		})
		return
	}
	c.JSON(200, newPerson)
}

// don't forget to pass person id and address id in PUT json request
func UpdatePerson(c *gin.Context) {
	var person models.Person
	var address models.Address
	err := c.ShouldBindJSON(&person)
	// if person does not exist, return error
	personAux, errAux := models.LoadPersonByID(fmt.Sprint(person.ID))

	if personAux.ID == 0 || personAux.IsDel == 1 {
		c.JSON(400, gin.H{
			"error": "cant update person who does not exist",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	person, err = models.UpdatePerson(person, address)
	if err != nil && errAux != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind update: " + err.Error(),
		})
		return
	}
	c.JSON(200, person)
}

// delet person by ID
func DeletePersonID(c *gin.Context) {
	id := c.Param("id")
	err := models.DeletePersonById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete person: " + id,
		})
		return
	}
	c.Status(204)
}
