package zedgen_test

import (
	"runtime"
	"testing"
)

func TestZedPolicyGen(t *testing.T) {
	t.Parallel()
	if runtime.GOOS == "windows" {
		// Windows tests fail because the \n\r vs \n. It's not worth trying
		// to replace newlines for os tests. If people start using this tool on windows
		// and are seeing problems, then we can add build tags and figure it out.
		t.Skip("Skipping on windows")
	}
}
