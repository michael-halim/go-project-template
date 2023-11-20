package cmd

import (
	"github.com/michael-halim/go-transactions/internal/server"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start Http Service",
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer()
		server.Run()
	},
}
