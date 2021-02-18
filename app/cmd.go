package app

import (
	"fmt"
	//"hotel-booking.com/v1/core"
)

func Hello() {
	var arg string
	fmt.Scanf("%s", &arg)
	fmt.Printf("arg: %s", arg)

	//fmt.Printf("dt1: %s, dt2: %s", dt1, dt2)
	if arg == "ping" {
		//core.ConnectionPing()
	}
}
