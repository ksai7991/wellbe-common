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

type cShopImageFilterCategoryMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError)
    FakeUpdate func(*model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CShopImageFilterCategory, *errordef.LogicError)
}

func (lr cShopImageFilterCategoryMockRepository) CreateCShopImageFilterCategory(ctx *context.Context, cShopImageFilterCategory *model.CShopImageFilterCategory)  (*model.CShopImageFilterCategory, *errordef.LogicError) {
    return lr.FakeCreate(cShopImageFilterCategory)
}

func (lr cShopImageFilterCategoryMockRepository) UpdateCShopImageFilterCategory(ctx *context.Context, cShopImageFilterCategory *model.CShopImageFilterCategory)  (*model.CShopImageFilterCategory, *errordef.LogicError) {
    return lr.FakeUpdate(cShopImageFilterCategory)
}

func (lr cShopImageFilterCategoryMockRepository) DeleteCShopImageFilterCategory(ctx *context.Context, shopImageFilterCategoryCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(shopImageFilterCategoryCd, languageCd)
}

func (lr cShopImageFilterCategoryMockRepository)GetCShopImageFilterCategoryWithKey(ctx *context.Context, shopImageFilterCategoryCd int, languageCd int)  ([]*model.CShopImageFilterCategory, *errordef.LogicError) {
    return lr.FakeGet(shopImageFilterCategoryCd, languageCd)
}


type cShopImageFilterCategoryMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cShopImageFilterCategoryMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCShopImageFilterCategory(t *testing.T) {
    ctx := context.Background()
    repository := &cShopImageFilterCategoryMockRepository{
        FakeCreate: func(cShopImageFilterCategory *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError) {
            return cShopImageFilterCategory, nil
        },
    }
    numberUtil := &cShopImageFilterCategoryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopImageFilterCategoryService := NewCShopImageFilterCategoryService(repository, numberUtil)
    in_cShopImageFilterCategory := new(model.CShopImageFilterCategory)
    in_cShopImageFilterCategory.ShopImageFilterCategoryCd = 0
    in_cShopImageFilterCategory.LanguageCd = 1
    in_cShopImageFilterCategory.ShopImageFilterCategoryName = "dummy-ShopImageFilterCategoryName"
    out_cShopImageFilterCategory, err := cShopImageFilterCategoryService.CreateCShopImageFilterCategory(&ctx, in_cShopImageFilterCategory)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopImageFilterCategory.ShopImageFilterCategoryCd)
    assert.Equal(t, 1, out_cShopImageFilterCategory.LanguageCd)
    assert.Equal(t, "dummy-ShopImageFilterCategoryName", out_cShopImageFilterCategory.ShopImageFilterCategoryName)
    assert.NotNil(t, out_cShopImageFilterCategory.CreateDatetime)
    assert.NotEqual(t, "", out_cShopImageFilterCategory.CreateDatetime)
    assert.Equal(t, "CreateCShopImageFilterCategory", out_cShopImageFilterCategory.CreateFunction)
    assert.NotNil(t, out_cShopImageFilterCategory.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopImageFilterCategory.UpdateDatetime)
    assert.Equal(t, "CreateCShopImageFilterCategory", out_cShopImageFilterCategory.UpdateFunction)
}

func TestUpdateCShopImageFilterCategory(t *testing.T) {
    ctx := context.Background()
    repository := &cShopImageFilterCategoryMockRepository{
        FakeUpdate: func(cShopImageFilterCategory *model.CShopImageFilterCategory) (*model.CShopImageFilterCategory, *errordef.LogicError) {
            return cShopImageFilterCategory, nil
        },
        FakeGet: func(shopImageFilterCategoryCd int, languageCd int) ([]*model.CShopImageFilterCategory, *errordef.LogicError) {
            return []*model.CShopImageFilterCategory{&model.CShopImageFilterCategory{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cShopImageFilterCategoryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopImageFilterCategoryService := NewCShopImageFilterCategoryService(repository, numberUtil)
    in_cShopImageFilterCategory := new(model.CShopImageFilterCategory)
    in_cShopImageFilterCategory.ShopImageFilterCategoryCd = 0
    in_cShopImageFilterCategory.LanguageCd = 1
    in_cShopImageFilterCategory.ShopImageFilterCategoryName = "dummy-ShopImageFilterCategoryName"
    out_cShopImageFilterCategory, err := cShopImageFilterCategoryService.UpdateCShopImageFilterCategory(&ctx, in_cShopImageFilterCategory)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopImageFilterCategory.ShopImageFilterCategoryCd)
    assert.Equal(t, 1, out_cShopImageFilterCategory.LanguageCd)
    assert.Equal(t, "dummy-ShopImageFilterCategoryName", out_cShopImageFilterCategory.ShopImageFilterCategoryName)
    assert.NotNil(t, out_cShopImageFilterCategory.CreateDatetime)
    assert.Equal(t, "", out_cShopImageFilterCategory.CreateDatetime)
    assert.Equal(t, "", out_cShopImageFilterCategory.CreateFunction)
    assert.NotNil(t, out_cShopImageFilterCategory.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopImageFilterCategory.UpdateDatetime)
    assert.Equal(t, "UpdateCShopImageFilterCategory", out_cShopImageFilterCategory.UpdateFunction)
}
