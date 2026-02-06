package zedgen

import (
	"fmt"

	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
)

func Generate(schema string) (string, error) {
	var prefix string // TODO: What is the prefix for?
	compiled, err := compiler.Compile(compiler.InputSchema{
		Source:       "policy.zed",
		SchemaString: schema,
	}, compiler.ObjectTypePrefix(prefix))
	if err != nil {
		return "", fmt.Errorf("compile schema: %w", err)
	}

	var _ = compiled

	return "", nil
}
