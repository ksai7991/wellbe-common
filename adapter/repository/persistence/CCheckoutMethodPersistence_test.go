package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCCheckoutMethodCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCheckoutMethodPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cCheckoutMethodPersistence.CreateCCheckoutMethod(&ctx, &model.CCheckoutMethod{
                                                        CheckoutMethodCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutMethodName: "dummy-CheckoutMethodName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cCheckoutMethods, _ := cCheckoutMethodPersistence.GetCCheckoutMethodWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCheckoutMethods[0].CheckoutMethodCd, 0)
    assert.Equal(t, cCheckoutMethods[0].LanguageCd, 1)
    assert.Equal(t, cCheckoutMethods[0].CheckoutMethodName, "dummy-CheckoutMethodName")
    assert.Equal(t, cCheckoutMethods[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cCheckoutMethods[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cCheckoutMethods[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cCheckoutMethods[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCCheckoutMethodUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCheckoutMethodPersistence := NewPersistence(tr)
    ctx := context.Background()
    cCheckoutMethodPersistence.CreateCCheckoutMethod(&ctx, &model.CCheckoutMethod{
                                                        CheckoutMethodCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutMethodName: "dummy-CheckoutMethodName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cCheckoutMethodPersistence.UpdateCCheckoutMethod(&ctx, &model.CCheckoutMethod{
                                                        CheckoutMethodCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutMethodName: "dummy-CheckoutMethodName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cCheckoutMethods, _ := cCheckoutMethodPersistence.GetCCheckoutMethodWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCheckoutMethods[0].CheckoutMethodName, "dummy-CheckoutMethodName2")
    assert.Equal(t, cCheckoutMethods[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cCheckoutMethods[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cCheckoutMethods[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cCheckoutMethods[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
