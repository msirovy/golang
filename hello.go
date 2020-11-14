package main

// need to use everything imported
import (
	"fmt"
	"log"
)

func sayHello(message string) (msg string, err error) {
	// validate input length
	if len(message) == 0 {
		// return empty string and error message
		return "", fmt.Errorf("Empty input given")
	}
	return message, nil
}

func main() {
	// it is good to inform users that your program have started correctly
	log.Println("Program starting")

	// Declare array of messages
	var test_msg [3]string
	test_msg[0] = "Minimalism is a sexy"
	// ship message 1 to test error handling
	test_msg[2] = "Hello, world"

	// iterate all messages
	for id, m := range test_msg {

		// return response and err status to validate
		msg, err := sayHello(m)

		// if err is nil, we can print output
		if err == nil {
			fmt.Println(id, " : ", msg)

			// else we log fatal error
		} else {
			log.Fatal(id, " - Something went wrong :-( ")
		}

	} // iterate message end

}
