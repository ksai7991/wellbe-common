package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCPayoutItemCategoryCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cPayoutItemCategoryPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cPayoutItemCategoryPersistence.CreateCPayoutItemCategory(&ctx, &model.CPayoutItemCategory{
                                                        PayoutItemCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        PayoutItemCategoryName: "dummy-PayoutItemCategoryName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cPayoutItemCategorys, _ := cPayoutItemCategoryPersistence.GetCPayoutItemCategoryWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cPayoutItemCategorys[0].PayoutItemCategoryCd, 0)
    assert.Equal(t, cPayoutItemCategorys[0].LanguageCd, 1)
    assert.Equal(t, cPayoutItemCategorys[0].PayoutItemCategoryName, "dummy-PayoutItemCategoryName")
    assert.Equal(t, cPayoutItemCategorys[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cPayoutItemCategorys[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cPayoutItemCategorys[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cPayoutItemCategorys[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCPayoutItemCategoryUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cPayoutItemCategoryPersistence := NewPersistence(tr)
    ctx := context.Background()
    cPayoutItemCategoryPersistence.CreateCPayoutItemCategory(&ctx, &model.CPayoutItemCategory{
                                                        PayoutItemCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        PayoutItemCategoryName: "dummy-PayoutItemCategoryName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cPayoutItemCategoryPersistence.UpdateCPayoutItemCategory(&ctx, &model.CPayoutItemCategory{
                                                        PayoutItemCategoryCd: 0,
                                                        LanguageCd: 11,
                                                        PayoutItemCategoryName: "dummy-PayoutItemCategoryName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cPayoutItemCategorys, _ := cPayoutItemCategoryPersistence.GetCPayoutItemCategoryWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cPayoutItemCategorys[0].LanguageCd, 11)
    assert.Equal(t, cPayoutItemCategorys[0].PayoutItemCategoryName, "dummy-PayoutItemCategoryName2")
    assert.Equal(t, cPayoutItemCategorys[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cPayoutItemCategorys[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cPayoutItemCategorys[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cPayoutItemCategorys[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
