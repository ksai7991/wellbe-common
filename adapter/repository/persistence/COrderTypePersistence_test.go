package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCOrderTypeCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cOrderTypePersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cOrderTypePersistence.CreateCOrderType(&ctx, &model.COrderType{
                                                        OrderTypeCd: 0,
                                                        LanguageCd: 1,
                                                        OrderTypeName: "dummy-OrderTypeName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cOrderTypes, _ := cOrderTypePersistence.GetCOrderTypeWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cOrderTypes[0].OrderTypeCd, 0)
    assert.Equal(t, cOrderTypes[0].LanguageCd, 1)
    assert.Equal(t, cOrderTypes[0].OrderTypeName, "dummy-OrderTypeName")
    assert.Equal(t, cOrderTypes[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cOrderTypes[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cOrderTypes[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cOrderTypes[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCOrderTypeUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cOrderTypePersistence := NewPersistence(tr)
    ctx := context.Background()
    cOrderTypePersistence.CreateCOrderType(&ctx, &model.COrderType{
                                                        OrderTypeCd: 0,
                                                        LanguageCd: 1,
                                                        OrderTypeName: "dummy-OrderTypeName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cOrderTypePersistence.UpdateCOrderType(&ctx, &model.COrderType{
                                                        OrderTypeCd: 0,
                                                        LanguageCd: 11,
                                                        OrderTypeName: "dummy-OrderTypeName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cOrderTypes, _ := cOrderTypePersistence.GetCOrderTypeWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cOrderTypes[0].LanguageCd, 11)
    assert.Equal(t, cOrderTypes[0].OrderTypeName, "dummy-OrderTypeName2")
    assert.Equal(t, cOrderTypes[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cOrderTypes[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cOrderTypes[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cOrderTypes[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
