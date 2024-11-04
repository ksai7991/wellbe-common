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

type cBookingMethodMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError)
    FakeUpdate func(*model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CBookingMethod, *errordef.LogicError)
}

func (lr cBookingMethodMockRepository) CreateCBookingMethod(ctx *context.Context, cBookingMethod *model.CBookingMethod)  (*model.CBookingMethod, *errordef.LogicError) {
    return lr.FakeCreate(cBookingMethod)
}

func (lr cBookingMethodMockRepository) UpdateCBookingMethod(ctx *context.Context, cBookingMethod *model.CBookingMethod)  (*model.CBookingMethod, *errordef.LogicError) {
    return lr.FakeUpdate(cBookingMethod)
}

func (lr cBookingMethodMockRepository) DeleteCBookingMethod(ctx *context.Context, bookingMethodCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(bookingMethodCd, languageCd)
}

func (lr cBookingMethodMockRepository)GetCBookingMethodWithKey(ctx *context.Context, bookingMethodCd int, languageCd int)  ([]*model.CBookingMethod, *errordef.LogicError) {
    return lr.FakeGet(bookingMethodCd, languageCd)
}


type cBookingMethodMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cBookingMethodMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCBookingMethod(t *testing.T) {
    ctx := context.Background()
    repository := &cBookingMethodMockRepository{
        FakeCreate: func(cBookingMethod *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError) {
            return cBookingMethod, nil
        },
    }
    numberUtil := &cBookingMethodMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBookingMethodService := NewCBookingMethodService(repository, numberUtil)
    in_cBookingMethod := new(model.CBookingMethod)
    in_cBookingMethod.BookingMethodCd = 0
    in_cBookingMethod.LanguageCd = 1
    in_cBookingMethod.BookingMethodName = "dummy-BookingMethodName"
    out_cBookingMethod, err := cBookingMethodService.CreateCBookingMethod(&ctx, in_cBookingMethod)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBookingMethod.BookingMethodCd)
    assert.Equal(t, 1, out_cBookingMethod.LanguageCd)
    assert.Equal(t, "dummy-BookingMethodName", out_cBookingMethod.BookingMethodName)
    assert.NotNil(t, out_cBookingMethod.CreateDatetime)
    assert.NotEqual(t, "", out_cBookingMethod.CreateDatetime)
    assert.Equal(t, "CreateCBookingMethod", out_cBookingMethod.CreateFunction)
    assert.NotNil(t, out_cBookingMethod.UpdateDatetime)
    assert.NotEqual(t, "", out_cBookingMethod.UpdateDatetime)
    assert.Equal(t, "CreateCBookingMethod", out_cBookingMethod.UpdateFunction)
}

func TestUpdateCBookingMethod(t *testing.T) {
    ctx := context.Background()
    repository := &cBookingMethodMockRepository{
        FakeUpdate: func(cBookingMethod *model.CBookingMethod) (*model.CBookingMethod, *errordef.LogicError) {
            return cBookingMethod, nil
        },
        FakeGet: func(bookingMethodCd int, languageCd int) ([]*model.CBookingMethod, *errordef.LogicError) {
            return []*model.CBookingMethod{&model.CBookingMethod{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cBookingMethodMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cBookingMethodService := NewCBookingMethodService(repository, numberUtil)
    in_cBookingMethod := new(model.CBookingMethod)
    in_cBookingMethod.BookingMethodCd = 0
    in_cBookingMethod.LanguageCd = 1
    in_cBookingMethod.BookingMethodName = "dummy-BookingMethodName"
    out_cBookingMethod, err := cBookingMethodService.UpdateCBookingMethod(&ctx, in_cBookingMethod)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cBookingMethod.BookingMethodCd)
    assert.Equal(t, 1, out_cBookingMethod.LanguageCd)
    assert.Equal(t, "dummy-BookingMethodName", out_cBookingMethod.BookingMethodName)
    assert.NotNil(t, out_cBookingMethod.CreateDatetime)
    assert.Equal(t, "", out_cBookingMethod.CreateDatetime)
    assert.Equal(t, "", out_cBookingMethod.CreateFunction)
    assert.NotNil(t, out_cBookingMethod.UpdateDatetime)
    assert.NotEqual(t, "", out_cBookingMethod.UpdateDatetime)
    assert.Equal(t, "UpdateCBookingMethod", out_cBookingMethod.UpdateFunction)
}
