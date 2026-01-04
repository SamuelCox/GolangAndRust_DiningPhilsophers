package main

import (
	"fmt"
	"time"
)

type Fork struct {
	id int
}

type Philosopher struct {
	name string
	leftFork chan Fork
	rightFork chan Fork
}

func (p Philosopher) dine(host chan struct{}) {
	for {
		host <- struct{}{}
		leftFork := <-p.leftFork
		rightFork := <-p.rightFork

		fmt.Printf("%s is eating\n", p.name)
		time.Sleep(250 * time.Millisecond)

		p.leftFork <- Fork { id: leftFork.id }
		p.rightFork <- Fork { id: rightFork.id }
		fmt.Printf("%s is done eating\n", p.name)
		<-host

		fmt.Printf("%s is thinking\n", p.name)
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%s is done thinking\n", p.name)
	}
}

func main() {
    fmt.Println("Starting Dining philosophers")

	philosopherNames := []string{"Dijkstra", "Turing", "Lovelace", "Torvalds", "Hopper"}
	forks := make([]chan Fork, len(philosopherNames))
	for i := range forks {
		forks[i] = make(chan Fork, 1)
		forks[i] <- Fork { id: i }
	}

	philosophers := make([]Philosopher, len(philosopherNames))
	for i, name := range philosopherNames {
		philosophers[i] = Philosopher {
			name: name,
			leftFork: forks[i],
			rightFork: forks[(i+1) % len(forks)],
		}
	}

	host := make(chan struct{}, len(philosophers) - 1)

	for _, philosopher := range philosophers {
		go philosopher.dine(host)
	}

	select {}
	
}