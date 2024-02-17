package passes

import (
	"github.com/khulnasoft/plugin-validator/pkg/analysis"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/archive"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/archivename"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/backenddebug"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/binarypermissions"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/brokenlinks"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/coderules"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/gomanifest"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/gosec"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/jargon"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/jssourcemap"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/legacybuilder"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/legacyplatform"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/license"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/logos"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/manifest"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/metadata"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/metadatapaths"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/metadatavalid"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/modulejs"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/org"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/osvscanner"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/packagejson"
	tzapgpt "github.com/khulnasoft/plugin-validator/pkg/analysis/passes/tzap-gpt"

	// "github.com/khulnasoft/plugin-validator/pkg/analysis/passes/osvscanner"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/pluginname"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/published"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/readme"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/restrictivedep"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/screenshots"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/signature"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/sourcecode"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/templatereadme"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/trackingscripts"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/typesuffix"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/unsafesvg"
	"github.com/khulnasoft/plugin-validator/pkg/analysis/passes/version"
)

var Analyzers = []*analysis.Analyzer{
	archive.Analyzer,
	archivename.Analyzer,
	brokenlinks.Analyzer,
	binarypermissions.Analyzer,
	coderules.Analyzer,
	gosec.Analyzer,
	gomanifest.Analyzer,
	jargon.Analyzer,
	jssourcemap.Analyzer,
	legacyplatform.Analyzer,
	legacybuilder.Analyzer,
	logos.Analyzer,
	license.Analyzer,
	manifest.Analyzer,
	metadata.Analyzer,
	metadatapaths.Analyzer,
	metadatavalid.Analyzer,
	modulejs.Analyzer,
	org.Analyzer,
	osvscanner.Analyzer,
	packagejson.Analyzer,
	pluginname.Analyzer,
	published.Analyzer,
	readme.Analyzer,
	restrictivedep.Analyzer,
	screenshots.Analyzer,
	signature.Analyzer,
	sourcecode.Analyzer,
	templatereadme.Analyzer,
	trackingscripts.Analyzer,
	typesuffix.Analyzer,
	tzapgpt.Analyzer,
	unsafesvg.Analyzer,
	version.Analyzer,
	backenddebug.Analyzer,
}
