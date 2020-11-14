/*
   Play with struct and new types

   Author: msirovy@gmail.com
*/

package main

import (
	"fmt"
	"log"
)

// define new struct
type User struct {
	Name    string
	Surname string
	Age     uint8
}

func main() {
	log.Println("Program started...")

	/*
		make preinicialize map where keys are string type 
		and valueas are my own predefined struc User
	*/
	users := make(map[string]User)

	users["kaja"] = User{"Karel", "Patočka", 42}
	users["vilda"] = User{"Viliem", "Novak", 34}


	if len(users) > 1 {
		for nick, usr := range users {
			fmt.Println(nick)
			fmt.Printf("Name: %s\nSurname: %s\nAge: %d\n\n", usr.Name, usr.Surname, usr.Age)
		}
	} else {
		// Println is super cool, could print whole array or struct
		fmt.Println(users)
	}

}
