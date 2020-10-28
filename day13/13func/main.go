package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	// x := 5
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	// x := 5
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {

	fmt.Println(f1())

	fmt.Println("--------->")
	fmt.Println(f2())

	fmt.Println("--------->")
	fmt.Println(f3())

	fmt.Println("--------->")
	fmt.Println(f4())
}