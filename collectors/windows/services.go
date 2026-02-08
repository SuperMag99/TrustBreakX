package windows

import (
	"os/exec"
	"strings"
)

type ServiceInfo struct {
	Name       string
	RunAs      string
	BinaryPath string
}

func CollectServices() ([]ServiceInfo, error) {
	// Step 1: enumerate services
	cmd := exec.Command("sc", "query", "state=", "all")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	var services []ServiceInfo

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "SERVICE_NAME:") {
			continue
		}

		name := strings.TrimSpace(strings.TrimPrefix(line, "SERVICE_NAME:"))

		// Step 2: query service config
		qc := exec.Command("sc", "qc", name)
		cfg, err := qc.Output()
		if err != nil {
			continue
		}

		var svc ServiceInfo
		svc.Name = name

		for _, l := range strings.Split(string(cfg), "\n") {
			l = strings.TrimSpace(l)

			if strings.HasPrefix(l, "BINARY_PATH_NAME") {
				svc.BinaryPath = NormalizePath(strings.SplitN(l, ":", 2)[1])
			}

			if strings.HasPrefix(l, "SERVICE_START_NAME") {
				svc.RunAs = strings.TrimSpace(strings.SplitN(l, ":", 2)[1])
			}
		}

		if svc.BinaryPath != "" && svc.RunAs != "" {
			services = append(services, svc)
		}
	}

	return services, nil
}
