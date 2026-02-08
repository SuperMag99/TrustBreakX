package windows

import (
	"os"
	"strings"
)

type PathEntry struct {
	Directory string
}

func CollectSystemPATH() []PathEntry {
	sysPath := os.Getenv("PATH")
	parts := strings.Split(sysPath, ";")

	seen := make(map[string]bool)
	var paths []PathEntry

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" || strings.Contains(p, "%") {
			continue
		}

		p = NormalizePath(p)
		info, err := os.Stat(p)
		if err != nil || !info.IsDir() {
			continue
		}

		l := strings.ToLower(p)
		if l == `c:\windows` || l == `c:\windows\system32` {
			continue
		}

		if seen[l] {
			continue
		}
		seen[l] = true

		paths = append(paths, PathEntry{Directory: p})
	}

	return paths
}
