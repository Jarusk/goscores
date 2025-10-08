package main

import (
	"github.com/jarusk/goscores/internal/providers"
	"github.com/spf13/cobra"
)

var baseballCmd = &cobra.Command{
	Use:   "baseball",
	Short: "get baseball scores",
	Run: func(cmd *cobra.Command, args []string) {
		run(providers.Baseball, providers.MLB, "")
	},
}

var hockeyCmd = &cobra.Command{
	Use:   "hockey",
	Short: "get hockey scores",
	Run: func(cmd *cobra.Command, args []string) {
		run(providers.Hockey, providers.NHL, "")
	},
}
