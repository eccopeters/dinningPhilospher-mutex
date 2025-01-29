package main

import (
	"fmt"
	"math/rand"

	"sync"
	"time"
)

var ch = make(chan int, 3)

type fork struct{ sync.Mutex }

type philosopher struct {
	id                  int
	leftFork, rightFork *fork
}

// Goes from thinking to hungry to eating and done eating then starts over.
// Adapt the pause values to increased or decrease contentions
// around the forks.
func (p philosopher) eat(n int) {

	say("thinking", p.id)
	randomPause(2)

	say("hungry", p.id)
	p.leftFork.Lock()
	p.rightFork.Lock()
	randomPause(5)

	say("eating", p.id)

	p.rightFork.Unlock()
	p.leftFork.Unlock()

	say("done eating", p.id)

	ch <- n

}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func say(action string, id int) {
	fmt.Printf("Philosopher #%d is %s\n", id+1, action)
}

func init() {
	// Random seed
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	// How many philosophers and forks
	count := 5

	// Create forks
	forks := make([]*fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(fork)
	}

	// Create philospoher, assign them 2 forks and send them to the dining table
	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count]}
		for j := 0; j < 3; j++ {
			go philosophers[i].eat(j)
		}
	}

	philoEat := make(chan int, 3)
	<-philoEat

}