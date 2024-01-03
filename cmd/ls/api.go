package ls

import (
	"fmt"

	"github.com/deifyed/gpctl/pkg/gopro"
	"github.com/spf13/cobra"
)

type Options struct {
	TargetDeviceIndex int
}

func RunE(opts *Options) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var targetPath string

		if len(args) == 0 {
			targetPath = "/"
		}

		deviceAddress, err := getDeviceAddressByIndex(opts.TargetDeviceIndex)
		if err != nil {
			return fmt.Errorf("getting device address: %w", err)
		}

		gopro.ListDirectory(deviceAddress, targetPath)

		return nil
	}
}

func getDeviceAddressByIndex(index int) (string, error) {
	devices, err := gopro.GetDeviceAddresses()
	if err != nil {
		return "", fmt.Errorf("getting device addresses: %w", err)
	}

	return devices[index], nil
}
