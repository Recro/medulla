package models

//Represents a user table field
type UserTableField struct {
	ID uint
	UserTableId uint `json:"userTableId" binding:"required"`
	FieldName string `json:"fieldName" binding:"required"`
	FieldType string `json:"fieldType" binding:"required"`
}
