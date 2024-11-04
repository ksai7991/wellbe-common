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

type cBillingContentMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CBillingContent) (*model.CBillingContent, *errordef.LogicError)
    FakeUpdate func(*model.CBillingContent) (*model.CBillingContent, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CBillingContent, *errordef.LogicError)
}

func (lr cBillingContentMockRepository) CreateCBillingContent(ctx *context.Context, cBillingContent *model.CBillingContent)  (*model.CBillingContent, *errordef.LogicError) {
    return lr.FakeCreate(cBillingContent)
}

func (lr cBillingContentMockRepository) UpdateCBillingContent(ctx *context.Context, cBillingContent *model.CBillingContent)  (*model.CBillingContent, *errordef.LogicError) {
    return lr.FakeUpdate(cBillingContent)
}

func (lr cBillingContentMockRepository) DeleteCBillingContent(ctx *context.Context, billingContentCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(billingContentCd, languageCd)
}

func (lr cBillingContentMockRepository)GetCBillingContentWithKey(ctx *context.Context, billingContentCd int, languageCd int)  ([]*model.CBillingContent, *errordef.LogicError) {
    return lr.FakeGet(billingContentCd, languageCd)
}


type cBillingContentMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cBillingContentMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCBillingContent(t *testing.T) {
    ctx := context.Background()
    repository := &cBillingContentMockRepository{
        FakeCreate: func(cBillingContent *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError) {
            return cBillingContent, nil
        },
    }
    numberUtil := &cBillingContentMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBillingContentService := NewCBillingContentService(repository, numberUtil)
    in_cBillingContent := new(model.CBillingContent)
    in_cBillingContent.BillingContentCd = 0
    in_cBillingContent.LanguageCd = 1
    in_cBillingContent.BillingContentName = "dummy-BillingContentName"
    out_cBillingContent, err := cBillingContentService.CreateCBillingContent(&ctx, in_cBillingContent)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBillingContent.BillingContentCd)
    assert.Equal(t, 1, out_cBillingContent.LanguageCd)
    assert.Equal(t, "dummy-BillingContentName", out_cBillingContent.BillingContentName)
    assert.NotNil(t, out_cBillingContent.CreateDatetime)
    assert.NotEqual(t, "", out_cBillingContent.CreateDatetime)
    assert.Equal(t, "CreateCBillingContent", out_cBillingContent.CreateFunction)
    assert.NotNil(t, out_cBillingContent.UpdateDatetime)
    assert.NotEqual(t, "", out_cBillingContent.UpdateDatetime)
    assert.Equal(t, "CreateCBillingContent", out_cBillingContent.UpdateFunction)
}

func TestUpdateCBillingContent(t *testing.T) {
    ctx := context.Background()
    repository := &cBillingContentMockRepository{
        FakeUpdate: func(cBillingContent *model.CBillingContent) (*model.CBillingContent, *errordef.LogicError) {
            return cBillingContent, nil
        },
        FakeGet: func(billingContentCd int, languageCd int) ([]*model.CBillingContent, *errordef.LogicError) {
            return []*model.CBillingContent{&model.CBillingContent{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cBillingContentMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBillingContentService := NewCBillingContentService(repository, numberUtil)
    in_cBillingContent := new(model.CBillingContent)
    in_cBillingContent.BillingContentCd = 0
    in_cBillingContent.LanguageCd = 1
    in_cBillingContent.BillingContentName = "dummy-BillingContentName"
    out_cBillingContent, err := cBillingContentService.UpdateCBillingContent(&ctx, in_cBillingContent)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBillingContent.BillingContentCd)
    assert.Equal(t, 1, out_cBillingContent.LanguageCd)
    assert.Equal(t, "dummy-BillingContentName", out_cBillingContent.BillingContentName)
    assert.NotNil(t, out_cBillingContent.CreateDatetime)
    assert.Equal(t, "", out_cBillingContent.CreateDatetime)
    assert.Equal(t, "", out_cBillingContent.CreateFunction)
    assert.NotNil(t, out_cBillingContent.UpdateDatetime)
    assert.NotEqual(t, "", out_cBillingContent.UpdateDatetime)
    assert.Equal(t, "UpdateCBillingContent", out_cBillingContent.UpdateFunction)
}
