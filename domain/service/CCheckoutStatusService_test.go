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

type cCheckoutStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError)
    FakeUpdate func(*model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CCheckoutStatus, *errordef.LogicError)
}

func (lr cCheckoutStatusMockRepository) CreateCCheckoutStatus(ctx *context.Context, cCheckoutStatus *model.CCheckoutStatus)  (*model.CCheckoutStatus, *errordef.LogicError) {
    return lr.FakeCreate(cCheckoutStatus)
}

func (lr cCheckoutStatusMockRepository) UpdateCCheckoutStatus(ctx *context.Context, cCheckoutStatus *model.CCheckoutStatus)  (*model.CCheckoutStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cCheckoutStatus)
}

func (lr cCheckoutStatusMockRepository) DeleteCCheckoutStatus(ctx *context.Context, checkoutStatusCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(checkoutStatusCd, languageCd)
}

func (lr cCheckoutStatusMockRepository)GetCCheckoutStatusWithKey(ctx *context.Context, checkoutStatusCd int, languageCd int)  ([]*model.CCheckoutStatus, *errordef.LogicError) {
    return lr.FakeGet(checkoutStatusCd, languageCd)
}


type cCheckoutStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cCheckoutStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCCheckoutStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cCheckoutStatusMockRepository{
        FakeCreate: func(cCheckoutStatus *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError) {
            return cCheckoutStatus, nil
        },
    }
    numberUtil := &cCheckoutStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCheckoutStatusService := NewCCheckoutStatusService(repository, numberUtil)
    in_cCheckoutStatus := new(model.CCheckoutStatus)
    in_cCheckoutStatus.CheckoutStatusCd = 0
    in_cCheckoutStatus.LanguageCd = 1
    in_cCheckoutStatus.CheckoutStatusName = "dummy-CheckoutStatusName"
    out_cCheckoutStatus, err := cCheckoutStatusService.CreateCCheckoutStatus(&ctx, in_cCheckoutStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCheckoutStatus.CheckoutStatusCd)
    assert.Equal(t, 1, out_cCheckoutStatus.LanguageCd)
    assert.Equal(t, "dummy-CheckoutStatusName", out_cCheckoutStatus.CheckoutStatusName)
    assert.NotNil(t, out_cCheckoutStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cCheckoutStatus.CreateDatetime)
    assert.Equal(t, "CreateCCheckoutStatus", out_cCheckoutStatus.CreateFunction)
    assert.NotNil(t, out_cCheckoutStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cCheckoutStatus.UpdateDatetime)
    assert.Equal(t, "CreateCCheckoutStatus", out_cCheckoutStatus.UpdateFunction)
}

func TestUpdateCCheckoutStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cCheckoutStatusMockRepository{
        FakeUpdate: func(cCheckoutStatus *model.CCheckoutStatus) (*model.CCheckoutStatus, *errordef.LogicError) {
            return cCheckoutStatus, nil
        },
        FakeGet: func(checkoutStatusCd int, languageCd int) ([]*model.CCheckoutStatus, *errordef.LogicError) {
            return []*model.CCheckoutStatus{&model.CCheckoutStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cCheckoutStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCheckoutStatusService := NewCCheckoutStatusService(repository, numberUtil)
    in_cCheckoutStatus := new(model.CCheckoutStatus)
    in_cCheckoutStatus.CheckoutStatusCd = 0
    in_cCheckoutStatus.LanguageCd = 1
    in_cCheckoutStatus.CheckoutStatusName = "dummy-CheckoutStatusName"
    out_cCheckoutStatus, err := cCheckoutStatusService.UpdateCCheckoutStatus(&ctx, in_cCheckoutStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCheckoutStatus.CheckoutStatusCd)
    assert.Equal(t, 1, out_cCheckoutStatus.LanguageCd)
    assert.Equal(t, "dummy-CheckoutStatusName", out_cCheckoutStatus.CheckoutStatusName)
    assert.NotNil(t, out_cCheckoutStatus.CreateDatetime)
    assert.Equal(t, "", out_cCheckoutStatus.CreateDatetime)
    assert.Equal(t, "", out_cCheckoutStatus.CreateFunction)
    assert.NotNil(t, out_cCheckoutStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cCheckoutStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCCheckoutStatus", out_cCheckoutStatus.UpdateFunction)
}
