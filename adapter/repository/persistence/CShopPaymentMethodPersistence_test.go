package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCShopPaymentMethodCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopPaymentMethodPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cShopPaymentMethodPersistence.CreateCShopPaymentMethod(&ctx, &model.CShopPaymentMethod{
                                                        ShopPaymentMethodCd: 0,
                                                        LanguageCd: 1,
                                                        ShopPaymentName: "dummy-ShopPaymentName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cShopPaymentMethods, _ := cShopPaymentMethodPersistence.GetCShopPaymentMethodWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopPaymentMethods[0].ShopPaymentMethodCd, 0)
    assert.Equal(t, cShopPaymentMethods[0].LanguageCd, 1)
    assert.Equal(t, cShopPaymentMethods[0].ShopPaymentName, "dummy-ShopPaymentName")
    assert.Equal(t, cShopPaymentMethods[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cShopPaymentMethods[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cShopPaymentMethods[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cShopPaymentMethods[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCShopPaymentMethodUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopPaymentMethodPersistence := NewPersistence(tr)
    ctx := context.Background()
    cShopPaymentMethodPersistence.CreateCShopPaymentMethod(&ctx, &model.CShopPaymentMethod{
                                                        ShopPaymentMethodCd: 0,
                                                        LanguageCd: 1,
                                                        ShopPaymentName: "dummy-ShopPaymentName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cShopPaymentMethodPersistence.UpdateCShopPaymentMethod(&ctx, &model.CShopPaymentMethod{
                                                        ShopPaymentMethodCd: 0,
                                                        LanguageCd: 11,
                                                        ShopPaymentName: "dummy-ShopPaymentName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cShopPaymentMethods, _ := cShopPaymentMethodPersistence.GetCShopPaymentMethodWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopPaymentMethods[0].LanguageCd, 11)
    assert.Equal(t, cShopPaymentMethods[0].ShopPaymentName, "dummy-ShopPaymentName2")
    assert.Equal(t, cShopPaymentMethods[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cShopPaymentMethods[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cShopPaymentMethods[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cShopPaymentMethods[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
