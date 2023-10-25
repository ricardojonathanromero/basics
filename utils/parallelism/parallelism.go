package parallelism

import (
	"sync"
)

type ParallelTask interface {
	RunInParallel(to To, next Next)
	Wait()
}

type parallelTaskImpl struct {
	wg *sync.WaitGroup
}

type Next func(args any)
type To func(args ...interface{}) any

func (pt *parallelTaskImpl) RunInParallel(to To, next Next) {
	pt.wg.Add(1)

	// exec any function call
	go func() {
		defer pt.wg.Done()
		next(to())
	}()
}

func (pt *parallelTaskImpl) Wait() {
	pt.wg.Wait()
}

func NewParallel() ParallelTask {
	return &parallelTaskImpl{
		wg: &sync.WaitGroup{},
	}
}
