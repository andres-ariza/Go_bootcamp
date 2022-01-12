package main

import "fmt"

func main() {
	var (
		temperatura string  = "15ºC"
		humedad     float32 = 0.78
		presion     int     = 1022
	)
	fmt.Println("Temperatura: ", temperatura, "\nHumedad: ", humedad, "\nPresion: ", presion)
}
