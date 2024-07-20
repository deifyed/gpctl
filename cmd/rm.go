package cmd

import (
	"github.com/deifyed/gpctl/cmd/rm"
	"github.com/spf13/cobra"
)

var rmOpts = rm.Options{
	TargetDeviceIndex: 0,
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a file or directory on the GoPro",
	Args:  cobra.ExactArgs(1),
	RunE:  rm.RunE(&rmOpts),
}

func init() {
	rootCmd.AddCommand(rmCmd)

	rmCmd.Flags().IntVarP(&rmOpts.TargetDeviceIndex, "device", "d", 0, "Index of the device to use (default is 0)")
	rmCmd.Flags().BoolVarP(&rmOpts.Progressbar, "progressbar", "p", true, "Show progressbar (default is true)")
}
