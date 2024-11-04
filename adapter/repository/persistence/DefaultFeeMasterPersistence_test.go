package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestDefaultFeeMasterCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    defaultFeeMasterPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := defaultFeeMasterPersistence.CreateDefaultFeeMaster(&ctx, &model.DefaultFeeMaster{
                                                        Id: "dummy-Id",
                                                        FeeRate: 1.2,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    defaultFeeMasters, _ := defaultFeeMasterPersistence.GetDefaultFeeMasterWithKey(&ctx, "dummy-Id")
    tr.Rollback(&ctx)
    assert.Equal(t, defaultFeeMasters[0].Id, "dummy-Id")
    assert.Equal(t, defaultFeeMasters[0].FeeRate, 1.2)
    assert.Equal(t, defaultFeeMasters[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, defaultFeeMasters[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, defaultFeeMasters[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, defaultFeeMasters[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestDefaultFeeMasterUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    defaultFeeMasterPersistence := NewPersistence(tr)
    ctx := context.Background()
    defaultFeeMasterPersistence.CreateDefaultFeeMaster(&ctx, &model.DefaultFeeMaster{
                                                        Id: "dummy-Id",
                                                        FeeRate: 1.2,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := defaultFeeMasterPersistence.UpdateDefaultFeeMaster(&ctx, &model.DefaultFeeMaster{
                                                        Id: "dummy-Id",
                                                        FeeRate: 11.2,
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    defaultFeeMasters, _ := defaultFeeMasterPersistence.GetDefaultFeeMasterWithKey(&ctx, "dummy-Id")
    tr.Rollback(&ctx)
    assert.Equal(t, defaultFeeMasters[0].FeeRate, 11.2)
    assert.Equal(t, defaultFeeMasters[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, defaultFeeMasters[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, defaultFeeMasters[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, defaultFeeMasters[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
