package rm

import (
	"fmt"

	"github.com/deifyed/gpctl/pkg/gopro"
	"github.com/spf13/cobra"
)

type Options struct {
	TargetDeviceIndex int
}

func RunE(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		deviceAddress, err := gopro.GetDeviceAddressByIndex(cmd.Context(), opts.TargetDeviceIndex)
		if err != nil {
			return fmt.Errorf("getting device address: %w", err)
		}

		err = gopro.RemoveFile(deviceAddress, targetPath)
		if err != nil {
			return fmt.Errorf("removing file: %w", err)
		}

		return nil
	}
}
