package generators

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hary654321/nuclei/v2/pkg/catalog/disk"
	"github.com/stretchr/testify/require"
)

func TestLoadPayloads(t *testing.T) {
	tempdir, err := os.MkdirTemp("", "templates-*")
	require.NoError(t, err, "could not create temp dir")
	defer os.RemoveAll(tempdir)

	generator := &PayloadGenerator{catalog: disk.NewCatalog(tempdir)}

	fullpath := filepath.Join(tempdir, "payloads.txt")
	err = os.WriteFile(fullpath, []byte("test\nanother"), 0777)
	require.NoError(t, err, "could not write payload")

	// Test sandbox
	t.Run("templates-directory", func(t *testing.T) {
		values, err := generator.loadPayloads(map[string]interface{}{
			"new": fullpath,
		}, "/test", tempdir, true)
		require.NoError(t, err, "could not load payloads")
		require.Equal(t, map[string][]string{"new": {"test", "another"}}, values, "could not get values")
	})
	t.Run("template-directory", func(t *testing.T) {
		values, err := generator.loadPayloads(map[string]interface{}{
			"new": fullpath,
		}, filepath.Join(tempdir, "test.yaml"), "/test", true)
		require.NoError(t, err, "could not load payloads")
		require.Equal(t, map[string][]string{"new": {"test", "another"}}, values, "could not get values")
	})
	t.Run("no-sandbox-unix", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			return
		}
		_, err := generator.loadPayloads(map[string]interface{}{
			"new": "/etc/passwd",
		}, "/random", "/test", false)
		require.NoError(t, err, "could load payloads")
	})
	t.Run("invalid", func(t *testing.T) {
		values, err := generator.loadPayloads(map[string]interface{}{
			"new": "/etc/passwd",
		}, "/random", "/test", true)
		require.Error(t, err, "could load payloads")
		require.Equal(t, 0, len(values), "could get values")

		values, err = generator.loadPayloads(map[string]interface{}{
			"new": fullpath,
		}, "/random", "/test", true)
		require.Error(t, err, "could load payloads")
		require.Equal(t, 0, len(values), "could get values")
	})
}
