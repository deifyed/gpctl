package cmd

import (
	"github.com/deifyed/gpctl/cmd/devices"
	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List available devices",
	RunE:  devices.RunE(),
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
