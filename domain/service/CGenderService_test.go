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

type cGenderMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CGender) (*model.CGender, *errordef.LogicError)
    FakeUpdate func(*model.CGender) (*model.CGender, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CGender, *errordef.LogicError)
}

func (lr cGenderMockRepository) CreateCGender(ctx *context.Context, cGender *model.CGender)  (*model.CGender, *errordef.LogicError) {
    return lr.FakeCreate(cGender)
}

func (lr cGenderMockRepository) UpdateCGender(ctx *context.Context, cGender *model.CGender)  (*model.CGender, *errordef.LogicError) {
    return lr.FakeUpdate(cGender)
}

func (lr cGenderMockRepository) DeleteCGender(ctx *context.Context, genderCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(genderCd, languageCd)
}

func (lr cGenderMockRepository)GetCGenderWithKey(ctx *context.Context, genderCd int, languageCd int)  ([]*model.CGender, *errordef.LogicError) {
    return lr.FakeGet(genderCd, languageCd)
}


type cGenderMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cGenderMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCGender(t *testing.T) {
    ctx := context.Background()
    repository := &cGenderMockRepository{
        FakeCreate: func(cGender *model.CGender) (*model.CGender, *errordef.LogicError) {
            return cGender, nil
        },
    }
    numberUtil := &cGenderMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cGenderService := NewCGenderService(repository, numberUtil)
    in_cGender := new(model.CGender)
    in_cGender.GenderCd = 0
    in_cGender.LanguageCd = 1
    in_cGender.GenderName = "dummy-GenderName"
    in_cGender.GenderAbbreviation = "dummy-GenderAbbreviation"
    out_cGender, err := cGenderService.CreateCGender(&ctx, in_cGender)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cGender.GenderCd)
    assert.Equal(t, 1, out_cGender.LanguageCd)
    assert.Equal(t, "dummy-GenderName", out_cGender.GenderName)
    assert.Equal(t, "dummy-GenderAbbreviation", out_cGender.GenderAbbreviation)
    assert.NotNil(t, out_cGender.CreateDatetime)
    assert.NotEqual(t, "", out_cGender.CreateDatetime)
    assert.Equal(t, "CreateCGender", out_cGender.CreateFunction)
    assert.NotNil(t, out_cGender.UpdateDatetime)
    assert.NotEqual(t, "", out_cGender.UpdateDatetime)
    assert.Equal(t, "CreateCGender", out_cGender.UpdateFunction)
}

func TestUpdateCGender(t *testing.T) {
    ctx := context.Background()
    repository := &cGenderMockRepository{
        FakeUpdate: func(cGender *model.CGender) (*model.CGender, *errordef.LogicError) {
            return cGender, nil
        },
        FakeGet: func(genderCd int, languageCd int) ([]*model.CGender, *errordef.LogicError) {
            return []*model.CGender{&model.CGender{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cGenderMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cGenderService := NewCGenderService(repository, numberUtil)
    in_cGender := new(model.CGender)
    in_cGender.GenderCd = 0
    in_cGender.LanguageCd = 1
    in_cGender.GenderName = "dummy-GenderName"
    in_cGender.GenderAbbreviation = "dummy-GenderAbbreviation"
    out_cGender, err := cGenderService.UpdateCGender(&ctx, in_cGender)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cGender.GenderCd)
    assert.Equal(t, 1, out_cGender.LanguageCd)
    assert.Equal(t, "dummy-GenderName", out_cGender.GenderName)
    assert.Equal(t, "dummy-GenderAbbreviation", out_cGender.GenderAbbreviation)
    assert.NotNil(t, out_cGender.CreateDatetime)
    assert.Equal(t, "", out_cGender.CreateDatetime)
    assert.Equal(t, "", out_cGender.CreateFunction)
    assert.NotNil(t, out_cGender.UpdateDatetime)
    assert.NotEqual(t, "", out_cGender.UpdateDatetime)
    assert.Equal(t, "UpdateCGender", out_cGender.UpdateFunction)
}
