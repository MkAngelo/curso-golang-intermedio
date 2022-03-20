package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x * 2
	}() // Para llamar a la funcion son los parentesis
	fmt.Println(y)

	// Ejemplo con Goroutines
	c := make(chan int)
	go func() {
		fmt.Println("Starting Function")
		time.Sleep(5 * time.Second)
		fmt.Println("end")
		c <- 1
	}()
	<-c
}
