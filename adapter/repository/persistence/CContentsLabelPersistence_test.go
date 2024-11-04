package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCContentsLabelCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cContentsLabelPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cContentsLabelPersistence.CreateCContentsLabel(&ctx, &model.CContentsLabel{
                                                        ContentsLabelCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsCategoryCd: 2,
                                                        ContentsLabelName: "dummy-ContentsLabelName",
                                                        ContentsLabelUrl: "dummy-ContentsLabelUrl",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cContentsLabels, _ := cContentsLabelPersistence.GetCContentsLabelWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cContentsLabels[0].ContentsLabelCd, 0)
    assert.Equal(t, cContentsLabels[0].LanguageCd, 1)
    assert.Equal(t, cContentsLabels[0].ContentsCategoryCd, 2)
    assert.Equal(t, cContentsLabels[0].ContentsLabelName, "dummy-ContentsLabelName")
    assert.Equal(t, cContentsLabels[0].ContentsLabelUrl, "dummy-ContentsLabelUrl")
    assert.Equal(t, cContentsLabels[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cContentsLabels[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cContentsLabels[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cContentsLabels[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCContentsLabelUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cContentsLabelPersistence := NewPersistence(tr)
    ctx := context.Background()
    cContentsLabelPersistence.CreateCContentsLabel(&ctx, &model.CContentsLabel{
                                                        ContentsLabelCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsCategoryCd: 2,
                                                        ContentsLabelName: "dummy-ContentsLabelName",
                                                        ContentsLabelUrl: "dummy-ContentsLabelUrl",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cContentsLabelPersistence.UpdateCContentsLabel(&ctx, &model.CContentsLabel{
                                                        ContentsLabelCd: 0,
                                                        LanguageCd: 1,
                                                        ContentsCategoryCd: 12,
                                                        ContentsLabelName: "dummy-ContentsLabelName2",
                                                        ContentsLabelUrl: "dummy-ContentsLabelUrl2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cContentsLabels, _ := cContentsLabelPersistence.GetCContentsLabelWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cContentsLabels[0].ContentsCategoryCd, 12)
    assert.Equal(t, cContentsLabels[0].ContentsLabelName, "dummy-ContentsLabelName2")
    assert.Equal(t, cContentsLabels[0].ContentsLabelUrl, "dummy-ContentsLabelUrl2")
    assert.Equal(t, cContentsLabels[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cContentsLabels[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cContentsLabels[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cContentsLabels[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
