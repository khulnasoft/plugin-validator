package templatereadme

import (
	"path/filepath"
	"testing"

	"github.com/khulnasoft/plugin-validator/pkg/analysis"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/readme"
	"github.com/khulnasoft/plugin-validator/pkg/testpassinterceptor"
	"github.com/stretchr/testify/require"
)

func TestTemplateReadme(t *testing.T) {
	var interceptor testpassinterceptor.TestPassInterceptor
	readmeContent := []byte(`# Khulnasoft Panel Plugin Template`)
	pass := &analysis.Pass{
		RootDir: filepath.Join("./"),
		ResultOf: map[*analysis.Analyzer]interface{}{
			readme.Analyzer: (readmeContent),
		},
		Report: interceptor.ReportInterceptor(),
	}

	_, err := Analyzer.Run(pass)
	require.NoError(t, err)
	require.Len(t, interceptor.Diagnostics, 1)
	require.Equal(t, interceptor.Diagnostics[0].Title, "README.md: uses README from template")
}

func TestTemplateReadmeLowerCase(t *testing.T) {
	var interceptor testpassinterceptor.TestPassInterceptor
	readmeContent := []byte(`# Khulnasoft panel Plugin Template`)
	pass := &analysis.Pass{
		RootDir: filepath.Join("./"),
		ResultOf: map[*analysis.Analyzer]interface{}{
			readme.Analyzer: (readmeContent),
		},
		Report: interceptor.ReportInterceptor(),
	}

	_, err := Analyzer.Run(pass)
	require.NoError(t, err)
	require.Len(t, interceptor.Diagnostics, 1)
	require.Equal(t, interceptor.Diagnostics[0].Title, "README.md: uses README from template")
}
