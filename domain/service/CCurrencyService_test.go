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

type cCurrencyMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CCurrency) (*model.CCurrency, *errordef.LogicError)
    FakeUpdate func(*model.CCurrency) (*model.CCurrency, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CCurrency, *errordef.LogicError)
}

func (lr cCurrencyMockRepository) CreateCCurrency(ctx *context.Context, cCurrency *model.CCurrency)  (*model.CCurrency, *errordef.LogicError) {
    return lr.FakeCreate(cCurrency)
}

func (lr cCurrencyMockRepository) UpdateCCurrency(ctx *context.Context, cCurrency *model.CCurrency)  (*model.CCurrency, *errordef.LogicError) {
    return lr.FakeUpdate(cCurrency)
}

func (lr cCurrencyMockRepository) DeleteCCurrency(ctx *context.Context, currencyCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(currencyCd, languageCd)
}

func (lr cCurrencyMockRepository)GetCCurrencyWithKey(ctx *context.Context, currencyCd int, languageCd int)  ([]*model.CCurrency, *errordef.LogicError) {
    return lr.FakeGet(currencyCd, languageCd)
}


type cCurrencyMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cCurrencyMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCCurrency(t *testing.T) {
    ctx := context.Background()
    repository := &cCurrencyMockRepository{
        FakeCreate: func(cCurrency *model.CCurrency) (*model.CCurrency, *errordef.LogicError) {
            return cCurrency, nil
        },
    }
    numberUtil := &cCurrencyMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCurrencyService := NewCCurrencyService(repository, numberUtil)
    in_cCurrency := new(model.CCurrency)
    in_cCurrency.CurrencyCd = 0
    in_cCurrency.LanguageCd = 1
    in_cCurrency.CurrencyName = "dummy-CurrencyName"
    in_cCurrency.CurrencyCdIso = "XXX"
    in_cCurrency.SignificantDigit = 4
    out_cCurrency, err := cCurrencyService.CreateCCurrency(&ctx, in_cCurrency)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCurrency.CurrencyCd)
    assert.Equal(t, 1, out_cCurrency.LanguageCd)
    assert.Equal(t, "dummy-CurrencyName", out_cCurrency.CurrencyName)
    assert.Equal(t, "XXX", out_cCurrency.CurrencyCdIso)
    assert.Equal(t, 4, out_cCurrency.SignificantDigit)
    assert.NotNil(t, out_cCurrency.CreateDatetime)
    assert.NotEqual(t, "", out_cCurrency.CreateDatetime)
    assert.Equal(t, "CreateCCurrency", out_cCurrency.CreateFunction)
    assert.NotNil(t, out_cCurrency.UpdateDatetime)
    assert.NotEqual(t, "", out_cCurrency.UpdateDatetime)
    assert.Equal(t, "CreateCCurrency", out_cCurrency.UpdateFunction)
}

func TestUpdateCCurrency(t *testing.T) {
    ctx := context.Background()
    repository := &cCurrencyMockRepository{
        FakeUpdate: func(cCurrency *model.CCurrency) (*model.CCurrency, *errordef.LogicError) {
            return cCurrency, nil
        },
        FakeGet: func(currencyCd int, languageCd int) ([]*model.CCurrency, *errordef.LogicError) {
            return []*model.CCurrency{&model.CCurrency{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cCurrencyMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCurrencyService := NewCCurrencyService(repository, numberUtil)
    in_cCurrency := new(model.CCurrency)
    in_cCurrency.CurrencyCd = 0
    in_cCurrency.LanguageCd = 1
    in_cCurrency.CurrencyName = "dummy-CurrencyName"
    in_cCurrency.CurrencyCdIso = "XXX"
    in_cCurrency.SignificantDigit = 4
    out_cCurrency, err := cCurrencyService.UpdateCCurrency(&ctx, in_cCurrency)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCurrency.CurrencyCd)
    assert.Equal(t, 1, out_cCurrency.LanguageCd)
    assert.Equal(t, "dummy-CurrencyName", out_cCurrency.CurrencyName)
    assert.Equal(t, "XXX", out_cCurrency.CurrencyCdIso)
    assert.Equal(t, 4, out_cCurrency.SignificantDigit)
    assert.NotNil(t, out_cCurrency.CreateDatetime)
    assert.Equal(t, "", out_cCurrency.CreateDatetime)
    assert.Equal(t, "", out_cCurrency.CreateFunction)
    assert.NotNil(t, out_cCurrency.UpdateDatetime)
    assert.NotEqual(t, "", out_cCurrency.UpdateDatetime)
    assert.Equal(t, "UpdateCCurrency", out_cCurrency.UpdateFunction)
}
