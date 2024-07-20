package rm

import (
	"fmt"
	"path"

	"github.com/deifyed/gpctl/pkg/gopro"
	"github.com/gobwas/glob"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

type Options struct {
	TargetDeviceIndex int
	Progressbar       bool
}

func RunE(opts *Options) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		deviceAddress, err := gopro.GetDeviceAddressByIndex(cmd.Context(), opts.TargetDeviceIndex)
		if err != nil {
			return fmt.Errorf("getting device address: %w", err)
		}

		realSources, err := getRealSource(deviceAddress, targetPath)
		if err != nil {
			return fmt.Errorf("getting real source: %w", err)
		}

		var progressBar progressBarI = &dummyProgressbar{}

		if opts.Progressbar {
			progressBar = progressbar.Default(int64(len(realSources)))
		}

		for _, sourceItem := range realSources {
			err = gopro.RemoveFile(deviceAddress, sourceItem)
			if err != nil {
				return fmt.Errorf("removing file: %w", err)
			}

			progressBar.Add(1)
		}

		return nil
	}
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

type progressBarI interface {
	Add(int) error
}

type dummyProgressbar struct{}

func (dummyProgressbar) Add(int) error { return nil }
