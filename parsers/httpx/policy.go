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

// Per-tool args allowlist for httpx.
//
// Coverage rationale:
//   - Detection / filter flags safe and documented.
//   - Output flags (`-o`, `-output`, `-store-response*`, `-srd`) DENIED.
//   - Probe-modification flags that allow caller-controlled arbitrary
//     payloads (`-X` custom method, `-body`) are allowed but value-bound
//     so callers cannot inject newline-separated multi-request payloads.
package httpx

import (
	"fmt"
	"strings"

	"github.com/zeroroot-ai/gibson-executor/internal/policy"
	"github.com/zeroroot-ai/gibson-executor/internal/registry"
)

// noNewlines rejects values containing '\n' / '\r' (HTTP request smuggling
// guard: httpx threads `-body` directly into the request body).
func noNewlines(value string) error {
	if strings.ContainsAny(value, "\r\n") {
		return fmt.Errorf("value must not contain newlines")
	}
	return nil
}

var argsPolicy = policy.ArgsPolicy{
	// Probe-shape flags.
	"-method":         policy.AllowEnum("GET", "POST", "HEAD", "OPTIONS", "PUT", "DELETE"),
	"-X":              policy.AllowEnum("GET", "POST", "HEAD", "OPTIONS", "PUT", "DELETE"),
	"-path":           noNewlines,
	"-paths":          noNewlines,
	"-ports":          policy.AllowAny,
	"-p":              policy.AllowAny,
	"-body":           noNewlines,
	"-H":              noNewlines,
	"-status-code":    policy.AllowAny,
	"-sc":             nil,
	"-content-length": nil,
	"-cl":             nil,
	"-content-type":   nil,
	"-ct":             nil,
	"-title":          nil,
	"-tech-detect":    nil,
	"-td":             nil,
	"-web-server":     nil,
	"-server":         nil,

	// Filtering.
	"-mc": policy.AllowAny,
	"-fc": policy.AllowAny,
	"-ms": policy.AllowAny,
	"-fs": policy.AllowAny,

	// Throughput.
	"-threads":    policy.AllowAny,
	"-t":          policy.AllowAny,
	"-rate-limit": policy.AllowAny,
	"-rl":         policy.AllowAny,
	"-timeout":    policy.AllowAny,
	"-retries":    policy.AllowAny,

	// Boolean diagnostics.
	"-silent":           nil,
	"-no-color":         nil,
	"-stats":            nil,
	"-follow-redirects": nil,
	"-fr":               nil,
}

func init() {
	registry.RegisterArgsPolicy(toolName, argsPolicy)
}
