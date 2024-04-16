package cp

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

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

		deviceAddress, err := gopro.GetDeviceAddressByIndex(cmd.Context(), opts.TargetDeviceIndex)
		if err != nil {
			return fmt.Errorf("getting device address: %w", err)
		}

		realDestination, err := getRealDestination(fs, source, destination)
		if err != nil {
			return fmt.Errorf("getting real destination: %w", err)
		}

		err = cp(fs, deviceAddress, source, realDestination[0])
		if err != nil {
			return fmt.Errorf("copying file: %w", err)
		}

		return nil
	}
}

func getRealDestination(fs *afero.Afero, source string, destination string) ([]string, error) {
	var err error
	realDestination := destination

	if !path.IsAbs(realDestination) {
		realDestination, err = filepath.Abs(destination)
		if err != nil {
			return []string{}, fmt.Errorf("getting absolute destination: %w", err)
		}
	}

	isDir, err := fs.IsDir(realDestination)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return []string{}, fmt.Errorf("checking if destination is directory: %w", err)
	}

	if isDir {
		return []string{path.Join(realDestination, path.Base(source))}, nil
	}

	return []string{realDestination}, nil
}

func cp(fs *afero.Afero, deviceAddress string, source string, destination string) error {
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
