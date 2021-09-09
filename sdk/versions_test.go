package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetVersionSuccess(t *testing.T) {
	var versions map[string]APIVersions
	versions, err := basicClient.GetVersionMap("vmware_vsphere", "esxi")
	require.Nil(t, err)
	assert.Greater(t, len(versions), 1, "Expected response to contain at least 1 item")
	assert.Contains(t, versions, "7.0.0")
	assert.Contains(t, versions, "6.7.0")
	assert.Contains(t, versions, "6.5.0")
}

func TestGetVersionMapInvalidSubProduct(t *testing.T) {
	var versions map[string]APIVersions
	versions, err := basicClient.GetVersionMap("vmware_tools", "dummy")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorInvalidSubProduct)
	assert.Empty(t, versions, "Expected response to be empty")
}

func TestGetVersionInvalidSlug(t *testing.T) {
	var versions map[string]APIVersions
	versions, err := basicClient.GetVersionMap("mware_tools", "vmtools")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorInvalidSlug)
	assert.Empty(t, versions, "Expected response to be empty")
}

func TestFindVersion(t *testing.T) {
	var foundVersion APIVersions
	foundVersion, err = basicClient.FindVersion("vmware_tools", "vmtools", "11.1.1")
	assert.Nil(t, err)
	assert.NotEmpty(t, foundVersion.Code, "Expected response not to be empty")
}

func TestFindVersionInvalidSlug(t *testing.T) {
	var foundVersion APIVersions
	foundVersion, err = basicClient.FindVersion("mware_tools", "vmtools", "11.1.1")
	assert.ErrorIs(t, err, ErrorInvalidSlug)
	assert.Empty(t, foundVersion.Code, "Expected response to be empty")
}

func TestFindVersionInvalidVersion(t *testing.T) {
	var foundVersion APIVersions
	foundVersion, err = basicClient.FindVersion("vmware_tools", "vmtools", "666")
	assert.ErrorIs(t, err, ErrorInvalidVersion)
	assert.Empty(t, foundVersion.Code, "Expected response to be empty")
}

func TestFindVersionInvalidSubProduct(t *testing.T) {
	var foundVersion APIVersions
	foundVersion, err = basicClient.FindVersion("vmware_tools", "tools", "11.1.1")
	assert.ErrorIs(t, err, ErrorInvalidSubProduct)
	assert.Empty(t, foundVersion.Code, "Expected response to be empty")
}

func TestFindVersionMinorGlob(t *testing.T) {
	var foundVersion APIVersions
	foundVersion, err = basicClient.FindVersion("vmware_tools", "vmtools", "10.2.*")
	assert.Nil(t, err)
	assert.Equal(t, foundVersion.Code, "VMTOOLS1025")
}

func TestFindVersionOnlyGlob(t *testing.T) {
	var foundVersion APIVersions
	foundVersion, err = basicClient.FindVersion("vmware_tools", "vmtools", "*")
	assert.Nil(t, err)
	assert.NotEmpty(t, foundVersion.Code)
}

func TestGetVersionArraySuccess(t *testing.T) {
	var versions []string
	versions, err := basicClient.GetVersionSlice("vmware_tools", "vmtools")
	assert.Nil(t, err)
	assert.Greater(t, len(versions), 10, "Expected response to contain at least 10 items")
}