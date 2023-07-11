package helloworld

import "fmt"

type HelloApp struct {
        name string
}

func NewHelloApp(name string) *HelloApp {
        return &HelloApp{
                name: name,
        }
}

func (h *HelloApp) SayHello() {
        fmt.Println("Hello", h.name)
}
