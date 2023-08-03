package server

import (
	"context"

	"github.com/kale-amruta/kopiasdkpoc/serverapi"
	"github.com/kopia/kopia/repo"
)

func handleCurrentUser(ctx context.Context, rc requestContext) (interface{}, *apiError) {
	return serverapi.CurrentUserResponse{
		Username: repo.GetDefaultUserName(ctx),
		Hostname: repo.GetDefaultHostName(ctx),
	}, nil
}
