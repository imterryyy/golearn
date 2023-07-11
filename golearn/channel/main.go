package channel

import (
	"fmt"
	"time"
)

func testGoroutine(goroutineName string) {
        fmt.Println(goroutineName)
}

func RunChannel() {
        fmt.Println("Start channel")
        chan1 := make(chan string)

        // send data to channel
        go func() {
                chan1 <- "Terry"
        }()

        // receive data from channel
        fmt.Println(<-chan1)

        go testGoroutine("second goroutine")
        testGoroutine("main goroutine")

        time.Sleep(time.Second)
}
