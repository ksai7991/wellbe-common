package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCBillingStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBillingStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cBillingStatusPersistence.CreateCBillingStatus(&ctx, &model.CBillingStatus{
                                                        BillingStatusCd: 0,
                                                        LanguageCd: 1,
                                                        BillingStatusName: "dummy-BillingStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cBillingStatuss, _ := cBillingStatusPersistence.GetCBillingStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBillingStatuss[0].BillingStatusCd, 0)
    assert.Equal(t, cBillingStatuss[0].LanguageCd, 1)
    assert.Equal(t, cBillingStatuss[0].BillingStatusName, "dummy-BillingStatusName")
    assert.Equal(t, cBillingStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cBillingStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cBillingStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cBillingStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCBillingStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBillingStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cBillingStatusPersistence.CreateCBillingStatus(&ctx, &model.CBillingStatus{
                                                        BillingStatusCd: 0,
                                                        LanguageCd: 1,
                                                        BillingStatusName: "dummy-BillingStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cBillingStatusPersistence.UpdateCBillingStatus(&ctx, &model.CBillingStatus{
                                                        BillingStatusCd: 0,
                                                        LanguageCd: 1,
                                                        BillingStatusName: "dummy-BillingStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cBillingStatuss, _ := cBillingStatusPersistence.GetCBillingStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBillingStatuss[0].BillingStatusName, "dummy-BillingStatusName2")
    assert.Equal(t, cBillingStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cBillingStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cBillingStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cBillingStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
