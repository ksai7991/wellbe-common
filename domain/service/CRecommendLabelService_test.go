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

type cRecommendLabelMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError)
    FakeUpdate func(*model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CRecommendLabel, *errordef.LogicError)
}

func (lr cRecommendLabelMockRepository) CreateCRecommendLabel(ctx *context.Context, cRecommendLabel *model.CRecommendLabel)  (*model.CRecommendLabel, *errordef.LogicError) {
    return lr.FakeCreate(cRecommendLabel)
}

func (lr cRecommendLabelMockRepository) UpdateCRecommendLabel(ctx *context.Context, cRecommendLabel *model.CRecommendLabel)  (*model.CRecommendLabel, *errordef.LogicError) {
    return lr.FakeUpdate(cRecommendLabel)
}

func (lr cRecommendLabelMockRepository) DeleteCRecommendLabel(ctx *context.Context, recommendLabelCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(recommendLabelCd, languageCd)
}

func (lr cRecommendLabelMockRepository)GetCRecommendLabelWithKey(ctx *context.Context, recommendLabelCd int, languageCd int)  ([]*model.CRecommendLabel, *errordef.LogicError) {
    return lr.FakeGet(recommendLabelCd, languageCd)
}


type cRecommendLabelMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cRecommendLabelMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCRecommendLabel(t *testing.T) {
    ctx := context.Background()
    repository := &cRecommendLabelMockRepository{
        FakeCreate: func(cRecommendLabel *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError) {
            return cRecommendLabel, nil
        },
    }
    numberUtil := &cRecommendLabelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cRecommendLabelService := NewCRecommendLabelService(repository, numberUtil)
    in_cRecommendLabel := new(model.CRecommendLabel)
    in_cRecommendLabel.RecommendLabelCd = 0
    in_cRecommendLabel.LanguageCd = 1
    in_cRecommendLabel.RecommendLabelName = "dummy-RecommendLabelName"
    out_cRecommendLabel, err := cRecommendLabelService.CreateCRecommendLabel(&ctx, in_cRecommendLabel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cRecommendLabel.RecommendLabelCd)
    assert.Equal(t, 1, out_cRecommendLabel.LanguageCd)
    assert.Equal(t, "dummy-RecommendLabelName", out_cRecommendLabel.RecommendLabelName)
    assert.NotNil(t, out_cRecommendLabel.CreateDatetime)
    assert.NotEqual(t, "", out_cRecommendLabel.CreateDatetime)
    assert.Equal(t, "CreateCRecommendLabel", out_cRecommendLabel.CreateFunction)
    assert.NotNil(t, out_cRecommendLabel.UpdateDatetime)
    assert.NotEqual(t, "", out_cRecommendLabel.UpdateDatetime)
    assert.Equal(t, "CreateCRecommendLabel", out_cRecommendLabel.UpdateFunction)
}

func TestUpdateCRecommendLabel(t *testing.T) {
    ctx := context.Background()
    repository := &cRecommendLabelMockRepository{
        FakeUpdate: func(cRecommendLabel *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError) {
            return cRecommendLabel, nil
        },
        FakeGet: func(recommendLabelCd int, languageCd int) ([]*model.CRecommendLabel, *errordef.LogicError) {
            return []*model.CRecommendLabel{&model.CRecommendLabel{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cRecommendLabelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cRecommendLabelService := NewCRecommendLabelService(repository, numberUtil)
    in_cRecommendLabel := new(model.CRecommendLabel)
    in_cRecommendLabel.RecommendLabelCd = 0
    in_cRecommendLabel.LanguageCd = 1
    in_cRecommendLabel.RecommendLabelName = "dummy-RecommendLabelName"
    out_cRecommendLabel, err := cRecommendLabelService.UpdateCRecommendLabel(&ctx, in_cRecommendLabel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cRecommendLabel.RecommendLabelCd)
    assert.Equal(t, 1, out_cRecommendLabel.LanguageCd)
    assert.Equal(t, "dummy-RecommendLabelName", out_cRecommendLabel.RecommendLabelName)
    assert.NotNil(t, out_cRecommendLabel.CreateDatetime)
    assert.Equal(t, "", out_cRecommendLabel.CreateDatetime)
    assert.Equal(t, "", out_cRecommendLabel.CreateFunction)
    assert.NotNil(t, out_cRecommendLabel.UpdateDatetime)
    assert.NotEqual(t, "", out_cRecommendLabel.UpdateDatetime)
    assert.Equal(t, "UpdateCRecommendLabel", out_cRecommendLabel.UpdateFunction)
}
