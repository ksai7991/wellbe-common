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

type cBookingStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError)
    FakeUpdate func(*model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CBookingStatus, *errordef.LogicError)
}

func (lr cBookingStatusMockRepository) CreateCBookingStatus(ctx *context.Context, cBookingStatus *model.CBookingStatus)  (*model.CBookingStatus, *errordef.LogicError) {
    return lr.FakeCreate(cBookingStatus)
}

func (lr cBookingStatusMockRepository) UpdateCBookingStatus(ctx *context.Context, cBookingStatus *model.CBookingStatus)  (*model.CBookingStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cBookingStatus)
}

func (lr cBookingStatusMockRepository) DeleteCBookingStatus(ctx *context.Context, bookingStatusCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(bookingStatusCd, languageCd)
}

func (lr cBookingStatusMockRepository)GetCBookingStatusWithKey(ctx *context.Context, bookingStatusCd int, languageCd int)  ([]*model.CBookingStatus, *errordef.LogicError) {
    return lr.FakeGet(bookingStatusCd, languageCd)
}


type cBookingStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cBookingStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCBookingStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cBookingStatusMockRepository{
        FakeCreate: func(cBookingStatus *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError) {
            return cBookingStatus, nil
        },
    }
    numberUtil := &cBookingStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBookingStatusService := NewCBookingStatusService(repository, numberUtil)
    in_cBookingStatus := new(model.CBookingStatus)
    in_cBookingStatus.BookingStatusCd = 0
    in_cBookingStatus.LanguageCd = 1
    in_cBookingStatus.BookingStatusName = "dummy-BookingStatusName"
    out_cBookingStatus, err := cBookingStatusService.CreateCBookingStatus(&ctx, in_cBookingStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBookingStatus.BookingStatusCd)
    assert.Equal(t, 1, out_cBookingStatus.LanguageCd)
    assert.Equal(t, "dummy-BookingStatusName", out_cBookingStatus.BookingStatusName)
    assert.NotNil(t, out_cBookingStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cBookingStatus.CreateDatetime)
    assert.Equal(t, "CreateCBookingStatus", out_cBookingStatus.CreateFunction)
    assert.NotNil(t, out_cBookingStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cBookingStatus.UpdateDatetime)
    assert.Equal(t, "CreateCBookingStatus", out_cBookingStatus.UpdateFunction)
}

func TestUpdateCBookingStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cBookingStatusMockRepository{
        FakeUpdate: func(cBookingStatus *model.CBookingStatus) (*model.CBookingStatus, *errordef.LogicError) {
            return cBookingStatus, nil
        },
        FakeGet: func(bookingStatusCd int, languageCd int) ([]*model.CBookingStatus, *errordef.LogicError) {
            return []*model.CBookingStatus{&model.CBookingStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cBookingStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBookingStatusService := NewCBookingStatusService(repository, numberUtil)
    in_cBookingStatus := new(model.CBookingStatus)
    in_cBookingStatus.BookingStatusCd = 0
    in_cBookingStatus.LanguageCd = 1
    in_cBookingStatus.BookingStatusName = "dummy-BookingStatusName"
    out_cBookingStatus, err := cBookingStatusService.UpdateCBookingStatus(&ctx, in_cBookingStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBookingStatus.BookingStatusCd)
    assert.Equal(t, 1, out_cBookingStatus.LanguageCd)
    assert.Equal(t, "dummy-BookingStatusName", out_cBookingStatus.BookingStatusName)
    assert.NotNil(t, out_cBookingStatus.CreateDatetime)
    assert.Equal(t, "", out_cBookingStatus.CreateDatetime)
    assert.Equal(t, "", out_cBookingStatus.CreateFunction)
    assert.NotNil(t, out_cBookingStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cBookingStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCBookingStatus", out_cBookingStatus.UpdateFunction)
}
