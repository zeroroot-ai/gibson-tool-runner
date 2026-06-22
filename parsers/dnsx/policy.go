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

// Per-tool args allowlist for dnsx.
//
// Coverage rationale:
//   - Record-type selectors are safe.
//   - Output flags (`-o`, `-output`) DENIED — runner reads JSON on stdout.
//   - Resolver-list (`-r`) DENIED unless paired with PathUnder validator;
//     allowing arbitrary resolver lists could exfiltrate via DNS to a
//     caller-controlled server. Today no caller needs to override the
//     default resolvers, so this is left out of the allowlist.
package dnsx

import (
	"github.com/zeroroot-ai/gibson-tool-runner/internal/policy"
	"github.com/zeroroot-ai/gibson-tool-runner/internal/registry"
)

var argsPolicy = policy.ArgsPolicy{
	// Record-type selectors (boolean form).
	"-a":     nil,
	"-aaaa":  nil,
	"-cname": nil,
	"-ns":    nil,
	"-txt":   nil,
	"-mx":    nil,
	"-soa":   nil,
	"-ptr":   nil,
	"-srv":   nil,
	"-caa":   nil,
	"-axfr":  nil,

	// Throughput / behaviour.
	"-rate-limit": policy.AllowAny,
	"-rl":         policy.AllowAny,
	"-c":          policy.AllowAny,
	"-t":          policy.AllowAny,
	"-retry":      policy.AllowAny,

	// Boolean diagnostics.
	"-silent":    nil,
	"-stats":     nil,
	"-resp":      nil,
	"-resp-only": nil,
}

func init() {
	registry.RegisterArgsPolicy(toolName, argsPolicy)
}
