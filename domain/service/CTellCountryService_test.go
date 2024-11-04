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

type cTellCountryMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CTellCountry) (*model.CTellCountry, *errordef.LogicError)
    FakeUpdate func(*model.CTellCountry) (*model.CTellCountry, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CTellCountry, *errordef.LogicError)
}

func (lr cTellCountryMockRepository) CreateCTellCountry(ctx *context.Context, cTellCountry *model.CTellCountry)  (*model.CTellCountry, *errordef.LogicError) {
    return lr.FakeCreate(cTellCountry)
}

func (lr cTellCountryMockRepository) UpdateCTellCountry(ctx *context.Context, cTellCountry *model.CTellCountry)  (*model.CTellCountry, *errordef.LogicError) {
    return lr.FakeUpdate(cTellCountry)
}

func (lr cTellCountryMockRepository) DeleteCTellCountry(ctx *context.Context, languageCd int, tellCountryCd int)  *errordef.LogicError {
    return lr.FakeDelete(languageCd, tellCountryCd)
}

func (lr cTellCountryMockRepository)GetCTellCountryWithKey(ctx *context.Context, languageCd int, tellCountryCd int)  ([]*model.CTellCountry, *errordef.LogicError) {
    return lr.FakeGet(languageCd, tellCountryCd)
}


type cTellCountryMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cTellCountryMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCTellCountry(t *testing.T) {
    ctx := context.Background()
    repository := &cTellCountryMockRepository{
        FakeCreate: func(cTellCountry *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError) {
            return cTellCountry, nil
        },
    }
    numberUtil := &cTellCountryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cTellCountryService := NewCTellCountryService(repository, numberUtil)
    in_cTellCountry := new(model.CTellCountry)
    in_cTellCountry.LanguageCd = 0
    in_cTellCountry.TellCountryCd = 1
    in_cTellCountry.CountryName = "dummy-CountryName"
    in_cTellCountry.CountryNo = "dummy-CountryNo"
    out_cTellCountry, err := cTellCountryService.CreateCTellCountry(&ctx, in_cTellCountry)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cTellCountry.LanguageCd)
    assert.Equal(t, 1, out_cTellCountry.TellCountryCd)
    assert.Equal(t, "dummy-CountryName", out_cTellCountry.CountryName)
    assert.Equal(t, "dummy-CountryNo", out_cTellCountry.CountryNo)
    assert.NotNil(t, out_cTellCountry.CreateDatetime)
    assert.NotEqual(t, "", out_cTellCountry.CreateDatetime)
    assert.Equal(t, "CreateCTellCountry", out_cTellCountry.CreateFunction)
    assert.NotNil(t, out_cTellCountry.UpdateDatetime)
    assert.NotEqual(t, "", out_cTellCountry.UpdateDatetime)
    assert.Equal(t, "CreateCTellCountry", out_cTellCountry.UpdateFunction)
}

func TestUpdateCTellCountry(t *testing.T) {
    ctx := context.Background()
    repository := &cTellCountryMockRepository{
        FakeUpdate: func(cTellCountry *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError) {
            return cTellCountry, nil
        },
        FakeGet: func(languageCd int, tellCountryCd int) ([]*model.CTellCountry, *errordef.LogicError) {
            return []*model.CTellCountry{&model.CTellCountry{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cTellCountryMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cTellCountryService := NewCTellCountryService(repository, numberUtil)
    in_cTellCountry := new(model.CTellCountry)
    in_cTellCountry.LanguageCd = 0
    in_cTellCountry.TellCountryCd = 1
    in_cTellCountry.CountryName = "dummy-CountryName"
    in_cTellCountry.CountryNo = "dummy-CountryNo"
    out_cTellCountry, err := cTellCountryService.UpdateCTellCountry(&ctx, in_cTellCountry)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cTellCountry.LanguageCd)
    assert.Equal(t, 1, out_cTellCountry.TellCountryCd)
    assert.Equal(t, "dummy-CountryName", out_cTellCountry.CountryName)
    assert.Equal(t, "dummy-CountryNo", out_cTellCountry.CountryNo)
    assert.NotNil(t, out_cTellCountry.CreateDatetime)
    assert.Equal(t, "", out_cTellCountry.CreateDatetime)
    assert.Equal(t, "", out_cTellCountry.CreateFunction)
    assert.NotNil(t, out_cTellCountry.UpdateDatetime)
    assert.NotEqual(t, "", out_cTellCountry.UpdateDatetime)
    assert.Equal(t, "UpdateCTellCountry", out_cTellCountry.UpdateFunction)
}
