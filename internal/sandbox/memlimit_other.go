//go:build !linux && !darwin

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

package sandbox

import "os/exec"

// applyMemoryLimit is a no-op on platforms where ulimit -v is not reliably
// available.  Process-group isolation (Setpgid) and output caps (LimitReader)
// still apply; only the RLIMIT_AS enforcement is absent.
func applyMemoryLimit(_ *exec.Cmd, _ uint64) {}
