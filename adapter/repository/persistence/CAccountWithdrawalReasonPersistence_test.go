package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCAccountWithdrawalReasonCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cAccountWithdrawalReasonPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cAccountWithdrawalReasonPersistence.CreateCAccountWithdrawalReason(&ctx, &model.CAccountWithdrawalReason{
                                                        AccountWithdrawalReasonCd: 0,
                                                        LanguageCd: 1,
                                                        AccountWithdrawalReasonName: "dummy-AccountWithdrawalReasonName",
                                                        AccountWithdrawalReasonAbbreviation: "dummy-AccountWithdrawalReasonAbbreviation",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cAccountWithdrawalReasons, _ := cAccountWithdrawalReasonPersistence.GetCAccountWithdrawalReasonWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cAccountWithdrawalReasons[0].AccountWithdrawalReasonCd, 0)
    assert.Equal(t, cAccountWithdrawalReasons[0].LanguageCd, 1)
    assert.Equal(t, cAccountWithdrawalReasons[0].AccountWithdrawalReasonName, "dummy-AccountWithdrawalReasonName")
    assert.Equal(t, cAccountWithdrawalReasons[0].AccountWithdrawalReasonAbbreviation, "dummy-AccountWithdrawalReasonAbbreviation")
    assert.Equal(t, cAccountWithdrawalReasons[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cAccountWithdrawalReasons[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cAccountWithdrawalReasons[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cAccountWithdrawalReasons[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCAccountWithdrawalReasonUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cAccountWithdrawalReasonPersistence := NewPersistence(tr)
    ctx := context.Background()
    cAccountWithdrawalReasonPersistence.CreateCAccountWithdrawalReason(&ctx, &model.CAccountWithdrawalReason{
                                                        AccountWithdrawalReasonCd: 0,
                                                        LanguageCd: 1,
                                                        AccountWithdrawalReasonName: "dummy-AccountWithdrawalReasonName",
                                                        AccountWithdrawalReasonAbbreviation: "dummy-AccountWithdrawalReasonAbbreviation",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cAccountWithdrawalReasonPersistence.UpdateCAccountWithdrawalReason(&ctx, &model.CAccountWithdrawalReason{
                                                        AccountWithdrawalReasonCd: 0,
                                                        LanguageCd: 1,
                                                        AccountWithdrawalReasonName: "dummy-AccountWithdrawalReasonName2",
                                                        AccountWithdrawalReasonAbbreviation: "dummy-AccountWithdrawalReasonAbbreviation2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cAccountWithdrawalReasons, _ := cAccountWithdrawalReasonPersistence.GetCAccountWithdrawalReasonWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cAccountWithdrawalReasons[0].AccountWithdrawalReasonName, "dummy-AccountWithdrawalReasonName2")
    assert.Equal(t, cAccountWithdrawalReasons[0].AccountWithdrawalReasonAbbreviation, "dummy-AccountWithdrawalReasonAbbreviation2")
    assert.Equal(t, cAccountWithdrawalReasons[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cAccountWithdrawalReasons[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cAccountWithdrawalReasons[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cAccountWithdrawalReasons[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
