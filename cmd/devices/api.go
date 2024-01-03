package devices

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/deifyed/gpctl/pkg/gopro"
	"github.com/spf13/cobra"
)

var styling = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))

func RunE() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		devices, err := gopro.GetDeviceAddresses(cmd.Context())
		if err != nil {
			return fmt.Errorf("getting device addresses: %w", err)
		}

		for index, device := range devices {
			fmt.Printf("%s - %s\n", styling.Render(fmt.Sprintf("%d", index)), styling.Render(device))
		}

		return nil
	}
}
