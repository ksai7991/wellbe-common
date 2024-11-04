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

type cServiceMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CService) (*model.CService, *errordef.LogicError)
    FakeUpdate func(*model.CService) (*model.CService, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CService, *errordef.LogicError)
}

func (lr cServiceMockRepository) CreateCService(ctx *context.Context, cService *model.CService)  (*model.CService, *errordef.LogicError) {
    return lr.FakeCreate(cService)
}

func (lr cServiceMockRepository) UpdateCService(ctx *context.Context, cService *model.CService)  (*model.CService, *errordef.LogicError) {
    return lr.FakeUpdate(cService)
}

func (lr cServiceMockRepository) DeleteCService(ctx *context.Context, serviceCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(serviceCd, languageCd)
}

func (lr cServiceMockRepository)GetCServiceWithKey(ctx *context.Context, serviceCd int, languageCd int)  ([]*model.CService, *errordef.LogicError) {
    return lr.FakeGet(serviceCd, languageCd)
}


type cServiceMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cServiceMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCService(t *testing.T) {
    ctx := context.Background()
    repository := &cServiceMockRepository{
        FakeCreate: func(cService *model.CService) (*model.CService, *errordef.LogicError) {
            return cService, nil
        },
    }
    numberUtil := &cServiceMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cServiceService := NewCServiceService(repository, numberUtil)
    in_cService := new(model.CService)
    in_cService.ServiceCd = 0
    in_cService.LanguageCd = 1
    in_cService.ServiceName = "dummy-ServiceName"
    out_cService, err := cServiceService.CreateCService(&ctx, in_cService)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cService.ServiceCd)
    assert.Equal(t, 1, out_cService.LanguageCd)
    assert.Equal(t, "dummy-ServiceName", out_cService.ServiceName)
    assert.NotNil(t, out_cService.CreateDatetime)
    assert.NotEqual(t, "", out_cService.CreateDatetime)
    assert.Equal(t, "CreateCService", out_cService.CreateFunction)
    assert.NotNil(t, out_cService.UpdateDatetime)
    assert.NotEqual(t, "", out_cService.UpdateDatetime)
    assert.Equal(t, "CreateCService", out_cService.UpdateFunction)
}

func TestUpdateCService(t *testing.T) {
    ctx := context.Background()
    repository := &cServiceMockRepository{
        FakeUpdate: func(cService *model.CService) (*model.CService, *errordef.LogicError) {
            return cService, nil
        },
        FakeGet: func(serviceCd int, languageCd int) ([]*model.CService, *errordef.LogicError) {
            return []*model.CService{&model.CService{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cServiceMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cServiceService := NewCServiceService(repository, numberUtil)
    in_cService := new(model.CService)
    in_cService.ServiceCd = 0
    in_cService.LanguageCd = 1
    in_cService.ServiceName = "dummy-ServiceName"
    out_cService, err := cServiceService.UpdateCService(&ctx, in_cService)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cService.ServiceCd)
    assert.Equal(t, 1, out_cService.LanguageCd)
    assert.Equal(t, "dummy-ServiceName", out_cService.ServiceName)
    assert.NotNil(t, out_cService.CreateDatetime)
    assert.Equal(t, "", out_cService.CreateDatetime)
    assert.Equal(t, "", out_cService.CreateFunction)
    assert.NotNil(t, out_cService.UpdateDatetime)
    assert.NotEqual(t, "", out_cService.UpdateDatetime)
    assert.Equal(t, "UpdateCService", out_cService.UpdateFunction)
}
