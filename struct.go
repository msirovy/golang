/*
   Play with struct and new types

   Author: msirovy@gmail.com
*/

package main

import (
	"fmt"
	"log"
)

// define new type
type Datum string

// define new struc
type User struct {
	Name    string
	Surname string
	Age     uint8
}

func main() {
	log.Println("Program started...")
	// define date of type Datum
	var date Datum = "2020-11-02"

	// define slice of type User
	var users []User

	users = append(users, User{"Karel", "Patočka", 42})
	users = append(users, User{"Vilda", "Novak", 34})

	fmt.Println(date)

	if len(users) > 1 {
		for _, usr := range users {
			fmt.Printf("Name: %s\nSurname: %s\nAge: %d\n\n", usr.Name, usr.Surname, usr.Age)
		}
	} else {
		// Println is super cool, could print whole array or struct
		fmt.Println(users)
	}

}
