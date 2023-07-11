package golearn

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloWorld = &cobra.Command{
        Use: "hello_world",
        Short: "Hello world program",
        Long: "Hello world program",
        Run: func(cmd * cobra.Command, args [] string) {
                fmt.Println("Hello world from hello_world command")
        },
}


func init() {
        rootCmd.AddCommand(helloWorld)
}
