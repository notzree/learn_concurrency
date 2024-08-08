package solutions

import (
	"fmt"
	"hash/fnv"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

type CheckpointSynchronizationParams struct {
	NumberOfAssemblies int      //number of completed parts to build
	Parts              []string //may be some sort of tree or hashmap in the future
	MedianBuildTime    time.Duration
}

func (s *Solver) CheckpointSynchronization(params CheckpointSynchronizationParams) {
	assembly := sync.WaitGroup{}

	for n := params.NumberOfAssemblies; n > 0; n-- {
		assembly.Add(len(params.Parts)) // need to wait for all parts to finish building
		for _, part := range params.Parts {
			assemblyTime := getRandomTime(params.MedianBuildTime, part)
			go build(part, assemblyTime, &assembly)
		}
		assembly.Wait() // assembly.wait() will block until all parts are built, aka synchronize the assembly
		fmt.Println("assembly complete")
	}

}

func build(part string, buildTime time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("working on part", part)
	time.Sleep(buildTime)
	fmt.Println("finished building", part)
}

func getRandomTime(medianTime time.Duration, input string) time.Duration {
	h := fnv.New64a()
	h.Write([]byte(input))
	rg := rand.New(rand.NewSource(uint64(h.Sum64())))
	return medianTime/2 + time.Duration(rg.Int63n(int64(medianTime)))
}
