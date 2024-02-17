package lockfile

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExpandPackageCircular(t *testing.T) {
	t.Parallel()
	aLockfile := filepath.Join("..", "testdata", "node", "circular-yarn", "yarn.lock")
	packages, err := ParseYarnLock(aLockfile)
	require.NoError(t, err)

	// this one would cause an infinite loop
	expandedCircular, err := ExpandPackage("@jest/test-sequencer", packages)
	require.NoError(t, err)
	require.Len(t, expandedCircular.Dependencies, 416)
}

func TestExpandKhulnasoftPackagesFromYarn(t *testing.T) {
	t.Parallel()
	aLockfile := filepath.Join("..", "testdata", "node", "circular-yarn", "yarn.lock")
	packages, err := ParseYarnLock(aLockfile)
	require.NoError(t, err)

	expandedKhulnasoftData, err := ExpandPackage("@grafana/data", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftData.Dependencies, 54)

	expandedKhulnasoftE2E, err := ExpandPackage("@grafana/e2e", packages)
	require.Error(t, err)
	require.Nil(t, expandedKhulnasoftE2E)

	expandedKhulnasoftRuntime, err := ExpandPackage("@grafana/runtime", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftRuntime.Dependencies, 352)

	expandedKhulnasoftToolkit, err := ExpandPackage("@grafana/toolkit", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftToolkit.Dependencies, 1307)

	expandedKhulnasoftUI, err := ExpandPackage("@grafana/ui", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftUI.Dependencies, 350)
}

func TestExpandKhulnasoftPackagesFromNPM(t *testing.T) {
	t.Parallel()
	aLockfile := filepath.Join("..", "testdata", "node", "critical-npm", "package-lock.json")
	packages, err := ParseNpmLock(aLockfile)
	require.NoError(t, err)

	expandedKhulnasoftData, err := ExpandPackage("@grafana/data", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftData.Dependencies, 61)

	expandedKhulnasoftToolkit, err := ExpandPackage("@grafana/toolkit", packages)
	require.Equal(t, err.Error(), "package not found: @grafana/toolkit")
	require.Nil(t, expandedKhulnasoftToolkit)

	expandedKhulnasoftE2E, err := ExpandPackage("@grafana/e2e", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftE2E.Dependencies, 422)

	expandedKhulnasoftRuntime, err := ExpandPackage("@grafana/runtime", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftRuntime.Dependencies, 404)

	_, err = ExpandPackage("@grafana/toolkit", packages)
	require.Error(t, err)

	expandedKhulnasoftUI, err := ExpandPackage("@grafana/ui", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftUI.Dependencies, 389)
}

func TestExpandKhulnasoftPackagesFromPnpm(t *testing.T) {
	t.Parallel()
	aLockfile := filepath.Join("..", "testdata", "node", "critical-pnpm", "pnpm-lock.yaml")
	packages, err := ParsePnpmLock(aLockfile)
	require.NoError(t, err)

	expandedKhulnasoftData, err := ExpandPackage("@grafana/data", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftData.Dependencies, 64)

	expandedKhulnasoftToolkit, err := ExpandPackage("@grafana/toolkit", packages)
	require.Equal(t, err.Error(), "package not found: @grafana/toolkit")
	require.Nil(t, expandedKhulnasoftToolkit)

	expandedKhulnasoftE2E, err := ExpandPackage("@grafana/e2e", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftE2E.Dependencies, 485)

	expandedKhulnasoftRuntime, err := ExpandPackage("@grafana/runtime", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftRuntime.Dependencies, 439)

	_, err = ExpandPackage("@grafana/toolkit", packages)
	require.Error(t, err)

	expandedKhulnasoftUI, err := ExpandPackage("@grafana/ui", packages)
	require.NoError(t, err)
	require.Len(t, expandedKhulnasoftUI.Dependencies, 424)
}
