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

type cCountryMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CCountry) (*model.CCountry, *errordef.LogicError)
    FakeUpdate func(*model.CCountry) (*model.CCountry, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CCountry, *errordef.LogicError)
}

func (lr cCountryMockRepository) CreateCCountry(ctx *context.Context, cCountry *model.CCountry)  (*model.CCountry, *errordef.LogicError) {
    return lr.FakeCreate(cCountry)
}

func (lr cCountryMockRepository) UpdateCCountry(ctx *context.Context, cCountry *model.CCountry)  (*model.CCountry, *errordef.LogicError) {
    return lr.FakeUpdate(cCountry)
}

func (lr cCountryMockRepository) DeleteCCountry(ctx *context.Context, countryCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(countryCd, languageCd)
}

func (lr cCountryMockRepository)GetCCountryWithKey(ctx *context.Context, countryCd int, languageCd int)  ([]*model.CCountry, *errordef.LogicError) {
    return lr.FakeGet(countryCd, languageCd)
}


type cCountryMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cCountryMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCCountry(t *testing.T) {
    ctx := context.Background()
    repository := &cCountryMockRepository{
        FakeCreate: func(cCountry *model.CCountry) (*model.CCountry, *errordef.LogicError) {
            return cCountry, nil
        },
    }
    numberUtil := &cCountryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCountryService := NewCCountryService(repository, numberUtil)
    in_cCountry := new(model.CCountry)
    in_cCountry.CountryCd = 0
    in_cCountry.LanguageCd = 1
    in_cCountry.CountryName = "dummy-CountryName"
    in_cCountry.CountryCdIso = "XX"
    out_cCountry, err := cCountryService.CreateCCountry(&ctx, in_cCountry)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCountry.CountryCd)
    assert.Equal(t, 1, out_cCountry.LanguageCd)
    assert.Equal(t, "dummy-CountryName", out_cCountry.CountryName)
    assert.Equal(t, "XX", out_cCountry.CountryCdIso)
    assert.NotNil(t, out_cCountry.CreateDatetime)
    assert.NotEqual(t, "", out_cCountry.CreateDatetime)
    assert.Equal(t, "CreateCCountry", out_cCountry.CreateFunction)
    assert.NotNil(t, out_cCountry.UpdateDatetime)
    assert.NotEqual(t, "", out_cCountry.UpdateDatetime)
    assert.Equal(t, "CreateCCountry", out_cCountry.UpdateFunction)
}

func TestUpdateCCountry(t *testing.T) {
    ctx := context.Background()
    repository := &cCountryMockRepository{
        FakeUpdate: func(cCountry *model.CCountry) (*model.CCountry, *errordef.LogicError) {
            return cCountry, nil
        },
        FakeGet: func(countryCd int, languageCd int) ([]*model.CCountry, *errordef.LogicError) {
            return []*model.CCountry{&model.CCountry{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cCountryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCountryService := NewCCountryService(repository, numberUtil)
    in_cCountry := new(model.CCountry)
    in_cCountry.CountryCd = 0
    in_cCountry.LanguageCd = 1
    in_cCountry.CountryName = "dummy-CountryName"
    in_cCountry.CountryCdIso = "XX"
    out_cCountry, err := cCountryService.UpdateCCountry(&ctx, in_cCountry)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCountry.CountryCd)
    assert.Equal(t, 1, out_cCountry.LanguageCd)
    assert.Equal(t, "dummy-CountryName", out_cCountry.CountryName)
    assert.Equal(t, "XX", out_cCountry.CountryCdIso)
    assert.NotNil(t, out_cCountry.CreateDatetime)
    assert.Equal(t, "", out_cCountry.CreateDatetime)
    assert.Equal(t, "", out_cCountry.CreateFunction)
    assert.NotNil(t, out_cCountry.UpdateDatetime)
    assert.NotEqual(t, "", out_cCountry.UpdateDatetime)
    assert.Equal(t, "UpdateCCountry", out_cCountry.UpdateFunction)
}
