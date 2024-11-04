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

type cShopEquipmentMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError)
    FakeUpdate func(*model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CShopEquipment, *errordef.LogicError)
}

func (lr cShopEquipmentMockRepository) CreateCShopEquipment(ctx *context.Context, cShopEquipment *model.CShopEquipment)  (*model.CShopEquipment, *errordef.LogicError) {
    return lr.FakeCreate(cShopEquipment)
}

func (lr cShopEquipmentMockRepository) UpdateCShopEquipment(ctx *context.Context, cShopEquipment *model.CShopEquipment)  (*model.CShopEquipment, *errordef.LogicError) {
    return lr.FakeUpdate(cShopEquipment)
}

func (lr cShopEquipmentMockRepository) DeleteCShopEquipment(ctx *context.Context, shopEquipmentCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(shopEquipmentCd, languageCd)
}

func (lr cShopEquipmentMockRepository)GetCShopEquipmentWithKey(ctx *context.Context, shopEquipmentCd int, languageCd int)  ([]*model.CShopEquipment, *errordef.LogicError) {
    return lr.FakeGet(shopEquipmentCd, languageCd)
}


type cShopEquipmentMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cShopEquipmentMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCShopEquipment(t *testing.T) {
    ctx := context.Background()
    repository := &cShopEquipmentMockRepository{
        FakeCreate: func(cShopEquipment *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError) {
            return cShopEquipment, nil
        },
    }
    numberUtil := &cShopEquipmentMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopEquipmentService := NewCShopEquipmentService(repository, numberUtil)
    in_cShopEquipment := new(model.CShopEquipment)
    in_cShopEquipment.ShopEquipmentCd = 0
    in_cShopEquipment.LanguageCd = 1
    in_cShopEquipment.ShopEquipmentName = "dummy-ShopEquipmentName"
    in_cShopEquipment.UnitName = "dummy-UnitName"
    out_cShopEquipment, err := cShopEquipmentService.CreateCShopEquipment(&ctx, in_cShopEquipment)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopEquipment.ShopEquipmentCd)
    assert.Equal(t, 1, out_cShopEquipment.LanguageCd)
    assert.Equal(t, "dummy-ShopEquipmentName", out_cShopEquipment.ShopEquipmentName)
    assert.Equal(t, "dummy-UnitName", out_cShopEquipment.UnitName)
    assert.NotNil(t, out_cShopEquipment.CreateDatetime)
    assert.NotEqual(t, "", out_cShopEquipment.CreateDatetime)
    assert.Equal(t, "CreateCShopEquipment", out_cShopEquipment.CreateFunction)
    assert.NotNil(t, out_cShopEquipment.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopEquipment.UpdateDatetime)
    assert.Equal(t, "CreateCShopEquipment", out_cShopEquipment.UpdateFunction)
}

func TestUpdateCShopEquipment(t *testing.T) {
    ctx := context.Background()
    repository := &cShopEquipmentMockRepository{
        FakeUpdate: func(cShopEquipment *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError) {
            return cShopEquipment, nil
        },
        FakeGet: func(shopEquipmentCd int, languageCd int) ([]*model.CShopEquipment, *errordef.LogicError) {
            return []*model.CShopEquipment{&model.CShopEquipment{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cShopEquipmentMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopEquipmentService := NewCShopEquipmentService(repository, numberUtil)
    in_cShopEquipment := new(model.CShopEquipment)
    in_cShopEquipment.ShopEquipmentCd = 0
    in_cShopEquipment.LanguageCd = 1
    in_cShopEquipment.ShopEquipmentName = "dummy-ShopEquipmentName"
    in_cShopEquipment.UnitName = "dummy-UnitName"
    out_cShopEquipment, err := cShopEquipmentService.UpdateCShopEquipment(&ctx, in_cShopEquipment)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopEquipment.ShopEquipmentCd)
    assert.Equal(t, 1, out_cShopEquipment.LanguageCd)
    assert.Equal(t, "dummy-ShopEquipmentName", out_cShopEquipment.ShopEquipmentName)
    assert.Equal(t, "dummy-UnitName", out_cShopEquipment.UnitName)
    assert.NotNil(t, out_cShopEquipment.CreateDatetime)
    assert.Equal(t, "", out_cShopEquipment.CreateDatetime)
    assert.Equal(t, "", out_cShopEquipment.CreateFunction)
    assert.NotNil(t, out_cShopEquipment.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopEquipment.UpdateDatetime)
    assert.Equal(t, "UpdateCShopEquipment", out_cShopEquipment.UpdateFunction)
}
