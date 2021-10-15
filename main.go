package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(5)*time.Second)
	defer cancel() // 无论何时， 都应该手动执行 cancel 方法

	go job(ctx)

	for {
	}
}

func isContinue(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return false
	default:
		return true
	}
}

func jobInit(ctx context.Context) {
	if !isContinue(ctx) {
		return
	}

	// do init
}

// 一个任务
func job(ctx context.Context) {

	go func() {
		fmt.Println("start job")
		// init
		time.Sleep(5 * time.Second)
		// main process
		time.Sleep(10 * time.Second)
		// clean
		time.Sleep(5 * time.Second)

		fmt.Println("20s passed")
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("ctx.Done with err: %v\n", ctx.Err())
		return
	}
}

// 控制信号
func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
