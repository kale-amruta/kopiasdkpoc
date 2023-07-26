package server_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	github.com/kale-amruta/kopiasdkpoc/apiclient"
	github.com/kale-amruta/kopiasdkpoc/repotesting"
	github.com/kale-amruta/kopiasdkpoc/serverapi"
	github.com/kale-amruta/kopiasdkpoc/servertesting"
	github.com/kale-amruta/kopiasdkpoc/testutil"
)

func TestPathsAPI(t *testing.T) {
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

	dir0 := testutil.TempDirectory(t)

	req := &serverapi.ResolvePathRequest{
		Path: dir0,
	}
	resp := &serverapi.ResolvePathResponse{}
	require.NoError(t, cli.Post(ctx, "paths/resolve", req, resp))

	require.Equal(t, env.LocalPathSourceInfo(dir0), resp.SourceInfo)
}
