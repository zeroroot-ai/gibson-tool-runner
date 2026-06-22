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

// Per-tool args allowlist for subfinder.
//
// Coverage rationale:
//   - Source-control flags safe.
//   - Output flags (`-o`) DENIED — runner consumes JSON on stdout.
//   - Resolver-list (`-r`, `-rl`) DENIED — see dnsx policy rationale.
package subfinder

import (
	"github.com/zeroroot-ai/gibson-tool-runner/internal/policy"
	"github.com/zeroroot-ai/gibson-tool-runner/internal/registry"
)

var argsPolicy = policy.ArgsPolicy{
	"-sources":         policy.AllowAny,
	"-s":               policy.AllowAny,
	"-exclude-sources": policy.AllowAny,
	"-es":              policy.AllowAny,
	"-all":             nil,
	"-recursive":       nil,
	"-active":          nil,
	"-timeout":         policy.AllowAny,
	"-rate-limit":      policy.AllowAny,
	"-rl":              policy.AllowAny,
	"-t":               policy.AllowAny,
	"-silent":          nil,
	"-stats":           nil,
	"-no-color":        nil,
}

func init() {
	registry.RegisterArgsPolicy(toolName, argsPolicy)
}
