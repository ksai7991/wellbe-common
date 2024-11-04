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

type cBillingStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError)
    FakeUpdate func(*model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CBillingStatus, *errordef.LogicError)
}

func (lr cBillingStatusMockRepository) CreateCBillingStatus(ctx *context.Context, cBillingStatus *model.CBillingStatus)  (*model.CBillingStatus, *errordef.LogicError) {
    return lr.FakeCreate(cBillingStatus)
}

func (lr cBillingStatusMockRepository) UpdateCBillingStatus(ctx *context.Context, cBillingStatus *model.CBillingStatus)  (*model.CBillingStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cBillingStatus)
}

func (lr cBillingStatusMockRepository) DeleteCBillingStatus(ctx *context.Context, billingStatusCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(billingStatusCd, languageCd)
}

func (lr cBillingStatusMockRepository)GetCBillingStatusWithKey(ctx *context.Context, billingStatusCd int, languageCd int)  ([]*model.CBillingStatus, *errordef.LogicError) {
    return lr.FakeGet(billingStatusCd, languageCd)
}


type cBillingStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cBillingStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCBillingStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cBillingStatusMockRepository{
        FakeCreate: func(cBillingStatus *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError) {
            return cBillingStatus, nil
        },
    }
    numberUtil := &cBillingStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBillingStatusService := NewCBillingStatusService(repository, numberUtil)
    in_cBillingStatus := new(model.CBillingStatus)
    in_cBillingStatus.BillingStatusCd = 0
    in_cBillingStatus.LanguageCd = 1
    in_cBillingStatus.BillingStatusName = "dummy-BillingStatusName"
    out_cBillingStatus, err := cBillingStatusService.CreateCBillingStatus(&ctx, in_cBillingStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBillingStatus.BillingStatusCd)
    assert.Equal(t, 1, out_cBillingStatus.LanguageCd)
    assert.Equal(t, "dummy-BillingStatusName", out_cBillingStatus.BillingStatusName)
    assert.NotNil(t, out_cBillingStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cBillingStatus.CreateDatetime)
    assert.Equal(t, "CreateCBillingStatus", out_cBillingStatus.CreateFunction)
    assert.NotNil(t, out_cBillingStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cBillingStatus.UpdateDatetime)
    assert.Equal(t, "CreateCBillingStatus", out_cBillingStatus.UpdateFunction)
}

func TestUpdateCBillingStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cBillingStatusMockRepository{
        FakeUpdate: func(cBillingStatus *model.CBillingStatus) (*model.CBillingStatus, *errordef.LogicError) {
            return cBillingStatus, nil
        },
        FakeGet: func(billingStatusCd int, languageCd int) ([]*model.CBillingStatus, *errordef.LogicError) {
            return []*model.CBillingStatus{&model.CBillingStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cBillingStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBillingStatusService := NewCBillingStatusService(repository, numberUtil)
    in_cBillingStatus := new(model.CBillingStatus)
    in_cBillingStatus.BillingStatusCd = 0
    in_cBillingStatus.LanguageCd = 1
    in_cBillingStatus.BillingStatusName = "dummy-BillingStatusName"
    out_cBillingStatus, err := cBillingStatusService.UpdateCBillingStatus(&ctx, in_cBillingStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBillingStatus.BillingStatusCd)
    assert.Equal(t, 1, out_cBillingStatus.LanguageCd)
    assert.Equal(t, "dummy-BillingStatusName", out_cBillingStatus.BillingStatusName)
    assert.NotNil(t, out_cBillingStatus.CreateDatetime)
    assert.Equal(t, "", out_cBillingStatus.CreateDatetime)
    assert.Equal(t, "", out_cBillingStatus.CreateFunction)
    assert.NotNil(t, out_cBillingStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cBillingStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCBillingStatus", out_cBillingStatus.UpdateFunction)
}
