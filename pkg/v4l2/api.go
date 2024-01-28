package v4l2

import (
	"bytes"
	"fmt"
	"os/exec"
)

type InjectModuleOptions struct {
	Label       string
	DeviceIndex int
}

func InjectModule(opts InjectModuleOptions) error {
	cmd := exec.Command(
		"modprobe",
		"v4l2loopback",
		"exclusive_caps=1",
		fmt.Sprintf("card_label='%s'", opts.Label),
		fmt.Sprintf("video_nr=%d", opts.DeviceIndex),
	)

	var stderr bytes.Buffer

	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		err = fmt.Errorf("%s: %w", stderr.String(), err)

		return fmt.Errorf("running command: %w", err)
	}

	return nil
}

func EjectModule() error {
	cmd := exec.Command("modprobe", "-rf", "v4l2loopback")

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("running command: %w", err)
	}

	return nil
}
