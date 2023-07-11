package worker

import (
	"log"
	"sync"
)

type Task interface {
        Execute(num int) error
        OnFailure(error)
}

type Pool struct {
        numWorkers      int
        tasks           chan Task
        start           sync.Once
        end             sync.Once
        quit            chan string
}

func NewPool(numWorkers int) *Pool {
        tasks := make(chan Task, numWorkers)
        return &Pool{
                numWorkers:     numWorkers,
                tasks:          tasks,
                start:          sync.Once{},
                end:            sync.Once{},
                quit:           make(chan string),
        }
}

func (p *Pool) Start() {
        p.start.Do(func() {
                log.Print("Start working pool")
                p.startWorkers()
        })
}

func (p *Pool) Stop() {
        p.end.Do(func () {
                log.Print("Stop working pool")
                close(p.quit)
        })
}

func (p *Pool) AddWork(t Task) {
        select {
        case p.tasks <- t:
        case <- p.quit:
        }
}

func (p *Pool) startWorkers() {
        for i := 0; i < p.numWorkers; i++ {
                go func(workerIndex int) {
                        log.Printf("Worker %d", workerIndex)
                        for {
                                select {
                                case <- p.quit:
                                        log.Printf("Stop worker %d", workerIndex)
                                        return
                                case task, ok := <- p.tasks:
                                        if !ok {
                                                log.Printf("Stop worker %d", workerIndex)
                                        }

                                        if err := task.Execute(workerIndex); err != nil {
                                                task.OnFailure(err)
                                        }
                                }
                        }
                }(i)
        }
}
