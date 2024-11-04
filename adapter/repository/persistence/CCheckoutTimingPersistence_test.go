package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCCheckoutTimingCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCheckoutTimingPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cCheckoutTimingPersistence.CreateCCheckoutTiming(&ctx, &model.CCheckoutTiming{
                                                        CheckoutTimingCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutTimingName: "dummy-CheckoutTimingName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cCheckoutTimings, _ := cCheckoutTimingPersistence.GetCCheckoutTimingWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCheckoutTimings[0].CheckoutTimingCd, 0)
    assert.Equal(t, cCheckoutTimings[0].LanguageCd, 1)
    assert.Equal(t, cCheckoutTimings[0].CheckoutTimingName, "dummy-CheckoutTimingName")
    assert.Equal(t, cCheckoutTimings[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cCheckoutTimings[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cCheckoutTimings[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cCheckoutTimings[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCCheckoutTimingUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCheckoutTimingPersistence := NewPersistence(tr)
    ctx := context.Background()
    cCheckoutTimingPersistence.CreateCCheckoutTiming(&ctx, &model.CCheckoutTiming{
                                                        CheckoutTimingCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutTimingName: "dummy-CheckoutTimingName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cCheckoutTimingPersistence.UpdateCCheckoutTiming(&ctx, &model.CCheckoutTiming{
                                                        CheckoutTimingCd: 0,
                                                        LanguageCd: 1,
                                                        CheckoutTimingName: "dummy-CheckoutTimingName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cCheckoutTimings, _ := cCheckoutTimingPersistence.GetCCheckoutTimingWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCheckoutTimings[0].CheckoutTimingName, "dummy-CheckoutTimingName2")
    assert.Equal(t, cCheckoutTimings[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cCheckoutTimings[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cCheckoutTimings[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cCheckoutTimings[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
