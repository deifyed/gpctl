package gopro

import (
	"context"
	"errors"
	"fmt"
	"net"
	"regexp"
	"time"
)

var ErrNoDevicesFound = errors.New("no devices found")

// Borrowed with love and modifications from https://github.com/KonradIT/mmt
var deviceAddressRegexp = regexp.MustCompile(`172.2\d.\d\d\d.5\d`)

func GetDeviceAddresses(ctx context.Context) ([]string, error) {
	ctx, cancelCtx := context.WithTimeout(ctx, 2*time.Second)
	defer cancelCtx()

	ipsFound := []string{}

	ifaces, err := net.Interfaces()
	if err != nil {
		return ipsFound, err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, a := range addrs {
			ipv4Addr := a.(*net.IPNet).IP.To4()

			if deviceAddressRegexp.MatchString(ipv4Addr.String()) {
				correctIP := ipv4Addr.String()[:len(ipv4Addr.String())-1] + "1"

				ipsFound = append(ipsFound, correctIP)
			}
		}
	}

	return ipsFound, nil
}

func GetDeviceAddressByIndex(ctx context.Context, index int) (string, error) {
	devices, err := GetDeviceAddresses(ctx)
	if err != nil {
		return "", fmt.Errorf("getting device addresses: %w", err)
	}

	if len(devices) == 0 {
		return "", ErrNoDevicesFound
	}

	return devices[index], nil
}
