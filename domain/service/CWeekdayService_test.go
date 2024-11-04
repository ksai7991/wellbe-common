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

type cWeekdayMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CWeekday) (*model.CWeekday, *errordef.LogicError)
    FakeUpdate func(*model.CWeekday) (*model.CWeekday, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.CWeekday, *errordef.LogicError)
}

func (lr cWeekdayMockRepository) CreateCWeekday(ctx *context.Context, cWeekday *model.CWeekday)  (*model.CWeekday, *errordef.LogicError) {
    return lr.FakeCreate(cWeekday)
}

func (lr cWeekdayMockRepository) UpdateCWeekday(ctx *context.Context, cWeekday *model.CWeekday)  (*model.CWeekday, *errordef.LogicError) {
    return lr.FakeUpdate(cWeekday)
}

func (lr cWeekdayMockRepository) DeleteCWeekday(ctx *context.Context, weekdayCd int)  *errordef.LogicError {
    return lr.FakeDelete(weekdayCd)
}

func (lr cWeekdayMockRepository)GetCWeekdayWithKey(ctx *context.Context, weekdayCd int)  ([]*model.CWeekday, *errordef.LogicError) {
    return lr.FakeGet(weekdayCd)
}


type cWeekdayMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cWeekdayMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCWeekday(t *testing.T) {
    ctx := context.Background()
    repository := &cWeekdayMockRepository{
        FakeCreate: func(cWeekday *model.CWeekday) (*model.CWeekday, *errordef.LogicError) {
            return cWeekday, nil
        },
    }
    numberUtil := &cWeekdayMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cWeekdayService := NewCWeekdayService(repository, numberUtil)
    in_cWeekday := new(model.CWeekday)
    in_cWeekday.WeekdayCd = 0
    in_cWeekday.LanguageCd = 1
    in_cWeekday.WeekdayName = "dummy-WeekdayName"
    in_cWeekday.WeekdayAbbreviation = "dummy-WeekdayAbbreviation"
    out_cWeekday, err := cWeekdayService.CreateCWeekday(&ctx, in_cWeekday)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cWeekday.WeekdayCd)
    assert.Equal(t, 1, out_cWeekday.LanguageCd)
    assert.Equal(t, "dummy-WeekdayName", out_cWeekday.WeekdayName)
    assert.Equal(t, "dummy-WeekdayAbbreviation", out_cWeekday.WeekdayAbbreviation)
    assert.NotNil(t, out_cWeekday.CreateDatetime)
    assert.NotEqual(t, "", out_cWeekday.CreateDatetime)
    assert.Equal(t, "CreateCWeekday", out_cWeekday.CreateFunction)
    assert.NotNil(t, out_cWeekday.UpdateDatetime)
    assert.NotEqual(t, "", out_cWeekday.UpdateDatetime)
    assert.Equal(t, "CreateCWeekday", out_cWeekday.UpdateFunction)
}

func TestUpdateCWeekday(t *testing.T) {
    ctx := context.Background()
    repository := &cWeekdayMockRepository{
        FakeUpdate: func(cWeekday *model.CWeekday) (*model.CWeekday, *errordef.LogicError) {
            return cWeekday, nil
        },
        FakeGet: func(weekdayCd int) ([]*model.CWeekday, *errordef.LogicError) {
            return []*model.CWeekday{&model.CWeekday{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cWeekdayMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cWeekdayService := NewCWeekdayService(repository, numberUtil)
    in_cWeekday := new(model.CWeekday)
    in_cWeekday.WeekdayCd = 0
    in_cWeekday.LanguageCd = 1
    in_cWeekday.WeekdayName = "dummy-WeekdayName"
    in_cWeekday.WeekdayAbbreviation = "dummy-WeekdayAbbreviation"
    out_cWeekday, err := cWeekdayService.UpdateCWeekday(&ctx, in_cWeekday)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cWeekday.WeekdayCd)
    assert.Equal(t, 1, out_cWeekday.LanguageCd)
    assert.Equal(t, "dummy-WeekdayName", out_cWeekday.WeekdayName)
    assert.Equal(t, "dummy-WeekdayAbbreviation", out_cWeekday.WeekdayAbbreviation)
    assert.NotNil(t, out_cWeekday.CreateDatetime)
    assert.Equal(t, "", out_cWeekday.CreateDatetime)
    assert.Equal(t, "", out_cWeekday.CreateFunction)
    assert.NotNil(t, out_cWeekday.UpdateDatetime)
    assert.NotEqual(t, "", out_cWeekday.UpdateDatetime)
    assert.Equal(t, "UpdateCWeekday", out_cWeekday.UpdateFunction)
}
