package zedgen_test

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/Emyrk/zedgen"
	"github.com/stretchr/testify/require"
)

var updateGoldenFiles = flag.Bool("update", false, "Update golden files")

func TestZedPolicyGen(t *testing.T) {
	t.Parallel()
	if runtime.GOOS == "windows" {
		// Windows tests fail because the \n\r vs \n. It's not worth trying
		// to replace newlines for os tests. If people start using this tool on windows
		// and are seeing problems, then we can add build tags and figure it out.
		t.Skip("Skipping on windows")
	}

	const testDataDir = "testdata"
	files, err := os.ReadDir(testDataDir)
	require.NoErrorf(t, err, "read dir: %s", testDataDir)

	for _, f := range files {
		if !f.IsDir() {
			// Only test directories
			continue
		}

		t.Run(f.Name(), func(t *testing.T) {
			t.Parallel()

			dir := filepath.Join(testDataDir, f.Name())
			schema, err := FindFileWithExtension(dir, "zed")
			require.NoError(t, err)

			outPath, err := FindFileWithExtension(dir, "go")
			require.NoError(t, err)

			expected, err := os.ReadFile(filepath.Join(dir, outPath))
			require.NoError(t, err)

			schemaData, err := os.ReadFile(filepath.Join(dir, schema))
			require.NoError(t, err)

			outputData, err := zedgen.Generate(string(schemaData), zedgen.Options{
				Package:        "policy",
				SchemaFileName: schema,
			})
			require.NoError(t, err)

			if *updateGoldenFiles {
				err = os.WriteFile(filepath.Join(dir, outPath), []byte(outputData), 0644)
				require.NoError(t, err)
				return
			}

			require.Equal(t, string(expected), outputData)
		})
	}

}

func FindFileWithExtension(dir string, extension string) (string, error) {
	if len(extension) == 0 {
		return "", errors.New("empty extension")
	}
	if extension[0] != '.' {
		extension = "." + extension
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("read dir %q: %w", dir, err)
	}
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == extension {
			return entry.Name(), nil
		}
	}
	return "", fmt.Errorf("file with extension %q not found in %q", extension, dir)
}
