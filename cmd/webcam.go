package cmd

import (
	"github.com/spf13/cobra"
)

// webcamCmd represents the webcam command
var webcamCmd = &cobra.Command{
	Use:   "webcam",
	Short: "Administrate webcam functionality",
}

func init() {
	rootCmd.AddCommand(webcamCmd)
}
