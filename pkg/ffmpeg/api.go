package ffmpeg

import (
	"bytes"
	"fmt"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

type ExposeOptions struct {
	DeviceIndex int
	Username    string
}

func Expose(opts ExposeOptions) error {
	cmd := exec.Command("ffmpeg",
		"-nostdin",
		"-threads", "1",
		"-i", "udp://@:8554?overrun_nonfatal=1&fifo_size=50000000",
		"-f:v", "mpegts",
		"-fflags", "nobuffer",
		"-vf", "format=yuv420p",
		"-f", "v4l2",
		fmt.Sprintf("/dev/video%d", opts.DeviceIndex),
	)

	credentials, err := userCredentials(opts.Username)
	if err != nil {
		return fmt.Errorf("acquiring user credentials: %w", err)
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{Credential: &credentials}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("%s: %w", stderr.String(), err)

		return fmt.Errorf("running command: %w", err)
	}

	return nil
}

func userCredentials(name string) (syscall.Credential, error) {
	targetUser, err := user.Lookup(name)
	if err != nil {
		return syscall.Credential{}, fmt.Errorf("getting user %s: %w", name, err)
	}

	uid, err := strconv.ParseUint(targetUser.Uid, 10, 32)
	if err != nil {
		return syscall.Credential{}, fmt.Errorf("parsing uid: %w", err)
	}

	gid, err := strconv.ParseUint(targetUser.Gid, 10, 32)
	if err != nil {
		return syscall.Credential{}, fmt.Errorf("parsing gid: %w", err)
	}

	return syscall.Credential{
		Uid: uint32(uid),
		Gid: uint32(gid),
	}, nil
}
