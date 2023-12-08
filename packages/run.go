package main

import (
	"fmt"
	"lib/tools"
)

func main() {
	fmt.Println(tools.GetEnv("X", "NIC!"))
}
