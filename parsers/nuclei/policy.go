// Copyright 2026 zero-day.ai
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Per-tool args allowlist for nuclei.
//
// Coverage rationale:
//   - Severity, tag, type filters: documented and safe.
//   - Template selection (`-t`/`--templates`) is DENIED here because the
//     runner's nuclei.go already pipes templates from req.Options
//     ("templates"). A caller-supplied -t would let an attacker reference
//     templates outside the runner's pinned template directory and
//     potentially execute arbitrary nuclei templates the operator has
//     not vetted.
//   - Output flags (`-o`, `-output`, `-store-resp`, `-store-resp-dir`)
//     are DENIED — the runner consumes JSONL on stdout and never writes
//     to disk on the caller's behalf.
package nuclei

import (
	"github.com/zeroroot-ai/gibson-executor/internal/policy"
	"github.com/zeroroot-ai/gibson-executor/internal/registry"
)

var argsPolicy = policy.ArgsPolicy{
	// Filtering by severity / tags / template type.
	"-severity": policy.AllowAny,
	"-s":        policy.AllowAny,
	"-tags":     policy.AllowAny,
	"-itags":    policy.AllowAny,
	"-etags":    policy.AllowAny,
	"-type":     policy.AllowAny,
	"-author":   policy.AllowAny,

	// Throughput controls.
	"-rate-limit":  policy.AllowAny,
	"-bulk-size":   policy.AllowAny,
	"-concurrency": policy.AllowAny,
	"-c":           policy.AllowAny,
	"-timeout":     policy.AllowAny,
	"-retries":     policy.AllowAny,

	// Boolean diagnostics that don't affect security posture.
	"-stats":                nil,
	"-silent":               nil,
	"-no-color":             nil,
	"-disable-update-check": nil,
	"-include-rr":           nil,
}

func init() {
	registry.RegisterArgsPolicy(toolName, argsPolicy)
}
