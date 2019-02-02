package main

import (
	"fmt"

	"github.com/nolanprewit1/employee_directory_api/database"

	"github.com/gin-gonic/gin"
)

// CREATE
func postContact(c *gin.Context) {
	var contact Directory
	c.BindJSON(&contact)
	if err := database.DbConnection.Create(&contact).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, "error")
	} else {
		c.JSON(200, contact)
	}
}

// READ
func getContacts(c *gin.Context) {
	var contacts []Directory
	if err := database.DbConnection.Find(&contacts).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, "error")
	} else {
		c.JSON(200, contacts)
	}
}

// READ
func getContact(c *gin.Context) {
	id := c.Params.ByName("id")
	var contact Directory
	if err := database.DbConnection.Where("id = ?", id).First(&contact).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, "error")
	} else {
		c.JSON(200, contact)
	}
}

// UPDATE
func updateContact(c *gin.Context) {
	id := c.Params.ByName("id")
	var contact Directory
	if err := database.DbConnection.Where("id = ?", id).First(&contact).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, "error")
	} else {
		c.BindJSON(&contact)
		database.DbConnection.Save(&contact)
		c.JSON(200, contact)
	}
}

// DELETE
func deleteContact(c *gin.Context) {
	id := c.Params.ByName("id")
	var contact Directory
	if err := database.DbConnection.Where("id = ?", id).Delete(&contact).Error; err != nil {
		fmt.Println(err)
		c.JSON(404, "error")
	} else {
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	}
}
