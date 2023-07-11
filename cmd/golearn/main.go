package golearn

import (
        "os"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = *&cobra.Command{
        Use: "golearn",
        Short: "Golearn is learn program to learn golang from scratch",
        Run: runApp,
}

func runApp(cmd *cobra.Command, args [] string) {
        fmt.Println("Hello world from Golearn")
}

func Execute() {
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
