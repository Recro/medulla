package controllers

import (
	"net/http"

	"github.com/Recro/medulla/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/ompluscator/dynamic-struct"
	"gorm.io/gorm/schema"
)

// GetUserTables gets the user tables
func GetUserTables(c *gin.Context) {
	var userTables []models.UserTable
	models.DB.Find(&userTables)

	c.JSON(http.StatusOK, userTables)
}

// CreateUserTable creates a user table
func CreateUserTable(c *gin.Context) {
	var input models.UserTable
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&input)

	newTable := dynamicstruct.NewStruct().
		AddField("ID", 0, `json:"int"`).
		Build().
		New()

	var ns = schema.NamingStrategy{}

	models.DB.Table(ns.TableName(input.TableName)).AutoMigrate(newTable)

	c.JSON(http.StatusOK, input)
}
