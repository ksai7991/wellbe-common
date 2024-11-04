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

type cBookingChanelMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError)
    FakeUpdate func(*model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CBookingChanel, *errordef.LogicError)
}

func (lr cBookingChanelMockRepository) CreateCBookingChanel(ctx *context.Context, cBookingChanel *model.CBookingChanel)  (*model.CBookingChanel, *errordef.LogicError) {
    return lr.FakeCreate(cBookingChanel)
}

func (lr cBookingChanelMockRepository) UpdateCBookingChanel(ctx *context.Context, cBookingChanel *model.CBookingChanel)  (*model.CBookingChanel, *errordef.LogicError) {
    return lr.FakeUpdate(cBookingChanel)
}

func (lr cBookingChanelMockRepository) DeleteCBookingChanel(ctx *context.Context, bookingChanelCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(bookingChanelCd, languageCd)
}

func (lr cBookingChanelMockRepository)GetCBookingChanelWithKey(ctx *context.Context, bookingChanelCd int, languageCd int)  ([]*model.CBookingChanel, *errordef.LogicError) {
    return lr.FakeGet(bookingChanelCd, languageCd)
}


type cBookingChanelMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cBookingChanelMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCBookingChanel(t *testing.T) {
    ctx := context.Background()
    repository := &cBookingChanelMockRepository{
        FakeCreate: func(cBookingChanel *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError) {
            return cBookingChanel, nil
        },
    }
    numberUtil := &cBookingChanelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBookingChanelService := NewCBookingChanelService(repository, numberUtil)
    in_cBookingChanel := new(model.CBookingChanel)
    in_cBookingChanel.BookingChanelCd = 0
    in_cBookingChanel.LanguageCd = 1
    in_cBookingChanel.BookingChanelName = "dummy-BookingChanelName"
    out_cBookingChanel, err := cBookingChanelService.CreateCBookingChanel(&ctx, in_cBookingChanel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBookingChanel.BookingChanelCd)
    assert.Equal(t, 1, out_cBookingChanel.LanguageCd)
    assert.Equal(t, "dummy-BookingChanelName", out_cBookingChanel.BookingChanelName)
    assert.NotNil(t, out_cBookingChanel.CreateDatetime)
    assert.NotEqual(t, "", out_cBookingChanel.CreateDatetime)
    assert.Equal(t, "CreateCBookingChanel", out_cBookingChanel.CreateFunction)
    assert.NotNil(t, out_cBookingChanel.UpdateDatetime)
    assert.NotEqual(t, "", out_cBookingChanel.UpdateDatetime)
    assert.Equal(t, "CreateCBookingChanel", out_cBookingChanel.UpdateFunction)
}

func TestUpdateCBookingChanel(t *testing.T) {
    ctx := context.Background()
    repository := &cBookingChanelMockRepository{
        FakeUpdate: func(cBookingChanel *model.CBookingChanel) (*model.CBookingChanel, *errordef.LogicError) {
            return cBookingChanel, nil
        },
        FakeGet: func(bookingChanelCd int, languageCd int) ([]*model.CBookingChanel, *errordef.LogicError) {
            return []*model.CBookingChanel{&model.CBookingChanel{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cBookingChanelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBookingChanelService := NewCBookingChanelService(repository, numberUtil)
    in_cBookingChanel := new(model.CBookingChanel)
    in_cBookingChanel.BookingChanelCd = 0
    in_cBookingChanel.LanguageCd = 1
    in_cBookingChanel.BookingChanelName = "dummy-BookingChanelName"
    out_cBookingChanel, err := cBookingChanelService.UpdateCBookingChanel(&ctx, in_cBookingChanel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBookingChanel.BookingChanelCd)
    assert.Equal(t, 1, out_cBookingChanel.LanguageCd)
    assert.Equal(t, "dummy-BookingChanelName", out_cBookingChanel.BookingChanelName)
    assert.NotNil(t, out_cBookingChanel.CreateDatetime)
    assert.Equal(t, "", out_cBookingChanel.CreateDatetime)
    assert.Equal(t, "", out_cBookingChanel.CreateFunction)
    assert.NotNil(t, out_cBookingChanel.UpdateDatetime)
    assert.NotEqual(t, "", out_cBookingChanel.UpdateDatetime)
    assert.Equal(t, "UpdateCBookingChanel", out_cBookingChanel.UpdateFunction)
}
