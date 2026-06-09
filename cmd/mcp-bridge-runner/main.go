// mcp-bridge-runner is the generic MCP-bridge runner image entry point
// (ADR-0048 Option 1). It resolves the connector manifest from the
// environment and hands control to the OSS SDK's mcpbridge: capability-grant
// registration, GetCredential secret resolution, vendor MCP server spawn
// (npx/uvx as a stdio subprocess), tools/list discovery, and the
// PollWork↔tools/call dispatch loop. One runner image serves every
// package-distributed connector; there is no per-connector build.
//
// Environment:
//
//	GIBSON_CONNECTOR_MANIFEST_B64   base64 connector YAML (hosted/setec path)
//	GIBSON_CONNECTOR_MANIFEST_PATH  manifest file path (local/dev)
//	GIBSON_URL                      platform base URL (read by plugin.Serve)
//	GIBSON_BOOTSTRAP_TOKEN          one-time registration token (first run)
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/zeroroot-ai/gibson-tool-runner/internal/bridgerunner"
	"github.com/zeroroot-ai/sdk/mcpbridge"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))

	path, err := bridgerunner.ManifestPath(os.Getenv, os.TempDir())
	if err != nil {
		slog.Error("mcp-bridge-runner: resolve connector manifest", "err", err)
		os.Exit(1)
	}

	// mcpbridge.Run blocks until SIGTERM/SIGINT (handled inside plugin.Serve)
	// or a fatal error. Exit code 75 (rotation-restart) propagates via
	// plugin.Serve's own os.Exit; any other error lands here.
	if err := mcpbridge.Run(context.Background(), path); err != nil {
		slog.Error("mcp-bridge-runner: bridge terminated", "err", err)
		os.Exit(1)
	}
}
