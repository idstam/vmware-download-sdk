package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDetailsSuccess(t *testing.T) {
	var dlgDetails DlgDetails
	dlgDetails, err = basicClient.GetDlgDetails("VMTOOLS1130", "1073")
	assert.Nil(t, err)
	assert.NotEmpty(t, dlgDetails.DownloadDetails, "Expected response to no be empty")
}

func TestGetDetailsInvalidProductId(t *testing.T) {
	var dlgDetails DlgDetails
	dlgDetails, err := basicClient.GetDlgDetails("VMTOOLS1130", "6666666")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorDlgDetailsInputs)
	assert.Empty(t, dlgDetails, "Expected response to be empty")
}

func TestGetDetailsInvalidDownloadGroup(t *testing.T) {
	var dlgDetails DlgDetails
	dlgDetails, err := basicClient.GetDlgDetails("VMTOOLS666", "1073")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorDlgDetailsInputs)
	assert.Empty(t, dlgDetails, "Expected response to be empty")
}

func TestFindDlgDetailsSuccess(t *testing.T) {
	err = ensureLogin(t)
	require.Nil(t, err)

	var downloadDetails FoundDownload
	downloadDetails, err := authenticatedClient.FindDlgDetails("VMTOOLS1130", "1073", "VMware-Tools-darwin-*.tar.gz")
	assert.Nil(t, err)
	assert.NotEmpty(t, downloadDetails.DownloadDetails.FileName, "Expected response to not be empty")
}

func TestFindDlgDetailsMultipleGlob(t *testing.T) {
	err = ensureLogin(t)
	require.Nil(t, err)

	var downloadDetails FoundDownload
	downloadDetails, err := authenticatedClient.FindDlgDetails("VMTOOLS1130", "1073", "double*glob*")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorMultipleGlob)
	assert.Empty(t, downloadDetails.DownloadDetails.FileName, "Expected response to be empty")
}

func TestFindDlgDetailsNoGlob(t *testing.T) {
	err = ensureLogin(t)
	require.Nil(t, err)

	var downloadDetails FoundDownload
	downloadDetails, err := authenticatedClient.FindDlgDetails("VMTOOLS1130", "1073", "no.glob")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorNoGlob)
	assert.Empty(t, downloadDetails.DownloadDetails.FileName, "Expected response to be empty")
}

func TestFindDlgDetailsNoMatch(t *testing.T) {
	err = ensureLogin(t)
	require.Nil(t, err)

	var downloadDetails FoundDownload
	downloadDetails, err := authenticatedClient.FindDlgDetails("VMTOOLS1130", "1073", "invalid*glob")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorNoMatchingFiles)
	assert.Empty(t, downloadDetails.DownloadDetails.FileName, "Expected response to be empty")
}

func TestFindDlgDetailsMultipleMatch(t *testing.T) {
	err = ensureLogin(t)
	require.Nil(t, err)

	var downloadDetails FoundDownload
	downloadDetails, err := authenticatedClient.FindDlgDetails("VMTOOLS1130", "1073", "VMware*.gz")
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorMultipleMatchingFiles)
	assert.Empty(t, downloadDetails.DownloadDetails.FileName, "Expected response to be empty")
}
