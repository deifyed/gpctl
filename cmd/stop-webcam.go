package cmd

import (
	stopwebcam "github.com/deifyed/gpctl/cmd/stop-webcam"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmdOpts stopwebcam.Options
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Deactivate webcam mode",
	RunE:  stopwebcam.RunE(&stopCmdOpts),
}

func init() {
	webcamCmd.AddCommand(stopCmd)

	stopCmd.Flags().IntVarP(&stopCmdOpts.TargetDeviceIndex, "device", "d", 0, "Index of the device to use (default is 0)")
}
