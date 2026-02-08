package windows

import (
	"os/exec"
	"path/filepath"
	"strings"
)

type WritableResult struct {
	IsWritable bool
	Reasons    []string
}

func isLowPrivPrincipal(name string) (bool, string) {
	n := strings.ToLower(name)

	switch {
	case strings.Contains(n, "everyone"):
		return true, "Everyone"
	case strings.Contains(n, "authenticated users"):
		return true, "Authenticated Users"
	case strings.HasSuffix(n, "\\users"):
		return true, "Users"
	default:
		return false, ""
	}
}

func evaluatePath(path string) WritableResult {
	var res WritableResult

	path = NormalizePath(path)

	cmd := exec.Command("icacls", path)
	out, err := cmd.Output()
	if err != nil {
		return res
	}

	lines := strings.Split(string(out), "\n")

	for _, line := range lines {
		l := strings.TrimSpace(line)
		if l == "" {
			continue
		}

		// Example:
		// BUILTIN\Users:(I)(M)
		// NT AUTHORITY\Authenticated Users:(I)(F)

		parts := strings.SplitN(l, ":", 2)
		if len(parts) != 2 {
			continue
		}

		principal := strings.TrimSpace(parts[0])
		perms := parts[1]

		ok, label := isLowPrivPrincipal(principal)
		if !ok {
			continue
		}

		// Dangerous permissions
		if strings.Contains(perms, "(F)") ||
			strings.Contains(perms, "(M)") ||
			strings.Contains(perms, "(W)") {

			res.IsWritable = true
			res.Reasons = append(res.Reasons, label+" has write permissions")
		}
	}

	return res
}

// FILE + PARENT DIRECTORY (services & tasks)
func IsPathWritableByLowPriv(path string) (WritableResult, error) {
	fileRes := evaluatePath(path)
	if fileRes.IsWritable {
		return fileRes, nil
	}

	parent := filepath.Dir(path)
	dirRes := evaluatePath(parent)
	if dirRes.IsWritable {
		dirRes.Reasons = append(dirRes.Reasons, "Parent directory writable")
		return dirRes, nil
	}

	return WritableResult{IsWritable: false}, nil
}

// DIRECTORY ONLY (PATH hijacking)
func IsDirectoryWritableByLowPriv(dir string) (WritableResult, error) {
	return evaluatePath(dir), nil
}

func NormalizePath(p string) string {
	p = strings.TrimSpace(p)
	p = strings.Trim(p, "\"")
	return p
}
