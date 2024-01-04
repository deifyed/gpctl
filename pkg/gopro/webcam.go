package gopro

import (
	"fmt"
	"net/http"
)

func StartWebcam(hostAddress string) error {
	resp, err := http.Get(fmt.Sprintf("http://%s/gopro/webcam/start", hostAddress))
	if err != nil {
		return fmt.Errorf("calling API: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}

func StopWebcam(hostAddress string) error {
	resp, err := http.Get(fmt.Sprintf("http://%s/gopro/webcam/stop", hostAddress))
	if err != nil {
		return fmt.Errorf("calling API: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}

func ExitWebcamMode(hostAddress string) error {
	resp, err := http.Get(fmt.Sprintf("http://%s/gopro/webcam/exit", hostAddress))
	if err != nil {
		return fmt.Errorf("calling API: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
