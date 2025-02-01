
package main

import (
        "fmt"
        "math/rand"
        "sync"
        "time"
)

type Philosopher struct {
        id    int
        left  *sync.Mutex
        right *sync.Mutex
}

func (p *Philosopher) eat() {
        fmt.Printf("Philosopher %d is eating\n", p.id)
        time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Simulate eating
        fmt.Printf("Philosopher %d finished eating\n", p.id)

}

func (p *Philosopher) think() {
        fmt.Printf("Philosopher %d is thinking\n", p.id)
        time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Simulate thinking
}

func (p *Philosopher) dine() {
        for {
                p.think()

                fmt.Printf("Philosopher %d is trying to eat\n", p.id)

                // Acquire locks -  Important: Order matters to prevent deadlock!
                p.left.Lock()
                fmt.Printf("Philosopher %d acquired left fork\n", p.id)
                p.right.Lock() // If this blocks, the left is held!
                fmt.Printf("Philosopher %d acquired right fork\n", p.id)

                p.eat()

                // Release locks - Order doesn't matter here, but good practice to reverse.
                p.right.Unlock()
                fmt.Printf("Philosopher %d released right fork\n", p.id)
                p.left.Unlock()
                fmt.Printf("Philosopher %d released left fork\n", p.id)


        }
}

func main() {
        rand.Seed(time.Now().UnixNano()) // Initialize random seed

        numPhilosophers := 5
        forks := make([]*sync.Mutex, numPhilosophers)
        philosophers := make([]*Philosopher, numPhilosophers)

        for i := 0; i < numPhilosophers; i++ {
                forks[i] = &sync.Mutex{}
        }

        for i := 0; i < numPhilosophers; i++ {
                philosophers[i] = &Philosopher{
                        id:    i,
                        left:  forks[i],
                        right: forks[(i+1)%numPhilosophers], // Circular - last philosopher's right is first fork
                }
        }

        var wg sync.WaitGroup
        for i := 0; i < numPhilosophers; i++ {
                wg.Add(1)
                go func(p *Philosopher) {
                        defer wg.Done()
                        p.dine()
                }(philosophers[i])
        }

        wg.Wait()
        fmt.Println("Dining simulation finished.")
}