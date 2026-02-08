package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"trustbreakx/collectors/windows"
	"trustbreakx/core"
)

func isSystemAccount(runAs string) bool {
	r := strings.ToUpper(strings.TrimSpace(runAs))
	return r == "SYSTEM" || (strings.Contains(r, "SYSTEM") && strings.Contains(r, "AUTHORITY"))
}

func severityRank(s core.Severity) int {
	switch s {
	case core.SeverityCritical:
		return 1
	default:
		return 99
	}
}

func main() {
	core.PrintBanner()

	engine := core.NewEngine()

	// SERVICES
	services, _ := windows.CollectServices()
	for _, s := range services {
		if strings.ToUpper(strings.TrimSpace(s.RunAs)) != "LOCALSYSTEM" {
			continue
		}

		res, err := windows.IsPathWritableByLowPriv(s.BinaryPath)
		if err == nil && res.IsWritable {
			engine.AddFinding(core.Finding{
				ID:         "service|" + s.Name + "|" + s.BinaryPath,
				Severity:   core.SeverityCritical,
				Category:   "Service Trust",
				Title:      "SYSTEM service trusts writable binary",
				ObjectName: s.Name,
				ObjectPath: s.BinaryPath,
				TrustChain: []string{"Low-priv User", "Writable Path", "SYSTEM Service"},
				Reasons:    res.Reasons,
			})
		}
	}

	// TASKS
	tasks, _ := windows.CollectScheduledTasks()
	for _, t := range tasks {
		if !isSystemAccount(t.RunAs) {
			continue
		}

		res, err := windows.IsPathWritableByLowPriv(t.ActionPath)
		if err == nil && res.IsWritable {
			engine.AddFinding(core.Finding{
				ID:         "task|" + t.Name + "|" + t.ActionPath,
				Severity:   core.SeverityCritical,
				Category:   "Scheduled Task Trust",
				Title:      "SYSTEM scheduled task trusts writable action",
				ObjectName: t.Name,
				ObjectPath: t.ActionPath,
				TrustChain: []string{"Low-priv User", "Writable Path", "SYSTEM Task"},
				Reasons:    res.Reasons,
			})
		}
	}

	// PATH
	for _, p := range windows.CollectSystemPATH() {
		res, err := windows.IsDirectoryWritableByLowPriv(p.Directory)
		if err == nil && res.IsWritable {
			engine.AddFinding(core.Finding{
				ID:         "path|" + p.Directory,
				Severity:   core.SeverityCritical,
				Category:   "PATH Trust",
				Title:      "Writable SYSTEM PATH directory can hijack execution",
				ObjectName: "PATH",
				ObjectPath: p.Directory,
				TrustChain: []string{"Low-priv User", "Writable PATH directory", "SYSTEM execution"},
				Reasons:    res.Reasons,
			})
		}
	}

	findings := engine.List()

	// 🔥 SEVERITY SORT
	sort.SliceStable(findings, func(i, j int) bool {
		return severityRank(findings[i].Severity) < severityRank(findings[j].Severity)
	})

	// PRINT FINDINGS
	categoryCount := map[string]int{}
	severityCount := map[core.Severity]int{}

	for _, f := range findings {
		severityCount[f.Severity]++
		categoryCount[f.Category]++

		fmt.Println(core.Critical(
			fmt.Sprintf("[%s] %s %s", f.Severity, core.SymCritical, f.Title),
		))
		fmt.Println(core.Dim(" Category: ") + f.Category)

		switch f.Category {
		case "Service Trust":
			fmt.Println(core.Dim(" Object  : ") + core.SymService + " " + f.ObjectName)
		case "Scheduled Task Trust":
			fmt.Println(core.Dim(" Object  : ") + core.SymTask + " " + f.ObjectName)
		case "PATH Trust":
			fmt.Println(core.Dim(" Object  : ") + core.SymPath + " PATH")
		}

		fmt.Println(core.Dim(" Path    : ") + f.ObjectPath)
		fmt.Println(core.Dim(" Trust   : ") + strings.Join(f.TrustChain, " "+core.SymTrust+" "))
		fmt.Println(core.Dim(" Reasons:"))
		for _, r := range f.Reasons {
			fmt.Println("   -", r)
		}
		fmt.Println()
	}

	// 📊 SUMMARY FOOTER
	fmt.Println(core.Info("Summary"))
	fmt.Println(core.Dim("────────────────────────────"))
	fmt.Println(" Total Findings :", len(findings))
	for sev, c := range severityCount {
		fmt.Println(" ", sev, ":", c)
	}
	for cat, c := range categoryCount {
		fmt.Println(" ", cat, ":", c)
	}
	fmt.Println()

	ts := time.Now().Format("2006-01-02_15-04-05")
	out := fmt.Sprintf("output/trustbreakx_%s.json", ts)
	engine.ExportJSON(out)

	fmt.Println(core.Dim("[+] Findings exported to ") + out)
}
