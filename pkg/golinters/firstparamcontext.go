package golinters

import (
	"github.com/golangci/golangci-lint/pkg/goanalysis"
	"github.com/liujtm/first_param_ctx_linter/firstparamcontext"
	"golang.org/x/tools/go/analysis"
)

func NewfirstparamcontextCheck() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"firstparamcontext",
		"Checks that functions first param type is Context",
		[]*analysis.Analyzer{firstparamcontext.Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
