package host

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

type HostDetails struct {
	Hostname     string
	OS           string
	Architecture string
	Interfaces   []NetworkInterface
}

type NetworkInterface struct {
	Name string
	MAC  string
}

func GetHostDetails() (*HostDetails, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("failed to get hostname: %w", err)
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get network interfaces: %w", err)
	}

	var nets []NetworkInterface
	for _, iface := range interfaces {
		// Skip interfaces with empty MACs
		if iface.HardwareAddr.String() == "" {
			continue
		}
		nets = append(nets, NetworkInterface{
			Name: iface.Name,
			MAC:  iface.HardwareAddr.String(),
		})
	}

	return &HostDetails{
		Hostname:     hostname,
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		Interfaces:   nets,
	}, nil
}
