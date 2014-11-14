package workgroup

import (
	"sync"

	"github.com/willglynn/hatt/manifest"
)

type Workgroup struct {
	input     chan string
	output    chan Result
	waitGroup sync.WaitGroup
}

type Result struct {
	Path  string
	Error error
	File  *manifest.File
}

func New(n int) *Workgroup {
	wg := &Workgroup{
		input:  make(chan string),
		output: make(chan Result),
	}

	for i := 0; i < n; i++ {
		go wg.work()
	}
	wg.waitGroup.Add(n)

	// automatically close the output channel when the workers are done
	go func() {
		wg.Wait()
		close(wg.output)
	}()

	return wg
}

func (wg *Workgroup) work() {
	defer wg.waitGroup.Done()

	for {
		path, ok := <-wg.input
		if !ok {
			return
		}

		file, err := manifest.NewFileFromPath(path)
		wg.output <- Result{
			Path:  path,
			Error: err,
			File:  file,
		}
	}
}

func (wg *Workgroup) Input() chan<- string {
	return wg.input
}

func (wg *Workgroup) Output() <-chan Result {
	return wg.output
}

func (wg *Workgroup) Wait() {
	wg.waitGroup.Wait()
}
