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

type cLanguageMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CLanguage) (*model.CLanguage, *errordef.LogicError)
    FakeUpdate func(*model.CLanguage) (*model.CLanguage, *errordef.LogicError)
    FakeDelete func(int) *errordef.LogicError
    FakeGet func(int) ([]*model.CLanguage, *errordef.LogicError)
}

func (lr cLanguageMockRepository) CreateCLanguage(ctx *context.Context, cLanguage *model.CLanguage)  (*model.CLanguage, *errordef.LogicError) {
    return lr.FakeCreate(cLanguage)
}

func (lr cLanguageMockRepository) UpdateCLanguage(ctx *context.Context, cLanguage *model.CLanguage)  (*model.CLanguage, *errordef.LogicError) {
    return lr.FakeUpdate(cLanguage)
}

func (lr cLanguageMockRepository) DeleteCLanguage(ctx *context.Context, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(languageCd)
}

func (lr cLanguageMockRepository)GetCLanguageWithKey(ctx *context.Context, languageCd int)  ([]*model.CLanguage, *errordef.LogicError) {
    return lr.FakeGet(languageCd)
}


type cLanguageMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cLanguageMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCLanguage(t *testing.T) {
    ctx := context.Background()
    repository := &cLanguageMockRepository{
        FakeCreate: func(cLanguage *model.CLanguage) (*model.CLanguage, *errordef.LogicError) {
            return cLanguage, nil
        },
    }
    numberUtil := &cLanguageMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cLanguageService := NewCLanguageService(repository, numberUtil)
    in_cLanguage := new(model.CLanguage)
    in_cLanguage.LanguageCd = 0
    in_cLanguage.LanguageCharCd = "XX"
    in_cLanguage.LanguageName = "dummy-LanguageName"
    in_cLanguage.SortNumber = 3
    out_cLanguage, err := cLanguageService.CreateCLanguage(&ctx, in_cLanguage)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cLanguage.LanguageCd)
    assert.Equal(t, "XX", out_cLanguage.LanguageCharCd)
    assert.Equal(t, "dummy-LanguageName", out_cLanguage.LanguageName)
    assert.Equal(t, 3, out_cLanguage.SortNumber)
    assert.NotNil(t, out_cLanguage.CreateDatetime)
    assert.NotEqual(t, "", out_cLanguage.CreateDatetime)
    assert.Equal(t, "CreateCLanguage", out_cLanguage.CreateFunction)
    assert.NotNil(t, out_cLanguage.UpdateDatetime)
    assert.NotEqual(t, "", out_cLanguage.UpdateDatetime)
    assert.Equal(t, "CreateCLanguage", out_cLanguage.UpdateFunction)
}

func TestUpdateCLanguage(t *testing.T) {
    ctx := context.Background()
    repository := &cLanguageMockRepository{
        FakeUpdate: func(cLanguage *model.CLanguage) (*model.CLanguage, *errordef.LogicError) {
            return cLanguage, nil
        },
        FakeGet: func(languageCd int) ([]*model.CLanguage, *errordef.LogicError) {
            return []*model.CLanguage{&model.CLanguage{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cLanguageMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cLanguageService := NewCLanguageService(repository, numberUtil)
    in_cLanguage := new(model.CLanguage)
    in_cLanguage.LanguageCd = 0
    in_cLanguage.LanguageCharCd = "XX"
    in_cLanguage.LanguageName = "dummy-LanguageName"
    in_cLanguage.SortNumber = 3
    out_cLanguage, err := cLanguageService.UpdateCLanguage(&ctx, in_cLanguage)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cLanguage.LanguageCd)
    assert.Equal(t, "XX", out_cLanguage.LanguageCharCd)
    assert.Equal(t, "dummy-LanguageName", out_cLanguage.LanguageName)
    assert.Equal(t, 3, out_cLanguage.SortNumber)
    assert.NotNil(t, out_cLanguage.CreateDatetime)
    assert.Equal(t, "", out_cLanguage.CreateDatetime)
    assert.Equal(t, "", out_cLanguage.CreateFunction)
    assert.NotNil(t, out_cLanguage.UpdateDatetime)
    assert.NotEqual(t, "", out_cLanguage.UpdateDatetime)
    assert.Equal(t, "UpdateCLanguage", out_cLanguage.UpdateFunction)
}
