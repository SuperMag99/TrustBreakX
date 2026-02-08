# TrustBreakX 🔓

**TrustBreakX** is a professional-grade Windows security assessment tool designed to identify **real SYSTEM-level trust breaks** caused by misconfigured execution paths.  
It focuses on **practical privilege escalation opportunities**, not theoretical vulnerabilities or missing patches.

TrustBreakX helps security engineers, red teams, blue teams, and defenders quickly answer one critical question:

> Can a low-privileged user realistically gain SYSTEM execution on this machine?

![TrustBreakX Hero Screenshot](https://github.com/SuperMag99/TrustBreakX/Screenshot%1.png)

---

## 🎯 Features

- **SYSTEM Service Trust Analysis**  
  Detects Windows services running as `LocalSystem` that execute binaries from user-writable paths.

- **SYSTEM Scheduled Task Analysis**  
  Identifies scheduled tasks running as `SYSTEM` with writable execution paths.

- **SYSTEM PATH Hijacking Detection**  
  Finds writable directories in the SYSTEM execution PATH that allow binary planting and command hijacking.

- **High-Confidence Findings Only**  
  No CVE scanning, no patch checks, no noisy ACL dumps — only real, exploitable trust relationships.

- **CLI-Friendly + Automation Ready**  
  Clean terminal output with structured JSON export for reporting and automation.

---

## 🚀 Getting Started

Follow the steps below to run TrustBreakX locally on a Windows system.

### 1. Download the Repository

Clone the repository using Git:
```powershell
git clone https://github.com/SuperMag99/TrustBreakX.git
cd TrustBreakX

### 2. Install Go (If Not Installed)

TrustBreakX is written in Go.

Download Go from:
https://go.dev/dl/

Install **Go 1.21 or newer**.

### 3. Run TrustBreakX

⚠️ PowerShell must be run as **Administrator**.

From the project root directory:
```powershell
go run ./cmd/trustbreakx

---

## 🧭 Security and Vulnerabilities

- **Security:** Refer to [SECURITY.md](./SECURITY.md).

## ⚖️ Legal & Ethical Disclaimer

This tool is intended for defensive security, auditing, penetration testing, and educational use only.
TrustBreakX does not exploit vulnerabilities, perform attacks, or modify system state.
It only analyzes existing Windows trust relationships and execution paths.
Users are solely responsible for ensuring their use complies with local laws and organizational policies.

## License Summary

This project is licensed under a **Non-Commercial Attribution License**. Key points:

1. ✅ **Free to use for personal, educational, and research purposes.**
2. ✅ **Any modification or derivative work must credit to the author.
3. ❌ **Commercial use, sale, licensing, or any use intended to generate revenue is strictly prohibited without prior written permission.**
4. ⚠️ **No warranty**: Use at your own risk.
5. ⚖️ **Legal protection**: Unauthorized commercial use or failure to credit the author may result in legal action.

For full license details, see the `LICENSE` file. [LICENSE](./LICENSE).

---

## 👤 Maintainer

🔗 **GitHub**: [https://github.com/SuperMag99](https://github.com/SuperMag99)  
🔗 **LinkedIn**: [https://www.linkedin.com/in/mag99/](https://www.linkedin.com/in/mag99/)

---

*All trademarks and service names mentioned in this project are the property of their respective owners.*
