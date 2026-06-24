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

package httpx

import (
	"strings"
	"testing"

	"github.com/zeroroot-ai/gibson-executor/internal/policy"
	"github.com/zeroroot-ai/gibson-executor/internal/registry"
)

func TestPolicy_DropsOutputFlag(t *testing.T) {
	p, _ := registry.LookupArgsPolicy(toolName)
	_, dropped, err := policy.ApplyArgs([]string{"-o", "/etc/passwd"}, p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dropped) != 1 || dropped[0].Flag != "-o" {
		t.Fatalf("expected -o dropped, got %+v", dropped)
	}
}

// TestPolicy_RejectsNewlineInBody asserts the noNewlines validator
// catches HTTP request smuggling attempts.
func TestPolicy_RejectsNewlineInBody(t *testing.T) {
	p, _ := registry.LookupArgsPolicy(toolName)
	_, _, err := policy.ApplyArgs([]string{"-body", "GET / HTTP/1.1\r\nHost: x"}, p)
	if err == nil {
		t.Fatal("expected newline rejection")
	}
	if !strings.Contains(err.Error(), "newline") {
		t.Fatalf("expected newline error, got %v", err)
	}
}

// TestPolicy_RejectsArbitraryMethod ensures the -X enum is enforced.
func TestPolicy_RejectsArbitraryMethod(t *testing.T) {
	p, _ := registry.LookupArgsPolicy(toolName)
	_, _, err := policy.ApplyArgs([]string{"-X", "TRACE"}, p)
	if err == nil {
		t.Fatal("expected TRACE rejection")
	}
}
