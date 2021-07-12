// +build integration

package api

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/newrelic/newrelic-cli/internal/config"
)

var (
	testConfig = `{
		"*":{
			"loglevel":"debug",
			"plugindir": ".newrelic/plugins",
			"prereleasefeatures": "NOT_ASKED",
			"sendusagedata": "NOT_ASKED",
		}
	}`

	testCredentials = `{
		"default": {
			"apiKey": "testApiKey",
			"insightsInsertKey": "testInsightsInsertKey",
			"region": "testRegion",
			"accountID": 12345,
			"licenseKey": "testLicenseKey"
		},
	}`
)

func TestGetActiveProfileValues(t *testing.T) {
	dir, err := ioutil.TempDir("", "newrelic-cli.config_test.*")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	config.Init(dir)

	err = ioutil.WriteFile(filepath.Join(dir, config.CredentialsFileName), []byte(testCredentials), 0644)
	require.NoError(t, err)

	os.Unsetenv("NEW_RELIC_API_KEY")
	os.Unsetenv("NEW_RELIC_LICENSE_KEY")
	os.Unsetenv("NEW_RELIC_INSIGHTS_INSERT_KEY")
	os.Unsetenv("NEW_RELIC_REGION")
	os.Unsetenv("NEW_RELIC_ACCOUNT_ID")

	config.InitializeCredentialsStore()

	require.Equal(t, "testApiKey", GetActiveProfileString("apiKey"))
	require.Equal(t, "testInsightsInsertKey", GetActiveProfileString("insightsInsertKey"))
	require.Equal(t, "testRegion", GetActiveProfileString("region"))
	require.Equal(t, "testLicenseKey", GetActiveProfileString("licenseKey"))
	require.Equal(t, 12345, GetActiveProfileAccountID())
}

func TestGetActiveProfileValues_EnvVarOverride(t *testing.T) {
	dir, err := ioutil.TempDir("", "newrelic-cli.config_test.*")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	config.Init(dir)

	err = ioutil.WriteFile(filepath.Join(dir, config.CredentialsFileName), []byte(testCredentials), 0644)
	require.NoError(t, err)

	os.Setenv("NEW_RELIC_API_KEY", "apiKeyOverride")
	os.Setenv("NEW_RELIC_LICENSE_KEY", "licenseKeyOverride")
	os.Setenv("NEW_RELIC_INSIGHTS_INSERT_KEY", "insightsInsertKeyOverride")
	os.Setenv("NEW_RELIC_REGION", "regionOverride")
	os.Setenv("NEW_RELIC_ACCOUNT_ID", "67890")

	config.InitializeCredentialsStore()

	require.Equal(t, "apiKeyOverride", GetActiveProfileString("apiKey"))
	require.Equal(t, "insightsInsertKeyOverride", GetActiveProfileString("insightsInsertKey"))
	require.Equal(t, "regionOverride", GetActiveProfileString("region"))
	require.Equal(t, "licenseKeyOverride", GetActiveProfileString("licenseKey"))
	require.Equal(t, 67890, GetActiveProfileAccountID())
}

func TestGetConfigValues(t *testing.T) {
	dir, err := ioutil.TempDir("", "newrelic-cli.config_test.*")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	config.Init(dir)

	err = ioutil.WriteFile(filepath.Join(dir, config.ConfigFileName), []byte(testConfig), 0644)
	require.NoError(t, err)

	os.Unsetenv("NEW_RELIC_CLI_LOG_LEVEL")
	os.Unsetenv("NEW_RELIC_CLI_PLUGIN_DIR")
	os.Unsetenv("NEW_RELIC_CLI_PRERELEASEFEATURES")
	os.Unsetenv("NEW_RELIC_CLI_SENDUSAGEDATA")

	config.InitializeConfigStore()

	require.Equal(t, "debug", GetConfigString("loglevel"))
	require.Equal(t, ".newrelic/plugins", GetConfigString("plugindir"))
	require.Equal(t, config.TernaryValues.Unknown, GetConfigTernary("prereleasefeatures"))
	require.Equal(t, config.TernaryValues.Unknown, GetConfigTernary("sendusagedata"))
}

func TestGetConfigValues_EnvVarOverride(t *testing.T) {
	dir, err := ioutil.TempDir("", "newrelic-cli.config_test.*")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	config.Init(dir)

	err = ioutil.WriteFile(filepath.Join(dir, config.CredentialsFileName), []byte(testConfig), 0644)
	require.NoError(t, err)

	os.Setenv("NEW_RELIC_CLI_LOG_LEVEL", "trace")
	os.Setenv("NEW_RELIC_CLI_PLUGIN_DIR", "/tmp")
	os.Setenv("NEW_RELIC_CLI_PRERELEASEFEATURES", "ALLOW")
	os.Setenv("NEW_RELIC_CLI_SENDUSAGEDATA", "ALLOW")

	config.InitializeConfigStore()

	require.Equal(t, "trace", GetConfigString("loglevel"))
	require.Equal(t, "/tmp", GetConfigString("plugindir"))
	require.Equal(t, config.TernaryValues.Allow, GetConfigTernary("prereleasefeatures"))
	require.Equal(t, config.TernaryValues.Allow, GetConfigTernary("sendusagedata"))
}

func TestGetConfigValues_DefaultValues(t *testing.T) {
	dir, err := ioutil.TempDir("", "newrelic-cli.config_test.*")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	config.Init(dir)

	os.Unsetenv("NEW_RELIC_CLI_LOG_LEVEL")
	os.Unsetenv("NEW_RELIC_CLI_PLUGIN_DIR")
	os.Unsetenv("NEW_RELIC_CLI_PRERELEASEFEATURES")
	os.Unsetenv("NEW_RELIC_CLI_SENDUSAGEDATA")

	config.InitializeConfigStore()

	require.Equal(t, "info", GetConfigString("loglevel"))
	require.Equal(t, filepath.Join(dir, config.DefaultPluginDir), GetConfigString("plugindir"))
	require.Equal(t, config.TernaryValues.Unknown, GetConfigTernary("prereleasefeatures"))
	require.Equal(t, config.TernaryValues.Unknown, GetConfigTernary("sendusagedata"))
}

func TestRemoveProfile(t *testing.T) {
	dir, err := ioutil.TempDir("", "newrelic-cli.config_test.*")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	config.Init(dir)

	filename := filepath.Join(dir, config.CredentialsFileName)
	err = ioutil.WriteFile(filename, []byte(testCredentials), 0644)
	require.NoError(t, err)

	os.Setenv("NEW_RELIC_API_KEY", "apiKeyOverride")
	os.Setenv("NEW_RELIC_LICENSE_KEY", "licenseKeyOverride")
	os.Setenv("NEW_RELIC_INSIGHTS_INSERT_KEY", "insightsInsertKeyOverride")
	os.Setenv("NEW_RELIC_REGION", "regionOverride")
	os.Setenv("NEW_RELIC_ACCOUNT_ID", "67890")

	config.InitializeCredentialsStore()
	err = RemoveProfile("default")
	require.NoError(t, err)

	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err)
	require.Regexp(t, regexp.MustCompile(`{\s*}`), string(data))
}
