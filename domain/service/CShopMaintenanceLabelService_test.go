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

type cShopMaintenanceLabelMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError)
    FakeUpdate func(*model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError)
}

func (lr cShopMaintenanceLabelMockRepository) CreateCShopMaintenanceLabel(ctx *context.Context, cShopMaintenanceLabel *model.CShopMaintenanceLabel)  (*model.CShopMaintenanceLabel, *errordef.LogicError) {
    return lr.FakeCreate(cShopMaintenanceLabel)
}

func (lr cShopMaintenanceLabelMockRepository) UpdateCShopMaintenanceLabel(ctx *context.Context, cShopMaintenanceLabel *model.CShopMaintenanceLabel)  (*model.CShopMaintenanceLabel, *errordef.LogicError) {
    return lr.FakeUpdate(cShopMaintenanceLabel)
}

func (lr cShopMaintenanceLabelMockRepository) DeleteCShopMaintenanceLabel(ctx *context.Context, shopMaintenanceLabelCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(shopMaintenanceLabelCd, languageCd)
}

func (lr cShopMaintenanceLabelMockRepository)GetCShopMaintenanceLabelWithKey(ctx *context.Context, shopMaintenanceLabelCd int, languageCd int)  ([]*model.CShopMaintenanceLabel, *errordef.LogicError) {
    return lr.FakeGet(shopMaintenanceLabelCd, languageCd)
}


type cShopMaintenanceLabelMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cShopMaintenanceLabelMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCShopMaintenanceLabel(t *testing.T) {
    ctx := context.Background()
    repository := &cShopMaintenanceLabelMockRepository{
        FakeCreate: func(cShopMaintenanceLabel *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError) {
            return cShopMaintenanceLabel, nil
        },
    }
    numberUtil := &cShopMaintenanceLabelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopMaintenanceLabelService := NewCShopMaintenanceLabelService(repository, numberUtil)
    in_cShopMaintenanceLabel := new(model.CShopMaintenanceLabel)
    in_cShopMaintenanceLabel.ShopMaintenanceLabelCd = 0
    in_cShopMaintenanceLabel.LanguageCd = 1
    in_cShopMaintenanceLabel.ShopMaintenanceLabelName = "dummy-ShopMaintenanceLabelName"
    out_cShopMaintenanceLabel, err := cShopMaintenanceLabelService.CreateCShopMaintenanceLabel(&ctx, in_cShopMaintenanceLabel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopMaintenanceLabel.ShopMaintenanceLabelCd)
    assert.Equal(t, 1, out_cShopMaintenanceLabel.LanguageCd)
    assert.Equal(t, "dummy-ShopMaintenanceLabelName", out_cShopMaintenanceLabel.ShopMaintenanceLabelName)
    assert.NotNil(t, out_cShopMaintenanceLabel.CreateDatetime)
    assert.NotEqual(t, "", out_cShopMaintenanceLabel.CreateDatetime)
    assert.Equal(t, "CreateCShopMaintenanceLabel", out_cShopMaintenanceLabel.CreateFunction)
    assert.NotNil(t, out_cShopMaintenanceLabel.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopMaintenanceLabel.UpdateDatetime)
    assert.Equal(t, "CreateCShopMaintenanceLabel", out_cShopMaintenanceLabel.UpdateFunction)
}

func TestUpdateCShopMaintenanceLabel(t *testing.T) {
    ctx := context.Background()
    repository := &cShopMaintenanceLabelMockRepository{
        FakeUpdate: func(cShopMaintenanceLabel *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError) {
            return cShopMaintenanceLabel, nil
        },
        FakeGet: func(shopMaintenanceLabelCd int, languageCd int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError) {
            return []*model.CShopMaintenanceLabel{&model.CShopMaintenanceLabel{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cShopMaintenanceLabelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopMaintenanceLabelService := NewCShopMaintenanceLabelService(repository, numberUtil)
    in_cShopMaintenanceLabel := new(model.CShopMaintenanceLabel)
    in_cShopMaintenanceLabel.ShopMaintenanceLabelCd = 0
    in_cShopMaintenanceLabel.LanguageCd = 1
    in_cShopMaintenanceLabel.ShopMaintenanceLabelName = "dummy-ShopMaintenanceLabelName"
    out_cShopMaintenanceLabel, err := cShopMaintenanceLabelService.UpdateCShopMaintenanceLabel(&ctx, in_cShopMaintenanceLabel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopMaintenanceLabel.ShopMaintenanceLabelCd)
    assert.Equal(t, 1, out_cShopMaintenanceLabel.LanguageCd)
    assert.Equal(t, "dummy-ShopMaintenanceLabelName", out_cShopMaintenanceLabel.ShopMaintenanceLabelName)
    assert.NotNil(t, out_cShopMaintenanceLabel.CreateDatetime)
    assert.Equal(t, "", out_cShopMaintenanceLabel.CreateDatetime)
    assert.Equal(t, "", out_cShopMaintenanceLabel.CreateFunction)
    assert.NotNil(t, out_cShopMaintenanceLabel.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopMaintenanceLabel.UpdateDatetime)
    assert.Equal(t, "UpdateCShopMaintenanceLabel", out_cShopMaintenanceLabel.UpdateFunction)
}
