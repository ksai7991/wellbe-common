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

type cReviewCategoryMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError)
    FakeUpdate func(*model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CReviewCategory, *errordef.LogicError)
}

func (lr cReviewCategoryMockRepository) CreateCReviewCategory(ctx *context.Context, cReviewCategory *model.CReviewCategory)  (*model.CReviewCategory, *errordef.LogicError) {
    return lr.FakeCreate(cReviewCategory)
}

func (lr cReviewCategoryMockRepository) UpdateCReviewCategory(ctx *context.Context, cReviewCategory *model.CReviewCategory)  (*model.CReviewCategory, *errordef.LogicError) {
    return lr.FakeUpdate(cReviewCategory)
}

func (lr cReviewCategoryMockRepository) DeleteCReviewCategory(ctx *context.Context, reviewCategoryCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(reviewCategoryCd, languageCd)
}

func (lr cReviewCategoryMockRepository)GetCReviewCategoryWithKey(ctx *context.Context, reviewCategoryCd int, languageCd int)  ([]*model.CReviewCategory, *errordef.LogicError) {
    return lr.FakeGet(reviewCategoryCd, languageCd)
}


type cReviewCategoryMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cReviewCategoryMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCReviewCategory(t *testing.T) {
    ctx := context.Background()
    repository := &cReviewCategoryMockRepository{
        FakeCreate: func(cReviewCategory *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError) {
            return cReviewCategory, nil
        },
    }
    numberUtil := &cReviewCategoryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cReviewCategoryService := NewCReviewCategoryService(repository, numberUtil)
    in_cReviewCategory := new(model.CReviewCategory)
    in_cReviewCategory.ReviewCategoryCd = 0
    in_cReviewCategory.LanguageCd = 1
    in_cReviewCategory.ReviewCategoryName = "dummy-ReviewCategoryName"
    out_cReviewCategory, err := cReviewCategoryService.CreateCReviewCategory(&ctx, in_cReviewCategory)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cReviewCategory.ReviewCategoryCd)
    assert.Equal(t, 1, out_cReviewCategory.LanguageCd)
    assert.Equal(t, "dummy-ReviewCategoryName", out_cReviewCategory.ReviewCategoryName)
    assert.NotNil(t, out_cReviewCategory.CreateDatetime)
    assert.NotEqual(t, "", out_cReviewCategory.CreateDatetime)
    assert.Equal(t, "CreateCReviewCategory", out_cReviewCategory.CreateFunction)
    assert.NotNil(t, out_cReviewCategory.UpdateDatetime)
    assert.NotEqual(t, "", out_cReviewCategory.UpdateDatetime)
    assert.Equal(t, "CreateCReviewCategory", out_cReviewCategory.UpdateFunction)
}

func TestUpdateCReviewCategory(t *testing.T) {
    ctx := context.Background()
    repository := &cReviewCategoryMockRepository{
        FakeUpdate: func(cReviewCategory *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError) {
            return cReviewCategory, nil
        },
        FakeGet: func(reviewCategoryCd int, languageCd int) ([]*model.CReviewCategory, *errordef.LogicError) {
            return []*model.CReviewCategory{&model.CReviewCategory{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cReviewCategoryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cReviewCategoryService := NewCReviewCategoryService(repository, numberUtil)
    in_cReviewCategory := new(model.CReviewCategory)
    in_cReviewCategory.ReviewCategoryCd = 0
    in_cReviewCategory.LanguageCd = 1
    in_cReviewCategory.ReviewCategoryName = "dummy-ReviewCategoryName"
    out_cReviewCategory, err := cReviewCategoryService.UpdateCReviewCategory(&ctx, in_cReviewCategory)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cReviewCategory.ReviewCategoryCd)
    assert.Equal(t, 1, out_cReviewCategory.LanguageCd)
    assert.Equal(t, "dummy-ReviewCategoryName", out_cReviewCategory.ReviewCategoryName)
    assert.NotNil(t, out_cReviewCategory.CreateDatetime)
    assert.Equal(t, "", out_cReviewCategory.CreateDatetime)
    assert.Equal(t, "", out_cReviewCategory.CreateFunction)
    assert.NotNil(t, out_cReviewCategory.UpdateDatetime)
    assert.NotEqual(t, "", out_cReviewCategory.UpdateDatetime)
    assert.Equal(t, "UpdateCReviewCategory", out_cReviewCategory.UpdateFunction)
}
