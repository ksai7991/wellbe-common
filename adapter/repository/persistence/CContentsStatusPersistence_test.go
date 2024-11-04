package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCContentsStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cContentsStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cContentsStatusPersistence.CreateCContentsStatus(&ctx, &model.CContentsStatus{
                                                        ContentsStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsStatusName: "dummy-ContentsStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cContentsStatuss, _ := cContentsStatusPersistence.GetCContentsStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cContentsStatuss[0].ContentsStatusCd, 0)
    assert.Equal(t, cContentsStatuss[0].LanguageCd, 1)
    assert.Equal(t, cContentsStatuss[0].ContentsStatusName, "dummy-ContentsStatusName")
    assert.Equal(t, cContentsStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cContentsStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cContentsStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cContentsStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCContentsStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cContentsStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cContentsStatusPersistence.CreateCContentsStatus(&ctx, &model.CContentsStatus{
                                                        ContentsStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsStatusName: "dummy-ContentsStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cContentsStatusPersistence.UpdateCContentsStatus(&ctx, &model.CContentsStatus{
                                                        ContentsStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsStatusName: "dummy-ContentsStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cContentsStatuss, _ := cContentsStatusPersistence.GetCContentsStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cContentsStatuss[0].ContentsStatusName, "dummy-ContentsStatusName2")
    assert.Equal(t, cContentsStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cContentsStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cContentsStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cContentsStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
