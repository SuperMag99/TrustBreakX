package core

import (
	"encoding/json"
	"os"
)

type Engine struct {
	findings map[string]Finding
}

func NewEngine() *Engine {
	return &Engine{
		findings: make(map[string]Finding),
	}
}

func (e *Engine) AddFinding(f Finding) {
	e.findings[f.ID] = f
}

func (e *Engine) List() []Finding {
	out := make([]Finding, 0, len(e.findings))
	for _, f := range e.findings {
		out = append(out, f)
	}
	return out
}

func (e *Engine) ExportJSON(path string) error {
	data, err := json.MarshalIndent(e.List(), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
