package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/terryyyz/golearn/golearn/worker"
)

type Task struct {}

func (t *Task) Execute(num int) error {
        log.Printf("Execute task by worker %d", num)
        return nil
}

func (t *Task) OnFailure(e error) {
        log.Printf("Execute Error")
}

var workerCmd = &cobra.Command{
        Use: "worker",
        Short: "Worker",
        Run: func(cmd *cobra.Command, args [] string) {
                workPool := worker.NewPool(4)
                workPool.Start()

                task := &Task{}

                for i := 0; i < 100; i++ {
                        workPool.AddWork(task)
                        time.Sleep(time.Second)
                }
        }, 
}

