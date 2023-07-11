package cmd

import (
	"github.com/spf13/cobra"
        "github.com/terryyyz/golearn/golearn/channel"
)

var channelCmd = &cobra.Command{
        Use: "channel",
        Short: "Channel",
        Run: func(cmd *cobra.Command, args [] string) {
                channel.RunChannel()
        }, 
}
