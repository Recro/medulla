package models

//Represents a user table
type UserTable struct {
	ID uint
	TableName string `json:"tableName" binding:"required"`
}
