package startwebcam

import (
	"context"
	"errors"
	"fmt"

	"github.com/deifyed/gpctl/pkg/ffmpeg"
	"github.com/deifyed/gpctl/pkg/gopro"
	"github.com/deifyed/gpctl/pkg/v4l2"
	"github.com/spf13/cobra"
)

const defaultDeviceLabel = "GoPro"
const defaultVideoDeviceIndex = 42

type Options struct {
	TargetDeviceIndex int
	Username          string
}

func RunE(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if opts.Username == "" {
			return errors.New("missing username")
		}

		deviceAddress, err := getDeviceAddressByIndex(cmd.Context(), opts.TargetDeviceIndex)
		if err != nil {
			return fmt.Errorf("getting device address: %w", err)
		}

		err = v4l2.InjectModule(v4l2.InjectModuleOptions{
			Label:       defaultDeviceLabel,
			DeviceIndex: defaultVideoDeviceIndex,
		})
		if err != nil {
			return fmt.Errorf("injecting v4l2loopback module: %w", err)
		}

		err = gopro.StartWebcam(deviceAddress)
		if err != nil {
			return fmt.Errorf("starting webcam: %w", err)
		}

		err = ffmpeg.Expose(ffmpeg.ExposeOptions{
			DeviceIndex: defaultVideoDeviceIndex,
			Username:    opts.Username,
		})
		if err != nil {
			return fmt.Errorf("exposing video device: %w", err)
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
