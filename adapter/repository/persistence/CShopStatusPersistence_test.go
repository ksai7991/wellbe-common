package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCShopStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cShopStatusPersistence.CreateCShopStatus(&ctx, &model.CShopStatus{
                                                        ShopStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ShopStatusName: "dummy-ShopStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cShopStatuss, _ := cShopStatusPersistence.GetCShopStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopStatuss[0].ShopStatusCd, 0)
    assert.Equal(t, cShopStatuss[0].LanguageCd, 1)
    assert.Equal(t, cShopStatuss[0].ShopStatusName, "dummy-ShopStatusName")
    assert.Equal(t, cShopStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cShopStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cShopStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cShopStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCShopStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cShopStatusPersistence.CreateCShopStatus(&ctx, &model.CShopStatus{
                                                        ShopStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ShopStatusName: "dummy-ShopStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cShopStatusPersistence.UpdateCShopStatus(&ctx, &model.CShopStatus{
                                                        ShopStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ShopStatusName: "dummy-ShopStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cShopStatuss, _ := cShopStatusPersistence.GetCShopStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopStatuss[0].ShopStatusName, "dummy-ShopStatusName2")
    assert.Equal(t, cShopStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cShopStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cShopStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cShopStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
