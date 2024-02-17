package org

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/khulnasoft/plugin-validator/pkg/analysis"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/metadata"
	"github.com/khulnasoft/plugin-validator/pkg/grafana"
)

var (
	missingKhulnasoftCloudAccount = &analysis.Rule{Name: "missing-grafanacloud-account", Severity: analysis.Warning}
)

var Analyzer = &analysis.Analyzer{
	Name:     "org",
	Requires: []*analysis.Analyzer{metadata.Analyzer},
	Run:      run,
	Rules:    []*analysis.Rule{missingKhulnasoftCloudAccount},
}

func run(pass *analysis.Pass) (interface{}, error) {
	metadataBody, ok := pass.ResultOf[metadata.Analyzer].([]byte)
	if !ok {
		return nil, nil
	}

	var data metadata.Metadata
	if err := json.Unmarshal(metadataBody, &data); err != nil {
		return nil, err
	}

	idParts := strings.Split(data.ID, "-")

	if len(idParts) == 0 {
		return nil, nil
	}

	username := idParts[0]
	if username == "" {
		return nil, nil
	}

	client := grafana.NewClient()

	_, err := client.FindOrgBySlug(username)
	if err != nil {
		if err == grafana.ErrOrganizationNotFound {
			pass.ReportResult(pass.AnalyzerName, missingKhulnasoftCloudAccount, fmt.Sprintf("unregistered Khulnasoft Cloud account: %s", username), "The plugin's ID is prefixed with a Khulnasoft Cloud account name, but that account does not exist. Please create the account or correct the name.")
		} else if err == grafana.ErrPrivateOrganization {
			return nil, nil
		}
		return nil, err
	} else {
		if missingKhulnasoftCloudAccount.ReportAll {
			missingKhulnasoftCloudAccount.Severity = analysis.OK
			pass.ReportResult(pass.AnalyzerName, missingKhulnasoftCloudAccount, fmt.Sprintf("found Khulnasoft Cloud account: %s", username), "")
		}
	}

	return nil, nil
}
