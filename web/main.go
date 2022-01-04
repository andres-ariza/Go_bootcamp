package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type creacion struct {
	Dia int
	Mes int
	Año int
}
type producto struct {
	Id        int
	Nombre    string
	Color     string
	Precio    float32
	Stock     int
	Codigo    string
	Publicado bool
	Creacion  creacion
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Andres",
		})
	})
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hola %s sos un crack", name)
	// })
	r.GET("/productos/GetAll", func(c *gin.Context) {
		// leo el archivo json
		jsonData, err1 := os.ReadFile("productos.json")
		if err1 != nil {
			log.Fatal(err1)
		}
		// se crea el puntero para unmarshall
		var p []producto
		err2 := json.Unmarshal(jsonData, &p)
		if err2 != nil {
			log.Fatal(err2)
		}
		c.String(http.StatusOK, "%+v\n", p)
	})
	r.Run()
}
