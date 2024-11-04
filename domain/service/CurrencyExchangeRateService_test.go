package service

import (
    "testing"
    "context"
    "github.com/stretchr/testify/assert"
    repository "wellbe-common/domain/repository"
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    number "wellbe-common/share/number"
)

type currencyExchangeRateMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError)
    FakeUpdate func(*model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CurrencyExchangeRate, *errordef.LogicError)
}

func (lr currencyExchangeRateMockRepository) CreateCurrencyExchangeRate(ctx *context.Context, currencyExchangeRate *model.CurrencyExchangeRate)  (*model.CurrencyExchangeRate, *errordef.LogicError) {
    return lr.FakeCreate(currencyExchangeRate)
}

func (lr currencyExchangeRateMockRepository) UpdateCurrencyExchangeRate(ctx *context.Context, currencyExchangeRate *model.CurrencyExchangeRate)  (*model.CurrencyExchangeRate, *errordef.LogicError) {
    return lr.FakeUpdate(currencyExchangeRate)
}

func (lr currencyExchangeRateMockRepository) DeleteCurrencyExchangeRate(ctx *context.Context, baseCurrencyCd int, targetCurrencyCd int)  *errordef.LogicError {
    return lr.FakeDelete(baseCurrencyCd, targetCurrencyCd)
}

func (lr currencyExchangeRateMockRepository)GetCurrencyExchangeRateWithKey(ctx *context.Context, baseCurrencyCd int, targetCurrencyCd int)  ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
    return lr.FakeGet(baseCurrencyCd, targetCurrencyCd)
}


type currencyExchangeRateMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr currencyExchangeRateMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCurrencyExchangeRate(t *testing.T) {
    ctx := context.Background()
    repository := &currencyExchangeRateMockRepository{
        FakeCreate: func(currencyExchangeRate *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError) {
            return currencyExchangeRate, nil
        },
    }
    numberUtil := &currencyExchangeRateMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    currencyExchangeRateService := NewCurrencyExchangeRateService(repository, numberUtil)
    in_currencyExchangeRate := new(model.CurrencyExchangeRate)
    in_currencyExchangeRate.BaseCurrencyCd = 0
    in_currencyExchangeRate.TargetCurrencyCd = 1
    in_currencyExchangeRate.PaireName = "XXXXXX"
    in_currencyExchangeRate.Rate = 3.2
    out_currencyExchangeRate, err := currencyExchangeRateService.CreateCurrencyExchangeRate(&ctx, in_currencyExchangeRate)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_currencyExchangeRate.BaseCurrencyCd)
    assert.Equal(t, 1, out_currencyExchangeRate.TargetCurrencyCd)
    assert.Equal(t, "XXXXXX", out_currencyExchangeRate.PaireName)
    assert.Equal(t, 3.2, out_currencyExchangeRate.Rate)
    assert.NotNil(t, out_currencyExchangeRate.CreateDatetime)
    assert.NotEqual(t, "", out_currencyExchangeRate.CreateDatetime)
    assert.Equal(t, "CreateCurrencyExchangeRate", out_currencyExchangeRate.CreateFunction)
    assert.NotNil(t, out_currencyExchangeRate.UpdateDatetime)
    assert.NotEqual(t, "", out_currencyExchangeRate.UpdateDatetime)
    assert.Equal(t, "CreateCurrencyExchangeRate", out_currencyExchangeRate.UpdateFunction)
}

func TestUpdateCurrencyExchangeRate(t *testing.T) {
    ctx := context.Background()
    repository := &currencyExchangeRateMockRepository{
        FakeUpdate: func(currencyExchangeRate *model.CurrencyExchangeRate) (*model.CurrencyExchangeRate, *errordef.LogicError) {
            return currencyExchangeRate, nil
        },
        FakeGet: func(baseCurrencyCd int, targetCurrencyCd int) ([]*model.CurrencyExchangeRate, *errordef.LogicError) {
            return []*model.CurrencyExchangeRate{&model.CurrencyExchangeRate{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &currencyExchangeRateMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    currencyExchangeRateService := NewCurrencyExchangeRateService(repository, numberUtil)
    in_currencyExchangeRate := new(model.CurrencyExchangeRate)
    in_currencyExchangeRate.BaseCurrencyCd = 0
    in_currencyExchangeRate.TargetCurrencyCd = 1
    in_currencyExchangeRate.PaireName = "XXXXXX"
    in_currencyExchangeRate.Rate = 3.2
    out_currencyExchangeRate, err := currencyExchangeRateService.UpdateCurrencyExchangeRate(&ctx, in_currencyExchangeRate)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_currencyExchangeRate.BaseCurrencyCd)
    assert.Equal(t, 1, out_currencyExchangeRate.TargetCurrencyCd)
    assert.Equal(t, "XXXXXX", out_currencyExchangeRate.PaireName)
    assert.Equal(t, 3.2, out_currencyExchangeRate.Rate)
    assert.NotNil(t, out_currencyExchangeRate.CreateDatetime)
    assert.Equal(t, "", out_currencyExchangeRate.CreateDatetime)
    assert.Equal(t, "", out_currencyExchangeRate.CreateFunction)
    assert.NotNil(t, out_currencyExchangeRate.UpdateDatetime)
    assert.NotEqual(t, "", out_currencyExchangeRate.UpdateDatetime)
    assert.Equal(t, "UpdateCurrencyExchangeRate", out_currencyExchangeRate.UpdateFunction)
}
