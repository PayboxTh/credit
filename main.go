package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	response, err := ntp.Query("0.asia.pool.ntp.org")
	if err != nil {
		log.Println("Eror:", err)
	}
	time := time.Now().Add(response.ClockOffset)
	fmt.Printf("Time is: %v \nOffset is: %v", time, response.ClockOffset)
}
