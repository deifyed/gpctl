package cmd

import (
	"github.com/deifyed/gpctl/cmd/ls"
	"github.com/spf13/cobra"
)

var lsOpts = ls.Options{
	TargetDeviceIndex: 0,
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files on the GoPro",
	RunE:  ls.RunE(&lsOpts),
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.Flags().IntVarP(&lsOpts.TargetDeviceIndex, "device", "d", 0, "Index of the device to use (default is 0)")
}
