package gopro

import (
	"fmt"
	"net/http"
)

func ListDirectory(hostAddress string, targetPath string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/gopro/media/list", hostAddress))
	if err != nil {
		return nil, fmt.Errorf("calling API: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil, nil
}
