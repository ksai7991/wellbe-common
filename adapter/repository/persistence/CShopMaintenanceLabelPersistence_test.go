package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCShopMaintenanceLabelCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopMaintenanceLabelPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cShopMaintenanceLabelPersistence.CreateCShopMaintenanceLabel(&ctx, &model.CShopMaintenanceLabel{
                                                        ShopMaintenanceLabelCd: 0,
                                                        LanguageCd: 1,
                                                        ShopMaintenanceLabelName: "dummy-ShopMaintenanceLabelName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cShopMaintenanceLabels, _ := cShopMaintenanceLabelPersistence.GetCShopMaintenanceLabelWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopMaintenanceLabels[0].ShopMaintenanceLabelCd, 0)
    assert.Equal(t, cShopMaintenanceLabels[0].LanguageCd, 1)
    assert.Equal(t, cShopMaintenanceLabels[0].ShopMaintenanceLabelName, "dummy-ShopMaintenanceLabelName")
    assert.Equal(t, cShopMaintenanceLabels[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cShopMaintenanceLabels[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cShopMaintenanceLabels[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cShopMaintenanceLabels[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCShopMaintenanceLabelUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopMaintenanceLabelPersistence := NewPersistence(tr)
    ctx := context.Background()
    cShopMaintenanceLabelPersistence.CreateCShopMaintenanceLabel(&ctx, &model.CShopMaintenanceLabel{
                                                        ShopMaintenanceLabelCd: 0,
                                                        LanguageCd: 1,
                                                        ShopMaintenanceLabelName: "dummy-ShopMaintenanceLabelName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cShopMaintenanceLabelPersistence.UpdateCShopMaintenanceLabel(&ctx, &model.CShopMaintenanceLabel{
                                                        ShopMaintenanceLabelCd: 0,
                                                        LanguageCd: 1,
                                                        ShopMaintenanceLabelName: "dummy-ShopMaintenanceLabelName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cShopMaintenanceLabels, _ := cShopMaintenanceLabelPersistence.GetCShopMaintenanceLabelWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopMaintenanceLabels[0].ShopMaintenanceLabelName, "dummy-ShopMaintenanceLabelName2")
    assert.Equal(t, cShopMaintenanceLabels[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cShopMaintenanceLabels[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cShopMaintenanceLabels[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cShopMaintenanceLabels[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
