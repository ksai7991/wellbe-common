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

type cAccountWithdrawalReasonMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError)
    FakeUpdate func(*model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError)
}

func (lr cAccountWithdrawalReasonMockRepository) CreateCAccountWithdrawalReason(ctx *context.Context, cAccountWithdrawalReason *model.CAccountWithdrawalReason)  (*model.CAccountWithdrawalReason, *errordef.LogicError) {
    return lr.FakeCreate(cAccountWithdrawalReason)
}

func (lr cAccountWithdrawalReasonMockRepository) UpdateCAccountWithdrawalReason(ctx *context.Context, cAccountWithdrawalReason *model.CAccountWithdrawalReason)  (*model.CAccountWithdrawalReason, *errordef.LogicError) {
    return lr.FakeUpdate(cAccountWithdrawalReason)
}

func (lr cAccountWithdrawalReasonMockRepository) DeleteCAccountWithdrawalReason(ctx *context.Context, accountWithdrawalReasonCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(accountWithdrawalReasonCd, languageCd)
}

func (lr cAccountWithdrawalReasonMockRepository)GetCAccountWithdrawalReasonWithKey(ctx *context.Context, accountWithdrawalReasonCd int, languageCd int)  ([]*model.CAccountWithdrawalReason, *errordef.LogicError) {
    return lr.FakeGet(accountWithdrawalReasonCd, languageCd)
}


type cAccountWithdrawalReasonMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cAccountWithdrawalReasonMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCAccountWithdrawalReason(t *testing.T) {
    ctx := context.Background()
    repository := &cAccountWithdrawalReasonMockRepository{
        FakeCreate: func(cAccountWithdrawalReason *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError) {
            return cAccountWithdrawalReason, nil
        },
    }
    numberUtil := &cAccountWithdrawalReasonMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cAccountWithdrawalReasonService := NewCAccountWithdrawalReasonService(repository, numberUtil)
    in_cAccountWithdrawalReason := new(model.CAccountWithdrawalReason)
    in_cAccountWithdrawalReason.AccountWithdrawalReasonCd = 0
    in_cAccountWithdrawalReason.LanguageCd = 1
    in_cAccountWithdrawalReason.AccountWithdrawalReasonName = "dummy-AccountWithdrawalReasonName"
    in_cAccountWithdrawalReason.AccountWithdrawalReasonAbbreviation = "dummy-AccountWithdrawalReasonAbbreviation"
    out_cAccountWithdrawalReason, err := cAccountWithdrawalReasonService.CreateCAccountWithdrawalReason(&ctx, in_cAccountWithdrawalReason)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cAccountWithdrawalReason.AccountWithdrawalReasonCd)
    assert.Equal(t, 1, out_cAccountWithdrawalReason.LanguageCd)
    assert.Equal(t, "dummy-AccountWithdrawalReasonName", out_cAccountWithdrawalReason.AccountWithdrawalReasonName)
    assert.Equal(t, "dummy-AccountWithdrawalReasonAbbreviation", out_cAccountWithdrawalReason.AccountWithdrawalReasonAbbreviation)
    assert.NotNil(t, out_cAccountWithdrawalReason.CreateDatetime)
    assert.NotEqual(t, "", out_cAccountWithdrawalReason.CreateDatetime)
    assert.Equal(t, "CreateCAccountWithdrawalReason", out_cAccountWithdrawalReason.CreateFunction)
    assert.NotNil(t, out_cAccountWithdrawalReason.UpdateDatetime)
    assert.NotEqual(t, "", out_cAccountWithdrawalReason.UpdateDatetime)
    assert.Equal(t, "CreateCAccountWithdrawalReason", out_cAccountWithdrawalReason.UpdateFunction)
}

func TestUpdateCAccountWithdrawalReason(t *testing.T) {
    ctx := context.Background()
    repository := &cAccountWithdrawalReasonMockRepository{
        FakeUpdate: func(cAccountWithdrawalReason *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError) {
            return cAccountWithdrawalReason, nil
        },
        FakeGet: func(accountWithdrawalReasonCd int, languageCd int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError) {
            return []*model.CAccountWithdrawalReason{&model.CAccountWithdrawalReason{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cAccountWithdrawalReasonMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cAccountWithdrawalReasonService := NewCAccountWithdrawalReasonService(repository, numberUtil)
    in_cAccountWithdrawalReason := new(model.CAccountWithdrawalReason)
    in_cAccountWithdrawalReason.AccountWithdrawalReasonCd = 0
    in_cAccountWithdrawalReason.LanguageCd = 1
    in_cAccountWithdrawalReason.AccountWithdrawalReasonName = "dummy-AccountWithdrawalReasonName"
    in_cAccountWithdrawalReason.AccountWithdrawalReasonAbbreviation = "dummy-AccountWithdrawalReasonAbbreviation"
    out_cAccountWithdrawalReason, err := cAccountWithdrawalReasonService.UpdateCAccountWithdrawalReason(&ctx, in_cAccountWithdrawalReason)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cAccountWithdrawalReason.AccountWithdrawalReasonCd)
    assert.Equal(t, 1, out_cAccountWithdrawalReason.LanguageCd)
    assert.Equal(t, "dummy-AccountWithdrawalReasonName", out_cAccountWithdrawalReason.AccountWithdrawalReasonName)
    assert.Equal(t, "dummy-AccountWithdrawalReasonAbbreviation", out_cAccountWithdrawalReason.AccountWithdrawalReasonAbbreviation)
    assert.NotNil(t, out_cAccountWithdrawalReason.CreateDatetime)
    assert.Equal(t, "", out_cAccountWithdrawalReason.CreateDatetime)
    assert.Equal(t, "", out_cAccountWithdrawalReason.CreateFunction)
    assert.NotNil(t, out_cAccountWithdrawalReason.UpdateDatetime)
    assert.NotEqual(t, "", out_cAccountWithdrawalReason.UpdateDatetime)
    assert.Equal(t, "UpdateCAccountWithdrawalReason", out_cAccountWithdrawalReason.UpdateFunction)
}
