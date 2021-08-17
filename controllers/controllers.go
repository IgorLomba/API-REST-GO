package controllers

import (
	"net/http"

	"github.com/IgorLomba/API-REST-GO/API-REST-GO/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, "API DARTHxIKE")
}

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

// get person by name (like)
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

// get person by address (like)
func GetPersonAddress(c *gin.Context) {
	id := c.Param("id")
	// log.Println("IDEZERAA", id)
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

/* func getPerson_Old(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.FormValue("id") != "" {
		id := r.FormValue("id")
		// log.Println("IDDDDDDDD", id)
		person, _ := models.LoadPersonByID(id)
		if person.ID == 0 {
			json.NewEncoder(w).Encode("ID NOT FOUND")
		} else {
			json.NewEncoder(w).Encode(person)
		}
	}

	if r.FormValue("name") != "" {
		name := r.FormValue("name")
		// log.Println("Name", name)
		person, _ := models.LoadPersonByName(name)
		if len(person) == 0 {
			json.NewEncoder(w).Encode("NAME NOT FOUND")
		} else {
			json.NewEncoder(w).Encode(person)
		}
	}

	if r.FormValue("address") != "" {
		address := r.FormValue("address")
		// log.Println("Address", address)
		person, _ := models.LoadPersonByAddress(address)
		if len(person) == 0 {
			json.NewEncoder(w).Encode("ADDRESS NOT FOUND")
		} else {
			json.NewEncoder(w).Encode(person)
		}
	}

}
*/

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
	// db := db.ConnectDb()
	var person models.Person
	var address models.Address
	err := c.ShouldBindJSON(&person)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	person, err = models.UpdatePerson(person, address)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind update: " + err.Error(),
		})
		return
	}
	c.JSON(200, person)
}

// get person by a ID
func DeletePersonID(c *gin.Context) {
	id := c.Param("id")
	// log.Println("IDEZERAA", id)
	err := models.DeletePersonById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete person: " + id,
		})
		return
	}
	c.Status(204)

}
