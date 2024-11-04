package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCContactStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cContactStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cContactStatusPersistence.CreateCContactStatus(&ctx, &model.CContactStatus{
                                                        ContactStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ContactStatusName: "dummy-ContactStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cContactStatuss, _ := cContactStatusPersistence.GetCContactStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cContactStatuss[0].ContactStatusCd, 0)
    assert.Equal(t, cContactStatuss[0].LanguageCd, 1)
    assert.Equal(t, cContactStatuss[0].ContactStatusName, "dummy-ContactStatusName")
    assert.Equal(t, cContactStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cContactStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cContactStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cContactStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCContactStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cContactStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cContactStatusPersistence.CreateCContactStatus(&ctx, &model.CContactStatus{
                                                        ContactStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ContactStatusName: "dummy-ContactStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cContactStatusPersistence.UpdateCContactStatus(&ctx, &model.CContactStatus{
                                                        ContactStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ContactStatusName: "dummy-ContactStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cContactStatuss, _ := cContactStatusPersistence.GetCContactStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cContactStatuss[0].ContactStatusName, "dummy-ContactStatusName2")
    assert.Equal(t, cContactStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cContactStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cContactStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cContactStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
