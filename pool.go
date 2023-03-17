package gpool

import (
	"sync"
)

type Pool struct {
	cap       int
	jobs      chan Job
	results   chan Result
	done      chan struct{}
	completed bool
}

func NewPool(cap int) *Pool {
	p := &Pool{cap: cap}
	p.jobs = make(chan Job, cap)
	p.results = make(chan Result, cap)
	return p
}

func (p *Pool) Start(tasks []any, procFunc TaskProcessorFn, resFunc ResultProcessorFn) {
	p.done = make(chan struct{})
	go p.allocate(tasks)
	go p.collect(resFunc)
	go p.process(procFunc)
	<-p.done
}

func (p *Pool) allocate(tasks []any) {
	defer close(p.jobs)
	for i, v := range tasks {
		p.jobs <- Job{id: i, task: v}
	}
}

//func (p *Pool) process(processor TaskProcessorFn) {
//	defer close(p.results)
//	eg := new(errgroup.Group)
//	eg.SetLimit(p.cap)
//	for job := range p.jobs {
//		eg.Go(wraper.RecoveredFn(func() error {
//			err := processor(job.task)
//			p.results <- Result{job, err}
//			return err
//		}))
//	}
//	_ = eg.Wait()
//}

func (p *Pool) process(processor TaskProcessorFn) {
	defer close(p.results)
	wg := sync.WaitGroup{}
	for i := 0; i < p.cap; i++ {
		wg.Add(1)
		go p.work(&wg, processor)
	}
	wg.Wait()
}

func (p *Pool) work(wg *sync.WaitGroup, processor TaskProcessorFn) {
	defer wg.Done()
	for job := range p.jobs {
		p.results <- Result{job, processor(job.task)}
	}
}

func (p *Pool) collect(proc ResultProcessorFn) {
	for result := range p.results {
		_ = proc(result)
	}
	p.done <- struct{}{}
	p.completed = true
}

func (p *Pool) IsCompleted() bool {
	return p.completed
}

type TaskProcessorFn func(task any) error

type ResultProcessorFn func(result Result) error

type Job struct {
	id   int
	task any
}

type Result struct {
	Job Job
	Err error
}
