package cmd

import (
	startwebcam "github.com/deifyed/gpctl/cmd/start-webcam"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmdOpts startwebcam.Options
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Activate the webcam",
	RunE:  startwebcam.RunE(&startCmdOpts),
}

func init() {
	webcamCmd.AddCommand(startCmd)

	startCmd.Flags().IntVarP(&startCmdOpts.TargetDeviceIndex, "device", "d", 0, "Index of the device to use (default is 0)")
}
