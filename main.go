package main

import (
	"fmt"

	"main.go/internal/http"
	"main.go/internal/nats"
)

func main() {
	fmt.Println("hi")
	go http.RunFiber()
	fmt.Println("1")
	nats.InitNats()
	fmt.Println("2")
	// nats.Sub()
}
