package main

import (
	"fmt"

	"main.go/internal/http"
)

func main() {
	fmt.Println("hi")
	http.RunFiber()
}
