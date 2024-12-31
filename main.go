package main

import (
	"net/http"

	"github.com/Yadier01/notes-backend/db"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	db.InitDB()
	r.GET("/ping", pong)
	r.Run(":8080") // listen and serve on localhost:8080

}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
