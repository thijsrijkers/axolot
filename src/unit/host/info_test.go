package host

import (
	"testing"
	"axolot/src/host"
)

func TestGetHostDetails(t *testing.T) {
	info, err := host.GetHostDetails()
	if err != nil {
		t.Fatalf("GetHostDetails returned an error: %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}

	if info.OS == "" {
		t.Error("Expected OS to be non-empty")
	}

	if info.Architecture == "" {
		t.Error("Expected Architecture to be non-empty")
	}

	for _, iface := range info.Interfaces {
		if iface.Name == "" {
			t.Error("Expected network interface name to be non-empty")
		}
		if iface.MAC == "" {
			t.Error("Expected MAC address to be non-empty for interface with entry")
		}
	}
}
