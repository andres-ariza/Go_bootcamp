package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type creacion struct {
	Dia int `json:"dia"`
	Mes int `json:"mes"`
	Año int `json:"año"`
}
type producto struct {
	Id        int      `json:"id"`
	Nombre    string   `json:"nombre"`
	Color     string   `json:"color"`
	Precio    float32  `json:"precio"`
	Stock     int      `json:"stock"`
	Codigo    string   `json:"codigo"`
	Publicado bool     `json:"publicado"`
	Creacion  creacion `json:"creacion"`
}

func HandlerGetName(c *gin.Context) {
	name := c.Param("nombre")
	c.JSON(200, gin.H{
		"message": "Holis " + name,
	})
}

func HandlerGetProductos(c *gin.Context) {
	jsonData, err := os.ReadFile("productos.json")
	if err != nil {
		log.Fatal(err)
	}
	// se crea el puntero para unmarshall
	var p []producto
	err = json.Unmarshal(jsonData, &p)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, p)
}

func HandlerPostProductos(c *gin.Context) {
	//ejemplo post
	body := c.Request.Body
	header := c.Request.Header
	metodo := c.Request.Method

	fmt.Println("¡He recibido algo!")
	fmt.Printf("\tMetodo: %s\n", metodo)
	fmt.Printf("\tContenido del header:\n")

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}
	fmt.Printf("\tEl body es un io.ReadCloser:(%s), y para trabajar con el vamos a tener que leerlo luego\n", body)
	c.String(200, "¡Lo recibí!") //Respondemos al cliente con 200 OK y un mensaje.
}
func HandlerGetProducto(c *gin.Context) {
	jsonData, err := os.ReadFile("productos.json")
	if err != nil {
		log.Fatal(err)
	}
	var p []producto
	err = json.Unmarshal(jsonData, &p)
	if err != nil {
		log.Fatal(err)
	}
	key, err := strconv.Atoi(c.Param("key"))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, p[key])
}

func main() {
	// ej2 - c1-m
	r := gin.Default()
	r.GET("/:nombre", HandlerGetName)
	r.GET("/productos", HandlerGetProductos)
	//c1 tarde ejemplo 1

	r.POST("/productos/", HandlerPostProductos)

	r.GET("productos/:producto", HandlerGetProducto)
	r.GET("productos/:id", func(c *gin.Context) {
		jsonData, err1 := os.ReadFile("productos.json")
		if err1 != nil {
			log.Fatal(err1)
		}
		var p []producto
		err2 := json.Unmarshal(jsonData, &p)
		if err2 != nil {
			log.Fatal(err2)
		}
		key, err3 := strconv.Atoi(c.Param("key"))
		if err3 != nil {
			log.Fatal(err3)
		}
		c.JSON(http.StatusOK, p[key])
	})

	r.Run()
}
