package listenevent

import "fmt"

type Listener struct {

}

func NewListener() *Listener {
        return &Listener{

        }
}

func (l *Listener) Run() {
        go fmt.Print("Hello")
}
