package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCConcernCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cConcernPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cConcernPersistence.CreateCConcern(&ctx, &model.CConcern{
                                                        ConcernCd: 0,
                                                        LanguageCd: 1,
                                                        ConcernName: "dummy-ConcernName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cConcerns, _ := cConcernPersistence.GetCConcernWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cConcerns[0].ConcernCd, 0)
    assert.Equal(t, cConcerns[0].LanguageCd, 1)
    assert.Equal(t, cConcerns[0].ConcernName, "dummy-ConcernName")
    assert.Equal(t, cConcerns[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cConcerns[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cConcerns[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cConcerns[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCConcernUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cConcernPersistence := NewPersistence(tr)
    ctx := context.Background()
    cConcernPersistence.CreateCConcern(&ctx, &model.CConcern{
                                                        ConcernCd: 0,
                                                        LanguageCd: 1,
                                                        ConcernName: "dummy-ConcernName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cConcernPersistence.UpdateCConcern(&ctx, &model.CConcern{
                                                        ConcernCd: 0,
                                                        LanguageCd: 1,
                                                        ConcernName: "dummy-ConcernName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cConcerns, _ := cConcernPersistence.GetCConcernWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cConcerns[0].ConcernName, "dummy-ConcernName2")
    assert.Equal(t, cConcerns[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cConcerns[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cConcerns[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cConcerns[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
