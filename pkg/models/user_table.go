package models

type UserTable struct {
	ID uint
	TableName string `json:"tableName" binding:"required"`
}
