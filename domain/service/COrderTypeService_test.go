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

type cOrderTypeMockRepository struct{
    repository.Repository
    FakeCreate func(*model.COrderType) (*model.COrderType, *errordef.LogicError)
    FakeUpdate func(*model.COrderType) (*model.COrderType, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.COrderType, *errordef.LogicError)
}

func (lr cOrderTypeMockRepository) CreateCOrderType(ctx *context.Context, cOrderType *model.COrderType)  (*model.COrderType, *errordef.LogicError) {
    return lr.FakeCreate(cOrderType)
}

func (lr cOrderTypeMockRepository) UpdateCOrderType(ctx *context.Context, cOrderType *model.COrderType)  (*model.COrderType, *errordef.LogicError) {
    return lr.FakeUpdate(cOrderType)
}

func (lr cOrderTypeMockRepository) DeleteCOrderType(ctx *context.Context, orderTypeCd int)  *errordef.LogicError {
    return lr.FakeDelete(orderTypeCd)
}

func (lr cOrderTypeMockRepository)GetCOrderTypeWithKey(ctx *context.Context, orderTypeCd int)  ([]*model.COrderType, *errordef.LogicError) {
    return lr.FakeGet(orderTypeCd)
}


type cOrderTypeMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cOrderTypeMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCOrderType(t *testing.T) {
    ctx := context.Background()
    repository := &cOrderTypeMockRepository{
        FakeCreate: func(cOrderType *model.COrderType) (*model.COrderType, *errordef.LogicError) {
            return cOrderType, nil
        },
    }
    numberUtil := &cOrderTypeMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cOrderTypeService := NewCOrderTypeService(repository, numberUtil)
    in_cOrderType := new(model.COrderType)
    in_cOrderType.OrderTypeCd = 0
    in_cOrderType.LanguageCd = 1
    in_cOrderType.OrderTypeName = "dummy-OrderTypeName"
    out_cOrderType, err := cOrderTypeService.CreateCOrderType(&ctx, in_cOrderType)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cOrderType.OrderTypeCd)
    assert.Equal(t, 1, out_cOrderType.LanguageCd)
    assert.Equal(t, "dummy-OrderTypeName", out_cOrderType.OrderTypeName)
    assert.NotNil(t, out_cOrderType.CreateDatetime)
    assert.NotEqual(t, "", out_cOrderType.CreateDatetime)
    assert.Equal(t, "CreateCOrderType", out_cOrderType.CreateFunction)
    assert.NotNil(t, out_cOrderType.UpdateDatetime)
    assert.NotEqual(t, "", out_cOrderType.UpdateDatetime)
    assert.Equal(t, "CreateCOrderType", out_cOrderType.UpdateFunction)
}

func TestUpdateCOrderType(t *testing.T) {
    ctx := context.Background()
    repository := &cOrderTypeMockRepository{
        FakeUpdate: func(cOrderType *model.COrderType) (*model.COrderType, *errordef.LogicError) {
            return cOrderType, nil
        },
        FakeGet: func(orderTypeCd int) ([]*model.COrderType, *errordef.LogicError) {
            return []*model.COrderType{&model.COrderType{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cOrderTypeMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cOrderTypeService := NewCOrderTypeService(repository, numberUtil)
    in_cOrderType := new(model.COrderType)
    in_cOrderType.OrderTypeCd = 0
    in_cOrderType.LanguageCd = 1
    in_cOrderType.OrderTypeName = "dummy-OrderTypeName"
    out_cOrderType, err := cOrderTypeService.UpdateCOrderType(&ctx, in_cOrderType)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cOrderType.OrderTypeCd)
    assert.Equal(t, 1, out_cOrderType.LanguageCd)
    assert.Equal(t, "dummy-OrderTypeName", out_cOrderType.OrderTypeName)
    assert.NotNil(t, out_cOrderType.CreateDatetime)
    assert.Equal(t, "", out_cOrderType.CreateDatetime)
    assert.Equal(t, "", out_cOrderType.CreateFunction)
    assert.NotNil(t, out_cOrderType.UpdateDatetime)
    assert.NotEqual(t, "", out_cOrderType.UpdateDatetime)
    assert.Equal(t, "UpdateCOrderType", out_cOrderType.UpdateFunction)
}
