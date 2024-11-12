package server

import (
	"github.com/frustra/bbcode"
)

func parseBBCode(rawText string) string {
	compiler := bbcode.NewCompiler(true, true)
	return compiler.Compile(rawText)
}
