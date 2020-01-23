package main

import (
	"fmt"
	"errors"
	"math"
)

type person struct {
	age int
	name string
}

func main() {
	fmt.Println("Hello, World")

	var x int = 5
	y := 8
	fmt.Println(add(x, y))

	array := []int{0,1,2,3,4,5,6,7}
	array = append(array, 20)

	verticies := make(map[string]int)

	verticies["triangle"] = 2
	verticies["square"] = 3
	verticies["dodecagon"] = 12

	delete(verticies, "triangle")

	fmt.Println(verticies)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	// for i < 5 {

	// }

	for index, value := range array {
		fmt.Println("index", index, "value", value)
	}

	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SquareRoot", result)
	}

	p := person{name: "Jake", age: 23}
	fmt.Println(p.age)

	// Pointers
	i := 7
	inc(&i)
	fmt.Println(i)

}

func inc(x *int) {
	*x++
}

func add(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negitive numbers")
	} 

	return math.Sqrt(x), nil
}