package cmd

import (
	"github.com/deifyed/gpctl/cmd/cp"
	"github.com/spf13/cobra"
)

var cpOpts cp.Options
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	RunE:  cp.RunE(fs, &cpOpts),
}

func init() {
	rootCmd.AddCommand(cpCmd)

	cpCmd.Flags().IntVarP(&lsOpts.TargetDeviceIndex, "device", "d", 0, "Index of the device to use (default is 0)")
}
