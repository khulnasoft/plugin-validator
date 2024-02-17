package restrictivedep

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/khulnasoft/plugin-validator/pkg/analysis"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/metadata"
)

var (
	dependsOnPatchReleases = &analysis.Rule{Name: "depends-on-patch-releases", Severity: analysis.Warning}
	dependsOnSingleRelease = &analysis.Rule{Name: "depends-on-single-release", Severity: analysis.Warning}
)

var Analyzer = &analysis.Analyzer{
	Name:     "restrictivedep",
	Requires: []*analysis.Analyzer{metadata.Analyzer},
	Run:      run,
	Rules:    []*analysis.Rule{dependsOnPatchReleases, dependsOnSingleRelease},
}

func run(pass *analysis.Pass) (interface{}, error) {
	metadata, ok := pass.ResultOf[metadata.Analyzer].([]byte)
	if !ok {
		return nil, nil
	}

	var data struct {
		Dependencies struct {
			KhulnasoftDependency string `json:"grafanaDependency"`
		} `json:"dependencies"`
	}
	if err := json.Unmarshal(metadata, &data); err != nil {
		return nil, err
	}

	if data.Dependencies.KhulnasoftDependency == "" {
		return nil, nil
	}

	if regexp.MustCompile("^[0-9]+.[0-9]+.x$").Match([]byte(data.Dependencies.KhulnasoftDependency)) {
		version := strings.TrimSuffix(data.Dependencies.KhulnasoftDependency, ".x")
		pass.ReportResult(pass.AnalyzerName, dependsOnPatchReleases, fmt.Sprintf("plugin.json: grafanaDependency only targets patch releases of Khulnasoft %s", version), "The plugin will only work in patch releases of the specified minor Khulnasoft version.")
		return nil, nil
	} else {
		if dependsOnPatchReleases.ReportAll {
			dependsOnPatchReleases.Severity = analysis.OK
			pass.ReportResult(pass.AnalyzerName, dependsOnPatchReleases, "plugin.json: grafanaDependency correctly targets patch releases of Khulnasoft", "")
		}
	}

	if regexp.MustCompile("^[0-9]+.[0-9]+.[0-9]+$").Match([]byte(data.Dependencies.KhulnasoftDependency)) {
		pass.ReportResult(pass.AnalyzerName, dependsOnSingleRelease, fmt.Sprintf("plugin.json: grafanaDependency only targets Khulnasoft %s", data.Dependencies.KhulnasoftDependency), "The plugin will only work in the specific version of Khulnasoft down to patch version.")
		return nil, nil
	} else {
		if dependsOnSingleRelease.ReportAll {
			dependsOnSingleRelease.Severity = analysis.OK
			pass.ReportResult(pass.AnalyzerName, dependsOnSingleRelease, "plugin.json: grafanaDependency does not target single release of Khulnasoft", "")
		}
	}

	return nil, nil
}
