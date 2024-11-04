package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCurrencyForPaymentCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    currencyForPaymentPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := currencyForPaymentPersistence.CreateCurrencyForPayment(&ctx, &model.CurrencyForPayment{
                                                        CurrencyCd: 0,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    currencyForPayments, _ := currencyForPaymentPersistence.GetCurrencyForPaymentWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, currencyForPayments[0].CurrencyCd, 0)
    assert.Equal(t, currencyForPayments[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, currencyForPayments[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, currencyForPayments[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, currencyForPayments[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCurrencyForPaymentUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    currencyForPaymentPersistence := NewPersistence(tr)
    ctx := context.Background()
    currencyForPaymentPersistence.CreateCurrencyForPayment(&ctx, &model.CurrencyForPayment{
                                                        CurrencyCd: 0,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := currencyForPaymentPersistence.UpdateCurrencyForPayment(&ctx, &model.CurrencyForPayment{
                                                        CurrencyCd: 0,
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    currencyForPayments, _ := currencyForPaymentPersistence.GetCurrencyForPaymentWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, currencyForPayments[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, currencyForPayments[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, currencyForPayments[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, currencyForPayments[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
