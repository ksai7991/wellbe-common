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

type cPayoutMethodMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError)
    FakeUpdate func(*model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.CPayoutMethod, *errordef.LogicError)
}

func (lr cPayoutMethodMockRepository) CreateCPayoutMethod(ctx *context.Context, cPayoutMethod *model.CPayoutMethod)  (*model.CPayoutMethod, *errordef.LogicError) {
    return lr.FakeCreate(cPayoutMethod)
}

func (lr cPayoutMethodMockRepository) UpdateCPayoutMethod(ctx *context.Context, cPayoutMethod *model.CPayoutMethod)  (*model.CPayoutMethod, *errordef.LogicError) {
    return lr.FakeUpdate(cPayoutMethod)
}

func (lr cPayoutMethodMockRepository) DeleteCPayoutMethod(ctx *context.Context, payoutMethodCd int)  *errordef.LogicError {
    return lr.FakeDelete(payoutMethodCd)
}

func (lr cPayoutMethodMockRepository)GetCPayoutMethodWithKey(ctx *context.Context, payoutMethodCd int)  ([]*model.CPayoutMethod, *errordef.LogicError) {
    return lr.FakeGet(payoutMethodCd)
}


type cPayoutMethodMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cPayoutMethodMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCPayoutMethod(t *testing.T) {
    ctx := context.Background()
    repository := &cPayoutMethodMockRepository{
        FakeCreate: func(cPayoutMethod *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError) {
            return cPayoutMethod, nil
        },
    }
    numberUtil := &cPayoutMethodMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cPayoutMethodService := NewCPayoutMethodService(repository, numberUtil)
    in_cPayoutMethod := new(model.CPayoutMethod)
    in_cPayoutMethod.PayoutMethodCd = 0
    in_cPayoutMethod.LanguageCd = 1
    in_cPayoutMethod.PayoutMethodName = "dummy-PayoutMethodName"
    out_cPayoutMethod, err := cPayoutMethodService.CreateCPayoutMethod(&ctx, in_cPayoutMethod)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cPayoutMethod.PayoutMethodCd)
    assert.Equal(t, 1, out_cPayoutMethod.LanguageCd)
    assert.Equal(t, "dummy-PayoutMethodName", out_cPayoutMethod.PayoutMethodName)
    assert.NotNil(t, out_cPayoutMethod.CreateDatetime)
    assert.NotEqual(t, "", out_cPayoutMethod.CreateDatetime)
    assert.Equal(t, "CreateCPayoutMethod", out_cPayoutMethod.CreateFunction)
    assert.NotNil(t, out_cPayoutMethod.UpdateDatetime)
    assert.NotEqual(t, "", out_cPayoutMethod.UpdateDatetime)
    assert.Equal(t, "CreateCPayoutMethod", out_cPayoutMethod.UpdateFunction)
}

func TestUpdateCPayoutMethod(t *testing.T) {
    ctx := context.Background()
    repository := &cPayoutMethodMockRepository{
        FakeUpdate: func(cPayoutMethod *model.CPayoutMethod) (*model.CPayoutMethod, *errordef.LogicError) {
            return cPayoutMethod, nil
        },
        FakeGet: func(payoutMethodCd int) ([]*model.CPayoutMethod, *errordef.LogicError) {
            return []*model.CPayoutMethod{&model.CPayoutMethod{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cPayoutMethodMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cPayoutMethodService := NewCPayoutMethodService(repository, numberUtil)
    in_cPayoutMethod := new(model.CPayoutMethod)
    in_cPayoutMethod.PayoutMethodCd = 0
    in_cPayoutMethod.LanguageCd = 1
    in_cPayoutMethod.PayoutMethodName = "dummy-PayoutMethodName"
    out_cPayoutMethod, err := cPayoutMethodService.UpdateCPayoutMethod(&ctx, in_cPayoutMethod)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cPayoutMethod.PayoutMethodCd)
    assert.Equal(t, 1, out_cPayoutMethod.LanguageCd)
    assert.Equal(t, "dummy-PayoutMethodName", out_cPayoutMethod.PayoutMethodName)
    assert.NotNil(t, out_cPayoutMethod.CreateDatetime)
    assert.Equal(t, "", out_cPayoutMethod.CreateDatetime)
    assert.Equal(t, "", out_cPayoutMethod.CreateFunction)
    assert.NotNil(t, out_cPayoutMethod.UpdateDatetime)
    assert.NotEqual(t, "", out_cPayoutMethod.UpdateDatetime)
    assert.Equal(t, "UpdateCPayoutMethod", out_cPayoutMethod.UpdateFunction)
}
