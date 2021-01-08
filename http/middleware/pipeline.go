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

// Run the contract with a final destination pipe holder.
// In this case, pipes are middlewares with a request and a response
func (p Pipeline) Then(controller inter.Controller) inter.Response {
	var holder inter.PipeHolder
	var holders []inter.PipeHolder
	nextHolder := 0
	pipes := p.Pipes

	stabilizeOrder(holders)
	pipes = reverse(pipes)

	for i, pipe := range pipes {
		// Clone pipe and disconnect the reference.
		pipe := pipe

		if i == 0 {
			// Give the last pipe holder a destination controller
			holder = func(request inter.Request) inter.Response {
				response := pipe.Handle(request, controller)
				// Ensure response has an application (needed when middleware returns a new response)
				response.SetApp(request.App())
				return response
			}
		} else {
			// Give other pipe holder the next pipe holder
			holder = func(request inter.Request) inter.Response {
				nextHolder--
				response := pipe.Handle(request, holders[nextHolder])
				// Ensure response has an application (needed when middleware returns a new response)
				response.SetApp(request.App())
				return response
			}
		}
		holders = append(holders, holder)
	}

	holders = setDefaultHolder(controller, holders)
	nextHolder = getNextIndex(holders)

	return holders[nextHolder](p.Passable)
}

// Ensure the pipe holders are stable sorted.
func stabilizeOrder(holders []inter.PipeHolder) {
	sort.SliceStable(holders, func(i, j int) bool {
		return true
	})
}

func setDefaultHolder(controller inter.Controller, pipeHolders []inter.PipeHolder) []inter.PipeHolder {
	// If no pipe holders can be generated because no pipes
	// are present, proceed directly to the destination
	if pipeHolders == nil {
		pipeHolders = append(pipeHolders, controller)
	}
	return pipeHolders
}

// calculate the next holder index
func getNextIndex(pipeHolders []inter.PipeHolder) int {
	return len(pipeHolders) - 1
}

func reverse(pipes []inter.HttpMiddleware) []inter.HttpMiddleware {
	for left, right := 0, len(pipes)-1; left < right; left, right = left+1, right-1 {
		pipes[left], pipes[right] = pipes[right], pipes[left]
	}

	return pipes
}
