package providervalidation_test

import (
	"testing"
	"time"

	github.com/kale-amruta/kopiasdkpoc/blobtesting"
	github.com/kale-amruta/kopiasdkpoc/providervalidation"
	github.com/kale-amruta/kopiasdkpoc/testlogging"
)

func TestProviderValidation(t *testing.T) {
	ctx := testlogging.Context(t)
	st := blobtesting.NewMapStorage(blobtesting.DataMap{}, nil, nil)
	opt := providervalidation.DefaultOptions
	opt.ConcurrencyTestDuration = 15 * time.Second
	providervalidation.ValidateProvider(ctx, st, opt)
}
