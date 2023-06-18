package main

import (
	"fmt"
	"net/http"
	"time"
)

func CheckServer(server string, c chan<- string) {
	_, err := http.Get(server)
	if err != nil {
		c <- server + " no se encuentra disponible"

		return
	}

	c <- server + " Esta funcionando :D"
}

func main() {
	start := time.Now()

	channel := make(chan string)

	servers := [...]string{
		"http://platzi.com",
		"http://google.com",
		"http://instagram.com",
		"http://facebook.com",
	}

	for _, server := range servers {
		go CheckServer(server, channel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	end := time.Since(start)

	fmt.Println("Ha tomado:", end)
}
