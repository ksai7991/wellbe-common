package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCShopImageFilterCategoryCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopImageFilterCategoryPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cShopImageFilterCategoryPersistence.CreateCShopImageFilterCategory(&ctx, &model.CShopImageFilterCategory{
                                                        ShopImageFilterCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ShopImageFilterCategoryName: "dummy-ShopImageFilterCategoryName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cShopImageFilterCategorys, _ := cShopImageFilterCategoryPersistence.GetCShopImageFilterCategoryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopImageFilterCategorys[0].ShopImageFilterCategoryCd, 0)
    assert.Equal(t, cShopImageFilterCategorys[0].LanguageCd, 1)
    assert.Equal(t, cShopImageFilterCategorys[0].ShopImageFilterCategoryName, "dummy-ShopImageFilterCategoryName")
    assert.Equal(t, cShopImageFilterCategorys[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cShopImageFilterCategorys[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cShopImageFilterCategorys[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cShopImageFilterCategorys[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCShopImageFilterCategoryUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopImageFilterCategoryPersistence := NewPersistence(tr)
    ctx := context.Background()
    cShopImageFilterCategoryPersistence.CreateCShopImageFilterCategory(&ctx, &model.CShopImageFilterCategory{
                                                        ShopImageFilterCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ShopImageFilterCategoryName: "dummy-ShopImageFilterCategoryName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cShopImageFilterCategoryPersistence.UpdateCShopImageFilterCategory(&ctx, &model.CShopImageFilterCategory{
                                                        ShopImageFilterCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ShopImageFilterCategoryName: "dummy-ShopImageFilterCategoryName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cShopImageFilterCategorys, _ := cShopImageFilterCategoryPersistence.GetCShopImageFilterCategoryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopImageFilterCategorys[0].ShopImageFilterCategoryName, "dummy-ShopImageFilterCategoryName2")
    assert.Equal(t, cShopImageFilterCategorys[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cShopImageFilterCategorys[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cShopImageFilterCategorys[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cShopImageFilterCategorys[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
