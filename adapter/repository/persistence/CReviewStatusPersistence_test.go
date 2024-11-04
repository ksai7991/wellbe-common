package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCReviewStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cReviewStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cReviewStatusPersistence.CreateCReviewStatus(&ctx, &model.CReviewStatus{
                                                        ReviewStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewStatusName: "dummy-ReviewStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cReviewStatuss, _ := cReviewStatusPersistence.GetCReviewStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cReviewStatuss[0].ReviewStatusCd, 0)
    assert.Equal(t, cReviewStatuss[0].LanguageCd, 1)
    assert.Equal(t, cReviewStatuss[0].ReviewStatusName, "dummy-ReviewStatusName")
    assert.Equal(t, cReviewStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cReviewStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cReviewStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cReviewStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCReviewStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cReviewStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cReviewStatusPersistence.CreateCReviewStatus(&ctx, &model.CReviewStatus{
                                                        ReviewStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewStatusName: "dummy-ReviewStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cReviewStatusPersistence.UpdateCReviewStatus(&ctx, &model.CReviewStatus{
                                                        ReviewStatusCd: 0,
                                                        LanguageCd: 1,
                                                        ReviewStatusName: "dummy-ReviewStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cReviewStatuss, _ := cReviewStatusPersistence.GetCReviewStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cReviewStatuss[0].ReviewStatusName, "dummy-ReviewStatusName2")
    assert.Equal(t, cReviewStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cReviewStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cReviewStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cReviewStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
