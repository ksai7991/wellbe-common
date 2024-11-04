package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCShopContractPlanItemCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopContractPlanItemPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cShopContractPlanItemPersistence.CreateCShopContractPlanItem(&ctx, &model.CShopContractPlanItem{
                                                        ShopContractPlanItemCd: 0,
                                                        LanguageCd: 1,
                                                        ShopContractPlanName: "dummy-ShopContractPlanName",
                                                        Unit: "dummy-Unit",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cShopContractPlanItems, _ := cShopContractPlanItemPersistence.GetCShopContractPlanItemWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopContractPlanItems[0].ShopContractPlanItemCd, 0)
    assert.Equal(t, cShopContractPlanItems[0].LanguageCd, 1)
    assert.Equal(t, cShopContractPlanItems[0].ShopContractPlanName, "dummy-ShopContractPlanName")
    assert.Equal(t, cShopContractPlanItems[0].Unit, "dummy-Unit")
    assert.Equal(t, cShopContractPlanItems[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cShopContractPlanItems[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cShopContractPlanItems[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cShopContractPlanItems[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCShopContractPlanItemUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopContractPlanItemPersistence := NewPersistence(tr)
    ctx := context.Background()
    cShopContractPlanItemPersistence.CreateCShopContractPlanItem(&ctx, &model.CShopContractPlanItem{
                                                        ShopContractPlanItemCd: 0,
                                                        LanguageCd: 1,
                                                        ShopContractPlanName: "dummy-ShopContractPlanName",
                                                        Unit: "dummy-Unit",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cShopContractPlanItemPersistence.UpdateCShopContractPlanItem(&ctx, &model.CShopContractPlanItem{
                                                        ShopContractPlanItemCd: 0,
                                                        LanguageCd: 1,
                                                        ShopContractPlanName: "dummy-ShopContractPlanName2",
                                                        Unit: "dummy-Unit2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cShopContractPlanItems, _ := cShopContractPlanItemPersistence.GetCShopContractPlanItemWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopContractPlanItems[0].ShopContractPlanName, "dummy-ShopContractPlanName2")
    assert.Equal(t, cShopContractPlanItems[0].Unit, "dummy-Unit2")
    assert.Equal(t, cShopContractPlanItems[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cShopContractPlanItems[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cShopContractPlanItems[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cShopContractPlanItems[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
