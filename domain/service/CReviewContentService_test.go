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

type cReviewContentMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CReviewContent) (*model.CReviewContent, *errordef.LogicError)
    FakeUpdate func(*model.CReviewContent) (*model.CReviewContent, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CReviewContent, *errordef.LogicError)
}

func (lr cReviewContentMockRepository) CreateCReviewContent(ctx *context.Context, cReviewContent *model.CReviewContent)  (*model.CReviewContent, *errordef.LogicError) {
    return lr.FakeCreate(cReviewContent)
}

func (lr cReviewContentMockRepository) UpdateCReviewContent(ctx *context.Context, cReviewContent *model.CReviewContent)  (*model.CReviewContent, *errordef.LogicError) {
    return lr.FakeUpdate(cReviewContent)
}

func (lr cReviewContentMockRepository) DeleteCReviewContent(ctx *context.Context, reviewContentCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(reviewContentCd, languageCd)
}

func (lr cReviewContentMockRepository)GetCReviewContentWithKey(ctx *context.Context, reviewContentCd int, languageCd int)  ([]*model.CReviewContent, *errordef.LogicError) {
    return lr.FakeGet(reviewContentCd, languageCd)
}


type cReviewContentMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cReviewContentMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCReviewContent(t *testing.T) {
    ctx := context.Background()
    repository := &cReviewContentMockRepository{
        FakeCreate: func(cReviewContent *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError) {
            return cReviewContent, nil
        },
    }
    numberUtil := &cReviewContentMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cReviewContentService := NewCReviewContentService(repository, numberUtil)
    in_cReviewContent := new(model.CReviewContent)
    in_cReviewContent.ReviewContentCd = 0
    in_cReviewContent.LanguageCd = 1
    in_cReviewContent.ReviewCategoryCd = 2
    in_cReviewContent.ReviewContentName = "dummy-ReviewContentName"
    out_cReviewContent, err := cReviewContentService.CreateCReviewContent(&ctx, in_cReviewContent)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cReviewContent.ReviewContentCd)
    assert.Equal(t, 1, out_cReviewContent.LanguageCd)
    assert.Equal(t, 2, out_cReviewContent.ReviewCategoryCd)
    assert.Equal(t, "dummy-ReviewContentName", out_cReviewContent.ReviewContentName)
    assert.NotNil(t, out_cReviewContent.CreateDatetime)
    assert.NotEqual(t, "", out_cReviewContent.CreateDatetime)
    assert.Equal(t, "CreateCReviewContent", out_cReviewContent.CreateFunction)
    assert.NotNil(t, out_cReviewContent.UpdateDatetime)
    assert.NotEqual(t, "", out_cReviewContent.UpdateDatetime)
    assert.Equal(t, "CreateCReviewContent", out_cReviewContent.UpdateFunction)
}

func TestUpdateCReviewContent(t *testing.T) {
    ctx := context.Background()
    repository := &cReviewContentMockRepository{
        FakeUpdate: func(cReviewContent *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError) {
            return cReviewContent, nil
        },
        FakeGet: func(reviewContentCd int, languageCd int) ([]*model.CReviewContent, *errordef.LogicError) {
            return []*model.CReviewContent{&model.CReviewContent{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cReviewContentMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cReviewContentService := NewCReviewContentService(repository, numberUtil)
    in_cReviewContent := new(model.CReviewContent)
    in_cReviewContent.ReviewContentCd = 0
    in_cReviewContent.LanguageCd = 1
    in_cReviewContent.ReviewCategoryCd = 2
    in_cReviewContent.ReviewContentName = "dummy-ReviewContentName"
    out_cReviewContent, err := cReviewContentService.UpdateCReviewContent(&ctx, in_cReviewContent)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cReviewContent.ReviewContentCd)
    assert.Equal(t, 1, out_cReviewContent.LanguageCd)
    assert.Equal(t, 2, out_cReviewContent.ReviewCategoryCd)
    assert.Equal(t, "dummy-ReviewContentName", out_cReviewContent.ReviewContentName)
    assert.NotNil(t, out_cReviewContent.CreateDatetime)
    assert.Equal(t, "", out_cReviewContent.CreateDatetime)
    assert.Equal(t, "", out_cReviewContent.CreateFunction)
    assert.NotNil(t, out_cReviewContent.UpdateDatetime)
    assert.NotEqual(t, "", out_cReviewContent.UpdateDatetime)
    assert.Equal(t, "UpdateCReviewContent", out_cReviewContent.UpdateFunction)
}
