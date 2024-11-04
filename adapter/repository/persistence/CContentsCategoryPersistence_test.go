package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCContentsCategoryCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cContentsCategoryPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cContentsCategoryPersistence.CreateCContentsCategory(&ctx, &model.CContentsCategory{
                                                        ContentsCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsCategoryName: "dummy-ContentsCategoryName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cContentsCategorys, _ := cContentsCategoryPersistence.GetCContentsCategoryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cContentsCategorys[0].ContentsCategoryCd, 0)
    assert.Equal(t, cContentsCategorys[0].LanguageCd, 1)
    assert.Equal(t, cContentsCategorys[0].ContentsCategoryName, "dummy-ContentsCategoryName")
    assert.Equal(t, cContentsCategorys[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cContentsCategorys[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cContentsCategorys[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cContentsCategorys[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCContentsCategoryUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cContentsCategoryPersistence := NewPersistence(tr)
    ctx := context.Background()
    cContentsCategoryPersistence.CreateCContentsCategory(&ctx, &model.CContentsCategory{
                                                        ContentsCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsCategoryName: "dummy-ContentsCategoryName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cContentsCategoryPersistence.UpdateCContentsCategory(&ctx, &model.CContentsCategory{
                                                        ContentsCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsCategoryName: "dummy-ContentsCategoryName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cContentsCategorys, _ := cContentsCategoryPersistence.GetCContentsCategoryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cContentsCategorys[0].ContentsCategoryName, "dummy-ContentsCategoryName2")
    assert.Equal(t, cContentsCategorys[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cContentsCategorys[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cContentsCategorys[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cContentsCategorys[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
