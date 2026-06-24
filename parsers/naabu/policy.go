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

// Per-tool args allowlist for naabu.
//
// Coverage rationale:
//   - Port spec / scan-mode flags safe and documented.
//   - Output flags (`-o`, `-output`, `-csv`, `-json` already pinned) are
//     DENIED. The runner consumes JSON on stdout via -json -silent.
package naabu

import (
	"github.com/zeroroot-ai/gibson-executor/internal/policy"
	"github.com/zeroroot-ai/gibson-executor/internal/registry"
)

var argsPolicy = policy.ArgsPolicy{
	// Port specification.
	"-p":             policy.AllowAny,
	"-top-ports":     policy.AllowAny,
	"-exclude-ports": policy.AllowAny,

	// Scan-mode flags.
	"-scan-type": policy.AllowEnum("s", "c", "syn", "connect"),
	"-s":         policy.AllowEnum("s", "c", "syn", "connect"),
	"-Pn":        nil,
	"-sn":        nil,

	// Throughput.
	"-rate":    policy.AllowAny,
	"-c":       policy.AllowAny,
	"-timeout": policy.AllowAny,
	"-retries": policy.AllowAny,

	// Boolean diagnostics.
	"-stats":  nil,
	"-silent": nil,
	"-verify": nil,
}

func init() {
	registry.RegisterArgsPolicy(toolName, argsPolicy)
}
