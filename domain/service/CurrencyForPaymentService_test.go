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

type currencyForPaymentMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError)
    FakeUpdate func(*model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.CurrencyForPayment, *errordef.LogicError)
}

func (lr currencyForPaymentMockRepository) CreateCurrencyForPayment(ctx *context.Context, currencyForPayment *model.CurrencyForPayment)  (*model.CurrencyForPayment, *errordef.LogicError) {
    return lr.FakeCreate(currencyForPayment)
}

func (lr currencyForPaymentMockRepository) UpdateCurrencyForPayment(ctx *context.Context, currencyForPayment *model.CurrencyForPayment)  (*model.CurrencyForPayment, *errordef.LogicError) {
    return lr.FakeUpdate(currencyForPayment)
}

func (lr currencyForPaymentMockRepository) DeleteCurrencyForPayment(ctx *context.Context, currencyCd int)  *errordef.LogicError {
    return lr.FakeDelete(currencyCd)
}

func (lr currencyForPaymentMockRepository)GetCurrencyForPaymentWithKey(ctx *context.Context, currencyCd int)  ([]*model.CurrencyForPayment, *errordef.LogicError) {
    return lr.FakeGet(currencyCd)
}


type currencyForPaymentMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr currencyForPaymentMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCurrencyForPayment(t *testing.T) {
    ctx := context.Background()
    repository := &currencyForPaymentMockRepository{
        FakeCreate: func(currencyForPayment *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError) {
            return currencyForPayment, nil
        },
    }
    numberUtil := &currencyForPaymentMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    currencyForPaymentService := NewCurrencyForPaymentService(repository, numberUtil)
    in_currencyForPayment := new(model.CurrencyForPayment)
    in_currencyForPayment.CurrencyCd = 0
    out_currencyForPayment, err := currencyForPaymentService.CreateCurrencyForPayment(&ctx, in_currencyForPayment)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_currencyForPayment.CurrencyCd)
    assert.NotNil(t, out_currencyForPayment.CreateDatetime)
    assert.NotEqual(t, "", out_currencyForPayment.CreateDatetime)
    assert.Equal(t, "CreateCurrencyForPayment", out_currencyForPayment.CreateFunction)
    assert.NotNil(t, out_currencyForPayment.UpdateDatetime)
    assert.NotEqual(t, "", out_currencyForPayment.UpdateDatetime)
    assert.Equal(t, "CreateCurrencyForPayment", out_currencyForPayment.UpdateFunction)
}

func TestUpdateCurrencyForPayment(t *testing.T) {
    ctx := context.Background()
    repository := &currencyForPaymentMockRepository{
        FakeUpdate: func(currencyForPayment *model.CurrencyForPayment) (*model.CurrencyForPayment, *errordef.LogicError) {
            return currencyForPayment, nil
        },
        FakeGet: func(currencyCd int) ([]*model.CurrencyForPayment, *errordef.LogicError) {
            return []*model.CurrencyForPayment{&model.CurrencyForPayment{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &currencyForPaymentMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    currencyForPaymentService := NewCurrencyForPaymentService(repository, numberUtil)
    in_currencyForPayment := new(model.CurrencyForPayment)
    in_currencyForPayment.CurrencyCd = 0
    out_currencyForPayment, err := currencyForPaymentService.UpdateCurrencyForPayment(&ctx, in_currencyForPayment)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_currencyForPayment.CurrencyCd)
    assert.NotNil(t, out_currencyForPayment.CreateDatetime)
    assert.Equal(t, "", out_currencyForPayment.CreateDatetime)
    assert.Equal(t, "", out_currencyForPayment.CreateFunction)
    assert.NotNil(t, out_currencyForPayment.UpdateDatetime)
    assert.NotEqual(t, "", out_currencyForPayment.UpdateDatetime)
    assert.Equal(t, "UpdateCurrencyForPayment", out_currencyForPayment.UpdateFunction)
}
