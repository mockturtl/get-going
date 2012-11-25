package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func SleepRand(index int) {
	r := rand.Intn(1000)
	fmt.Printf("#%d Sleeping for %d ms...\n", index, r)
	time.Sleep(time.Duration(r) * time.Millisecond)
}

func WordCount(s string, c chan map[string]int, index int) {
	arr := strings.Fields(s)
	out := make(map[string]int)
	for _, v := range arr {
		out[v]++
	}
	fmt.Printf("#%d %q\n", index, arr)
	SleepRand(index) // sleepy routines finish last!
	c <- out
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	tests := []string{
		"I am learning Go!",
		"The quick brown fox jumped over the lazy dog.",
		"I ate a donut.  Then I ate another donut.",
		"A man a plan a canal panama.",
	}

	c := make(chan map[string]int, len(tests))

	for k, v := range tests {
		go WordCount(v, c, k) // race!
	}

	for i := 0; i < len(tests); i++ {
		fmt.Println(<-c)
	}
}
