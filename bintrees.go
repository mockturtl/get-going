package main

import (
	"code.google.com/p/go-tour/tree"
	"log"
	"math/rand"
	"time"
)

func block() {
	time.Sleep(time.Duration(rand.Intn(500) * 1e6))
}

// see code.google.com/p/go-tour/tree/tree.go
// tree#insert nests tuples with `<`, left to right, in the usual numeric order
// tree#String prints the tuple representation
// 
// read values left to right
// sends 1..10 regardless of branch and leaf structure
func doWalk(t *tree.Tree, ch chan int) {
	block() //  test
	if t.Left != nil {
		doWalk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		doWalk(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	doWalk(t, ch)
	close(ch)
}

// k is the ratio of t2's to t1's values
func Same(t1, t2 *tree.Tree, k int) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < 10; i++ {
		n1, ok1 := <-c1 // sleep blocks channel read
		n2, ok2 := <-c2
		log.Println("comparing nodes ", n1, n2)
		if n1 != n2 || // order never changes, see above
			ok1 != ok2 {
			return false
		}
	}
	return true
}

func dump(t *tree.Tree) {
	c := make(chan int)
	go Walk(t, c)
	for i := 0; i < 10; i++ {
		log.Println(<-c)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	t1, t2 := tree.New(1), tree.New(2)

	log.Println("t1: ", t1)
	//dump(t1)

	log.Println("t2: ", t2)
	//dump(t2)

	log.Println("t1 == t1? ", Same(t1, t1, 1))
	log.Println("t2 == t2? ", Same(t2, t2, 1))
	log.Println("t1 == t2? ", Same(t1, t2, 2))
}
