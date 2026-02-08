package windows

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type TaskInfo struct {
	Name       string
	RunAs      string
	ActionPath string
}

func isExecutable(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".exe", ".bat", ".cmd", ".ps1":
		return true
	default:
		return false
	}
}

func resolveActionPath(raw string) string {
	raw = NormalizePath(raw)
	if raw == "" {
		return ""
	}

	if strings.HasPrefix(raw, "\"") {
		if end := strings.Index(raw[1:], "\""); end >= 0 {
			raw = raw[1 : end+1]
		}
	}

	expanded := os.ExpandEnv(raw)
	parts := strings.Fields(expanded)

	for i := len(parts); i > 0; i-- {
		candidate := strings.Join(parts[:i], " ")
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() && isExecutable(candidate) {
			return candidate
		}
	}

	return ""
}

func CollectScheduledTasks() ([]TaskInfo, error) {
	cmd := exec.Command("schtasks.exe", "/query", "/fo", "LIST", "/v")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	var tasks []TaskInfo
	var cur TaskInfo

	flush := func() {
		if cur.Name != "" {
			cur.ActionPath = resolveActionPath(cur.ActionPath)
			if cur.ActionPath != "" {
				tasks = append(tasks, cur)
			}
		}
		cur = TaskInfo{}
	}

	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		switch key {
		case "TaskName":
			if cur.Name != "" {
				flush()
			}
			cur.Name = val
		case "Run As User":
			cur.RunAs = strings.ToUpper(val)
		case "Task To Run":
			cur.ActionPath = val
		}
	}

	flush()
	return tasks, nil
}
