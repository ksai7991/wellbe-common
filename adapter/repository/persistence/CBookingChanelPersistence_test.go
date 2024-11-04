package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCBookingChanelCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBookingChanelPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cBookingChanelPersistence.CreateCBookingChanel(&ctx, &model.CBookingChanel{
                                                        BookingChanelCd: 0,
                                                        LanguageCd: 1,
                                                        BookingChanelName: "dummy-BookingChanelName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cBookingChanels, _ := cBookingChanelPersistence.GetCBookingChanelWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBookingChanels[0].BookingChanelCd, 0)
    assert.Equal(t, cBookingChanels[0].LanguageCd, 1)
    assert.Equal(t, cBookingChanels[0].BookingChanelName, "dummy-BookingChanelName")
    assert.Equal(t, cBookingChanels[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cBookingChanels[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cBookingChanels[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cBookingChanels[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCBookingChanelUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cBookingChanelPersistence := NewPersistence(tr)
    ctx := context.Background()
    cBookingChanelPersistence.CreateCBookingChanel(&ctx, &model.CBookingChanel{
                                                        BookingChanelCd: 0,
                                                        LanguageCd: 1,
                                                        BookingChanelName: "dummy-BookingChanelName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cBookingChanelPersistence.UpdateCBookingChanel(&ctx, &model.CBookingChanel{
                                                        BookingChanelCd: 0,
                                                        LanguageCd: 1,
                                                        BookingChanelName: "dummy-BookingChanelName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cBookingChanels, _ := cBookingChanelPersistence.GetCBookingChanelWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cBookingChanels[0].BookingChanelName, "dummy-BookingChanelName2")
    assert.Equal(t, cBookingChanels[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cBookingChanels[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cBookingChanels[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cBookingChanels[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
