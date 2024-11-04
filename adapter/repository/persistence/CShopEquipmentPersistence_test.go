package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCShopEquipmentCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopEquipmentPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cShopEquipmentPersistence.CreateCShopEquipment(&ctx, &model.CShopEquipment{
                                                        ShopEquipmentCd: 0,
                                                        LanguageCd: 1,
                                                        ShopEquipmentName: "dummy-ShopEquipmentName",
                                                        UnitName: "dummy-UnitName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cShopEquipments, _ := cShopEquipmentPersistence.GetCShopEquipmentWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopEquipments[0].ShopEquipmentCd, 0)
    assert.Equal(t, cShopEquipments[0].LanguageCd, 1)
    assert.Equal(t, cShopEquipments[0].ShopEquipmentName, "dummy-ShopEquipmentName")
    assert.Equal(t, cShopEquipments[0].UnitName, "dummy-UnitName")
    assert.Equal(t, cShopEquipments[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cShopEquipments[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cShopEquipments[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cShopEquipments[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCShopEquipmentUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cShopEquipmentPersistence := NewPersistence(tr)
    ctx := context.Background()
    cShopEquipmentPersistence.CreateCShopEquipment(&ctx, &model.CShopEquipment{
                                                        ShopEquipmentCd: 0,
                                                        LanguageCd: 1,
                                                        ShopEquipmentName: "dummy-ShopEquipmentName",
                                                        UnitName: "dummy-UnitName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cShopEquipmentPersistence.UpdateCShopEquipment(&ctx, &model.CShopEquipment{
                                                        ShopEquipmentCd: 0,
                                                        LanguageCd: 1,
                                                        ShopEquipmentName: "dummy-ShopEquipmentName2",
                                                        UnitName: "dummy-UnitName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cShopEquipments, _ := cShopEquipmentPersistence.GetCShopEquipmentWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cShopEquipments[0].ShopEquipmentName, "dummy-ShopEquipmentName2")
    assert.Equal(t, cShopEquipments[0].UnitName, "dummy-UnitName2")
    assert.Equal(t, cShopEquipments[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cShopEquipments[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cShopEquipments[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cShopEquipments[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
