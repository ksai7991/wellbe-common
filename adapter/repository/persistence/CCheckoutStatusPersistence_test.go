package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCCheckoutStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCheckoutStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cCheckoutStatusPersistence.CreateCCheckoutStatus(&ctx, &model.CCheckoutStatus{
                                                        CheckoutStatusCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutStatusName: "dummy-CheckoutStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cCheckoutStatuss, _ := cCheckoutStatusPersistence.GetCCheckoutStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCheckoutStatuss[0].CheckoutStatusCd, 0)
    assert.Equal(t, cCheckoutStatuss[0].LanguageCd, 1)
    assert.Equal(t, cCheckoutStatuss[0].CheckoutStatusName, "dummy-CheckoutStatusName")
    assert.Equal(t, cCheckoutStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cCheckoutStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cCheckoutStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cCheckoutStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCCheckoutStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCheckoutStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cCheckoutStatusPersistence.CreateCCheckoutStatus(&ctx, &model.CCheckoutStatus{
                                                        CheckoutStatusCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutStatusName: "dummy-CheckoutStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cCheckoutStatusPersistence.UpdateCCheckoutStatus(&ctx, &model.CCheckoutStatus{
                                                        CheckoutStatusCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutStatusName: "dummy-CheckoutStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cCheckoutStatuss, _ := cCheckoutStatusPersistence.GetCCheckoutStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCheckoutStatuss[0].CheckoutStatusName, "dummy-CheckoutStatusName2")
    assert.Equal(t, cCheckoutStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cCheckoutStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cCheckoutStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cCheckoutStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
