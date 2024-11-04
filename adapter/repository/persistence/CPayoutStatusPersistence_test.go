package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCPayoutStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cPayoutStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cPayoutStatusPersistence.CreateCPayoutStatus(&ctx, &model.CPayoutStatus{
                                                        PayoutStatusCd: 0,
                                                        LanguageCd: 1,
                                                        PayoutStatusName: "dummy-PayoutStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cPayoutStatuss, _ := cPayoutStatusPersistence.GetCPayoutStatusWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cPayoutStatuss[0].PayoutStatusCd, 0)
    assert.Equal(t, cPayoutStatuss[0].LanguageCd, 1)
    assert.Equal(t, cPayoutStatuss[0].PayoutStatusName, "dummy-PayoutStatusName")
    assert.Equal(t, cPayoutStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cPayoutStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cPayoutStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cPayoutStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCPayoutStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cPayoutStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cPayoutStatusPersistence.CreateCPayoutStatus(&ctx, &model.CPayoutStatus{
                                                        PayoutStatusCd: 0,
                                                        LanguageCd: 1,
                                                        PayoutStatusName: "dummy-PayoutStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cPayoutStatusPersistence.UpdateCPayoutStatus(&ctx, &model.CPayoutStatus{
                                                        PayoutStatusCd: 0,
                                                        LanguageCd: 11,
                                                        PayoutStatusName: "dummy-PayoutStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cPayoutStatuss, _ := cPayoutStatusPersistence.GetCPayoutStatusWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cPayoutStatuss[0].LanguageCd, 11)
    assert.Equal(t, cPayoutStatuss[0].PayoutStatusName, "dummy-PayoutStatusName2")
    assert.Equal(t, cPayoutStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cPayoutStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cPayoutStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cPayoutStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
