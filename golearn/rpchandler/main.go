package rpchandler

import (
	"log"
	"net"
	"net/rpc"
)

type RpcHandler struct {
}

func NewRpcHandler() *RpcHandler {
        return &RpcHandler{
        }
}

func (r *RpcHandler) Hello(request string, reply *string) error {
        *reply = "Hello" + request
        return nil
}

func (r *RpcHandler) Run() {
        rpc.RegisterName("RpcHandler", new(RpcHandler))
        listener, err := net.Listen("tcp", ":8000")

        if err != nil {
                log.Fatal("ListenTCP error:", err)
        }

        for {
                conn, err := listener.Accept()
                if err != nil {
                        log.Fatal("ListenTCP error:", err)
                }

                go rpc.ServeConn(conn)
        }
}
