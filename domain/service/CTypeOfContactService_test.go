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

type cTypeOfContactMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError)
    FakeUpdate func(*model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CTypeOfContact, *errordef.LogicError)
}

func (lr cTypeOfContactMockRepository) CreateCTypeOfContact(ctx *context.Context, cTypeOfContact *model.CTypeOfContact)  (*model.CTypeOfContact, *errordef.LogicError) {
    return lr.FakeCreate(cTypeOfContact)
}

func (lr cTypeOfContactMockRepository) UpdateCTypeOfContact(ctx *context.Context, cTypeOfContact *model.CTypeOfContact)  (*model.CTypeOfContact, *errordef.LogicError) {
    return lr.FakeUpdate(cTypeOfContact)
}

func (lr cTypeOfContactMockRepository) DeleteCTypeOfContact(ctx *context.Context, typeOfContactCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(typeOfContactCd, languageCd)
}

func (lr cTypeOfContactMockRepository)GetCTypeOfContactWithKey(ctx *context.Context, typeOfContactCd int, languageCd int)  ([]*model.CTypeOfContact, *errordef.LogicError) {
    return lr.FakeGet(typeOfContactCd, languageCd)
}


type cTypeOfContactMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cTypeOfContactMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCTypeOfContact(t *testing.T) {
    ctx := context.Background()
    repository := &cTypeOfContactMockRepository{
        FakeCreate: func(cTypeOfContact *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError) {
            return cTypeOfContact, nil
        },
    }
    numberUtil := &cTypeOfContactMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cTypeOfContactService := NewCTypeOfContactService(repository, numberUtil)
    in_cTypeOfContact := new(model.CTypeOfContact)
    in_cTypeOfContact.TypeOfContactCd = 0
    in_cTypeOfContact.LanguageCd = 1
    in_cTypeOfContact.TypeOfContactName = "dummy-TypeOfContactName"
    out_cTypeOfContact, err := cTypeOfContactService.CreateCTypeOfContact(&ctx, in_cTypeOfContact)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cTypeOfContact.TypeOfContactCd)
    assert.Equal(t, 1, out_cTypeOfContact.LanguageCd)
    assert.Equal(t, "dummy-TypeOfContactName", out_cTypeOfContact.TypeOfContactName)
    assert.NotNil(t, out_cTypeOfContact.CreateDatetime)
    assert.NotEqual(t, "", out_cTypeOfContact.CreateDatetime)
    assert.Equal(t, "CreateCTypeOfContact", out_cTypeOfContact.CreateFunction)
    assert.NotNil(t, out_cTypeOfContact.UpdateDatetime)
    assert.NotEqual(t, "", out_cTypeOfContact.UpdateDatetime)
    assert.Equal(t, "CreateCTypeOfContact", out_cTypeOfContact.UpdateFunction)
}

func TestUpdateCTypeOfContact(t *testing.T) {
    ctx := context.Background()
    repository := &cTypeOfContactMockRepository{
        FakeUpdate: func(cTypeOfContact *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError) {
            return cTypeOfContact, nil
        },
        FakeGet: func(typeOfContactCd int, languageCd int) ([]*model.CTypeOfContact, *errordef.LogicError) {
            return []*model.CTypeOfContact{&model.CTypeOfContact{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cTypeOfContactMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cTypeOfContactService := NewCTypeOfContactService(repository, numberUtil)
    in_cTypeOfContact := new(model.CTypeOfContact)
    in_cTypeOfContact.TypeOfContactCd = 0
    in_cTypeOfContact.LanguageCd = 1
    in_cTypeOfContact.TypeOfContactName = "dummy-TypeOfContactName"
    out_cTypeOfContact, err := cTypeOfContactService.UpdateCTypeOfContact(&ctx, in_cTypeOfContact)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cTypeOfContact.TypeOfContactCd)
    assert.Equal(t, 1, out_cTypeOfContact.LanguageCd)
    assert.Equal(t, "dummy-TypeOfContactName", out_cTypeOfContact.TypeOfContactName)
    assert.NotNil(t, out_cTypeOfContact.CreateDatetime)
    assert.Equal(t, "", out_cTypeOfContact.CreateDatetime)
    assert.Equal(t, "", out_cTypeOfContact.CreateFunction)
    assert.NotNil(t, out_cTypeOfContact.UpdateDatetime)
    assert.NotEqual(t, "", out_cTypeOfContact.UpdateDatetime)
    assert.Equal(t, "UpdateCTypeOfContact", out_cTypeOfContact.UpdateFunction)
}
