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

type cShopPaymentMethodMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError)
    FakeUpdate func(*model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.CShopPaymentMethod, *errordef.LogicError)
}

func (lr cShopPaymentMethodMockRepository) CreateCShopPaymentMethod(ctx *context.Context, cShopPaymentMethod *model.CShopPaymentMethod)  (*model.CShopPaymentMethod, *errordef.LogicError) {
    return lr.FakeCreate(cShopPaymentMethod)
}

func (lr cShopPaymentMethodMockRepository) UpdateCShopPaymentMethod(ctx *context.Context, cShopPaymentMethod *model.CShopPaymentMethod)  (*model.CShopPaymentMethod, *errordef.LogicError) {
    return lr.FakeUpdate(cShopPaymentMethod)
}

func (lr cShopPaymentMethodMockRepository) DeleteCShopPaymentMethod(ctx *context.Context, shopPaymentMethodCd int)  *errordef.LogicError {
    return lr.FakeDelete(shopPaymentMethodCd)
}

func (lr cShopPaymentMethodMockRepository)GetCShopPaymentMethodWithKey(ctx *context.Context, shopPaymentMethodCd int)  ([]*model.CShopPaymentMethod, *errordef.LogicError) {
    return lr.FakeGet(shopPaymentMethodCd)
}


type cShopPaymentMethodMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cShopPaymentMethodMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCShopPaymentMethod(t *testing.T) {
    ctx := context.Background()
    repository := &cShopPaymentMethodMockRepository{
        FakeCreate: func(cShopPaymentMethod *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError) {
            return cShopPaymentMethod, nil
        },
    }
    numberUtil := &cShopPaymentMethodMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopPaymentMethodService := NewCShopPaymentMethodService(repository, numberUtil)
    in_cShopPaymentMethod := new(model.CShopPaymentMethod)
    in_cShopPaymentMethod.ShopPaymentMethodCd = 0
    in_cShopPaymentMethod.LanguageCd = 1
    in_cShopPaymentMethod.ShopPaymentName = "dummy-ShopPaymentName"
    out_cShopPaymentMethod, err := cShopPaymentMethodService.CreateCShopPaymentMethod(&ctx, in_cShopPaymentMethod)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopPaymentMethod.ShopPaymentMethodCd)
    assert.Equal(t, 1, out_cShopPaymentMethod.LanguageCd)
    assert.Equal(t, "dummy-ShopPaymentName", out_cShopPaymentMethod.ShopPaymentName)
    assert.NotNil(t, out_cShopPaymentMethod.CreateDatetime)
    assert.NotEqual(t, "", out_cShopPaymentMethod.CreateDatetime)
    assert.Equal(t, "CreateCShopPaymentMethod", out_cShopPaymentMethod.CreateFunction)
    assert.NotNil(t, out_cShopPaymentMethod.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopPaymentMethod.UpdateDatetime)
    assert.Equal(t, "CreateCShopPaymentMethod", out_cShopPaymentMethod.UpdateFunction)
}

func TestUpdateCShopPaymentMethod(t *testing.T) {
    ctx := context.Background()
    repository := &cShopPaymentMethodMockRepository{
        FakeUpdate: func(cShopPaymentMethod *model.CShopPaymentMethod) (*model.CShopPaymentMethod, *errordef.LogicError) {
            return cShopPaymentMethod, nil
        },
        FakeGet: func(shopPaymentMethodCd int) ([]*model.CShopPaymentMethod, *errordef.LogicError) {
            return []*model.CShopPaymentMethod{&model.CShopPaymentMethod{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cShopPaymentMethodMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopPaymentMethodService := NewCShopPaymentMethodService(repository, numberUtil)
    in_cShopPaymentMethod := new(model.CShopPaymentMethod)
    in_cShopPaymentMethod.ShopPaymentMethodCd = 0
    in_cShopPaymentMethod.LanguageCd = 1
    in_cShopPaymentMethod.ShopPaymentName = "dummy-ShopPaymentName"
    out_cShopPaymentMethod, err := cShopPaymentMethodService.UpdateCShopPaymentMethod(&ctx, in_cShopPaymentMethod)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopPaymentMethod.ShopPaymentMethodCd)
    assert.Equal(t, 1, out_cShopPaymentMethod.LanguageCd)
    assert.Equal(t, "dummy-ShopPaymentName", out_cShopPaymentMethod.ShopPaymentName)
    assert.NotNil(t, out_cShopPaymentMethod.CreateDatetime)
    assert.Equal(t, "", out_cShopPaymentMethod.CreateDatetime)
    assert.Equal(t, "", out_cShopPaymentMethod.CreateFunction)
    assert.NotNil(t, out_cShopPaymentMethod.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopPaymentMethod.UpdateDatetime)
    assert.Equal(t, "UpdateCShopPaymentMethod", out_cShopPaymentMethod.UpdateFunction)
}
