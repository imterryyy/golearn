package cmd

import (
	"errors"

	"github.com/spf13/cobra"
        "github.com/terryyyz/golearn/golearn/helloworld"
)

var helloWorldCmd = &cobra.Command{
        Use: "hello_world",
        Short: "Hello world",
        Long: "Golearn is learn program to learn golang from scratch",
        Args: func(cmd *cobra.Command, args [] string) error {
                if (len(args) < 1) {
                        return errors.New("Missing name")
                }
                return nil
        },
        Run: func(cmd *cobra.Command, args [] string) {
                var helloWorldApp = helloworld.NewHelloApp(args[0])
                helloWorldApp.SayHello()
        }, 
}
