package stopwebcam

import (
	"fmt"

	"github.com/deifyed/gpctl/pkg/gopro"
	"github.com/deifyed/gpctl/pkg/v4l2"
	"github.com/spf13/cobra"
)

type Options struct {
	TargetDeviceIndex int
	ExitWebcamMode    bool
}

func RunE(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		deviceAddress, err := gopro.GetDeviceAddressByIndex(cmd.Context(), opts.TargetDeviceIndex)
		if err != nil {
			return fmt.Errorf("getting device address: %w", err)
		}

		err = gopro.StopWebcam(deviceAddress)
		if err != nil {
			return fmt.Errorf("starting webcam: %w", err)
		}

		if !opts.ExitWebcamMode {
			return nil
		}

		err = gopro.ExitWebcamMode(deviceAddress)
		if err != nil {
			return fmt.Errorf("exiting webcam mode: %w", err)
		}

		err = v4l2.EjectModule()
		if err != nil {
			return fmt.Errorf("ejecting v4l2loopback module: %w", err)
		}

		return nil
	}
}
