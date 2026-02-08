TrustBreakX 🔓

TrustBreakX is a professional-grade Windows security assessment tool designed to identify
real SYSTEM-level trust breaks caused by misconfigured execution paths.

It focuses on practical privilege escalation opportunities, not theoretical vulnerabilities
or missing patches.

TrustBreakX helps security engineers, red teams, blue teams, and defenders quickly answer one critical question:

Can a low-privileged user realistically gain SYSTEM execution on this machine?

🎯 Features

SYSTEM Service Trust Analysis
Detects Windows services running as LocalSystem that execute binaries from user-writable paths.

SYSTEM Scheduled Task Analysis
Identifies scheduled tasks running as SYSTEM with writable execution paths.

SYSTEM PATH Hijacking Detection
Finds writable directories in the SYSTEM execution PATH that allow binary planting and command hijacking.

High-Confidence Findings Only
No CVE scanning, no patch checks, no noisy ACL dumps — only real, exploitable trust relationships.

CLI-Friendly and Automation Ready
Clean terminal output with structured JSON export for reporting and automation.

🚀 Getting Started

Follow the steps below to run TrustBreakX locally on a Windows system.

1. Download the Repository

Clone the repository using Git:

git clone https://github.com/SuperMag99/TrustBreakX.git
cd TrustBreakX

2. Install Go (If Not Installed)

TrustBreakX is written in Go.

Download Go from:

https://go.dev/dl/


Install Go version 1.21 or newer.

3. Run TrustBreakX

PowerShell must be run as Administrator.

From the project root directory, run:

go run ./cmd/trustbreakx


Findings will be printed to the console and exported as JSON in the output directory.

🧭 Security and Vulnerabilities

Security policy and vulnerability reporting instructions are available in:

SECURITY.md

⚖️ Legal and Ethical Disclaimer

This tool is intended for defensive security, auditing, and educational use only.

TrustBreakX does not exploit vulnerabilities, perform attacks, or modify system state.
It only analyzes existing Windows trust relationships and execution paths.

Users are responsible for ensuring their usage complies with organizational policies and local laws.

📜 License Summary

This project is licensed under a Non-Commercial Attribution License.

Key points:

Free to use for personal, educational, and research purposes

Any modification or derivative work must credit the author

Commercial use, sale, or monetization is prohibited without written permission

No warranty — use at your own risk

Unauthorized commercial use may result in legal action

For full license details, see:

LICENSE

👤 Maintainer

Mohammad Ali Ghanem

GitHub: https://github.com/SuperMag99

LinkedIn: https://www.linkedin.com/in/mag99/
