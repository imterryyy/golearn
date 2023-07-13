package cmd

import (
        "os"
        "os/signal"
	"syscall"
	"github.com/spf13/cobra"
        "github.com/terryyyz/golearn/golearn/listenevent"
)

var listenerCmd = &cobra.Command{
        Use: "listen",
        Short: "Listen",
        Run: func(cmd *cobra.Command, args [] string) {
                listener := listenevent.NewListener()
                sigChan := make(chan os.Signal, 1)
                signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

                // Start your service logic here
                go listener.Run()

                // Wait for termination signal
                <-sigChan
        }, 
}
