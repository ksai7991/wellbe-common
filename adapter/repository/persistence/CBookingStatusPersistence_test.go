package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCBookingStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBookingStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cBookingStatusPersistence.CreateCBookingStatus(&ctx, &model.CBookingStatus{
                                                        BookingStatusCd: 0,
                                                        LanguageCd: 1,
                                                        BookingStatusName: "dummy-BookingStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cBookingStatuss, _ := cBookingStatusPersistence.GetCBookingStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBookingStatuss[0].BookingStatusCd, 0)
    assert.Equal(t, cBookingStatuss[0].LanguageCd, 1)
    assert.Equal(t, cBookingStatuss[0].BookingStatusName, "dummy-BookingStatusName")
    assert.Equal(t, cBookingStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cBookingStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cBookingStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cBookingStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCBookingStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBookingStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cBookingStatusPersistence.CreateCBookingStatus(&ctx, &model.CBookingStatus{
                                                        BookingStatusCd: 0,
                                                        LanguageCd: 1,
                                                        BookingStatusName: "dummy-BookingStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cBookingStatusPersistence.UpdateCBookingStatus(&ctx, &model.CBookingStatus{
                                                        BookingStatusCd: 0,
                                                        LanguageCd: 1,
                                                        BookingStatusName: "dummy-BookingStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cBookingStatuss, _ := cBookingStatusPersistence.GetCBookingStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBookingStatuss[0].BookingStatusName, "dummy-BookingStatusName2")
    assert.Equal(t, cBookingStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cBookingStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cBookingStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cBookingStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
