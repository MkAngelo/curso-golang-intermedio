package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8  // Version explicita
	y := 7 // Version implicita
	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil { // Cachar errores
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}
	m := make(map[string]int) // Llave valor [Map]
	m["Key"] = 6
	fmt.Println(m["Key"])
	s := []int{1, 2, 3} // Slice [Slice]
	for index, value := range s {
		fmt.Println(index, value)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index, value)
	}
	// c := make(chan int) // Canal
	// go doSomething(c)   // Gorutine
	// <-c // Hace que main espere a la Goroutine
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
