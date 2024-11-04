package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCMenuLabelCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cMenuLabelPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cMenuLabelPersistence.CreateCMenuLabel(&ctx, &model.CMenuLabel{
                                                        MenuLabelCd: 0,
                                                        LanguageCd: 1,
                                                        MenuLabelName: "dummy-MenuLabelName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cMenuLabels, _ := cMenuLabelPersistence.GetCMenuLabelWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cMenuLabels[0].MenuLabelCd, 0)
    assert.Equal(t, cMenuLabels[0].LanguageCd, 1)
    assert.Equal(t, cMenuLabels[0].MenuLabelName, "dummy-MenuLabelName")
    assert.Equal(t, cMenuLabels[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cMenuLabels[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cMenuLabels[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cMenuLabels[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCMenuLabelUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cMenuLabelPersistence := NewPersistence(tr)
    ctx := context.Background()
    cMenuLabelPersistence.CreateCMenuLabel(&ctx, &model.CMenuLabel{
                                                        MenuLabelCd: 0,
                                                        LanguageCd: 1,
                                                        MenuLabelName: "dummy-MenuLabelName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cMenuLabelPersistence.UpdateCMenuLabel(&ctx, &model.CMenuLabel{
                                                        MenuLabelCd: 0,
                                                        LanguageCd: 11,
                                                        MenuLabelName: "dummy-MenuLabelName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cMenuLabels, _ := cMenuLabelPersistence.GetCMenuLabelWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cMenuLabels[0].LanguageCd, 11)
    assert.Equal(t, cMenuLabels[0].MenuLabelName, "dummy-MenuLabelName2")
    assert.Equal(t, cMenuLabels[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cMenuLabels[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cMenuLabels[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cMenuLabels[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
