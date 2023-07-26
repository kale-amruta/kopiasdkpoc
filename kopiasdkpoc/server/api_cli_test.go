package server_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	github.com/kale-amruta/kopiasdkpoc/apiclient"
	github.com/kale-amruta/kopiasdkpoc/repotesting"
	github.com/kale-amruta/kopiasdkpoc/serverapi"
	github.com/kale-amruta/kopiasdkpoc/servertesting"
)

func TestCLIAPI(t *testing.T) {
	ctx, env := repotesting.NewEnvironment(t, repotesting.FormatNotImportant)
	srvInfo := servertesting.StartServer(t, env, false)

	cli, err := apiclient.NewKopiaAPIClient(apiclient.Options{
		BaseURL:                             srvInfo.BaseURL,
		TrustedServerCertificateFingerprint: srvInfo.TrustedServerCertificateFingerprint,
		Username:                            servertesting.TestUIUsername,
		Password:                            servertesting.TestUIPassword,
	})

	require.NoError(t, err)
	require.NoError(t, cli.FetchCSRFTokenForTesting(ctx))

	resp := &serverapi.CLIInfo{}
	require.NoError(t, cli.Get(ctx, "cli", nil, resp))

	exe, _ := os.Executable()

	require.Equal(t, exe+" --config-file="+env.ConfigFile(), resp.Executable)
}
