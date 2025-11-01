package main

import (
	"fmt"
)

func main() {
	c := make(chan string) // creates channel++
	go count("sheep", c)   // concurrent call with chanel as par

	for msg := range c { // Itera hasta que msg cierre
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
	}
	close(c) // closes channer it knows it wont recive more
}
