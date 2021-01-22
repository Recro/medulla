package models

// UserTable represents a user table
type UserTable struct {
	ID        uint
	TableName string `json:"tableName" binding:"required"`
}
