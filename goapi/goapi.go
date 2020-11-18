package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"io/ioutil"
)

// getnev like in Python
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// render static page
func serveFile(w http.ResponseWriter, r *http.Request) {
	log.Println(" Requested url ", r.URL.Path)

	// try to find file localy
	fc, err := ioutil.ReadFile("./public/" + r.URL.Path[1:])

	if err == nil {
		// if requested file has been found, return hic content
		fmt.Fprint(w, string(fc))
	} else {
		// else return 404
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not found!")
		log.Println(" ERR: Not found")
	}
}


// Render dynamic page
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(" Requested url ", r.URL.Path)

	// generate web response
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hi there!!!\nYou are on %s url\n", r.URL.Path[1:])
}

func main() {
	log.Println("Starting app on port ", getEnv("PORT", ":8080"))

	// Routing
	http.HandleFunc("/", serveFile)
	http.HandleFunc("/api", handler)

	log.Fatal(http.ListenAndServe(getEnv("PORT", ":8080"), nil))
}

