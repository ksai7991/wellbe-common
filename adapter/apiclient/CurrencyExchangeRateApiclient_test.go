package apiclient

import (
	"context"
	"testing"
	env "wellbe-common/settings/env"

	"github.com/stretchr/testify/assert"
)

func TestGetRate(t *testing.T) {
    env.EnvLoad("./../../")
    ctx := context.Background()
    api := NewApiclient()
    r, err := api.GetRate(&ctx, "USD")
    assert.Nil(t, err)
    assert.Equal(t, 3.673042, r.Rates["AED"])
}

func TestGetRateError(t *testing.T) {
    env.EnvLoad("./../../")
    ctx := context.Background()
    api := NewApiclient()
    _, err := api.GetRate(&ctx, "USDa")
    assert.NotNil(t, err)
}