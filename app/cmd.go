package app

import (
	"fmt"

	"github.com/diego3/hotel-booking-api/core"
)

// Hello represents a function test
func Hello() {
	var arg string
	fmt.Scanf("%s", &arg)
	fmt.Printf("arg: %s", arg)

	//fmt.Printf("dt1: %s, dt2: %s", dt1, dt2)
	if arg == "ping" {
		core.ConnectionPing()
	}
}
