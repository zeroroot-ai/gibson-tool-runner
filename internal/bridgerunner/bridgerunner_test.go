package bridgerunner

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"
)

const manifestYAML = `apiVersion: connector.gibson.zeroroot.ai/v1
kind: Connector
`

func getenvFrom(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

func TestManifestPath_B64_WritesFile(t *testing.T) {
	dir := t.TempDir()
	env := map[string]string{
		EnvManifestB64: base64.StdEncoding.EncodeToString([]byte(manifestYAML)),
	}

	path, err := ManifestPath(getenvFrom(env), dir)
	if err != nil {
		t.Fatalf("ManifestPath: %v", err)
	}
	if want := filepath.Join(dir, "connector.yaml"); path != want {
		t.Fatalf("path = %q, want %q", path, want)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read back: %v", err)
	}
	if string(got) != manifestYAML {
		t.Fatalf("content = %q, want %q", got, manifestYAML)
	}
	info, _ := os.Stat(path)
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("perm = %v, want 0600 (manifest may reference secrets)", info.Mode().Perm())
	}
}

func TestManifestPath_PathPassthrough(t *testing.T) {
	f := filepath.Join(t.TempDir(), "c.yaml")
	if err := os.WriteFile(f, []byte(manifestYAML), 0o600); err != nil {
		t.Fatal(err)
	}
	path, err := ManifestPath(getenvFrom(map[string]string{EnvManifestPath: f}), t.TempDir())
	if err != nil {
		t.Fatalf("ManifestPath: %v", err)
	}
	if path != f {
		t.Fatalf("path = %q, want %q", path, f)
	}
}

func TestManifestPath_PathMissingFile_Errors(t *testing.T) {
	_, err := ManifestPath(getenvFrom(map[string]string{EnvManifestPath: "/nope/c.yaml"}), t.TempDir())
	if err == nil {
		t.Fatal("expected error for missing manifest file")
	}
}

func TestManifestPath_BothSet_Errors(t *testing.T) {
	_, err := ManifestPath(getenvFrom(map[string]string{
		EnvManifestB64:  "eA==",
		EnvManifestPath: "/tmp/c.yaml",
	}), t.TempDir())
	if err == nil {
		t.Fatal("expected error when both env vars are set")
	}
}

func TestManifestPath_NeitherSet_Errors(t *testing.T) {
	_, err := ManifestPath(getenvFrom(map[string]string{}), t.TempDir())
	if err == nil {
		t.Fatal("expected error when neither env var is set")
	}
}

func TestManifestPath_BadBase64_Errors(t *testing.T) {
	_, err := ManifestPath(getenvFrom(map[string]string{EnvManifestB64: "!!!not-base64!!!"}), t.TempDir())
	if err == nil {
		t.Fatal("expected error for invalid base64")
	}
}
