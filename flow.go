package main

import (
	"fmt"
)

// stupid example but it works
func sum(x, y int) (sum int, err error){
	if x == 0 && y == 0 {
		return 0, fmt.Errorf("Numbers shouldn't be zeros")
	} else {
		return x + y, nil
	}
}

func main() {
	/* 
		I can allocate local variables in if, 
		but sum variable will not be available out of if
	*/
	if sum, err := sum(10, 21); err == nil {
		fmt.Println(sum)
	}

	// define empty var
	var age uint

	fmt.Println("Give me your age")
	// read input to age variable
	fmt.Scanln(&age)
	fmt.Println(age)
	switch {
	case age < 10:
		fmt.Println("You are children")
	case 10 < age && age < 18:
		fmt.Println("You are adolescent")
	case 18 < age:
		fmt.Println("You are adult")
	}

}
