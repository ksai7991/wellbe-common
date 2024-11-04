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

type cAgeRangeMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CAgeRange) (*model.CAgeRange, *errordef.LogicError)
    FakeUpdate func(*model.CAgeRange) (*model.CAgeRange, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CAgeRange, *errordef.LogicError)
}

func (lr cAgeRangeMockRepository) CreateCAgeRange(ctx *context.Context, cAgeRange *model.CAgeRange)  (*model.CAgeRange, *errordef.LogicError) {
    return lr.FakeCreate(cAgeRange)
}

func (lr cAgeRangeMockRepository) UpdateCAgeRange(ctx *context.Context, cAgeRange *model.CAgeRange)  (*model.CAgeRange, *errordef.LogicError) {
    return lr.FakeUpdate(cAgeRange)
}

func (lr cAgeRangeMockRepository) DeleteCAgeRange(ctx *context.Context, ageRangeCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(ageRangeCd, languageCd)
}

func (lr cAgeRangeMockRepository)GetCAgeRangeWithKey(ctx *context.Context, ageRangeCd int, languageCd int)  ([]*model.CAgeRange, *errordef.LogicError) {
    return lr.FakeGet(ageRangeCd, languageCd)
}


type cAgeRangeMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cAgeRangeMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCAgeRange(t *testing.T) {
    ctx := context.Background()
    repository := &cAgeRangeMockRepository{
        FakeCreate: func(cAgeRange *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError) {
            return cAgeRange, nil
        },
    }
    numberUtil := &cAgeRangeMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cAgeRangeService := NewCAgeRangeService(repository, numberUtil)
    in_cAgeRange := new(model.CAgeRange)
    in_cAgeRange.AgeRangeCd = 0
    in_cAgeRange.LanguageCd = 1
    in_cAgeRange.AgeRangeGender = "dummy-AgeRangeGender"
    in_cAgeRange.AgeRangeFrom = 3
    in_cAgeRange.AgeRangeTo = 4
    out_cAgeRange, err := cAgeRangeService.CreateCAgeRange(&ctx, in_cAgeRange)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cAgeRange.AgeRangeCd)
    assert.Equal(t, 1, out_cAgeRange.LanguageCd)
    assert.Equal(t, "dummy-AgeRangeGender", out_cAgeRange.AgeRangeGender)
    assert.Equal(t, 3, out_cAgeRange.AgeRangeFrom)
    assert.Equal(t, 4, out_cAgeRange.AgeRangeTo)
    assert.NotNil(t, out_cAgeRange.CreateDatetime)
    assert.NotEqual(t, "", out_cAgeRange.CreateDatetime)
    assert.Equal(t, "CreateCAgeRange", out_cAgeRange.CreateFunction)
    assert.NotNil(t, out_cAgeRange.UpdateDatetime)
    assert.NotEqual(t, "", out_cAgeRange.UpdateDatetime)
    assert.Equal(t, "CreateCAgeRange", out_cAgeRange.UpdateFunction)
}

func TestUpdateCAgeRange(t *testing.T) {
    ctx := context.Background()
    repository := &cAgeRangeMockRepository{
        FakeUpdate: func(cAgeRange *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError) {
            return cAgeRange, nil
        },
        FakeGet: func(ageRangeCd int, languageCd int) ([]*model.CAgeRange, *errordef.LogicError) {
            return []*model.CAgeRange{&model.CAgeRange{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cAgeRangeMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cAgeRangeService := NewCAgeRangeService(repository, numberUtil)
    in_cAgeRange := new(model.CAgeRange)
    in_cAgeRange.AgeRangeCd = 0
    in_cAgeRange.LanguageCd = 1
    in_cAgeRange.AgeRangeGender = "dummy-AgeRangeGender"
    in_cAgeRange.AgeRangeFrom = 3
    in_cAgeRange.AgeRangeTo = 4
    out_cAgeRange, err := cAgeRangeService.UpdateCAgeRange(&ctx, in_cAgeRange)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cAgeRange.AgeRangeCd)
    assert.Equal(t, 1, out_cAgeRange.LanguageCd)
    assert.Equal(t, "dummy-AgeRangeGender", out_cAgeRange.AgeRangeGender)
    assert.Equal(t, 3, out_cAgeRange.AgeRangeFrom)
    assert.Equal(t, 4, out_cAgeRange.AgeRangeTo)
    assert.NotNil(t, out_cAgeRange.CreateDatetime)
    assert.Equal(t, "", out_cAgeRange.CreateDatetime)
    assert.Equal(t, "", out_cAgeRange.CreateFunction)
    assert.NotNil(t, out_cAgeRange.UpdateDatetime)
    assert.NotEqual(t, "", out_cAgeRange.UpdateDatetime)
    assert.Equal(t, "UpdateCAgeRange", out_cAgeRange.UpdateFunction)
}
