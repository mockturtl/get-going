package main

import (
	"fmt"
	"strings"
)

func WordCount(s string, c chan map[string]int) {
	arr := strings.Fields(s)
	out := make(map[string]int)
	for _, v := range arr {
		out[v]++
	}
	fmt.Println(arr)
	c <- out
}

func main() {
	tests := []string{
		"I am learning Go!",
		"The quick brown fox jumped over the lazy dog.",
		"I ate a donut.  Then I ate another donut.",
		"A man a plan a canal panama.",
	}

	c := make(chan map[string]int, len(tests))

	for _, v := range tests {
		go WordCount(v, c) // races!
	}

	for i := 0; i < len(tests); i++ {
		fmt.Println(<-c)
	}
}
