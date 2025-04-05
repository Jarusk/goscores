package main

import (
	"github.com/spf13/cobra"
)

var hockeyCmd = &cobra.Command{
	Use:   "hockey",
	Short: "get hockey scores",
	Run: func(cmd *cobra.Command, args []string) {
		run("hockey", "nhl", "")
	},
}
