package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type person struct{
	XMLName xml.Name `xml:"person"`
	FirstName   string `xml:"firstName,attr"`
	LastName    string `xml:"lastName,attr"`

}

func IndexHandler(c *gin.Context) {
	c.XML(200, person{
		FirstName: "ali",
		LastName: "tehrani",
	})
}

func main() {
	router := gin.Default()

	router.GET("/", IndexHandler)

	
	router.Run(":8085")

}