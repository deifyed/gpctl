package gopro

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"
)

func ListDirectory(hostAddress string, targetPath string) ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/gopro/media/list", hostAddress))
	if err != nil {
		return nil, fmt.Errorf("calling API: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var listDirectoryResponse listDirectoryResponse

	err = json.Unmarshal(raw, &listDirectoryResponse)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling response: %w", err)
	}

	fs, err := modelFilesystem(listDirectoryResponse)
	if err != nil {
		return nil, fmt.Errorf("modeling filesystem: %w", err)
	}

	contents, err := fs.ReadDir(targetPath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %s: %w", targetPath, err)
	}

	return fileInfoAsStrings(contents), nil
}

type listDirectoryResponseDirectoryFile struct {
	Name string `json:"n"`
}

type listDirectoryResponseDirectory struct {
	Directory string                               `json:"d"`
	Files     []listDirectoryResponseDirectoryFile `json:"fs"`
}

type listDirectoryResponse struct {
	ID    string                           `json:"id"`
	Media []listDirectoryResponseDirectory `json:"media"`
}

func fileInfoAsStrings(fileInfos []fs.FileInfo) []string {
	var result []string

	for _, fileInfo := range fileInfos {
		result = append(result, fileInfo.Name())
	}

	return result
}

func ReadFile(hostAddress string, targetPath string) (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/videos/DCIM%s", hostAddress, targetPath))
	if err != nil {
		return nil, fmt.Errorf("calling API: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}
	return resp.Body, nil
}

func RemoveFile(hostAddress string, targetPath string) error {
	targetPath = strings.TrimLeft(targetPath, "/")

	resp, err := http.Get(fmt.Sprintf("http://%s/gopro/media/delete/file?path=%s", hostAddress, targetPath))
	if err != nil {
		return fmt.Errorf("calling API: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	return nil
}
