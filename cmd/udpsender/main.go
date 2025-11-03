package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(">")

		lines, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading line: %v", err)
		}

		_, err = conn.Write([]byte(lines))
		if err != nil {
			fmt.Printf("Error writing line: %v", err)
		}
	}

}
