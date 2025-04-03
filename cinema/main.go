package main

import (
	"fmt"

	"github.com/worachit1/cinema/movie"
	"github.com/worachit1/cinema/ticket"
)

func init() {
	fmt.Println("init: main!")
}

func main() {
	movieName := movie.FindName("tt4323314")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
