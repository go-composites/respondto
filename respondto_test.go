package respondto_test

import (
	"testing"

	"github.com/go-composites/respondto"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), respondto.Analyzer, "a")
}
