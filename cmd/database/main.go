package main

import (
	"github.com/Recro/medulla/pkg/controllers"
	"github.com/Recro/medulla/pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	models.Connect()

	app.GET("/user_tables", controllers.GetUserTables)
	app.POST("/user_tables", controllers.CreateUserTable)

	app.GET("/user_table_fields", controllers.GetUserTableFields)
	app.POST("/user_table_fields", controllers.CreateUserTableField)

	app.Run()
}
