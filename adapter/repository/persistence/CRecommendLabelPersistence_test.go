package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCRecommendLabelCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cRecommendLabelPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cRecommendLabelPersistence.CreateCRecommendLabel(&ctx, &model.CRecommendLabel{
                                                        RecommendLabelCd: 0,
                                                        LanguageCd: 1,
                                                        RecommendLabelName: "dummy-RecommendLabelName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cRecommendLabels, _ := cRecommendLabelPersistence.GetCRecommendLabelWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cRecommendLabels[0].RecommendLabelCd, 0)
    assert.Equal(t, cRecommendLabels[0].LanguageCd, 1)
    assert.Equal(t, cRecommendLabels[0].RecommendLabelName, "dummy-RecommendLabelName")
    assert.Equal(t, cRecommendLabels[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cRecommendLabels[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cRecommendLabels[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cRecommendLabels[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCRecommendLabelUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cRecommendLabelPersistence := NewPersistence(tr)
    ctx := context.Background()
    cRecommendLabelPersistence.CreateCRecommendLabel(&ctx, &model.CRecommendLabel{
                                                        RecommendLabelCd: 0,
                                                        LanguageCd: 1,
                                                        RecommendLabelName: "dummy-RecommendLabelName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cRecommendLabelPersistence.UpdateCRecommendLabel(&ctx, &model.CRecommendLabel{
                                                        RecommendLabelCd: 0,
                                                        LanguageCd: 1,
                                                        RecommendLabelName: "dummy-RecommendLabelName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cRecommendLabels, _ := cRecommendLabelPersistence.GetCRecommendLabelWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cRecommendLabels[0].RecommendLabelName, "dummy-RecommendLabelName2")
    assert.Equal(t, cRecommendLabels[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cRecommendLabels[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cRecommendLabels[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cRecommendLabels[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
