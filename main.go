package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type register struct{
	ID string `json:"id"`
	Name string `json:"name"`
}
var names = []register{
	{ID: "001", Name: "Ravi"},
	{ID: "002", Name: "Chetan"},
	{ID: "003", Name: "Sarthak"},
	{ID: "004", Name: "Tushar"},
}

func main(){
	router := gin.Default()
	router.GET("/items", getItems)
	router.GET("/items/:id", getItemByID)
	router.POST("/items", postItems)
	router.Run("localhost:8080")
}

func getItems(c *gin.Context){
	c.IndentedJSON(http.StatusOK, names)
}

func postItems(c *gin.Context){
	var newItem register
	if err := c.BindJSON(&newItem); err != nil {
		return
	}
	names = append(names, newItem)
	c.IndentedJSON(http.StatusCreated, newItem)
}

func getItemByID(c *gin.Context){
	id := c.Param("id")
	for _, a := range names {
		if a.ID ==id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
}
