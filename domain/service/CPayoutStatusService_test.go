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

type cPayoutStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError)
    FakeUpdate func(*model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.CPayoutStatus, *errordef.LogicError)
}

func (lr cPayoutStatusMockRepository) CreateCPayoutStatus(ctx *context.Context, cPayoutStatus *model.CPayoutStatus)  (*model.CPayoutStatus, *errordef.LogicError) {
    return lr.FakeCreate(cPayoutStatus)
}

func (lr cPayoutStatusMockRepository) UpdateCPayoutStatus(ctx *context.Context, cPayoutStatus *model.CPayoutStatus)  (*model.CPayoutStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cPayoutStatus)
}

func (lr cPayoutStatusMockRepository) DeleteCPayoutStatus(ctx *context.Context, payoutStatusCd int)  *errordef.LogicError {
    return lr.FakeDelete(payoutStatusCd)
}

func (lr cPayoutStatusMockRepository)GetCPayoutStatusWithKey(ctx *context.Context, payoutStatusCd int)  ([]*model.CPayoutStatus, *errordef.LogicError) {
    return lr.FakeGet(payoutStatusCd)
}


type cPayoutStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cPayoutStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCPayoutStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cPayoutStatusMockRepository{
        FakeCreate: func(cPayoutStatus *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError) {
            return cPayoutStatus, nil
        },
    }
    numberUtil := &cPayoutStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cPayoutStatusService := NewCPayoutStatusService(repository, numberUtil)
    in_cPayoutStatus := new(model.CPayoutStatus)
    in_cPayoutStatus.PayoutStatusCd = 0
    in_cPayoutStatus.LanguageCd = 1
    in_cPayoutStatus.PayoutStatusName = "dummy-PayoutStatusName"
    out_cPayoutStatus, err := cPayoutStatusService.CreateCPayoutStatus(&ctx, in_cPayoutStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cPayoutStatus.PayoutStatusCd)
    assert.Equal(t, 1, out_cPayoutStatus.LanguageCd)
    assert.Equal(t, "dummy-PayoutStatusName", out_cPayoutStatus.PayoutStatusName)
    assert.NotNil(t, out_cPayoutStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cPayoutStatus.CreateDatetime)
    assert.Equal(t, "CreateCPayoutStatus", out_cPayoutStatus.CreateFunction)
    assert.NotNil(t, out_cPayoutStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cPayoutStatus.UpdateDatetime)
    assert.Equal(t, "CreateCPayoutStatus", out_cPayoutStatus.UpdateFunction)
}

func TestUpdateCPayoutStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cPayoutStatusMockRepository{
        FakeUpdate: func(cPayoutStatus *model.CPayoutStatus) (*model.CPayoutStatus, *errordef.LogicError) {
            return cPayoutStatus, nil
        },
        FakeGet: func(payoutStatusCd int) ([]*model.CPayoutStatus, *errordef.LogicError) {
            return []*model.CPayoutStatus{&model.CPayoutStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cPayoutStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cPayoutStatusService := NewCPayoutStatusService(repository, numberUtil)
    in_cPayoutStatus := new(model.CPayoutStatus)
    in_cPayoutStatus.PayoutStatusCd = 0
    in_cPayoutStatus.LanguageCd = 1
    in_cPayoutStatus.PayoutStatusName = "dummy-PayoutStatusName"
    out_cPayoutStatus, err := cPayoutStatusService.UpdateCPayoutStatus(&ctx, in_cPayoutStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cPayoutStatus.PayoutStatusCd)
    assert.Equal(t, 1, out_cPayoutStatus.LanguageCd)
    assert.Equal(t, "dummy-PayoutStatusName", out_cPayoutStatus.PayoutStatusName)
    assert.NotNil(t, out_cPayoutStatus.CreateDatetime)
    assert.Equal(t, "", out_cPayoutStatus.CreateDatetime)
    assert.Equal(t, "", out_cPayoutStatus.CreateFunction)
    assert.NotNil(t, out_cPayoutStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cPayoutStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCPayoutStatus", out_cPayoutStatus.UpdateFunction)
}
