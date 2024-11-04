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

type cAreaMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CArea) (*model.CArea, *errordef.LogicError)
    FakeUpdate func(*model.CArea) (*model.CArea, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CArea, *errordef.LogicError)
}

func (lr cAreaMockRepository) CreateCArea(ctx *context.Context, cArea *model.CArea)  (*model.CArea, *errordef.LogicError) {
    return lr.FakeCreate(cArea)
}

func (lr cAreaMockRepository) UpdateCArea(ctx *context.Context, cArea *model.CArea)  (*model.CArea, *errordef.LogicError) {
    return lr.FakeUpdate(cArea)
}

func (lr cAreaMockRepository) DeleteCArea(ctx *context.Context, languageCd int, areaCd int)  *errordef.LogicError {
    return lr.FakeDelete(languageCd, areaCd)
}

func (lr cAreaMockRepository)GetCAreaWithKey(ctx *context.Context, languageCd int, areaCd int)  ([]*model.CArea, *errordef.LogicError) {
    return lr.FakeGet(languageCd, areaCd)
}


type cAreaMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cAreaMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCArea(t *testing.T) {
    ctx := context.Background()
    repository := &cAreaMockRepository{
        FakeCreate: func(cArea *model.CArea) (*model.CArea, *errordef.LogicError) {
            return cArea, nil
        },
    }
    numberUtil := &cAreaMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cAreaService := NewCAreaService(repository, numberUtil)
    in_cArea := new(model.CArea)
    in_cArea.LanguageCd = 0
    in_cArea.AreaCd = 1
    in_cArea.StateCd = 2
    in_cArea.AreaName = "dummy-AreaName"
    in_cArea.SearchAreaNameSeo = "dummy-SearchAreaNameSeo"
    in_cArea.WestLongitude = 5.2
    in_cArea.EastLongitude = 6.2
    in_cArea.NorthLatitude = 7.2
    in_cArea.SouthLatitude = 8.2
    in_cArea.SummaryAreaFlg = true
    out_cArea, err := cAreaService.CreateCArea(&ctx, in_cArea)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cArea.LanguageCd)
    assert.Equal(t, 1, out_cArea.AreaCd)
    assert.Equal(t, 2, out_cArea.StateCd)
    assert.Equal(t, "dummy-AreaName", out_cArea.AreaName)
    assert.Equal(t, "dummy-SearchAreaNameSeo", out_cArea.SearchAreaNameSeo)
    assert.Equal(t, 5.2, out_cArea.WestLongitude)
    assert.Equal(t, 6.2, out_cArea.EastLongitude)
    assert.Equal(t, 7.2, out_cArea.NorthLatitude)
    assert.Equal(t, 8.2, out_cArea.SouthLatitude)
    assert.Equal(t, true, out_cArea.SummaryAreaFlg)
    assert.NotNil(t, out_cArea.CreateDatetime)
    assert.NotEqual(t, "", out_cArea.CreateDatetime)
    assert.Equal(t, "CreateCArea", out_cArea.CreateFunction)
    assert.NotNil(t, out_cArea.UpdateDatetime)
    assert.NotEqual(t, "", out_cArea.UpdateDatetime)
    assert.Equal(t, "CreateCArea", out_cArea.UpdateFunction)
}

func TestUpdateCArea(t *testing.T) {
    ctx := context.Background()
    repository := &cAreaMockRepository{
        FakeUpdate: func(cArea *model.CArea) (*model.CArea, *errordef.LogicError) {
            return cArea, nil
        },
        FakeGet: func(languageCd int, areaCd int) ([]*model.CArea, *errordef.LogicError) {
            return []*model.CArea{&model.CArea{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cAreaMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cAreaService := NewCAreaService(repository, numberUtil)
    in_cArea := new(model.CArea)
    in_cArea.LanguageCd = 0
    in_cArea.AreaCd = 1
    in_cArea.StateCd = 2
    in_cArea.AreaName = "dummy-AreaName"
    in_cArea.SearchAreaNameSeo = "dummy-SearchAreaNameSeo"
    in_cArea.WestLongitude = 5.2
    in_cArea.EastLongitude = 6.2
    in_cArea.NorthLatitude = 7.2
    in_cArea.SouthLatitude = 8.2
    in_cArea.SummaryAreaFlg = true
    out_cArea, err := cAreaService.UpdateCArea(&ctx, in_cArea)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cArea.LanguageCd)
    assert.Equal(t, 1, out_cArea.AreaCd)
    assert.Equal(t, 2, out_cArea.StateCd)
    assert.Equal(t, "dummy-AreaName", out_cArea.AreaName)
    assert.Equal(t, "dummy-SearchAreaNameSeo", out_cArea.SearchAreaNameSeo)
    assert.Equal(t, 5.2, out_cArea.WestLongitude)
    assert.Equal(t, 6.2, out_cArea.EastLongitude)
    assert.Equal(t, 7.2, out_cArea.NorthLatitude)
    assert.Equal(t, 8.2, out_cArea.SouthLatitude)
    assert.Equal(t, true, out_cArea.SummaryAreaFlg)
    assert.NotNil(t, out_cArea.CreateDatetime)
    assert.Equal(t, "", out_cArea.CreateDatetime)
    assert.Equal(t, "", out_cArea.CreateFunction)
    assert.NotNil(t, out_cArea.UpdateDatetime)
    assert.NotEqual(t, "", out_cArea.UpdateDatetime)
    assert.Equal(t, "UpdateCArea", out_cArea.UpdateFunction)
}
