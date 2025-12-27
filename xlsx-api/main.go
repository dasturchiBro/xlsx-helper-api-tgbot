package main

import (
	"github.com/gin-gonic/gin"
	"xlsx/handlers"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "the platform is on"})
}

func main() {
	r := gin.Default()
	r.Static("/downloads", "./downloads")
	r.GET("/", )
	v1 := r.Group("/v1")
	{
		v1.POST("/xlsx", handlers.XLSXHandler)
	}

	r.Run()
}

//curl -X POST http://127.0.0.1:8080/v1/xlsx -H "Content-Type: application/json" -d "{\"header\":[\"11\",\"4\",\"1-BSB\"],\"criteria\":[\"Writing 10\",\"Grammar 5\", \"Vocabulary 5\"],\"students\":[[\"Fayozbek\"[1,2,3]],[\"Anton\"[31,32,23]]]}"