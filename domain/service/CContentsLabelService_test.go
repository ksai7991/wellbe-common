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

type cContentsLabelMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError)
    FakeUpdate func(*model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CContentsLabel, *errordef.LogicError)
}

func (lr cContentsLabelMockRepository) CreateCContentsLabel(ctx *context.Context, cContentsLabel *model.CContentsLabel)  (*model.CContentsLabel, *errordef.LogicError) {
    return lr.FakeCreate(cContentsLabel)
}

func (lr cContentsLabelMockRepository) UpdateCContentsLabel(ctx *context.Context, cContentsLabel *model.CContentsLabel)  (*model.CContentsLabel, *errordef.LogicError) {
    return lr.FakeUpdate(cContentsLabel)
}

func (lr cContentsLabelMockRepository) DeleteCContentsLabel(ctx *context.Context, contentsLabelCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(contentsLabelCd, languageCd)
}

func (lr cContentsLabelMockRepository)GetCContentsLabelWithKey(ctx *context.Context, contentsLabelCd int, languageCd int)  ([]*model.CContentsLabel, *errordef.LogicError) {
    return lr.FakeGet(contentsLabelCd, languageCd)
}


type cContentsLabelMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cContentsLabelMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCContentsLabel(t *testing.T) {
    ctx := context.Background()
    repository := &cContentsLabelMockRepository{
        FakeCreate: func(cContentsLabel *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError) {
            return cContentsLabel, nil
        },
    }
    numberUtil := &cContentsLabelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cContentsLabelService := NewCContentsLabelService(repository, numberUtil)
    in_cContentsLabel := new(model.CContentsLabel)
    in_cContentsLabel.ContentsLabelCd = 0
    in_cContentsLabel.LanguageCd = 1
    in_cContentsLabel.ContentsCategoryCd = 2
    in_cContentsLabel.ContentsLabelName = "dummy-ContentsLabelName"
    in_cContentsLabel.ContentsLabelUrl = "dummy-ContentsLabelUrl"
    out_cContentsLabel, err := cContentsLabelService.CreateCContentsLabel(&ctx, in_cContentsLabel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cContentsLabel.ContentsLabelCd)
    assert.Equal(t, 1, out_cContentsLabel.LanguageCd)
    assert.Equal(t, 2, out_cContentsLabel.ContentsCategoryCd)
    assert.Equal(t, "dummy-ContentsLabelName", out_cContentsLabel.ContentsLabelName)
    assert.Equal(t, "dummy-ContentsLabelUrl", out_cContentsLabel.ContentsLabelUrl)
    assert.NotNil(t, out_cContentsLabel.CreateDatetime)
    assert.NotEqual(t, "", out_cContentsLabel.CreateDatetime)
    assert.Equal(t, "CreateCContentsLabel", out_cContentsLabel.CreateFunction)
    assert.NotNil(t, out_cContentsLabel.UpdateDatetime)
    assert.NotEqual(t, "", out_cContentsLabel.UpdateDatetime)
    assert.Equal(t, "CreateCContentsLabel", out_cContentsLabel.UpdateFunction)
}

func TestUpdateCContentsLabel(t *testing.T) {
    ctx := context.Background()
    repository := &cContentsLabelMockRepository{
        FakeUpdate: func(cContentsLabel *model.CContentsLabel) (*model.CContentsLabel, *errordef.LogicError) {
            return cContentsLabel, nil
        },
        FakeGet: func(contentsLabelCd int, languageCd int) ([]*model.CContentsLabel, *errordef.LogicError) {
            return []*model.CContentsLabel{&model.CContentsLabel{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cContentsLabelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cContentsLabelService := NewCContentsLabelService(repository, numberUtil)
    in_cContentsLabel := new(model.CContentsLabel)
    in_cContentsLabel.ContentsLabelCd = 0
    in_cContentsLabel.LanguageCd = 1
    in_cContentsLabel.ContentsCategoryCd = 2
    in_cContentsLabel.ContentsLabelName = "dummy-ContentsLabelName"
    in_cContentsLabel.ContentsLabelUrl = "dummy-ContentsLabelUrl"
    out_cContentsLabel, err := cContentsLabelService.UpdateCContentsLabel(&ctx, in_cContentsLabel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cContentsLabel.ContentsLabelCd)
    assert.Equal(t, 1, out_cContentsLabel.LanguageCd)
    assert.Equal(t, 2, out_cContentsLabel.ContentsCategoryCd)
    assert.Equal(t, "dummy-ContentsLabelName", out_cContentsLabel.ContentsLabelName)
    assert.Equal(t, "dummy-ContentsLabelUrl", out_cContentsLabel.ContentsLabelUrl)
    assert.NotNil(t, out_cContentsLabel.CreateDatetime)
    assert.Equal(t, "", out_cContentsLabel.CreateDatetime)
    assert.Equal(t, "", out_cContentsLabel.CreateFunction)
    assert.NotNil(t, out_cContentsLabel.UpdateDatetime)
    assert.NotEqual(t, "", out_cContentsLabel.UpdateDatetime)
    assert.Equal(t, "UpdateCContentsLabel", out_cContentsLabel.UpdateFunction)
}
