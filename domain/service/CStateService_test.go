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

type cStateMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CState) (*model.CState, *errordef.LogicError)
    FakeUpdate func(*model.CState) (*model.CState, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CState, *errordef.LogicError)
}

func (lr cStateMockRepository) CreateCState(ctx *context.Context, cState *model.CState)  (*model.CState, *errordef.LogicError) {
    return lr.FakeCreate(cState)
}

func (lr cStateMockRepository) UpdateCState(ctx *context.Context, cState *model.CState)  (*model.CState, *errordef.LogicError) {
    return lr.FakeUpdate(cState)
}

func (lr cStateMockRepository) DeleteCState(ctx *context.Context, stateCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(stateCd, languageCd)
}

func (lr cStateMockRepository)GetCStateWithKey(ctx *context.Context, stateCd int, languageCd int)  ([]*model.CState, *errordef.LogicError) {
    return lr.FakeGet(stateCd, languageCd)
}


type cStateMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cStateMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCState(t *testing.T) {
    ctx := context.Background()
    repository := &cStateMockRepository{
        FakeCreate: func(cState *model.CState) (*model.CState, *errordef.LogicError) {
            return cState, nil
        },
    }
    numberUtil := &cStateMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cStateService := NewCStateService(repository, numberUtil)
    in_cState := new(model.CState)
    in_cState.StateCd = 0
    in_cState.LanguageCd = 1
    in_cState.CountryCd = 2
    in_cState.StateName = "dummy-StateName"
    in_cState.StateCdIso = "dummy-StateCdIso"
    in_cState.TimezoneIana = "dummy-TimezoneIana"
    out_cState, err := cStateService.CreateCState(&ctx, in_cState)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cState.StateCd)
    assert.Equal(t, 1, out_cState.LanguageCd)
    assert.Equal(t, 2, out_cState.CountryCd)
    assert.Equal(t, "dummy-StateName", out_cState.StateName)
    assert.Equal(t, "dummy-StateCdIso", out_cState.StateCdIso)
    assert.Equal(t, "dummy-TimezoneIana", out_cState.TimezoneIana)
    assert.NotNil(t, out_cState.CreateDatetime)
    assert.NotEqual(t, "", out_cState.CreateDatetime)
    assert.Equal(t, "CreateCState", out_cState.CreateFunction)
    assert.NotNil(t, out_cState.UpdateDatetime)
    assert.NotEqual(t, "", out_cState.UpdateDatetime)
    assert.Equal(t, "CreateCState", out_cState.UpdateFunction)
}

func TestUpdateCState(t *testing.T) {
    ctx := context.Background()
    repository := &cStateMockRepository{
        FakeUpdate: func(cState *model.CState) (*model.CState, *errordef.LogicError) {
            return cState, nil
        },
        FakeGet: func(stateCd int, languageCd int) ([]*model.CState, *errordef.LogicError) {
            return []*model.CState{&model.CState{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cStateMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cStateService := NewCStateService(repository, numberUtil)
    in_cState := new(model.CState)
    in_cState.StateCd = 0
    in_cState.LanguageCd = 1
    in_cState.CountryCd = 2
    in_cState.StateName = "dummy-StateName"
    in_cState.StateCdIso = "dummy-StateCdIso"
    in_cState.TimezoneIana = "dummy-TimezoneIana"
    out_cState, err := cStateService.UpdateCState(&ctx, in_cState)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cState.StateCd)
    assert.Equal(t, 1, out_cState.LanguageCd)
    assert.Equal(t, 2, out_cState.CountryCd)
    assert.Equal(t, "dummy-StateName", out_cState.StateName)
    assert.Equal(t, "dummy-StateCdIso", out_cState.StateCdIso)
    assert.Equal(t, "dummy-TimezoneIana", out_cState.TimezoneIana)
    assert.NotNil(t, out_cState.CreateDatetime)
    assert.Equal(t, "", out_cState.CreateDatetime)
    assert.Equal(t, "", out_cState.CreateFunction)
    assert.NotNil(t, out_cState.UpdateDatetime)
    assert.NotEqual(t, "", out_cState.UpdateDatetime)
    assert.Equal(t, "UpdateCState", out_cState.UpdateFunction)
}
