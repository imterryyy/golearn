package cmd

import (
        "os"
        "os/signal"
	"syscall"
	"github.com/spf13/cobra"
	"github.com/terryyyz/golearn/golearn/rpchandler"
)

var rpcHandlerCmd = &cobra.Command{
        Use: "rpc_handler",
        Short: "Rpc Handler",
        Run: func(cmd *cobra.Command, args [] string) {
                rpcHandler := rpchandler.NewRpcHandler()
                sigChan := make(chan os.Signal, 1)
                signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

                // Start your service logic here
                go rpcHandler.Run()

                // Wait for termination signal
                <-sigChan
        }, 
}
