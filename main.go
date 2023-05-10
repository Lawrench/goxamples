package main

import (
	"fmt"
	"math/rand"

	"goxamples/math_examples"
)

func main() {

	// Quick IsPrime demo
	n := rand.Int()
	fmt.Println("IsPrime: ", math_examples.IsPrime(n))
	fmt.Println()

	// Quick demo to find the smalled Clock Angle demo
	t := "3:45"
	angle, err := math_examples.GetAngle(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Least angle for %s is %v.\n", t, angle)

}
