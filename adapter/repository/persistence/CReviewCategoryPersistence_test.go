package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCReviewCategoryCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cReviewCategoryPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cReviewCategoryPersistence.CreateCReviewCategory(&ctx, &model.CReviewCategory{
                                                        ReviewCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewCategoryName: "dummy-ReviewCategoryName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cReviewCategorys, _ := cReviewCategoryPersistence.GetCReviewCategoryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cReviewCategorys[0].ReviewCategoryCd, 0)
    assert.Equal(t, cReviewCategorys[0].LanguageCd, 1)
    assert.Equal(t, cReviewCategorys[0].ReviewCategoryName, "dummy-ReviewCategoryName")
    assert.Equal(t, cReviewCategorys[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cReviewCategorys[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cReviewCategorys[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cReviewCategorys[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCReviewCategoryUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cReviewCategoryPersistence := NewPersistence(tr)
    ctx := context.Background()
    cReviewCategoryPersistence.CreateCReviewCategory(&ctx, &model.CReviewCategory{
                                                        ReviewCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewCategoryName: "dummy-ReviewCategoryName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cReviewCategoryPersistence.UpdateCReviewCategory(&ctx, &model.CReviewCategory{
                                                        ReviewCategoryCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewCategoryName: "dummy-ReviewCategoryName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cReviewCategorys, _ := cReviewCategoryPersistence.GetCReviewCategoryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cReviewCategorys[0].ReviewCategoryName, "dummy-ReviewCategoryName2")
    assert.Equal(t, cReviewCategorys[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cReviewCategorys[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cReviewCategorys[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cReviewCategorys[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
