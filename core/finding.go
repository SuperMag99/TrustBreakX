package core

type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
)

type Finding struct {
	ID         string   `json:"id"`
	Severity   Severity `json:"severity"`
	Category   string   `json:"category"`
	Title      string   `json:"title"`
	ObjectName string   `json:"object_name"`
	ObjectPath string   `json:"object_path"`
	TrustChain []string `json:"trust_chain"`
	Reasons    []string `json:"reasons"`
}
