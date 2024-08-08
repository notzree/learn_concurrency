package solutions

import (
	"fmt"
	"hash/fnv"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

type DiningProblemParams struct {
	Philosophers      [5]string
	Hunger            int           // number of times each philosopher needs to eat
	MedianTimeToEat   time.Duration // median time to eat
	MedianTimeToThink time.Duration // median time to think
}

func (s *Solver) DiningPhilosophers(params DiningProblemParams) {
	var dining = sync.WaitGroup{}
	forks := make([]sync.Mutex, len(params.Philosophers))
	dining.Add(len(params.Philosophers))
	start := time.Now()

	for i, philosopher := range params.Philosophers {
		h := fnv.New64a()
		h.Write([]byte(philosopher))
		rg := rand.New(rand.NewSource(uint64(h.Sum64())))
		actualTimeToEat := params.MedianTimeToEat/2 + time.Duration(rg.Int63n(int64(params.MedianTimeToEat)))
		actualTimeToThink := params.MedianTimeToThink/2 + time.Duration(rg.Int63n(int64(params.MedianTimeToThink)))
		fmt.Printf("%s takes %v ms to eat and %v ms to think", philosopher, actualTimeToEat.Milliseconds(), actualTimeToThink.Milliseconds())
		fmt.Println("")
		go eat(philosopher, actualTimeToEat, actualTimeToThink, params.Hunger, &dining, &forks[i], &forks[(i+1)%len(params.Philosophers)]) // plus 1 to convert to 1-based index
	}
	dining.Wait()
	since := time.Since(start)
	fmt.Printf("Everyone is done eating after %.2f seconds \n", since.Seconds())

}

func eat(philosopher string, timeToEat, timeToThink time.Duration, hunger int, dining *sync.WaitGroup, leftHand, rightHand *sync.Mutex) {
	fmt.Println(philosopher, "sat down")
	defer dining.Done()
	// while still hungry, perform these operations
	for h := hunger; h > 0; h-- {
		//try to acquire left and right forks
		fmt.Println(philosopher, "hungry")
		leftHand.Lock()
		rightHand.Lock()
		fmt.Println(philosopher, "eating")
		time.Sleep(timeToEat)
		// release left and right forks
		leftHand.Unlock()
		rightHand.Unlock()
		fmt.Println(philosopher, "thinking")
		time.Sleep(timeToThink)
	}
	fmt.Println(philosopher, "satisfied")

}
