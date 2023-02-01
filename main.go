package main

import (
	"fmt"
	"goxamples/numutils"
	"math/rand"
)

// go run main.go
// go test -v ./...
func main() {
	fmt.Println("Hello, World!")

	n := rand.Int()
	fmt.Println("IsPrime: ", numutils.IsPrime(n))
}
