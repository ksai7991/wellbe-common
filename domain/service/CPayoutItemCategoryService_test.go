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

type cPayoutItemCategoryMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError)
    FakeUpdate func(*model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.CPayoutItemCategory, *errordef.LogicError)
}

func (lr cPayoutItemCategoryMockRepository) CreateCPayoutItemCategory(ctx *context.Context, cPayoutItemCategory *model.CPayoutItemCategory)  (*model.CPayoutItemCategory, *errordef.LogicError) {
    return lr.FakeCreate(cPayoutItemCategory)
}

func (lr cPayoutItemCategoryMockRepository) UpdateCPayoutItemCategory(ctx *context.Context, cPayoutItemCategory *model.CPayoutItemCategory)  (*model.CPayoutItemCategory, *errordef.LogicError) {
    return lr.FakeUpdate(cPayoutItemCategory)
}

func (lr cPayoutItemCategoryMockRepository) DeleteCPayoutItemCategory(ctx *context.Context, payoutItemCategoryCd int)  *errordef.LogicError {
    return lr.FakeDelete(payoutItemCategoryCd)
}

func (lr cPayoutItemCategoryMockRepository)GetCPayoutItemCategoryWithKey(ctx *context.Context, payoutItemCategoryCd int)  ([]*model.CPayoutItemCategory, *errordef.LogicError) {
    return lr.FakeGet(payoutItemCategoryCd)
}


type cPayoutItemCategoryMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cPayoutItemCategoryMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCPayoutItemCategory(t *testing.T) {
    ctx := context.Background()
    repository := &cPayoutItemCategoryMockRepository{
        FakeCreate: func(cPayoutItemCategory *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError) {
            return cPayoutItemCategory, nil
        },
    }
    numberUtil := &cPayoutItemCategoryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cPayoutItemCategoryService := NewCPayoutItemCategoryService(repository, numberUtil)
    in_cPayoutItemCategory := new(model.CPayoutItemCategory)
    in_cPayoutItemCategory.PayoutItemCategoryCd = 0
    in_cPayoutItemCategory.LanguageCd = 1
    in_cPayoutItemCategory.PayoutItemCategoryName = "dummy-PayoutItemCategoryName"
    out_cPayoutItemCategory, err := cPayoutItemCategoryService.CreateCPayoutItemCategory(&ctx, in_cPayoutItemCategory)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cPayoutItemCategory.PayoutItemCategoryCd)
    assert.Equal(t, 1, out_cPayoutItemCategory.LanguageCd)
    assert.Equal(t, "dummy-PayoutItemCategoryName", out_cPayoutItemCategory.PayoutItemCategoryName)
    assert.NotNil(t, out_cPayoutItemCategory.CreateDatetime)
    assert.NotEqual(t, "", out_cPayoutItemCategory.CreateDatetime)
    assert.Equal(t, "CreateCPayoutItemCategory", out_cPayoutItemCategory.CreateFunction)
    assert.NotNil(t, out_cPayoutItemCategory.UpdateDatetime)
    assert.NotEqual(t, "", out_cPayoutItemCategory.UpdateDatetime)
    assert.Equal(t, "CreateCPayoutItemCategory", out_cPayoutItemCategory.UpdateFunction)
}

func TestUpdateCPayoutItemCategory(t *testing.T) {
    ctx := context.Background()
    repository := &cPayoutItemCategoryMockRepository{
        FakeUpdate: func(cPayoutItemCategory *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError) {
            return cPayoutItemCategory, nil
        },
        FakeGet: func(payoutItemCategoryCd int) ([]*model.CPayoutItemCategory, *errordef.LogicError) {
            return []*model.CPayoutItemCategory{&model.CPayoutItemCategory{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cPayoutItemCategoryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cPayoutItemCategoryService := NewCPayoutItemCategoryService(repository, numberUtil)
    in_cPayoutItemCategory := new(model.CPayoutItemCategory)
    in_cPayoutItemCategory.PayoutItemCategoryCd = 0
    in_cPayoutItemCategory.LanguageCd = 1
    in_cPayoutItemCategory.PayoutItemCategoryName = "dummy-PayoutItemCategoryName"
    out_cPayoutItemCategory, err := cPayoutItemCategoryService.UpdateCPayoutItemCategory(&ctx, in_cPayoutItemCategory)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cPayoutItemCategory.PayoutItemCategoryCd)
    assert.Equal(t, 1, out_cPayoutItemCategory.LanguageCd)
    assert.Equal(t, "dummy-PayoutItemCategoryName", out_cPayoutItemCategory.PayoutItemCategoryName)
    assert.NotNil(t, out_cPayoutItemCategory.CreateDatetime)
    assert.Equal(t, "", out_cPayoutItemCategory.CreateDatetime)
    assert.Equal(t, "", out_cPayoutItemCategory.CreateFunction)
    assert.NotNil(t, out_cPayoutItemCategory.UpdateDatetime)
    assert.NotEqual(t, "", out_cPayoutItemCategory.UpdateDatetime)
    assert.Equal(t, "UpdateCPayoutItemCategory", out_cPayoutItemCategory.UpdateFunction)
}
