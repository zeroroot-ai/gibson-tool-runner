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

package nmap

import (
	"strings"
	"testing"

	"github.com/zeroroot-ai/gibson-tool-runner/internal/policy"
	"github.com/zeroroot-ai/gibson-tool-runner/internal/registry"
)

// TestPolicy_DropsOutputFile_ON asserts that the canonical injection
// `-oN /etc/passwd` is rejected: the runner already pins -oX -, so allowing
// caller-supplied output flags is an arbitrary-file-write hole.
func TestPolicy_DropsOutputFile_ON(t *testing.T) {
	p, ok := registry.LookupArgsPolicy(toolName)
	if !ok || p == nil {
		t.Fatal("nmap policy not registered")
	}

	out, dropped, err := policy.ApplyArgs([]string{"-oN", "/etc/passwd", "-sV"}, p)
	if err != nil {
		t.Fatalf("unexpected validator error: %v", err)
	}

	// -sV is allowed; -oN must be dropped.
	if len(dropped) != 1 || dropped[0].Flag != "-oN" {
		t.Fatalf("expected -oN dropped, got %+v", dropped)
	}
	if dropped[0].Value != "/etc/passwd" {
		t.Fatalf("expected -oN value /etc/passwd, got %q", dropped[0].Value)
	}
	if !strings.Contains(dropped[0].Reason, "allowlist") {
		t.Fatalf("expected reason to mention allowlist, got %q", dropped[0].Reason)
	}
	// Filtered argv must contain -sV but not -oN.
	if len(out) != 1 || out[0] != "-sV" {
		t.Fatalf("expected only [-sV] after filter, got %v", out)
	}
}

// TestPolicy_AllowsSafeFlags asserts the documented happy path.
func TestPolicy_AllowsSafeFlags(t *testing.T) {
	p, _ := registry.LookupArgsPolicy(toolName)
	out, dropped, err := policy.ApplyArgs(
		[]string{"-sV", "-Pn", "-T4", "--top-ports", "100"},
		p,
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dropped) != 0 {
		t.Fatalf("expected nothing dropped, got %v", dropped)
	}
	if len(out) != 5 {
		t.Fatalf("expected 5 allowed, got %d (%v)", len(out), out)
	}
}
