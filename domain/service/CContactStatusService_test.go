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

type cContactStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CContactStatus) (*model.CContactStatus, *errordef.LogicError)
    FakeUpdate func(*model.CContactStatus) (*model.CContactStatus, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CContactStatus, *errordef.LogicError)
}

func (lr cContactStatusMockRepository) CreateCContactStatus(ctx *context.Context, cContactStatus *model.CContactStatus)  (*model.CContactStatus, *errordef.LogicError) {
    return lr.FakeCreate(cContactStatus)
}

func (lr cContactStatusMockRepository) UpdateCContactStatus(ctx *context.Context, cContactStatus *model.CContactStatus)  (*model.CContactStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cContactStatus)
}

func (lr cContactStatusMockRepository) DeleteCContactStatus(ctx *context.Context, contactStatusCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(contactStatusCd, languageCd)
}

func (lr cContactStatusMockRepository)GetCContactStatusWithKey(ctx *context.Context, contactStatusCd int, languageCd int)  ([]*model.CContactStatus, *errordef.LogicError) {
    return lr.FakeGet(contactStatusCd, languageCd)
}


type cContactStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cContactStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCContactStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cContactStatusMockRepository{
        FakeCreate: func(cContactStatus *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError) {
            return cContactStatus, nil
        },
    }
    numberUtil := &cContactStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cContactStatusService := NewCContactStatusService(repository, numberUtil)
    in_cContactStatus := new(model.CContactStatus)
    in_cContactStatus.ContactStatusCd = 0
    in_cContactStatus.LanguageCd = 1
    in_cContactStatus.ContactStatusName = "dummy-ContactStatusName"
    out_cContactStatus, err := cContactStatusService.CreateCContactStatus(&ctx, in_cContactStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cContactStatus.ContactStatusCd)
    assert.Equal(t, 1, out_cContactStatus.LanguageCd)
    assert.Equal(t, "dummy-ContactStatusName", out_cContactStatus.ContactStatusName)
    assert.NotNil(t, out_cContactStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cContactStatus.CreateDatetime)
    assert.Equal(t, "CreateCContactStatus", out_cContactStatus.CreateFunction)
    assert.NotNil(t, out_cContactStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cContactStatus.UpdateDatetime)
    assert.Equal(t, "CreateCContactStatus", out_cContactStatus.UpdateFunction)
}

func TestUpdateCContactStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cContactStatusMockRepository{
        FakeUpdate: func(cContactStatus *model.CContactStatus) (*model.CContactStatus, *errordef.LogicError) {
            return cContactStatus, nil
        },
        FakeGet: func(contactStatusCd int, languageCd int) ([]*model.CContactStatus, *errordef.LogicError) {
            return []*model.CContactStatus{&model.CContactStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cContactStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cContactStatusService := NewCContactStatusService(repository, numberUtil)
    in_cContactStatus := new(model.CContactStatus)
    in_cContactStatus.ContactStatusCd = 0
    in_cContactStatus.LanguageCd = 1
    in_cContactStatus.ContactStatusName = "dummy-ContactStatusName"
    out_cContactStatus, err := cContactStatusService.UpdateCContactStatus(&ctx, in_cContactStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cContactStatus.ContactStatusCd)
    assert.Equal(t, 1, out_cContactStatus.LanguageCd)
    assert.Equal(t, "dummy-ContactStatusName", out_cContactStatus.ContactStatusName)
    assert.NotNil(t, out_cContactStatus.CreateDatetime)
    assert.Equal(t, "", out_cContactStatus.CreateDatetime)
    assert.Equal(t, "", out_cContactStatus.CreateFunction)
    assert.NotNil(t, out_cContactStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cContactStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCContactStatus", out_cContactStatus.UpdateFunction)
}
