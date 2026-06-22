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

// Per-tool args allowlist for amass.
//
// Coverage rationale:
//   - Subcommand "enum" plus -passive / -json /dev/stdout already pinned.
//   - Output-redirect flags (`-o`, `-dir`) DENIED — runner reads JSON on
//     stdout. The amass `-config` flag is DENIED to prevent loading
//     attacker-controlled configurations.
package amass

import (
	"github.com/zeroroot-ai/gibson-tool-runner/internal/policy"
	"github.com/zeroroot-ai/gibson-tool-runner/internal/registry"
)

var argsPolicy = policy.ArgsPolicy{
	"-passive": nil,
	"-active":  nil,
	"-brute":   nil,
	"-w":       policy.AllowAny, // wordlist (still requires path validator
	// to be added before use; keep narrow)
	"-timeout":           policy.AllowAny,
	"-max-depth":         policy.AllowAny,
	"-min-for-recursive": policy.AllowAny,
	"-include":           policy.AllowAny,
	"-exclude":           policy.AllowAny,
	"-silent":            nil,
	"-nolocaldb":         nil,
}

func init() {
	registry.RegisterArgsPolicy(toolName, argsPolicy)
}
