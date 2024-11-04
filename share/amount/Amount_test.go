package amount

import (
	"testing"
	"wellbe-common/settings/env"
	"wellbe-common/share/commonsettings/constants/code/cCurrency"

	"github.com/stretchr/testify/assert"
)

func TestExchangeConversionJPYUSD(t *testing.T) {
    env.EnvLoad("./../../")
	a, _, err := ExchangeConversion(cCurrency.JPY, cCurrency.USD, float64(10000))
	assert.Nil(t, err)
	assert.Equal(t, float64(71), a)
}

func TestExchangeConversionVNDJPY(t *testing.T) {
    env.EnvLoad("./../../")
	a, _, err := ExchangeConversion(cCurrency.VND, cCurrency.JPY, float64(1659))
	assert.Nil(t, err)
	assert.Equal(t, float64(10), a)
}

func TestRoundWithCurrency(t *testing.T) {
    env.EnvLoad("./../../")
	a1, _ := RoundWithCurrency(float64(123.345), cCurrency.USD)
	a2, err := RoundWithCurrency(float64(123.344), cCurrency.USD)
	assert.Nil(t, err)
	assert.Equal(t, float64(123.35), a1)
	assert.Equal(t, float64(123.34), a2)
}