package cmd

import (
	"github.com/deifyed/gpctl/cmd/cp"
	"github.com/spf13/cobra"
)

var cpOpts cp.Options
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copies files from a GoPro device to the local machine",
	Args:  cobra.ExactArgs(2),
	RunE:  cp.RunE(fs, &cpOpts),
}

func init() {
	rootCmd.AddCommand(cpCmd)

	cpCmd.Flags().IntVarP(&cpOpts.TargetDeviceIndex, "device", "d", 0, "Index of the device to use (default is 0)")
	cpCmd.Flags().BoolVarP(&cpOpts.Progressbar, "progressbar", "p", true, "Show progressbar (default is true)")
}
