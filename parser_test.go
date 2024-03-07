package jrpcconsole

import (
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestX(t *testing.T) {
	parser := participle.MustBuild[Invocation]()

	if _, err := parser.ParseString("", "func()"); err != nil {
		panic(err)
	}

	// assert.Equal()
}
