package main

import (
	"fmt"
	"log"
	"reflect"
	"unsafe"
)

func throw(sum uint64) {
	err := fmt.Sprintf("`sum` rolled over: Sizeof(%s)=%d", reflect.TypeOf(sum), unsafe.Sizeof(sum))
	log.Panic(err)
}

func fibonacci() func() (int, uint64) {
	var a, b uint64 = 1, 1
	c := 0
	return func() (count int, sum uint64) {
		if c == 0 {
			sum = a
			c++
			return c, sum
		} else if c == 1 {
			sum = b
			c++
			return c, sum
		}
		sum = a + b
		//  2^64 ~= 1.84e19 < fib(94) ~= 1.97e19
		if sum < a || sum < b {
			throw(sum)
		}
		a, b = b, sum
		c++
		return c, sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 93; i++ {
		fmt.Println(f())
	}
}
