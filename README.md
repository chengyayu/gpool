# gpool
goroutine pool

## 结构

[]tasks -> (allocate fn) -> [jobs channel] -> (work fn) -> [result channel] -> (collect fn) -> [done channel]

## 使用

```go
  ......
  taskProcessor := func(task any) error {
		fmt.Printf("task processor got: %s\n", task)
		return nil
	}
	resultProcessor := func(res Result) error {
		fmt.Printf("result processor got: %v\n", res)
		return nil
	}

	strings := []string{"first", "second"}
	task := make([]interface{}, len(strings))
	for i, s := range tasks {
		tasks[i] = s
	}

	pool := NewPool(3)
	pool.Start(tasks, taskProcessor, resultProcessor)
  ......
```
