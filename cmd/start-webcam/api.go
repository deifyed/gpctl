package startwebcam

import "github.com/spf13/cobra"

type Options struct {
	TargetDeviceIndex int
}

func RunE(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return nil
	}
}
