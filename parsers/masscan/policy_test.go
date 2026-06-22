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

package masscan

import (
	"testing"

	"github.com/zeroroot-ai/gibson-tool-runner/internal/policy"
	"github.com/zeroroot-ai/gibson-tool-runner/internal/registry"
)

// TestPolicy_DropsOutputFile asserts the canonical injection
// `-oN /etc/passwd` is rejected. The runner pins -oJ - so any -o* flag is
// a guaranteed misuse signal.
func TestPolicy_DropsOutputFile(t *testing.T) {
	p, _ := registry.LookupArgsPolicy(toolName)
	_, dropped, err := policy.ApplyArgs([]string{"-oN", "/etc/passwd"}, p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dropped) != 1 || dropped[0].Flag != "-oN" {
		t.Fatalf("expected -oN dropped, got %+v", dropped)
	}
}

func TestPolicy_AllowsRateAndBanners(t *testing.T) {
	p, _ := registry.LookupArgsPolicy(toolName)
	out, dropped, err := policy.ApplyArgs([]string{"--rate", "10000", "--banners"}, p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dropped) != 0 {
		t.Fatalf("expected 0 dropped, got %v", dropped)
	}
	if len(out) != 3 {
		t.Fatalf("expected 3 args, got %v", out)
	}
}
