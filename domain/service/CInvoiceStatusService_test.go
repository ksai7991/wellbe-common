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

type cInvoiceStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError)
    FakeUpdate func(*model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CInvoiceStatus, *errordef.LogicError)
}

func (lr cInvoiceStatusMockRepository) CreateCInvoiceStatus(ctx *context.Context, cInvoiceStatus *model.CInvoiceStatus)  (*model.CInvoiceStatus, *errordef.LogicError) {
    return lr.FakeCreate(cInvoiceStatus)
}

func (lr cInvoiceStatusMockRepository) UpdateCInvoiceStatus(ctx *context.Context, cInvoiceStatus *model.CInvoiceStatus)  (*model.CInvoiceStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cInvoiceStatus)
}

func (lr cInvoiceStatusMockRepository) DeleteCInvoiceStatus(ctx *context.Context, invoiceStatusCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(invoiceStatusCd, languageCd)
}

func (lr cInvoiceStatusMockRepository)GetCInvoiceStatusWithKey(ctx *context.Context, invoiceStatusCd int, languageCd int)  ([]*model.CInvoiceStatus, *errordef.LogicError) {
    return lr.FakeGet(invoiceStatusCd, languageCd)
}


type cInvoiceStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cInvoiceStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCInvoiceStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cInvoiceStatusMockRepository{
        FakeCreate: func(cInvoiceStatus *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError) {
            return cInvoiceStatus, nil
        },
    }
    numberUtil := &cInvoiceStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cInvoiceStatusService := NewCInvoiceStatusService(repository, numberUtil)
    in_cInvoiceStatus := new(model.CInvoiceStatus)
    in_cInvoiceStatus.InvoiceStatusCd = 0
    in_cInvoiceStatus.LanguageCd = 1
    in_cInvoiceStatus.InvoiceStatusName = "dummy-InvoiceStatusName"
    out_cInvoiceStatus, err := cInvoiceStatusService.CreateCInvoiceStatus(&ctx, in_cInvoiceStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cInvoiceStatus.InvoiceStatusCd)
    assert.Equal(t, 1, out_cInvoiceStatus.LanguageCd)
    assert.Equal(t, "dummy-InvoiceStatusName", out_cInvoiceStatus.InvoiceStatusName)
    assert.NotNil(t, out_cInvoiceStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cInvoiceStatus.CreateDatetime)
    assert.Equal(t, "CreateCInvoiceStatus", out_cInvoiceStatus.CreateFunction)
    assert.NotNil(t, out_cInvoiceStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cInvoiceStatus.UpdateDatetime)
    assert.Equal(t, "CreateCInvoiceStatus", out_cInvoiceStatus.UpdateFunction)
}

func TestUpdateCInvoiceStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cInvoiceStatusMockRepository{
        FakeUpdate: func(cInvoiceStatus *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError) {
            return cInvoiceStatus, nil
        },
        FakeGet: func(invoiceStatusCd int, languageCd int) ([]*model.CInvoiceStatus, *errordef.LogicError) {
            return []*model.CInvoiceStatus{&model.CInvoiceStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cInvoiceStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cInvoiceStatusService := NewCInvoiceStatusService(repository, numberUtil)
    in_cInvoiceStatus := new(model.CInvoiceStatus)
    in_cInvoiceStatus.InvoiceStatusCd = 0
    in_cInvoiceStatus.LanguageCd = 1
    in_cInvoiceStatus.InvoiceStatusName = "dummy-InvoiceStatusName"
    out_cInvoiceStatus, err := cInvoiceStatusService.UpdateCInvoiceStatus(&ctx, in_cInvoiceStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cInvoiceStatus.InvoiceStatusCd)
    assert.Equal(t, 1, out_cInvoiceStatus.LanguageCd)
    assert.Equal(t, "dummy-InvoiceStatusName", out_cInvoiceStatus.InvoiceStatusName)
    assert.NotNil(t, out_cInvoiceStatus.CreateDatetime)
    assert.Equal(t, "", out_cInvoiceStatus.CreateDatetime)
    assert.Equal(t, "", out_cInvoiceStatus.CreateFunction)
    assert.NotNil(t, out_cInvoiceStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cInvoiceStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCInvoiceStatus", out_cInvoiceStatus.UpdateFunction)
}
