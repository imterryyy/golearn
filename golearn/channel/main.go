package channel

import (
	"fmt"
	"time"
)

func RunChannel() {
        fmt.Println("Start channel")
        chan1 := make(chan string)

        // use goroutine to send data to chan1
        go func() {
                chan1 <- "Terry"
        }()

        // main goroutine receive data from chan1
        fmt.Println(<-chan1)

        time.Sleep(time.Second)
}
