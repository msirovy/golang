/*
	request from testing API looks like this
	{
		"message": "success",
		"people": [
			{"name": "Mark Vande Hei", "craft": "ISS"},
			{"name": "Pyotr Dubrov", "craft": "ISS"},
			{"name": "Anton Shkaplerov", "craft": "ISS"},
			{"name": "Zhai Zhigang", "craft": "Shenzhou 13"},
			{"name": "Wang Yaping", "craft": "Shenzhou 13"},
			{"name": "Ye Guangfu", "craft": "Shenzhou 13"},
			{"name": "Raja Chari", "craft": "ISS"},
			{"name": "Tom Marshburn", "craft": "ISS"},
			{"name": "Kayla Barron", "craft": "ISS"},
			{"name": "Matthias Maurer", "craft": "ISS"}
		],
		"number": 10
	}


*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// sub-struct for array of people
type people struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

// struct for unmarshal response
type response struct {
	Number  int      `json:"number"`
	Peoples []people `json:"people"`
}

func main() {
	const url string = "http://api.open-notify.org/astros.json"
	fmt.Printf("Establish connection to %s.\n", url)

	// GET request to the given url
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Connection error\n", err)
	}
	fmt.Printf("Response status: %s\n", r.Status)

	// Read response and convert it to readable form
	raw_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Response parse error\n", err)
	}

	// Print readable response (unformated)
	fmt.Println("Response raw data:\n", string(raw_body))


    /*
        Parse response to predefined structure using unmarshal function

    */
    fmt.Println("\nResponse formated data:\n")
    // define variable of defined struct response
	response_formated := response{}

    // Parse raw_body to response_formated
	jsonErr := json.Unmarshal(raw_body, &response_formated)
	if jsonErr != nil {
		fmt.Println("json parse error: ", jsonErr)
	}

    // Print output using parsed data
	fmt.Println(response_formated.Number)
	for _, v := range response_formated.Peoples {
		fmt.Printf("---\n - %s\n - %s\n", v.Name, v.Craft)
	}
}

