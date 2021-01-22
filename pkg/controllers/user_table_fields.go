package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Recro/medulla/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/ompluscator/dynamic-struct"
	"gorm.io/gorm"
)

// GetUserTableFields gets the user table fields
func GetUserTableFields(c *gin.Context) {
	var userTableFields []models.UserTableField
	models.DB.Find(&userTableFields)

	c.JSON(http.StatusOK, userTableFields)
}

// CreateUserTableField creates a user table field
func CreateUserTableField(c *gin.Context) {
	var input models.UserTableField
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userTable models.UserTable
	err := models.DB.First(&userTable, input.UserTableId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&input)

	rows, err := models.DB.Model(&models.UserTableField{}).Where("user_table_id = ?", strconv.Itoa(int(userTable.ID))).Rows()
	defer rows.Close()

	tableUpdate := dynamicstruct.NewStruct().
		AddField("ID", 0, `json:"int"`)

	for rows.Next() {
		var userTableField models.UserTableField
		models.DB.ScanRows(rows, &userTableField)

		tableUpdate.AddField(userTableField.FieldName, 0, `json:"`+userTableField.FieldType+`"`)
	}

	models.DB.Table(userTable.TableName).AutoMigrate(tableUpdate.Build().New())

	c.JSON(http.StatusOK, input)
}
