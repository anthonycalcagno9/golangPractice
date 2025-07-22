package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// The context package in Go is essential for managing request-scoped values,
// cancellation signals, and timeouts across API boundaries and goroutines.
//
// WHY CONTEXT IS USEFUL:
// 1. Cancellation: Cancel long-running operations (HTTP requests, database queries)
// 2. Timeouts: Set deadlines for operations to prevent hanging
// 3. Request-scoped values: Pass user IDs, auth tokens, trace IDs through call chains
// 4. Graceful shutdown: Signal goroutines to stop cleanly
//
// COMMON USE CASES:
// - HTTP servers: Each request gets a context that can be cancelled if client disconnects
// - Database operations: Set timeouts to prevent long-running queries
// - Microservices: Propagate request metadata (trace IDs, user context)
// - Worker pools: Cancel all workers when shutting down
// - Fan-out patterns: Cancel multiple goroutines if one fails
//
// CONTEXT TYPES:
// - context.Background(): Root context, never cancelled
// - context.TODO(): When you're unsure what context to use
// - context.WithCancel(): Manually cancellable context
// - context.WithTimeout(): Automatically cancels after duration
// - context.WithDeadline(): Cancels at specific time
// - context.WithValue(): Carries request-scoped values
//
// BEST PRACTICES:
// - Always pass context as first parameter: func doWork(ctx context.Context, ...)
// - Don't store contexts in structs, pass them explicitly
// - Don't pass nil context, use context.Background() or context.TODO()
// - Check ctx.Done() in loops and long-running operations

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1 = ", ctx.Err())
	fmt.Println("num of goroutines before = ", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 500)
				fmt.Println("working on ", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2", ctx.Err())
	fmt.Println("num goroutines check 2 = ", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("context cancelled")

	time.Sleep(time.Millisecond * 500)
	fmt.Println("error check 3 (after cancel) = ", ctx.Err())
	fmt.Println("num goroutines after cancel = ", runtime.NumGoroutine())

}
