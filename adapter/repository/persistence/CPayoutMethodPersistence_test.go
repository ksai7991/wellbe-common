package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCPayoutMethodCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cPayoutMethodPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cPayoutMethodPersistence.CreateCPayoutMethod(&ctx, &model.CPayoutMethod{
                                                        PayoutMethodCd: 0,
                                                        LanguageCd: 1,
                                                        PayoutMethodName: "dummy-PayoutMethodName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cPayoutMethods, _ := cPayoutMethodPersistence.GetCPayoutMethodWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cPayoutMethods[0].PayoutMethodCd, 0)
    assert.Equal(t, cPayoutMethods[0].LanguageCd, 1)
    assert.Equal(t, cPayoutMethods[0].PayoutMethodName, "dummy-PayoutMethodName")
    assert.Equal(t, cPayoutMethods[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cPayoutMethods[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cPayoutMethods[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cPayoutMethods[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCPayoutMethodUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cPayoutMethodPersistence := NewPersistence(tr)
    ctx := context.Background()
    cPayoutMethodPersistence.CreateCPayoutMethod(&ctx, &model.CPayoutMethod{
                                                        PayoutMethodCd: 0,
                                                        LanguageCd: 1,
                                                        PayoutMethodName: "dummy-PayoutMethodName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cPayoutMethodPersistence.UpdateCPayoutMethod(&ctx, &model.CPayoutMethod{
                                                        PayoutMethodCd: 0,
                                                        LanguageCd: 11,
                                                        PayoutMethodName: "dummy-PayoutMethodName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cPayoutMethods, _ := cPayoutMethodPersistence.GetCPayoutMethodWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cPayoutMethods[0].LanguageCd, 11)
    assert.Equal(t, cPayoutMethods[0].PayoutMethodName, "dummy-PayoutMethodName2")
    assert.Equal(t, cPayoutMethods[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cPayoutMethods[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cPayoutMethods[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cPayoutMethods[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
