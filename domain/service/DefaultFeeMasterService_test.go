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

type defaultFeeMasterMockRepository struct{
    repository.Repository
    FakeCreate func(*model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError)
    FakeUpdate func(*model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError)
    FakeDelete func(string) *errordef.LogicError
    FakeGet func(string) ([]*model.DefaultFeeMaster, *errordef.LogicError)
}

func (lr defaultFeeMasterMockRepository) CreateDefaultFeeMaster(ctx *context.Context, defaultFeeMaster *model.DefaultFeeMaster)  (*model.DefaultFeeMaster, *errordef.LogicError) {
    return lr.FakeCreate(defaultFeeMaster)
}

func (lr defaultFeeMasterMockRepository) UpdateDefaultFeeMaster(ctx *context.Context, defaultFeeMaster *model.DefaultFeeMaster)  (*model.DefaultFeeMaster, *errordef.LogicError) {
    return lr.FakeUpdate(defaultFeeMaster)
}

func (lr defaultFeeMasterMockRepository) DeleteDefaultFeeMaster(ctx *context.Context, id string)  *errordef.LogicError {
    return lr.FakeDelete(id)
}

func (lr defaultFeeMasterMockRepository)GetDefaultFeeMasterWithKey(ctx *context.Context, id string)  ([]*model.DefaultFeeMaster, *errordef.LogicError) {
    return lr.FakeGet(id)
}


type defaultFeeMasterMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr defaultFeeMasterMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateDefaultFeeMaster(t *testing.T) {
    ctx := context.Background()
    repository := &defaultFeeMasterMockRepository{
        FakeCreate: func(defaultFeeMaster *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError) {
            return defaultFeeMaster, nil
        },
    }
    numberUtil := &defaultFeeMasterMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    defaultFeeMasterService := NewDefaultFeeMasterService(repository, numberUtil)
    in_defaultFeeMaster := new(model.DefaultFeeMaster)
    in_defaultFeeMaster.Id = "dummy-Id"
    in_defaultFeeMaster.FeeRate = 1.2
    out_defaultFeeMaster, err := defaultFeeMasterService.CreateDefaultFeeMaster(&ctx, in_defaultFeeMaster)
    assert.Nil(t, err)
    assert.Equal(t, "dummy-Id", out_defaultFeeMaster.Id)
    assert.Equal(t, 1.2, out_defaultFeeMaster.FeeRate)
    assert.NotNil(t, out_defaultFeeMaster.CreateDatetime)
    assert.NotEqual(t, "", out_defaultFeeMaster.CreateDatetime)
    assert.Equal(t, "CreateDefaultFeeMaster", out_defaultFeeMaster.CreateFunction)
    assert.NotNil(t, out_defaultFeeMaster.UpdateDatetime)
    assert.NotEqual(t, "", out_defaultFeeMaster.UpdateDatetime)
    assert.Equal(t, "CreateDefaultFeeMaster", out_defaultFeeMaster.UpdateFunction)
}

func TestUpdateDefaultFeeMaster(t *testing.T) {
    ctx := context.Background()
    repository := &defaultFeeMasterMockRepository{
        FakeUpdate: func(defaultFeeMaster *model.DefaultFeeMaster) (*model.DefaultFeeMaster, *errordef.LogicError) {
            return defaultFeeMaster, nil
        },
        FakeGet: func(id string) ([]*model.DefaultFeeMaster, *errordef.LogicError) {
            return []*model.DefaultFeeMaster{&model.DefaultFeeMaster{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &defaultFeeMasterMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    defaultFeeMasterService := NewDefaultFeeMasterService(repository, numberUtil)
    in_defaultFeeMaster := new(model.DefaultFeeMaster)
    in_defaultFeeMaster.Id = "dummy-Id"
    in_defaultFeeMaster.FeeRate = 1.2
    out_defaultFeeMaster, err := defaultFeeMasterService.UpdateDefaultFeeMaster(&ctx, in_defaultFeeMaster)
    assert.Nil(t, err)
    assert.Equal(t, "dummy-Id", out_defaultFeeMaster.Id)
    assert.Equal(t, 1.2, out_defaultFeeMaster.FeeRate)
    assert.NotNil(t, out_defaultFeeMaster.CreateDatetime)
    assert.Equal(t, "", out_defaultFeeMaster.CreateDatetime)
    assert.Equal(t, "", out_defaultFeeMaster.CreateFunction)
    assert.NotNil(t, out_defaultFeeMaster.UpdateDatetime)
    assert.NotEqual(t, "", out_defaultFeeMaster.UpdateDatetime)
    assert.Equal(t, "UpdateDefaultFeeMaster", out_defaultFeeMaster.UpdateFunction)
}
