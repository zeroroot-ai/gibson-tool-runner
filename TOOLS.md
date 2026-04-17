# Tool backlog

Curated list of CLI tools worth adding as parsers to `gibson-tool-runner`.
Ordered by category, with a rough priority (🟢 first, 🟡 second, 🔵 long tail)
and a note on the expected DiscoveryResult shape.

Adding a tool = one parser file under `parsers/<tool>/`, one apt/curl line
in the Dockerfile, one blank-import in `cmd/gibson-runner/main.go`,
one golden fixture. Pattern is established; most take <200 LOC.

## Shipping today (v0.1)

- 🟢 nmap — Host/Port/Service
- 🟢 httpx — Endpoint/Service/Technology
- 🟢 nuclei — Finding (CVSS + CVE + MITRE)
- 🟢 subfinder — Domain/Subdomain
- 🟢 amass — Domain/Subdomain/Host (dedup by IP)
- 🟢 dnsx — Host (A/AAAA)
- 🟢 naabu — Host/Port
- 🟢 masscan — Host/Port (JSON-array)

## Network discovery / port scanning

- 🟡 rustscan — wraps nmap; Host/Port
- 🟡 unicornscan — Host/Port
- 🔵 zmap — Host/Port (very fast, /0 capable)
- 🔵 hping3 — individual Host/Port probes, mostly CLI only

## Subdomain / asset enumeration

- 🟡 assetfinder — Subdomain
- 🟡 chaos-client (PD) — Subdomain from PD hosted data
- 🟡 shuffledns (PD) — Subdomain via wildcard-aware DNS brute
- 🔵 findomain — Subdomain
- 🔵 sublist3r — Subdomain (Python; showing age)

## DNS

- 🟡 massdns — Host (bulk resolver, Port/Protocol from answer type)
- 🟡 dnsenum — Host/Subdomain
- 🟡 fierce — Subdomain + Host
- 🔵 dig / host — utility, use generic runner

## Web fingerprinting & crawling

- 🟢 katana (PD) — Endpoint (crawler JSON)
- 🟢 whatweb — Technology (tag list)
- 🟡 hakrawler — Endpoint
- 🟡 waybackurls — Endpoint (historical URLs)
- 🟡 gau (getallurls) — Endpoint
- 🟡 gospider — Endpoint
- 🔵 paramspider — Endpoint

## Content / directory discovery

- 🟢 ffuf — Endpoint (JSON)
- 🟡 feroxbuster — Endpoint (JSON)
- 🟡 gobuster — Endpoint + Subdomain (dns+web modes)
- 🟡 dirsearch — Endpoint (JSON)
- 🔵 dirb — Endpoint

## Vulnerability scanners

- 🟢 nikto — Finding (XML → parse)
- 🟢 wpscan — Finding (JSON, WordPress-specific)
- 🔵 joomscan — Finding (Joomla)
- 🔵 droopescan — Finding (Drupal)
- 🔵 cmseek — Finding (generic CMS)

## TLS / certificate analysis

- 🟢 sslyze — Certificate + Finding (JSON)
- 🟢 tlsx (PD) — Certificate (JSON-lines)
- 🟢 testssl.sh — Finding (JSON, very thorough)
- 🟡 sslscan — Certificate + Finding
- 🔵 openssl — utility, generic runner

## SQL injection

- 🟢 sqlmap — Finding (XML results dir or JSON)
- 🔵 nosqlmap — Finding (NoSQL variants)

## XSS / client injection

- 🟢 dalfox — Finding (JSON, XSS-specific)
- 🟡 xsstrike — Finding (JSON)
- 🔵 paramspider — Endpoint (pairs with XSS tools)

## API / GraphQL

- 🟡 kiterunner — Endpoint (API discovery)
- 🟡 graphw00f — Technology (GraphQL fingerprint)
- 🔵 graphql-cop — Finding (GraphQL audit)
- 🔵 clairvoyance — Endpoint (GraphQL schema extraction)

## Secrets / credential discovery (static)

- 🟢 trufflehog — Finding (high-value, CVE-ish signals for leaked creds)
- 🟢 gitleaks — Finding (git-aware secret scan)
- 🟡 noseyparker — Finding (new OSS, dedup-aware)
- 🔵 detect-secrets — Finding (Yelp)
- 🔵 scanrepo — Finding

## Code / SAST

- 🟢 semgrep — Finding (SAST, JSON output)
- 🟡 bandit — Finding (Python-specific)
- 🟡 gosec — Finding (Go-specific)
- 🔵 brakeman — Finding (Rails)

## Cloud / IaC

- 🟢 trivy — Finding (container + IaC + SBOM; large surface)
- 🟢 grype — Finding (container vuln)
- 🟡 prowler — Finding (multi-cloud posture)
- 🟡 scoutsuite — Finding (multi-cloud)
- 🟡 checkov — Finding (IaC misconfig)
- 🟡 tfsec — Finding (Terraform-specific)
- 🟡 kics — Finding (IaC, Checkmarx OSS)
- 🔵 pacu — AWS-specific; complex state

## Kubernetes

- 🟡 kube-hunter — Finding (K8s-specific)
- 🟡 kubesec — Finding (manifest audit)
- 🟡 kubeaudit — Finding (cluster audit)
- 🟡 kube-bench — Finding (CIS benchmark)
- 🔵 kubescape — Finding (NSA/CISA frameworks)

## Supply chain / SBOM

- 🟢 syft — Technology + Evidence (SBOM generation)
- 🟡 osv-scanner — Finding (OSS Vulnerability DB)
- 🟡 cosign — Evidence (signature verification)
- 🔵 sigstore — Evidence

## OSINT / reconnaissance

- 🟡 theHarvester — Host + Subdomain + Evidence (emails)
- 🟡 spiderfoot — broad; may be better as a plugin rather than tool
- 🔵 sherlock — Evidence (username enumeration)
- 🔵 photon — Endpoint + Evidence
- 🔵 recon-ng — broad framework; probably plugin-shaped

## Network brute force (online)

- 🟡 hydra — Finding (credential-spray results)
- 🟡 medusa — Finding
- 🔵 patator — Finding (modular)
- 🔵 crowbar — Finding (RDP/VNC)

## Packet capture & inspection

- 🟡 tcpdump — Evidence (pcap files)
- 🟡 tshark — Finding + Evidence (JSON output)
- 🔵 termshark — terminal wireshark; unlikely in automation

## Identity / auth

- 🟡 kerbrute — Finding (AD user enumeration)
- 🟡 ldapsearch — Evidence (generic runner)
- 🔵 impacket (Secretsdump / GetADUsers / GetNPUsers) — Finding (AD escalation)
- 🔵 bloodhound-python — Evidence (ingest into BloodHound)

## File / forensics (light, microVM-safe)

- 🟡 binwalk — Evidence (embedded file signatures)
- 🟡 strings — Evidence
- 🟡 exiftool — Evidence (metadata)
- 🔵 yara — Finding (IOC matching)
- 🔵 radare2 / rizin — Evidence (binary analysis; heavy)
- 🔵 steghide — Evidence (steg)

## Container-specific

- 🟡 dive — Evidence (layer analysis)
- 🔵 anchore-grype alias of grype above
- 🔵 skopeo — utility (image inspection)

## Generic / long tail

- 🟡 `kali.raw_exec` / `shell` — fallback parser that runs any CLI and
  emits a single Evidence node with truncated stdout. Covers the
  600-odd Kali tools nobody writes a typed parser for. Schema:
  `{ command: string[], stdin?: string, timeout?: int }`. Useful for
  one-off probes and ad-hoc commands agents construct.

## Operator tooling (typed where possible)

- 🟡 kubectl — Evidence (JSON output from `-o json`; most operator commands fit)
- 🟡 helm — Evidence
- 🟡 aws / gcloud / az CLIs — Evidence (generic; tool-specific parsers later)
- 🟡 terraform — Evidence (`state show` / `plan -json`)
- 🟡 jq / yq / gron — utility (use generic runner)
- 🟡 git — Evidence (commit history queries)
- 🟡 curl / httpie — utility; prefer httpx for probe shape

## Explicitly NOT to ship

- metasploit-framework — heavy Ruby + Postgres runtime; doesn't fit the
  one-shot microVM model. Better as its own long-lived plugin.
- john / hashcat (offline cracking) — GPU-bound; not orchestrator-shaped.
- Burp / ZAP / Maltego / Ghidra / Wireshark GUI — meaningless headless.
- aircrack-ng / kismet — need real wireless hardware.
- volatility — needs pre-captured memory images; pipeline-shaped.

## Priority v0.2 ship list (next 10)

Ordered by LLM-agent value × parser complexity:

1. ffuf
2. katana
3. sslyze
4. tlsx
5. trufflehog
6. gitleaks
7. trivy
8. dalfox
9. sqlmap
10. kali.raw_exec fallback

These bring the image to ~18 typed parsers + the raw-exec fallback,
covering ~80% of realistic agent workflows in recon, web probing,
TLS, secrets, and vulnerability scanning.
