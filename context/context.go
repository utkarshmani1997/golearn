package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "./dummy")
	out, err := cmd.CombinedOutput()
	if err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		select {
		case <-ctx.Done():
			fmt.Println("Hi", ctx.Err(), "err", err)
			return
		default:
			fmt.Println("Error: ", err)
			return
		}
	}

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	default:
	}
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(out)
	}
}
