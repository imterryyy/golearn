package golearn

import (
        "os"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
        Use: "golearn",
        Short: "Golearn is learn program to learn golang from scratch",
        Long: "Golearn is learn program to learn golang from scratch",
}

func Execute() {
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

func init() {}
