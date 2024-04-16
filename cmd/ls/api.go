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
		targetPath := "/"

		if len(args) > 0 {
			targetPath = args[0]
		} else {
			fmt.Fprintln(cmd.OutOrStdout(), "No directory specified, using root '/'")
		}

		deviceAddress, err := gopro.GetDeviceAddressByIndex(cmd.Context(), opts.TargetDeviceIndex)
		if err != nil {
			return fmt.Errorf("getting device address: %w", err)
		}

		items, err := gopro.ListDirectory(deviceAddress, targetPath)
		if err != nil {
			return fmt.Errorf("listing directory %s: %w", targetPath, err)
		}

		for _, item := range items {
			fmt.Println(item)
		}

		return nil
	}
}
