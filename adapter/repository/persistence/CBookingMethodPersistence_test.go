package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCBookingMethodCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBookingMethodPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cBookingMethodPersistence.CreateCBookingMethod(&ctx, &model.CBookingMethod{
                                                        BookingMethodCd: 0,
                                                        LanguageCd: 1,
                                                        BookingMethodName: "dummy-BookingMethodName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cBookingMethods, _ := cBookingMethodPersistence.GetCBookingMethodWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBookingMethods[0].BookingMethodCd, 0)
    assert.Equal(t, cBookingMethods[0].LanguageCd, 1)
    assert.Equal(t, cBookingMethods[0].BookingMethodName, "dummy-BookingMethodName")
    assert.Equal(t, cBookingMethods[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cBookingMethods[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cBookingMethods[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cBookingMethods[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCBookingMethodUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBookingMethodPersistence := NewPersistence(tr)
    ctx := context.Background()
    cBookingMethodPersistence.CreateCBookingMethod(&ctx, &model.CBookingMethod{
                                                        BookingMethodCd: 0,
                                                        LanguageCd: 1,
                                                        BookingMethodName: "dummy-BookingMethodName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cBookingMethodPersistence.UpdateCBookingMethod(&ctx, &model.CBookingMethod{
                                                        BookingMethodCd: 0,
                                                        LanguageCd: 1,
                                                        BookingMethodName: "dummy-BookingMethodName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cBookingMethods, _ := cBookingMethodPersistence.GetCBookingMethodWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBookingMethods[0].BookingMethodName, "dummy-BookingMethodName2")
    assert.Equal(t, cBookingMethods[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cBookingMethods[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cBookingMethods[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cBookingMethods[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
