// Command respondto runs the respondto analyzer standalone or as a `go vet` tool:
//
//	go install github.com/go-composites/respondto/cmd/respondto@latest
//	go vet -vettool=$(which respondto) ./...
package main

import (
	"github.com/go-composites/respondto"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(respondto.Analyzer) }
