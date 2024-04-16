package cp

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/deifyed/gpctl/pkg/gopro"
	"github.com/gobwas/glob"
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

		realSources, err := getRealSource(deviceAddress, source)
		if err != nil {
			return fmt.Errorf("getting real source: %w", err)
		}

		for _, sourceItem := range realSources {
			realDestination, err := getRealDestination(fs, sourceItem, destination)
			if err != nil {
				return fmt.Errorf("getting real destination for %s: %w", destination, err)
			}

			err = cp(fs, deviceAddress, sourceItem, realDestination)
			if err != nil {
				return fmt.Errorf("copying file from %s to %s: %w", source, realDestination, err)
			}
		}

		return nil
	}
}

func getRealDestination(fs *afero.Afero, source string, destination string) (string, error) {
	var err error
	realDestination := destination

	if !path.IsAbs(realDestination) {
		realDestination, err = filepath.Abs(destination)
		if err != nil {
			return "", fmt.Errorf("getting absolute destination: %w", err)
		}
	}

	isDir, err := fs.IsDir(realDestination)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("checking if destination is directory: %w", err)
	}

	if isDir {
		return path.Join(realDestination, path.Base(source)), nil
	}

	return realDestination, nil
}

func getRealSource(deviceAddress string, source string) ([]string, error) {
	if path.Base(source) != "*" {
		return []string{source}, nil
	}

	g, err := glob.Compile(source)
	if err != nil {
		return []string{}, fmt.Errorf("compiling glob fn: %w", err)
	}

	parentDir := path.Dir(source)

	items, err := gopro.ListDirectory(deviceAddress, parentDir)
	if err != nil {
		return []string{}, fmt.Errorf("listing directory: %w", err)
	}

	sources := make([]string, 0)

	for _, item := range items {
		absoluteItem := path.Join(parentDir, item)

		if g.Match(absoluteItem) {
			sources = append(sources, absoluteItem)
		}
	}

	return sources, nil
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
