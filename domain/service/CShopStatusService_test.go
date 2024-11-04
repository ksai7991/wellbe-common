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

type cShopStatusMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CShopStatus) (*model.CShopStatus, *errordef.LogicError)
    FakeUpdate func(*model.CShopStatus) (*model.CShopStatus, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CShopStatus, *errordef.LogicError)
}

func (lr cShopStatusMockRepository) CreateCShopStatus(ctx *context.Context, cShopStatus *model.CShopStatus)  (*model.CShopStatus, *errordef.LogicError) {
    return lr.FakeCreate(cShopStatus)
}

func (lr cShopStatusMockRepository) UpdateCShopStatus(ctx *context.Context, cShopStatus *model.CShopStatus)  (*model.CShopStatus, *errordef.LogicError) {
    return lr.FakeUpdate(cShopStatus)
}

func (lr cShopStatusMockRepository) DeleteCShopStatus(ctx *context.Context, shopStatusCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(shopStatusCd, languageCd)
}

func (lr cShopStatusMockRepository)GetCShopStatusWithKey(ctx *context.Context, shopStatusCd int, languageCd int)  ([]*model.CShopStatus, *errordef.LogicError) {
    return lr.FakeGet(shopStatusCd, languageCd)
}


type cShopStatusMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cShopStatusMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCShopStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cShopStatusMockRepository{
        FakeCreate: func(cShopStatus *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError) {
            return cShopStatus, nil
        },
    }
    numberUtil := &cShopStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopStatusService := NewCShopStatusService(repository, numberUtil)
    in_cShopStatus := new(model.CShopStatus)
    in_cShopStatus.ShopStatusCd = 0
    in_cShopStatus.LanguageCd = 1
    in_cShopStatus.ShopStatusName = "dummy-ShopStatusName"
    out_cShopStatus, err := cShopStatusService.CreateCShopStatus(&ctx, in_cShopStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopStatus.ShopStatusCd)
    assert.Equal(t, 1, out_cShopStatus.LanguageCd)
    assert.Equal(t, "dummy-ShopStatusName", out_cShopStatus.ShopStatusName)
    assert.NotNil(t, out_cShopStatus.CreateDatetime)
    assert.NotEqual(t, "", out_cShopStatus.CreateDatetime)
    assert.Equal(t, "CreateCShopStatus", out_cShopStatus.CreateFunction)
    assert.NotNil(t, out_cShopStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopStatus.UpdateDatetime)
    assert.Equal(t, "CreateCShopStatus", out_cShopStatus.UpdateFunction)
}

func TestUpdateCShopStatus(t *testing.T) {
    ctx := context.Background()
    repository := &cShopStatusMockRepository{
        FakeUpdate: func(cShopStatus *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError) {
            return cShopStatus, nil
        },
        FakeGet: func(shopStatusCd int, languageCd int) ([]*model.CShopStatus, *errordef.LogicError) {
            return []*model.CShopStatus{&model.CShopStatus{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cShopStatusMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cShopStatusService := NewCShopStatusService(repository, numberUtil)
    in_cShopStatus := new(model.CShopStatus)
    in_cShopStatus.ShopStatusCd = 0
    in_cShopStatus.LanguageCd = 1
    in_cShopStatus.ShopStatusName = "dummy-ShopStatusName"
    out_cShopStatus, err := cShopStatusService.UpdateCShopStatus(&ctx, in_cShopStatus)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cShopStatus.ShopStatusCd)
    assert.Equal(t, 1, out_cShopStatus.LanguageCd)
    assert.Equal(t, "dummy-ShopStatusName", out_cShopStatus.ShopStatusName)
    assert.NotNil(t, out_cShopStatus.CreateDatetime)
    assert.Equal(t, "", out_cShopStatus.CreateDatetime)
    assert.Equal(t, "", out_cShopStatus.CreateFunction)
    assert.NotNil(t, out_cShopStatus.UpdateDatetime)
    assert.NotEqual(t, "", out_cShopStatus.UpdateDatetime)
    assert.Equal(t, "UpdateCShopStatus", out_cShopStatus.UpdateFunction)
}
