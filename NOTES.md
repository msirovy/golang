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
gofmt source_code.go
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

=== Strings


```golang
var name string = "Karel"
fmt.Printf("%s", name)

// string is array of unicode bits
fmt.Println(name[3])	// prints 101
fmt.Println(string(x[3]))	// prints e

```


Basic program structure
-----------------------



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
