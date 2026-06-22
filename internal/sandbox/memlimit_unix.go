//go:build linux || darwin

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

// applyMemoryLimit rewrites cmd to run under a shell wrapper that applies
// `ulimit -v <limitKiB>` before exec-ing the real binary.  This is the
// only portable stdlib-only mechanism for enforcing RLIMIT_AS in a child
// process: the Go syscall package's SysProcAttr does not expose an Rlimit
// slice on Linux, so we rely on the shell's built-in ulimit to set the
// limit in the child's address space before the tool binary is exec'd.
func applyMemoryLimit(cmd *exec.Cmd, limitKiB uint64) {
	shellWrapArgs(cmd, limitKiB)
}
