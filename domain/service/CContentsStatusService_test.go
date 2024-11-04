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

type cContentsStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError)
    FakeUpdate func(*model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CContentsStatus, *errordef.LogicError)
}

func (lr cContentsStatusMockRepository) CreateCContentsStatus(ctx *context.Context, cContentsStatus *model.CContentsStatus)  (*model.CContentsStatus, *errordef.LogicError) {
    return lr.FakeCreate(cContentsStatus)
}

func (lr cContentsStatusMockRepository) UpdateCContentsStatus(ctx *context.Context, cContentsStatus *model.CContentsStatus)  (*model.CContentsStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cContentsStatus)
}

func (lr cContentsStatusMockRepository) DeleteCContentsStatus(ctx *context.Context, contentsStatusCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(contentsStatusCd, languageCd)
}

func (lr cContentsStatusMockRepository)GetCContentsStatusWithKey(ctx *context.Context, contentsStatusCd int, languageCd int)  ([]*model.CContentsStatus, *errordef.LogicError) {
    return lr.FakeGet(contentsStatusCd, languageCd)
}


type cContentsStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cContentsStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCContentsStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cContentsStatusMockRepository{
        FakeCreate: func(cContentsStatus *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError) {
            return cContentsStatus, nil
        },
    }
    numberUtil := &cContentsStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cContentsStatusService := NewCContentsStatusService(repository, numberUtil)
    in_cContentsStatus := new(model.CContentsStatus)
    in_cContentsStatus.ContentsStatusCd = 0
    in_cContentsStatus.LanguageCd = 1
    in_cContentsStatus.ContentsStatusName = "dummy-ContentsStatusName"
    out_cContentsStatus, err := cContentsStatusService.CreateCContentsStatus(&ctx, in_cContentsStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cContentsStatus.ContentsStatusCd)
    assert.Equal(t, 1, out_cContentsStatus.LanguageCd)
    assert.Equal(t, "dummy-ContentsStatusName", out_cContentsStatus.ContentsStatusName)
    assert.NotNil(t, out_cContentsStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cContentsStatus.CreateDatetime)
    assert.Equal(t, "CreateCContentsStatus", out_cContentsStatus.CreateFunction)
    assert.NotNil(t, out_cContentsStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cContentsStatus.UpdateDatetime)
    assert.Equal(t, "CreateCContentsStatus", out_cContentsStatus.UpdateFunction)
}

func TestUpdateCContentsStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cContentsStatusMockRepository{
        FakeUpdate: func(cContentsStatus *model.CContentsStatus) (*model.CContentsStatus, *errordef.LogicError) {
            return cContentsStatus, nil
        },
        FakeGet: func(contentsStatusCd int, languageCd int) ([]*model.CContentsStatus, *errordef.LogicError) {
            return []*model.CContentsStatus{&model.CContentsStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cContentsStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cContentsStatusService := NewCContentsStatusService(repository, numberUtil)
    in_cContentsStatus := new(model.CContentsStatus)
    in_cContentsStatus.ContentsStatusCd = 0
    in_cContentsStatus.LanguageCd = 1
    in_cContentsStatus.ContentsStatusName = "dummy-ContentsStatusName"
    out_cContentsStatus, err := cContentsStatusService.UpdateCContentsStatus(&ctx, in_cContentsStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cContentsStatus.ContentsStatusCd)
    assert.Equal(t, 1, out_cContentsStatus.LanguageCd)
    assert.Equal(t, "dummy-ContentsStatusName", out_cContentsStatus.ContentsStatusName)
    assert.NotNil(t, out_cContentsStatus.CreateDatetime)
    assert.Equal(t, "", out_cContentsStatus.CreateDatetime)
    assert.Equal(t, "", out_cContentsStatus.CreateFunction)
    assert.NotNil(t, out_cContentsStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cContentsStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCContentsStatus", out_cContentsStatus.UpdateFunction)
}
