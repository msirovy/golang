package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	srv, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Panicln("FATAL:", err)
	}
	defer srv.Close()
	log.Println("Start: Listen on TCP port :9999")

	for {
		ch, err := srv.Accept()
		if err != nil {
			fmt.Println("ERR: ", err)
		}
		io.WriteString(ch, "Welcome there!\n##################\n\nTell me something interesting\n>>> ")

		scanner := bufio.NewScanner(ch)
		for scanner.Scan() {
			ln := scanner.Text()
			log.Println(ln)

			switch ln {
			case "quit":
				log.Println("Exiting, user wants it...")
				os.Exit(0)

			case "ahoj":
				io.WriteString(ch, "No nazdar...\n\n>>> ")

			default:
				io.WriteString(ch, "\n>>> ")
			}

		}
	}

}
