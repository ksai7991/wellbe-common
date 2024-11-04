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

type cConcernMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CConcern) (*model.CConcern, *errordef.LogicError)
    FakeUpdate func(*model.CConcern) (*model.CConcern, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CConcern, *errordef.LogicError)
}

func (lr cConcernMockRepository) CreateCConcern(ctx *context.Context, cConcern *model.CConcern)  (*model.CConcern, *errordef.LogicError) {
    return lr.FakeCreate(cConcern)
}

func (lr cConcernMockRepository) UpdateCConcern(ctx *context.Context, cConcern *model.CConcern)  (*model.CConcern, *errordef.LogicError) {
    return lr.FakeUpdate(cConcern)
}

func (lr cConcernMockRepository) DeleteCConcern(ctx *context.Context, concernCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(concernCd, languageCd)
}

func (lr cConcernMockRepository)GetCConcernWithKey(ctx *context.Context, concernCd int, languageCd int)  ([]*model.CConcern, *errordef.LogicError) {
    return lr.FakeGet(concernCd, languageCd)
}


type cConcernMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cConcernMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCConcern(t *testing.T) {
    ctx := context.Background()
    repository := &cConcernMockRepository{
        FakeCreate: func(cConcern *model.CConcern) (*model.CConcern, *errordef.LogicError) {
            return cConcern, nil
        },
    }
    numberUtil := &cConcernMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cConcernService := NewCConcernService(repository, numberUtil)
    in_cConcern := new(model.CConcern)
    in_cConcern.ConcernCd = 0
    in_cConcern.LanguageCd = 1
    in_cConcern.ConcernName = "dummy-ConcernName"
    out_cConcern, err := cConcernService.CreateCConcern(&ctx, in_cConcern)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cConcern.ConcernCd)
    assert.Equal(t, 1, out_cConcern.LanguageCd)
    assert.Equal(t, "dummy-ConcernName", out_cConcern.ConcernName)
    assert.NotNil(t, out_cConcern.CreateDatetime)
    assert.NotEqual(t, "", out_cConcern.CreateDatetime)
    assert.Equal(t, "CreateCConcern", out_cConcern.CreateFunction)
    assert.NotNil(t, out_cConcern.UpdateDatetime)
    assert.NotEqual(t, "", out_cConcern.UpdateDatetime)
    assert.Equal(t, "CreateCConcern", out_cConcern.UpdateFunction)
}

func TestUpdateCConcern(t *testing.T) {
    ctx := context.Background()
    repository := &cConcernMockRepository{
        FakeUpdate: func(cConcern *model.CConcern) (*model.CConcern, *errordef.LogicError) {
            return cConcern, nil
        },
        FakeGet: func(concernCd int, languageCd int) ([]*model.CConcern, *errordef.LogicError) {
            return []*model.CConcern{&model.CConcern{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cConcernMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cConcernService := NewCConcernService(repository, numberUtil)
    in_cConcern := new(model.CConcern)
    in_cConcern.ConcernCd = 0
    in_cConcern.LanguageCd = 1
    in_cConcern.ConcernName = "dummy-ConcernName"
    out_cConcern, err := cConcernService.UpdateCConcern(&ctx, in_cConcern)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cConcern.ConcernCd)
    assert.Equal(t, 1, out_cConcern.LanguageCd)
    assert.Equal(t, "dummy-ConcernName", out_cConcern.ConcernName)
    assert.NotNil(t, out_cConcern.CreateDatetime)
    assert.Equal(t, "", out_cConcern.CreateDatetime)
    assert.Equal(t, "", out_cConcern.CreateFunction)
    assert.NotNil(t, out_cConcern.UpdateDatetime)
    assert.NotEqual(t, "", out_cConcern.UpdateDatetime)
    assert.Equal(t, "UpdateCConcern", out_cConcern.UpdateFunction)
}
