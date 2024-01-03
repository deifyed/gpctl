package cp

import (
	"context"
	"fmt"

	"github.com/deifyed/gpctl/pkg/gopro"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type Options struct {
	TargetDeviceIndex int
}

func RunE(fs *afero.Afero, opts *Options) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		source := args[0]
		destination := args[1]

		deviceAddress, err := getDeviceAddressByIndex(cmd.Context(), opts.TargetDeviceIndex)
		if err != nil {
			return fmt.Errorf("getting device address: %w", err)
		}

		f, err := gopro.ReadFile(deviceAddress, source)
		if err != nil {
			return fmt.Errorf("reading file %s: %w", source, err)
		}

		defer f.Close()

		err = fs.WriteReader(destination, f)
		if err != nil {
			return fmt.Errorf("writing file %s: %w", destination, err)
		}

		return nil
	}
}

func getDeviceAddressByIndex(ctx context.Context, index int) (string, error) {
	devices, err := gopro.GetDeviceAddresses(ctx)
	if err != nil {
		return "", fmt.Errorf("getting device addresses: %w", err)
	}

	return devices[index], nil
}
