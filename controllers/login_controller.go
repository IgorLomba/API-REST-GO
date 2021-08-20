package controllers

// func Login(c *gin.Context) {
// 	var login models.Login
// 	err := c.ShouldBindJSON(&login)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "cannot bind JSON: " + err.Error(),
// 		})
// 		return
// 	}

// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": "cannot find user:",
// 		})
// 		return
// 	}

// 	// verify credentials
// 	// encode the recieved password and verify it
// 	if person.Password != services.SHA256Encoder(login.Password) {
// 		c.JSON(401, gin.H{
// 			"error": "invalid credentials",
// 		})
// 		return
// 	}

// 	token, err := services.NewJwtService().GenerateToken(person.ID)

// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	// got the token
// 	c.JSON(200, gin.H{
// 		"token": token,
// 	})

// }
