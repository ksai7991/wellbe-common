package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCReviewContentCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cReviewContentPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cReviewContentPersistence.CreateCReviewContent(&ctx, &model.CReviewContent{
                                                        ReviewContentCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewCategoryCd: 2,
                                                        ReviewContentName: "dummy-ReviewContentName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cReviewContents, _ := cReviewContentPersistence.GetCReviewContentWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cReviewContents[0].ReviewContentCd, 0)
    assert.Equal(t, cReviewContents[0].LanguageCd, 1)
    assert.Equal(t, cReviewContents[0].ReviewCategoryCd, 2)
    assert.Equal(t, cReviewContents[0].ReviewContentName, "dummy-ReviewContentName")
    assert.Equal(t, cReviewContents[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cReviewContents[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cReviewContents[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cReviewContents[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCReviewContentUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cReviewContentPersistence := NewPersistence(tr)
    ctx := context.Background()
    cReviewContentPersistence.CreateCReviewContent(&ctx, &model.CReviewContent{
                                                        ReviewContentCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewCategoryCd: 2,
                                                        ReviewContentName: "dummy-ReviewContentName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cReviewContentPersistence.UpdateCReviewContent(&ctx, &model.CReviewContent{
                                                        ReviewContentCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewCategoryCd: 12,
                                                        ReviewContentName: "dummy-ReviewContentName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cReviewContents, _ := cReviewContentPersistence.GetCReviewContentWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cReviewContents[0].ReviewCategoryCd, 12)
    assert.Equal(t, cReviewContents[0].ReviewContentName, "dummy-ReviewContentName2")
    assert.Equal(t, cReviewContents[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cReviewContents[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cReviewContents[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cReviewContents[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
