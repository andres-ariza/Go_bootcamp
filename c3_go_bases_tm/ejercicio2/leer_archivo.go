package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	// "strconv"
	"strings"
)

func main() {
	producto := "id: 111223, precio: 30012.00, cantidad: 1;\nid: 444321, precio: 1000000.00, cantidad: 4;\nid: 434321, precio: 50.50, cantidad: 1;\n"
	contenido := []byte(producto)
	err := ioutil.WriteFile("./listadoProductos.csv", contenido, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("listadoProductos.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	// var total float64

	fmt.Printf("%-20s%12s%12s\n", "ID", "Precio", "Cantidad")
	for _, record := range records {
		id := strings.Split(record[0], ":")[1]
		id = strings.TrimSpace(id)
		precio := strings.Split(record[1], ":")[1]
		precio = strings.TrimSpace(precio)
		cantidad := strings.Split(record[2], ":")[1]
		cantidad = strings.TrimSuffix(cantidad, ";")

		fmt.Printf("%-20s%12s%12s\n", id, precio, cantidad)
		// numero, _ = strconv.ParseFloat(precio, 64)
		// total += float64(numero)
		// fmt.Printf("%-32d\n", total)
	}
}
