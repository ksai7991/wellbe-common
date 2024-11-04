package wellbefee

import (
	"testing"
	"wellbe-common/settings/env"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultFeeMasterWithKey(t *testing.T) {
    env.EnvLoad("./../../")
	a, err := GetDefaultFeeMasterWithKey("1")
	assert.Nil(t, err)
	assert.Equal(t, float64(0.13), a.FeeRate)
}