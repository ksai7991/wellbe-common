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

type cContentsCategoryMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError)
    FakeUpdate func(*model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CContentsCategory, *errordef.LogicError)
}

func (lr cContentsCategoryMockRepository) CreateCContentsCategory(ctx *context.Context, cContentsCategory *model.CContentsCategory)  (*model.CContentsCategory, *errordef.LogicError) {
    return lr.FakeCreate(cContentsCategory)
}

func (lr cContentsCategoryMockRepository) UpdateCContentsCategory(ctx *context.Context, cContentsCategory *model.CContentsCategory)  (*model.CContentsCategory, *errordef.LogicError) {
    return lr.FakeUpdate(cContentsCategory)
}

func (lr cContentsCategoryMockRepository) DeleteCContentsCategory(ctx *context.Context, contentsCategoryCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(contentsCategoryCd, languageCd)
}

func (lr cContentsCategoryMockRepository)GetCContentsCategoryWithKey(ctx *context.Context, contentsCategoryCd int, languageCd int)  ([]*model.CContentsCategory, *errordef.LogicError) {
    return lr.FakeGet(contentsCategoryCd, languageCd)
}


type cContentsCategoryMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cContentsCategoryMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCContentsCategory(t *testing.T) {
    ctx := context.Background()
    repository := &cContentsCategoryMockRepository{
        FakeCreate: func(cContentsCategory *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError) {
            return cContentsCategory, nil
        },
    }
    numberUtil := &cContentsCategoryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cContentsCategoryService := NewCContentsCategoryService(repository, numberUtil)
    in_cContentsCategory := new(model.CContentsCategory)
    in_cContentsCategory.ContentsCategoryCd = 0
    in_cContentsCategory.LanguageCd = 1
    in_cContentsCategory.ContentsCategoryName = "dummy-ContentsCategoryName"
    out_cContentsCategory, err := cContentsCategoryService.CreateCContentsCategory(&ctx, in_cContentsCategory)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cContentsCategory.ContentsCategoryCd)
    assert.Equal(t, 1, out_cContentsCategory.LanguageCd)
    assert.Equal(t, "dummy-ContentsCategoryName", out_cContentsCategory.ContentsCategoryName)
    assert.NotNil(t, out_cContentsCategory.CreateDatetime)
    assert.NotEqual(t, "", out_cContentsCategory.CreateDatetime)
    assert.Equal(t, "CreateCContentsCategory", out_cContentsCategory.CreateFunction)
    assert.NotNil(t, out_cContentsCategory.UpdateDatetime)
    assert.NotEqual(t, "", out_cContentsCategory.UpdateDatetime)
    assert.Equal(t, "CreateCContentsCategory", out_cContentsCategory.UpdateFunction)
}

func TestUpdateCContentsCategory(t *testing.T) {
    ctx := context.Background()
    repository := &cContentsCategoryMockRepository{
        FakeUpdate: func(cContentsCategory *model.CContentsCategory) (*model.CContentsCategory, *errordef.LogicError) {
            return cContentsCategory, nil
        },
        FakeGet: func(contentsCategoryCd int, languageCd int) ([]*model.CContentsCategory, *errordef.LogicError) {
            return []*model.CContentsCategory{&model.CContentsCategory{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cContentsCategoryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cContentsCategoryService := NewCContentsCategoryService(repository, numberUtil)
    in_cContentsCategory := new(model.CContentsCategory)
    in_cContentsCategory.ContentsCategoryCd = 0
    in_cContentsCategory.LanguageCd = 1
    in_cContentsCategory.ContentsCategoryName = "dummy-ContentsCategoryName"
    out_cContentsCategory, err := cContentsCategoryService.UpdateCContentsCategory(&ctx, in_cContentsCategory)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cContentsCategory.ContentsCategoryCd)
    assert.Equal(t, 1, out_cContentsCategory.LanguageCd)
    assert.Equal(t, "dummy-ContentsCategoryName", out_cContentsCategory.ContentsCategoryName)
    assert.NotNil(t, out_cContentsCategory.CreateDatetime)
    assert.Equal(t, "", out_cContentsCategory.CreateDatetime)
    assert.Equal(t, "", out_cContentsCategory.CreateFunction)
    assert.NotNil(t, out_cContentsCategory.UpdateDatetime)
    assert.NotEqual(t, "", out_cContentsCategory.UpdateDatetime)
    assert.Equal(t, "UpdateCContentsCategory", out_cContentsCategory.UpdateFunction)
}
