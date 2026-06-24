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

// Per-tool args allowlist for masscan.
//
// Coverage rationale:
//   - Output flags (`-oN`, `-oX`, `-oJ`, `-oG`, `--output-filename`) DENIED.
//     The runner pins `-oJ -` so JSON reaches stdout; allowing caller-
//     supplied output paths is the canonical example of arbitrary file
//     write via tool injection.
//   - Throughput / scope flags safe.
package masscan

import (
	"github.com/zeroroot-ai/gibson-executor/internal/policy"
	"github.com/zeroroot-ai/gibson-executor/internal/registry"
)

var argsPolicy = policy.ArgsPolicy{
	"-p":          policy.AllowAny,
	"--ports":     policy.AllowAny,
	"--rate":      policy.AllowAny,
	"--top-ports": policy.AllowAny,
	"--exclude":   policy.AllowAny,
	"--banners":   nil,
	"--ping":      nil,
	"-v":          nil,
	"-vv":         nil,
	"--retries":   policy.AllowAny,
}

func init() {
	registry.RegisterArgsPolicy(toolName, argsPolicy)
}
