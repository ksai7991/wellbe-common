package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCStateCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cStatePersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cStatePersistence.CreateCState(&ctx, &model.CState{
                                                        StateCd: 0,
                                                        LanguageCd: 1,
                                                        CountryCd: 2,
                                                        StateName: "dummy-StateName",
                                                        StateCdIso: "dummy-StateCdIso",
                                                        TimezoneIana: "dummy-TimezoneIana",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cStates, _ := cStatePersistence.GetCStateWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cStates[0].StateCd, 0)
    assert.Equal(t, cStates[0].LanguageCd, 1)
    assert.Equal(t, cStates[0].CountryCd, 2)
    assert.Equal(t, cStates[0].StateName, "dummy-StateName")
    assert.Equal(t, cStates[0].StateCdIso, "dummy-StateCdIso")
    assert.Equal(t, cStates[0].TimezoneIana, "dummy-TimezoneIana")
    assert.Equal(t, cStates[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cStates[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cStates[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cStates[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCStateUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cStatePersistence := NewPersistence(tr)
    ctx := context.Background()
    cStatePersistence.CreateCState(&ctx, &model.CState{
                                                        StateCd: 0,
                                                        LanguageCd: 1,
                                                        CountryCd: 2,
                                                        StateName: "dummy-StateName",
                                                        StateCdIso: "dummy-StateCdIso",
                                                        TimezoneIana: "dummy-TimezoneIana",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cStatePersistence.UpdateCState(&ctx, &model.CState{
                                                        StateCd: 0,
                                                        LanguageCd: 1,
                                                        CountryCd: 12,
                                                        StateName: "dummy-StateName2",
                                                        StateCdIso: "dummy-StateCdIso2",
                                                        TimezoneIana: "dummy-TimezoneIana2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cStates, _ := cStatePersistence.GetCStateWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cStates[0].CountryCd, 12)
    assert.Equal(t, cStates[0].StateName, "dummy-StateName2")
    assert.Equal(t, cStates[0].StateCdIso, "dummy-StateCdIso2")
    assert.Equal(t, cStates[0].TimezoneIana, "dummy-TimezoneIana2")
    assert.Equal(t, cStates[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cStates[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cStates[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cStates[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
