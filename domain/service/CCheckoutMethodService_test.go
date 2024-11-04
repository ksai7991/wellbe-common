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

type cCheckoutMethodMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError)
    FakeUpdate func(*model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CCheckoutMethod, *errordef.LogicError)
}

func (lr cCheckoutMethodMockRepository) CreateCCheckoutMethod(ctx *context.Context, cCheckoutMethod *model.CCheckoutMethod)  (*model.CCheckoutMethod, *errordef.LogicError) {
    return lr.FakeCreate(cCheckoutMethod)
}

func (lr cCheckoutMethodMockRepository) UpdateCCheckoutMethod(ctx *context.Context, cCheckoutMethod *model.CCheckoutMethod)  (*model.CCheckoutMethod, *errordef.LogicError) {
    return lr.FakeUpdate(cCheckoutMethod)
}

func (lr cCheckoutMethodMockRepository) DeleteCCheckoutMethod(ctx *context.Context, checkoutMethodCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(checkoutMethodCd, languageCd)
}

func (lr cCheckoutMethodMockRepository)GetCCheckoutMethodWithKey(ctx *context.Context, checkoutMethodCd int, languageCd int)  ([]*model.CCheckoutMethod, *errordef.LogicError) {
    return lr.FakeGet(checkoutMethodCd, languageCd)
}


type cCheckoutMethodMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cCheckoutMethodMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCCheckoutMethod(t *testing.T) {
    ctx := context.Background()
    repository := &cCheckoutMethodMockRepository{
        FakeCreate: func(cCheckoutMethod *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError) {
            return cCheckoutMethod, nil
        },
    }
    numberUtil := &cCheckoutMethodMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCheckoutMethodService := NewCCheckoutMethodService(repository, numberUtil)
    in_cCheckoutMethod := new(model.CCheckoutMethod)
    in_cCheckoutMethod.CheckoutMethodCd = 0
    in_cCheckoutMethod.LanguageCd = 1
    in_cCheckoutMethod.CheckoutMethodName = "dummy-CheckoutMethodName"
    out_cCheckoutMethod, err := cCheckoutMethodService.CreateCCheckoutMethod(&ctx, in_cCheckoutMethod)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCheckoutMethod.CheckoutMethodCd)
    assert.Equal(t, 1, out_cCheckoutMethod.LanguageCd)
    assert.Equal(t, "dummy-CheckoutMethodName", out_cCheckoutMethod.CheckoutMethodName)
    assert.NotNil(t, out_cCheckoutMethod.CreateDatetime)
    assert.NotEqual(t, "", out_cCheckoutMethod.CreateDatetime)
    assert.Equal(t, "CreateCCheckoutMethod", out_cCheckoutMethod.CreateFunction)
    assert.NotNil(t, out_cCheckoutMethod.UpdateDatetime)
    assert.NotEqual(t, "", out_cCheckoutMethod.UpdateDatetime)
    assert.Equal(t, "CreateCCheckoutMethod", out_cCheckoutMethod.UpdateFunction)
}

func TestUpdateCCheckoutMethod(t *testing.T) {
    ctx := context.Background()
    repository := &cCheckoutMethodMockRepository{
        FakeUpdate: func(cCheckoutMethod *model.CCheckoutMethod) (*model.CCheckoutMethod, *errordef.LogicError) {
            return cCheckoutMethod, nil
        },
        FakeGet: func(checkoutMethodCd int, languageCd int) ([]*model.CCheckoutMethod, *errordef.LogicError) {
            return []*model.CCheckoutMethod{&model.CCheckoutMethod{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cCheckoutMethodMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCheckoutMethodService := NewCCheckoutMethodService(repository, numberUtil)
    in_cCheckoutMethod := new(model.CCheckoutMethod)
    in_cCheckoutMethod.CheckoutMethodCd = 0
    in_cCheckoutMethod.LanguageCd = 1
    in_cCheckoutMethod.CheckoutMethodName = "dummy-CheckoutMethodName"
    out_cCheckoutMethod, err := cCheckoutMethodService.UpdateCCheckoutMethod(&ctx, in_cCheckoutMethod)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCheckoutMethod.CheckoutMethodCd)
    assert.Equal(t, 1, out_cCheckoutMethod.LanguageCd)
    assert.Equal(t, "dummy-CheckoutMethodName", out_cCheckoutMethod.CheckoutMethodName)
    assert.NotNil(t, out_cCheckoutMethod.CreateDatetime)
    assert.Equal(t, "", out_cCheckoutMethod.CreateDatetime)
    assert.Equal(t, "", out_cCheckoutMethod.CreateFunction)
    assert.NotNil(t, out_cCheckoutMethod.UpdateDatetime)
    assert.NotEqual(t, "", out_cCheckoutMethod.UpdateDatetime)
    assert.Equal(t, "UpdateCCheckoutMethod", out_cCheckoutMethod.UpdateFunction)
}
