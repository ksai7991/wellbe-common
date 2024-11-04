package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCServiceCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cServicePersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cServicePersistence.CreateCService(&ctx, &model.CService{
                                                        ServiceCd: 0,
                                                        LanguageCd: 1,
                                                        ServiceName: "dummy-ServiceName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cServices, _ := cServicePersistence.GetCServiceWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cServices[0].ServiceCd, 0)
    assert.Equal(t, cServices[0].LanguageCd, 1)
    assert.Equal(t, cServices[0].ServiceName, "dummy-ServiceName")
    assert.Equal(t, cServices[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cServices[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cServices[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cServices[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCServiceUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cServicePersistence := NewPersistence(tr)
    ctx := context.Background()
    cServicePersistence.CreateCService(&ctx, &model.CService{
                                                        ServiceCd: 0,
                                                        LanguageCd: 1,
                                                        ServiceName: "dummy-ServiceName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cServicePersistence.UpdateCService(&ctx, &model.CService{
                                                        ServiceCd: 0,
                                                        LanguageCd: 1,
                                                        ServiceName: "dummy-ServiceName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cServices, _ := cServicePersistence.GetCServiceWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cServices[0].ServiceName, "dummy-ServiceName2")
    assert.Equal(t, cServices[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cServices[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cServices[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cServices[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
