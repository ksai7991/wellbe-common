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

type cCheckoutTimingMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError)
    FakeUpdate func(*model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CCheckoutTiming, *errordef.LogicError)
}

func (lr cCheckoutTimingMockRepository) CreateCCheckoutTiming(ctx *context.Context, cCheckoutTiming *model.CCheckoutTiming)  (*model.CCheckoutTiming, *errordef.LogicError) {
    return lr.FakeCreate(cCheckoutTiming)
}

func (lr cCheckoutTimingMockRepository) UpdateCCheckoutTiming(ctx *context.Context, cCheckoutTiming *model.CCheckoutTiming)  (*model.CCheckoutTiming, *errordef.LogicError) {
    return lr.FakeUpdate(cCheckoutTiming)
}

func (lr cCheckoutTimingMockRepository) DeleteCCheckoutTiming(ctx *context.Context, checkoutTimingCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(checkoutTimingCd, languageCd)
}

func (lr cCheckoutTimingMockRepository)GetCCheckoutTimingWithKey(ctx *context.Context, checkoutTimingCd int, languageCd int)  ([]*model.CCheckoutTiming, *errordef.LogicError) {
    return lr.FakeGet(checkoutTimingCd, languageCd)
}


type cCheckoutTimingMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cCheckoutTimingMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCCheckoutTiming(t *testing.T) {
    ctx := context.Background()
    repository := &cCheckoutTimingMockRepository{
        FakeCreate: func(cCheckoutTiming *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError) {
            return cCheckoutTiming, nil
        },
    }
    numberUtil := &cCheckoutTimingMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCheckoutTimingService := NewCCheckoutTimingService(repository, numberUtil)
    in_cCheckoutTiming := new(model.CCheckoutTiming)
    in_cCheckoutTiming.CheckoutTimingCd = 0
    in_cCheckoutTiming.LanguageCd = 1
    in_cCheckoutTiming.CheckoutTimingName = "dummy-CheckoutTimingName"
    out_cCheckoutTiming, err := cCheckoutTimingService.CreateCCheckoutTiming(&ctx, in_cCheckoutTiming)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCheckoutTiming.CheckoutTimingCd)
    assert.Equal(t, 1, out_cCheckoutTiming.LanguageCd)
    assert.Equal(t, "dummy-CheckoutTimingName", out_cCheckoutTiming.CheckoutTimingName)
    assert.NotNil(t, out_cCheckoutTiming.CreateDatetime)
    assert.NotEqual(t, "", out_cCheckoutTiming.CreateDatetime)
    assert.Equal(t, "CreateCCheckoutTiming", out_cCheckoutTiming.CreateFunction)
    assert.NotNil(t, out_cCheckoutTiming.UpdateDatetime)
    assert.NotEqual(t, "", out_cCheckoutTiming.UpdateDatetime)
    assert.Equal(t, "CreateCCheckoutTiming", out_cCheckoutTiming.UpdateFunction)
}

func TestUpdateCCheckoutTiming(t *testing.T) {
    ctx := context.Background()
    repository := &cCheckoutTimingMockRepository{
        FakeUpdate: func(cCheckoutTiming *model.CCheckoutTiming) (*model.CCheckoutTiming, *errordef.LogicError) {
            return cCheckoutTiming, nil
        },
        FakeGet: func(checkoutTimingCd int, languageCd int) ([]*model.CCheckoutTiming, *errordef.LogicError) {
            return []*model.CCheckoutTiming{&model.CCheckoutTiming{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cCheckoutTimingMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCheckoutTimingService := NewCCheckoutTimingService(repository, numberUtil)
    in_cCheckoutTiming := new(model.CCheckoutTiming)
    in_cCheckoutTiming.CheckoutTimingCd = 0
    in_cCheckoutTiming.LanguageCd = 1
    in_cCheckoutTiming.CheckoutTimingName = "dummy-CheckoutTimingName"
    out_cCheckoutTiming, err := cCheckoutTimingService.UpdateCCheckoutTiming(&ctx, in_cCheckoutTiming)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCheckoutTiming.CheckoutTimingCd)
    assert.Equal(t, 1, out_cCheckoutTiming.LanguageCd)
    assert.Equal(t, "dummy-CheckoutTimingName", out_cCheckoutTiming.CheckoutTimingName)
    assert.NotNil(t, out_cCheckoutTiming.CreateDatetime)
    assert.Equal(t, "", out_cCheckoutTiming.CreateDatetime)
    assert.Equal(t, "", out_cCheckoutTiming.CreateFunction)
    assert.NotNil(t, out_cCheckoutTiming.UpdateDatetime)
    assert.NotEqual(t, "", out_cCheckoutTiming.UpdateDatetime)
    assert.Equal(t, "UpdateCCheckoutTiming", out_cCheckoutTiming.UpdateFunction)
}
