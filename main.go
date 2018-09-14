package main

import (
	"fmt"
	"math/rand"
	//"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//f, err := os.Open("/dev/urandom")
	//check(err)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Hola Meli Devs ", rand.Intn(11))
}