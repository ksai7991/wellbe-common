package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCBillingContentCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBillingContentPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cBillingContentPersistence.CreateCBillingContent(&ctx, &model.CBillingContent{
                                                        BillingContentCd: 0,
                                                        LanguageCd: 1,
                                                        BillingContentName: "dummy-BillingContentName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cBillingContents, _ := cBillingContentPersistence.GetCBillingContentWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBillingContents[0].BillingContentCd, 0)
    assert.Equal(t, cBillingContents[0].LanguageCd, 1)
    assert.Equal(t, cBillingContents[0].BillingContentName, "dummy-BillingContentName")
    assert.Equal(t, cBillingContents[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cBillingContents[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cBillingContents[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cBillingContents[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCBillingContentUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBillingContentPersistence := NewPersistence(tr)
    ctx := context.Background()
    cBillingContentPersistence.CreateCBillingContent(&ctx, &model.CBillingContent{
                                                        BillingContentCd: 0,
                                                        LanguageCd: 1,
                                                        BillingContentName: "dummy-BillingContentName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cBillingContentPersistence.UpdateCBillingContent(&ctx, &model.CBillingContent{
                                                        BillingContentCd: 0,
                                                        LanguageCd: 1,
                                                        BillingContentName: "dummy-BillingContentName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cBillingContents, _ := cBillingContentPersistence.GetCBillingContentWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBillingContents[0].BillingContentName, "dummy-BillingContentName2")
    assert.Equal(t, cBillingContents[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cBillingContents[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cBillingContents[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cBillingContents[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
