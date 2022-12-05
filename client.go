package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Te rog adauga un host:port!")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		switch strings.TrimSpace(string(text)) {
		case "1":
			message1, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message1)
			message2, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message2)
			reader1 := bufio.NewReader(os.Stdin) //scrii numerele
			fmt.Print(">> ")
			text1, _ := reader1.ReadString('\n')
			fmt.Fprintf(c, text1+"\n")
			message3, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message3)
		case "2":
			message1, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message1)
			message2, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message2)
			reader1 := bufio.NewReader(os.Stdin) //scrii date
			fmt.Print(">> ")
			text1, _ := reader1.ReadString('\n')
			fmt.Fprintf(c, text1+"\n")
			message3, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message3)
		case "3":
			message1, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message1)
			message2, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message2)
			reader1 := bufio.NewReader(os.Stdin) //scrii date
			fmt.Print(">> ")
			text1, _ := reader1.ReadString('\n')
			fmt.Fprintf(c, text1+"\n")
			message3, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message3)
		case "4":
			message1, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message1)
			message2, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message2)
			reader1 := bufio.NewReader(os.Stdin) //scrii date
			fmt.Print(">> ")
			text1, _ := reader1.ReadString('\n')
			fmt.Fprintf(c, text1+"\n")
			message3, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message3)
		case "STOP":
			fmt.Println("Ai iesit de pe server!")
			return
		default:
			message, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message)
		}
	}
}
