package main

import (
	"fmt"
)

func PrettyDiff(a, b map[string]string) {
	/*
		Compare two map[string]string and pretty print results
	*/
	var (
		reset  = "\033[0m"
		red    = "\033[31m"
		green  = "\033[32m"
		yellow = "\033[33m"
		white  = "\033[97m"
	)

	// Iterate second map and detect changes
	for k, v := range b {
		if _, ok := a[k]; !ok {
			fmt.Printf("%s+ %s: %v%s\n", white, k, v, reset)
			delete(a, k)
		} else {
			switch {
			case a[k] == b[k]:
				delete(a, k)
				fmt.Printf("%s= %s: %s%s\n", green, k, v, reset)
			case a[k] != b[k]:
				fmt.Printf("%s~ %s: %s => %s%s\n", yellow, k, a[k], v, reset)
				delete(a, k)
			}
		}
	}

	// iterate uncompared keys in first map
	// and print them as removed in second map
	for k, v := range a {
		fmt.Printf("%s- %s: %s%s\n", red, k, v, reset)
	}
}

func main() {

	labels1 := map[string]string{
		"domain": "deomena.fejk.net",
		"backup": "daily",
		"name":   "Karel",
		"age":    "19",
	}

	labels2 := map[string]string{
		"domain": "deomena.fejk.net",
		"backup": "monthly",
		"type":   "local",
		"name":   "Karel",
	}

	PrettyDiff(labels1, labels2)
}
