package zedgen_test

import (
	"fmt"
	"testing"

	"github.com/Emyrk/zedgen"
	"github.com/stretchr/testify/require"
)

func TestTemplatesLoad(t *testing.T) {
	t.Parallel()

	tpls, err := zedgen.LoadTemplates()
	require.NoError(t, err)

	fmt.Println(tpls)
}
