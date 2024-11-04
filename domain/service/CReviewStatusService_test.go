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

type cReviewStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError)
    FakeUpdate func(*model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CReviewStatus, *errordef.LogicError)
}

func (lr cReviewStatusMockRepository) CreateCReviewStatus(ctx *context.Context, cReviewStatus *model.CReviewStatus)  (*model.CReviewStatus, *errordef.LogicError) {
    return lr.FakeCreate(cReviewStatus)
}

func (lr cReviewStatusMockRepository) UpdateCReviewStatus(ctx *context.Context, cReviewStatus *model.CReviewStatus)  (*model.CReviewStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cReviewStatus)
}

func (lr cReviewStatusMockRepository) DeleteCReviewStatus(ctx *context.Context, reviewStatusCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(reviewStatusCd, languageCd)
}

func (lr cReviewStatusMockRepository)GetCReviewStatusWithKey(ctx *context.Context, reviewStatusCd int, languageCd int)  ([]*model.CReviewStatus, *errordef.LogicError) {
    return lr.FakeGet(reviewStatusCd, languageCd)
}


type cReviewStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cReviewStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCReviewStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cReviewStatusMockRepository{
        FakeCreate: func(cReviewStatus *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError) {
            return cReviewStatus, nil
        },
    }
    numberUtil := &cReviewStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cReviewStatusService := NewCReviewStatusService(repository, numberUtil)
    in_cReviewStatus := new(model.CReviewStatus)
    in_cReviewStatus.ReviewStatusCd = 0
    in_cReviewStatus.LanguageCd = 1
    in_cReviewStatus.ReviewStatusName = "dummy-ReviewStatusName"
    out_cReviewStatus, err := cReviewStatusService.CreateCReviewStatus(&ctx, in_cReviewStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cReviewStatus.ReviewStatusCd)
    assert.Equal(t, 1, out_cReviewStatus.LanguageCd)
    assert.Equal(t, "dummy-ReviewStatusName", out_cReviewStatus.ReviewStatusName)
    assert.NotNil(t, out_cReviewStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cReviewStatus.CreateDatetime)
    assert.Equal(t, "CreateCReviewStatus", out_cReviewStatus.CreateFunction)
    assert.NotNil(t, out_cReviewStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cReviewStatus.UpdateDatetime)
    assert.Equal(t, "CreateCReviewStatus", out_cReviewStatus.UpdateFunction)
}

func TestUpdateCReviewStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cReviewStatusMockRepository{
        FakeUpdate: func(cReviewStatus *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError) {
            return cReviewStatus, nil
        },
        FakeGet: func(reviewStatusCd int, languageCd int) ([]*model.CReviewStatus, *errordef.LogicError) {
            return []*model.CReviewStatus{&model.CReviewStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cReviewStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cReviewStatusService := NewCReviewStatusService(repository, numberUtil)
    in_cReviewStatus := new(model.CReviewStatus)
    in_cReviewStatus.ReviewStatusCd = 0
    in_cReviewStatus.LanguageCd = 1
    in_cReviewStatus.ReviewStatusName = "dummy-ReviewStatusName"
    out_cReviewStatus, err := cReviewStatusService.UpdateCReviewStatus(&ctx, in_cReviewStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cReviewStatus.ReviewStatusCd)
    assert.Equal(t, 1, out_cReviewStatus.LanguageCd)
    assert.Equal(t, "dummy-ReviewStatusName", out_cReviewStatus.ReviewStatusName)
    assert.NotNil(t, out_cReviewStatus.CreateDatetime)
    assert.Equal(t, "", out_cReviewStatus.CreateDatetime)
    assert.Equal(t, "", out_cReviewStatus.CreateFunction)
    assert.NotNil(t, out_cReviewStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cReviewStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCReviewStatus", out_cReviewStatus.UpdateFunction)
}
