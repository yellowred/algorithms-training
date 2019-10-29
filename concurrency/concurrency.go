package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

const N = 1

var forks [5]*Philosopher

type Philosopher int

func (p Philosopher) pickLeftFork() {
	forks[leftFork(p)] = &p
	log.Println(fmt.Sprintf("[pickLeftFork] forks=%v", forks))
}

func (p Philosopher) pickRightFork() {
	forks[rightFork(p)] = &p
	log.Println(fmt.Sprintf("[pickRightFork] forks=%v", forks))
}

func (p Philosopher) putLeftFork() {
	forks[leftFork(p)] = nil
	log.Println(fmt.Sprintf("[putRightFork] forks=%v", forks))
}

func (p Philosopher) putRightFork() {
	forks[rightFork(p)] = nil
	log.Println(fmt.Sprintf("[putLeftFork] forks=%v", forks))
}

func (p Philosopher) eat() {
	if int(*forks[leftFork(p)]) == int(p) && int(*forks[rightFork(p)]) == int(p) {
		log.Printf("Eats: %v\n", p)
	} else {
		panic(fmt.Sprintf("Forks %d, %d are not taken for philosopher %d, forks=%v", leftFork(p), rightFork(p), p, forks))
	}
}

// ----
var forkMutexes [5]*sync.Mutex
var wg sync.WaitGroup

func main() {
	log.Println("Start")
	for i := 0; i < 5; i++ {
		forkMutexes[i] = &sync.Mutex{}
	}
	for i := 0; i < 5; i++ {
		wg.Add(N)

		go func(p Philosopher) {
			for l := 0; l < N; l++ {
				p.wantsToEat()
				wg.Done()
				time.Sleep(100 * time.Millisecond)
			}
		}(Philosopher(i))
	}
	wg.Wait()
}

func (p Philosopher) wantsToEat() {
	forkMutexes[leftFork(p)].Lock()
	p.pickLeftFork()
	forkMutexes[rightFork(p)].Lock()
	p.pickRightFork()
	p.eat()
	p.putLeftFork()
	forkMutexes[leftFork(p)].Unlock()
	p.putRightFork()
	forkMutexes[rightFork(p)].Unlock()
}

// ----
func leftFork(philosopher Philosopher) int {
	if philosopher == 0 {
		return 4
	}
	return int(philosopher)
}

func rightFork(philosopher Philosopher) int {
	if philosopher == 4 {
		return 0
	}
	return int(philosopher)
}
