GOlang
======


Good to know
------------
 - Source codes ending with \*\_test.go should be tests.



CLI notes
---------

```bash

go run source_code.go

# offline documentation for builtins, etc
go doc builtin	
go doc fmt.Println


# code normalization is there implemented
gofmt source_code.go > public_this.go
```
		

Types
-----


* **int**		- integer se znamenkem (could define 8 - 64)
* **uint** 	- integer bez znamenka (could define 8 - 64)
* **float32/64**		
* **complex32/64**	


```golang
package main

import "fmt"	

func main() {
	// declare a
	var a int = 109

	// change type and seve to b
	var b uint8 = a

	// > 109
	fmt.Printf("> %d\n", b)

	// >   109        add two spaces before the number
	fmt.Printf("> %5d\n", b)

	// > 00109				add two zeros before the number
	fmt.Printf("> %05d\n", b)

	// > 1101101			print it as a binary
	fmt.Printf("> %b\n", b)
}

```

* **string**   - array of unicode chars (allows all utf8 chars)


```golang
var name string = "Karel"
fmt.Printf("%s", name)

// string is array of unicode bits
fmt.Println(name[3])	// prints 101
fmt.Println(string(x[3]))	// prints e

```


* **array**  - the length and type must be defined during initialization

```golang
package main

import "fmt"

func main() {
  var arr1 [7]string

  arr1[0] = "go"
  arr1[2] = "is"
  arr1[4] = "the"
  arr1[6] = "best"

  for id, word := range arr1 {
    fmt.Println(id, " - ", word)
  }

}
```


* **slice** 	- let's say more flexible array (more [informations](https://blog.golang.org/slices-intro))
```golang
slc := []string{
	"go",
	"is",
	"the",
	"best"}

// Println is cool and can print whole array or slice
fmt.Println(slc)

```

* **map**  		- hashtable

```golang
users := make(map[string]User)

users["kaja"] = User{"Karel", "Patočka", 42}
users["vilda"] = User{"Viliem", "Novak", 34}
``` 


* **Own types and structures**
  golang allows to define own structures and types and is very strictly (it increases safety)

```golang
package main

// define new type
type Datum string

// define new struc
type User struct {
	Name    string
	Surname string
	Age     uint8
}

func main() {
	// define date of type Datum
	var date Datum = "2020-11-02"

	// define slice of type User
	var users []User

	users = append(users, User{"Karel", "Patočka", 42})
	users = append(users, User{"Vilda", "Novak", 34})
	
	fmt.Println(date)
	fmt.Println(users)

}

```

Basic program structure
-----------------------

This is example how the nice program could work

```golang
package main

// need to use everything imported
import (
	"fmt"
	"log"
)

func sayHello(message string) string {
	var err bool = false
	if len(message) == 0 {
		err = true
	}
	return message, err
}


func main() {
	// it is good to inform users that your program have started correctly
	log.Println("Program starting")

	msg, err := sayHello("Minimalism is a sexy!!!")	
	if err == nil {
		fmt.Println(msg)
	} else {
		log.Error("Something went wrong :-( ")
	}

}

```
