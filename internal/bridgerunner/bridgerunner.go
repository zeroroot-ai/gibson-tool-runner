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

// Package bridgerunner resolves the component manifest for the MCP-bridge
// runner image (ADR-0048 Option 1: one generic OSS runner image serves every
// package-distributed connector; the daemon launches it as a setec Sandbox).
//
// The manifest arrives one of two ways:
//
//   - GIBSON_CONNECTOR_MANIFEST_B64 — base64-encoded plugin manifest YAML (a
//     runtime: mcp-bridge plugin, ADR-0049). The
//     hosted path delivers it this way because a setec Sandbox launch carries
//     env, not volumes.
//   - GIBSON_CONNECTOR_MANIFEST_PATH — filesystem path. Local/dev runs.
//
// Exactly one must be set.
package bridgerunner

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// Env var names read by the runner.
const (
	EnvManifestB64  = "GIBSON_CONNECTOR_MANIFEST_B64"
	EnvManifestPath = "GIBSON_CONNECTOR_MANIFEST_PATH"
)

// ManifestPath resolves the connector manifest location from the environment.
// When the manifest is delivered inline (base64), it is materialised as
// plugin.yaml under writeDir with owner-only permissions and that path is
// returned. getenv is parameterised for tests (pass os.Getenv in production).
func ManifestPath(getenv func(string) string, writeDir string) (string, error) {
	b64 := getenv(EnvManifestB64)
	path := getenv(EnvManifestPath)

	switch {
	case b64 != "" && path != "":
		return "", fmt.Errorf("bridgerunner: %s and %s are both set; exactly one is required",
			EnvManifestB64, EnvManifestPath)
	case b64 == "" && path == "":
		return "", fmt.Errorf("bridgerunner: connector manifest missing; set %s (hosted) or %s (local)",
			EnvManifestB64, EnvManifestPath)
	case path != "":
		if _, err := os.Stat(path); err != nil {
			return "", fmt.Errorf("bridgerunner: manifest path %s: %w", path, err)
		}
		return path, nil
	}

	raw, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", fmt.Errorf("bridgerunner: decode %s: %w", EnvManifestB64, err)
	}
	out := filepath.Join(writeDir, "plugin.yaml")
	if err := os.WriteFile(out, raw, 0o600); err != nil {
		return "", fmt.Errorf("bridgerunner: write manifest to %s: %w", out, err)
	}
	return out, nil
}
