package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support/caller"
	"sort"
)

type Pipe interface {
	Handle(request inter.Request, next inter.Next) inter.Response
}

// noinspection GoNameStartsWithPackageName
type Pipeline struct {
	App inter.App

	// The object being passed through the contract.
	Passable inter.Request

	// the array of pipes.
	Pipes []inter.HttpMiddleware
}

func NewPipeline(app inter.App) Pipeline {
	return Pipeline{App: app}
}

func (p Pipeline) AppPath() string {
	return caller.Path()
}

// Set the object being sent through the contract.
func (p Pipeline) Send(passable inter.Request) Pipeline {
	p.Passable = passable

	return p
}

// Set the array of pipes.
func (p Pipeline) Through(pipes []inter.HttpMiddleware) Pipeline {
	p.Pipes = pipes

	return p
}

// Run the contract with a final destination callback.
func (p Pipeline) Then(destination inter.Next) inter.Response {

	var callbacks []func(data inter.Request) inter.Response
	sort.SliceStable(callbacks, func(i, j int) bool {
		return true
	})
	nextCallback := 0
	pipes := p.Pipes

	pipes = reverse(pipes)

	for i, pipe := range pipes {
		pipe := pipe
		if i == 0 {
			// Give the last callback a destination callback
			callback := func(data inter.Request) inter.Response {
				return pipe.Handle(data, destination)
			}
			callbacks = append(callbacks, callback)
		} else {
			// Give other callback the next callback
			callback := func(data inter.Request) inter.Response {
				nextCallback--
				return pipe.Handle(data, callbacks[nextCallback])
			}
			callbacks = append(callbacks, callback)
		}
	}

	// If no callbacks can be generated because no pipes
	// are present, proceed directly to the destination
	if callbacks == nil {
		callbacks = append(callbacks, destination)
	}

	nextCallback = len(callbacks) - 1

	return callbacks[nextCallback](p.Passable)
}

func reverse(pipes []inter.HttpMiddleware) []inter.HttpMiddleware {
	for left, right := 0, len(pipes)-1; left < right; left, right = left+1, right-1 {
		pipes[left], pipes[right] = pipes[right], pipes[left]
	}

	return pipes
}
