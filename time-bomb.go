package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Boom?")

	// see http://golang.org/pkg/time/ 
	tick := time.Tick(1e8)  // 100 ms
	boom := time.After(2e9) // 2000 ms

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BADA BOOM!")
			return
		default:
			fmt.Print("    .")
			time.Sleep(5e7) // 50 ms
		}
	}
	//  this never executes; return is exit 0
	fmt.Println("Big bada boom.")
}
