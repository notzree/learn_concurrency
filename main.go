package main

import (
	"time"

	"github.com/notzree/learn_concurrency/solutions"
)

func main() {
	solver := solutions.Solver{}
	// solver.DiningPhilosophers(solutions.DiningProblemParams{
	// 	Philosophers:      [5]string{"Aristotle", "Kant", "Spinoza", "Marx", "Russell"},
	// 	Hunger:            3,
	// 	MedianTimeToEat:   time.Second / 100,
	// 	MedianTimeToThink: time.Second / 100,
	// })
	solver.CheckpointSynchronization(solutions.CheckpointSynchronizationParams{
		NumberOfAssemblies: 3,
		Parts:              []string{"engine", "wheels", "chassis", "body", "paint"},
		MedianBuildTime:    time.Second / 100,
	})
}
