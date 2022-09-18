package main

import (
	"fmt"

	"main.go/internal/http"
	"main.go/internal/nats"
)

func main() {
	fmt.Println("hi")
	// go
	// defer http.RunFiber()
	fmt.Println("1")
	nats.InitNats()
	fmt.Println("2")
	nats.SubKon()
	http.RunFiber()
	// nats.Sub()
}
