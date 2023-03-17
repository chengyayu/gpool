package gpool

import (
	"fmt"
	"testing"
)

func TestPool_Start(t *testing.T) {
	taskProcessor := func(task any) error {
		fmt.Printf("task processor got: %s\n", task)
		return nil
	}
	resultProcessor := func(res Result) error {
		fmt.Printf("result processor got: %v\n", res)
		return nil
	}

	strings := []string{"first", "second"}
	tasks := make([]interface{}, len(strings))
	for i, s := range strings {
		tasks[i] = s
	}

	pool := NewPool(3)
	pool.Start(tasks, taskProcessor, resultProcessor)
}
