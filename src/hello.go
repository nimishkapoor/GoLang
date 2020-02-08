package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello, World!")
	var x int = 5
	y := 7
	var sum int = x + y
	fmt.Println(sum)
	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("Less than 2")
	} else {
		fmt.Println("Between 2 and 6")
	}
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := []int{1, 2, 3, 4, 5}
	c = append(c, 13)
	fmt.Println(c)

	v := make(map[string]int)

	v["triangle"] = 2
	v["square"] = 3
	v["dodecagon"] = 12

	fmt.Println(v)
	fmt.Println(v["triangle"])

	delete(v, "square")

	fmt.Println(v)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	j := 0

	for j < 5 { //while loop
		fmt.Println(j)
		j++
	}

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	for key, value := range v {
		fmt.Println("key", key, "value", value)
	}

	result := sumfun(10, 9)
	fmt.Println(result)

	fmt.Println(sqrt(25))
	fmt.Println(sqrt(-6))

	res, err := sqrt(25)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	p := person{name: "Nimish", age: 22}
	fmt.Println(p)
	fmt.Println(p.age)

	n := 7
	fmt.Println(&n)

	inc(n)
	fmt.Println(n)

	incNew(&n)
	fmt.Println(n)
}

func sumfun(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

func inc(x int) {
	x++
}

func incNew(x *int) {
	*x++
}
