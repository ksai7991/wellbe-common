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

type cShopContractPlanItemMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError)
    FakeUpdate func(*model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CShopContractPlanItem, *errordef.LogicError)
}

func (lr cShopContractPlanItemMockRepository) CreateCShopContractPlanItem(ctx *context.Context, cShopContractPlanItem *model.CShopContractPlanItem)  (*model.CShopContractPlanItem, *errordef.LogicError) {
    return lr.FakeCreate(cShopContractPlanItem)
}

func (lr cShopContractPlanItemMockRepository) UpdateCShopContractPlanItem(ctx *context.Context, cShopContractPlanItem *model.CShopContractPlanItem)  (*model.CShopContractPlanItem, *errordef.LogicError) {
    return lr.FakeUpdate(cShopContractPlanItem)
}

func (lr cShopContractPlanItemMockRepository) DeleteCShopContractPlanItem(ctx *context.Context, shopContractPlanItemCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(shopContractPlanItemCd, languageCd)
}

func (lr cShopContractPlanItemMockRepository)GetCShopContractPlanItemWithKey(ctx *context.Context, shopContractPlanItemCd int, languageCd int)  ([]*model.CShopContractPlanItem, *errordef.LogicError) {
    return lr.FakeGet(shopContractPlanItemCd, languageCd)
}


type cShopContractPlanItemMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cShopContractPlanItemMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCShopContractPlanItem(t *testing.T) {
    ctx := context.Background()
    repository := &cShopContractPlanItemMockRepository{
        FakeCreate: func(cShopContractPlanItem *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError) {
            return cShopContractPlanItem, nil
        },
    }
    numberUtil := &cShopContractPlanItemMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopContractPlanItemService := NewCShopContractPlanItemService(repository, numberUtil)
    in_cShopContractPlanItem := new(model.CShopContractPlanItem)
    in_cShopContractPlanItem.ShopContractPlanItemCd = 0
    in_cShopContractPlanItem.LanguageCd = 1
    in_cShopContractPlanItem.ShopContractPlanName = "dummy-ShopContractPlanName"
    in_cShopContractPlanItem.Unit = "dummy-Unit"
    out_cShopContractPlanItem, err := cShopContractPlanItemService.CreateCShopContractPlanItem(&ctx, in_cShopContractPlanItem)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopContractPlanItem.ShopContractPlanItemCd)
    assert.Equal(t, 1, out_cShopContractPlanItem.LanguageCd)
    assert.Equal(t, "dummy-ShopContractPlanName", out_cShopContractPlanItem.ShopContractPlanName)
    assert.Equal(t, "dummy-Unit", out_cShopContractPlanItem.Unit)
    assert.NotNil(t, out_cShopContractPlanItem.CreateDatetime)
    assert.NotEqual(t, "", out_cShopContractPlanItem.CreateDatetime)
    assert.Equal(t, "CreateCShopContractPlanItem", out_cShopContractPlanItem.CreateFunction)
    assert.NotNil(t, out_cShopContractPlanItem.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopContractPlanItem.UpdateDatetime)
    assert.Equal(t, "CreateCShopContractPlanItem", out_cShopContractPlanItem.UpdateFunction)
}

func TestUpdateCShopContractPlanItem(t *testing.T) {
    ctx := context.Background()
    repository := &cShopContractPlanItemMockRepository{
        FakeUpdate: func(cShopContractPlanItem *model.CShopContractPlanItem) (*model.CShopContractPlanItem, *errordef.LogicError) {
            return cShopContractPlanItem, nil
        },
        FakeGet: func(shopContractPlanItemCd int, languageCd int) ([]*model.CShopContractPlanItem, *errordef.LogicError) {
            return []*model.CShopContractPlanItem{&model.CShopContractPlanItem{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cShopContractPlanItemMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopContractPlanItemService := NewCShopContractPlanItemService(repository, numberUtil)
    in_cShopContractPlanItem := new(model.CShopContractPlanItem)
    in_cShopContractPlanItem.ShopContractPlanItemCd = 0
    in_cShopContractPlanItem.LanguageCd = 1
    in_cShopContractPlanItem.ShopContractPlanName = "dummy-ShopContractPlanName"
    in_cShopContractPlanItem.Unit = "dummy-Unit"
    out_cShopContractPlanItem, err := cShopContractPlanItemService.UpdateCShopContractPlanItem(&ctx, in_cShopContractPlanItem)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopContractPlanItem.ShopContractPlanItemCd)
    assert.Equal(t, 1, out_cShopContractPlanItem.LanguageCd)
    assert.Equal(t, "dummy-ShopContractPlanName", out_cShopContractPlanItem.ShopContractPlanName)
    assert.Equal(t, "dummy-Unit", out_cShopContractPlanItem.Unit)
    assert.NotNil(t, out_cShopContractPlanItem.CreateDatetime)
    assert.Equal(t, "", out_cShopContractPlanItem.CreateDatetime)
    assert.Equal(t, "", out_cShopContractPlanItem.CreateFunction)
    assert.NotNil(t, out_cShopContractPlanItem.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopContractPlanItem.UpdateDatetime)
    assert.Equal(t, "UpdateCShopContractPlanItem", out_cShopContractPlanItem.UpdateFunction)
}
