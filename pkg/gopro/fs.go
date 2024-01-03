package gopro

import (
	"fmt"

	"github.com/spf13/afero"
)

func modelFilesystem(resp listDirectoryResponse) (*afero.Afero, error) {
	fs := afero.Afero{Fs: afero.NewMemMapFs()}

	for _, directory := range resp.Media {
		currentDir := directory.Directory

		for _, f := range directory.Files {
			fs.WriteFile(fmt.Sprintf("/%s/%s", currentDir, f.Name), []byte(""), 0644)
		}
	}

	return &fs, nil
}
