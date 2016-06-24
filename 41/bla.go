package main

import "fmt"

func main() {
	var x [3]int= [3]int{1, 2, 3}
	y := [3]int{1, 2, 3}
	z := [...]int{1, 2, 3}

	for i, v := range x {
		fmt.Printf("%d: %d\n", i, v)
	}

	for i, v := range y {
		fmt.Printf("%d: %d\n", i, v)
	}

	for i, v := range z {
		fmt.Printf("%d: %d\n", i, v)
	}

	fmt.Printf("length x =  %d", len(x))

	fmt.Printf("%d", z[2:])

	var s []int
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\n", cap(s))
		s = append(s, 1)
	}




}
