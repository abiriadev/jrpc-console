package lib

import (
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
)

func TestX(t *testing.T) {
	parser := participle.MustBuild[Invocation]()

	if ast, err := parser.ParseString("", "func()"); err != nil {
		panic(err)
	} else {
		assert.Equal(t, "func", ast.Name)
	}
}
