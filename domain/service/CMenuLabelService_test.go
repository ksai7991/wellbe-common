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

type cMenuLabelMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError)
    FakeUpdate func(*model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.CMenuLabel, *errordef.LogicError)
}

func (lr cMenuLabelMockRepository) CreateCMenuLabel(ctx *context.Context, cMenuLabel *model.CMenuLabel)  (*model.CMenuLabel, *errordef.LogicError) {
    return lr.FakeCreate(cMenuLabel)
}

func (lr cMenuLabelMockRepository) UpdateCMenuLabel(ctx *context.Context, cMenuLabel *model.CMenuLabel)  (*model.CMenuLabel, *errordef.LogicError) {
    return lr.FakeUpdate(cMenuLabel)
}

func (lr cMenuLabelMockRepository) DeleteCMenuLabel(ctx *context.Context, menuLabelCd int)  *errordef.LogicError {
    return lr.FakeDelete(menuLabelCd)
}

func (lr cMenuLabelMockRepository)GetCMenuLabelWithKey(ctx *context.Context, menuLabelCd int)  ([]*model.CMenuLabel, *errordef.LogicError) {
    return lr.FakeGet(menuLabelCd)
}


type cMenuLabelMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cMenuLabelMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCMenuLabel(t *testing.T) {
    ctx := context.Background()
    repository := &cMenuLabelMockRepository{
        FakeCreate: func(cMenuLabel *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError) {
            return cMenuLabel, nil
        },
    }
    numberUtil := &cMenuLabelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cMenuLabelService := NewCMenuLabelService(repository, numberUtil)
    in_cMenuLabel := new(model.CMenuLabel)
    in_cMenuLabel.MenuLabelCd = 0
    in_cMenuLabel.LanguageCd = 1
    in_cMenuLabel.MenuLabelName = "dummy-MenuLabelName"
    out_cMenuLabel, err := cMenuLabelService.CreateCMenuLabel(&ctx, in_cMenuLabel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cMenuLabel.MenuLabelCd)
    assert.Equal(t, 1, out_cMenuLabel.LanguageCd)
    assert.Equal(t, "dummy-MenuLabelName", out_cMenuLabel.MenuLabelName)
    assert.NotNil(t, out_cMenuLabel.CreateDatetime)
    assert.NotEqual(t, "", out_cMenuLabel.CreateDatetime)
    assert.Equal(t, "CreateCMenuLabel", out_cMenuLabel.CreateFunction)
    assert.NotNil(t, out_cMenuLabel.UpdateDatetime)
    assert.NotEqual(t, "", out_cMenuLabel.UpdateDatetime)
    assert.Equal(t, "CreateCMenuLabel", out_cMenuLabel.UpdateFunction)
}

func TestUpdateCMenuLabel(t *testing.T) {
    ctx := context.Background()
    repository := &cMenuLabelMockRepository{
        FakeUpdate: func(cMenuLabel *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError) {
            return cMenuLabel, nil
        },
        FakeGet: func(menuLabelCd int) ([]*model.CMenuLabel, *errordef.LogicError) {
            return []*model.CMenuLabel{&model.CMenuLabel{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cMenuLabelMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cMenuLabelService := NewCMenuLabelService(repository, numberUtil)
    in_cMenuLabel := new(model.CMenuLabel)
    in_cMenuLabel.MenuLabelCd = 0
    in_cMenuLabel.LanguageCd = 1
    in_cMenuLabel.MenuLabelName = "dummy-MenuLabelName"
    out_cMenuLabel, err := cMenuLabelService.UpdateCMenuLabel(&ctx, in_cMenuLabel)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cMenuLabel.MenuLabelCd)
    assert.Equal(t, 1, out_cMenuLabel.LanguageCd)
    assert.Equal(t, "dummy-MenuLabelName", out_cMenuLabel.MenuLabelName)
    assert.NotNil(t, out_cMenuLabel.CreateDatetime)
    assert.Equal(t, "", out_cMenuLabel.CreateDatetime)
    assert.Equal(t, "", out_cMenuLabel.CreateFunction)
    assert.NotNil(t, out_cMenuLabel.UpdateDatetime)
    assert.NotEqual(t, "", out_cMenuLabel.UpdateDatetime)
    assert.Equal(t, "UpdateCMenuLabel", out_cMenuLabel.UpdateFunction)
}
