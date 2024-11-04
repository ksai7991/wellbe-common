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

type cTreatmentTimeRangeMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError)
    FakeUpdate func(*model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CTreatmentTimeRange, *errordef.LogicError)
}

func (lr cTreatmentTimeRangeMockRepository) CreateCTreatmentTimeRange(ctx *context.Context, cTreatmentTimeRange *model.CTreatmentTimeRange)  (*model.CTreatmentTimeRange, *errordef.LogicError) {
    return lr.FakeCreate(cTreatmentTimeRange)
}

func (lr cTreatmentTimeRangeMockRepository) UpdateCTreatmentTimeRange(ctx *context.Context, cTreatmentTimeRange *model.CTreatmentTimeRange)  (*model.CTreatmentTimeRange, *errordef.LogicError) {
    return lr.FakeUpdate(cTreatmentTimeRange)
}

func (lr cTreatmentTimeRangeMockRepository) DeleteCTreatmentTimeRange(ctx *context.Context, treatmentTimeCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(treatmentTimeCd, languageCd)
}

func (lr cTreatmentTimeRangeMockRepository)GetCTreatmentTimeRangeWithKey(ctx *context.Context, treatmentTimeCd int, languageCd int)  ([]*model.CTreatmentTimeRange, *errordef.LogicError) {
    return lr.FakeGet(treatmentTimeCd, languageCd)
}


type cTreatmentTimeRangeMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cTreatmentTimeRangeMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCTreatmentTimeRange(t *testing.T) {
    ctx := context.Background()
    repository := &cTreatmentTimeRangeMockRepository{
        FakeCreate: func(cTreatmentTimeRange *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError) {
            return cTreatmentTimeRange, nil
        },
    }
    numberUtil := &cTreatmentTimeRangeMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cTreatmentTimeRangeService := NewCTreatmentTimeRangeService(repository, numberUtil)
    in_cTreatmentTimeRange := new(model.CTreatmentTimeRange)
    in_cTreatmentTimeRange.TreatmentTimeCd = 0
    in_cTreatmentTimeRange.LanguageCd = 1
    in_cTreatmentTimeRange.TreatmentTimeName = "dummy-TreatmentTimeName"
    in_cTreatmentTimeRange.MinTime = 3
    in_cTreatmentTimeRange.MaxTime = 4
    out_cTreatmentTimeRange, err := cTreatmentTimeRangeService.CreateCTreatmentTimeRange(&ctx, in_cTreatmentTimeRange)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cTreatmentTimeRange.TreatmentTimeCd)
    assert.Equal(t, 1, out_cTreatmentTimeRange.LanguageCd)
    assert.Equal(t, "dummy-TreatmentTimeName", out_cTreatmentTimeRange.TreatmentTimeName)
    assert.Equal(t, 3, out_cTreatmentTimeRange.MinTime)
    assert.Equal(t, 4, out_cTreatmentTimeRange.MaxTime)
    assert.NotNil(t, out_cTreatmentTimeRange.CreateDatetime)
    assert.NotEqual(t, "", out_cTreatmentTimeRange.CreateDatetime)
    assert.Equal(t, "CreateCTreatmentTimeRange", out_cTreatmentTimeRange.CreateFunction)
    assert.NotNil(t, out_cTreatmentTimeRange.UpdateDatetime)
    assert.NotEqual(t, "", out_cTreatmentTimeRange.UpdateDatetime)
    assert.Equal(t, "CreateCTreatmentTimeRange", out_cTreatmentTimeRange.UpdateFunction)
}

func TestUpdateCTreatmentTimeRange(t *testing.T) {
    ctx := context.Background()
    repository := &cTreatmentTimeRangeMockRepository{
        FakeUpdate: func(cTreatmentTimeRange *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError) {
            return cTreatmentTimeRange, nil
        },
        FakeGet: func(treatmentTimeCd int, languageCd int) ([]*model.CTreatmentTimeRange, *errordef.LogicError) {
            return []*model.CTreatmentTimeRange{&model.CTreatmentTimeRange{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cTreatmentTimeRangeMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cTreatmentTimeRangeService := NewCTreatmentTimeRangeService(repository, numberUtil)
    in_cTreatmentTimeRange := new(model.CTreatmentTimeRange)
    in_cTreatmentTimeRange.TreatmentTimeCd = 0
    in_cTreatmentTimeRange.LanguageCd = 1
    in_cTreatmentTimeRange.TreatmentTimeName = "dummy-TreatmentTimeName"
    in_cTreatmentTimeRange.MinTime = 3
    in_cTreatmentTimeRange.MaxTime = 4
    out_cTreatmentTimeRange, err := cTreatmentTimeRangeService.UpdateCTreatmentTimeRange(&ctx, in_cTreatmentTimeRange)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cTreatmentTimeRange.TreatmentTimeCd)
    assert.Equal(t, 1, out_cTreatmentTimeRange.LanguageCd)
    assert.Equal(t, "dummy-TreatmentTimeName", out_cTreatmentTimeRange.TreatmentTimeName)
    assert.Equal(t, 3, out_cTreatmentTimeRange.MinTime)
    assert.Equal(t, 4, out_cTreatmentTimeRange.MaxTime)
    assert.NotNil(t, out_cTreatmentTimeRange.CreateDatetime)
    assert.Equal(t, "", out_cTreatmentTimeRange.CreateDatetime)
    assert.Equal(t, "", out_cTreatmentTimeRange.CreateFunction)
    assert.NotNil(t, out_cTreatmentTimeRange.UpdateDatetime)
    assert.NotEqual(t, "", out_cTreatmentTimeRange.UpdateDatetime)
    assert.Equal(t, "UpdateCTreatmentTimeRange", out_cTreatmentTimeRange.UpdateFunction)
}
