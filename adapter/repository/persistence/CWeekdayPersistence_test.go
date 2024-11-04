package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCWeekdayCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cWeekdayPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cWeekdayPersistence.CreateCWeekday(&ctx, &model.CWeekday{
                                                        WeekdayCd: 0,
                                                        LanguageCd: 1,
                                                        WeekdayName: "dummy-WeekdayName",
                                                        WeekdayAbbreviation: "dummy-WeekdayAbbreviation",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cWeekdays, _ := cWeekdayPersistence.GetCWeekdayWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cWeekdays[0].WeekdayCd, 0)
    assert.Equal(t, cWeekdays[0].LanguageCd, 1)
    assert.Equal(t, cWeekdays[0].WeekdayName, "dummy-WeekdayName")
    assert.Equal(t, cWeekdays[0].WeekdayAbbreviation, "dummy-WeekdayAbbreviation")
    assert.Equal(t, cWeekdays[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cWeekdays[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cWeekdays[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cWeekdays[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCWeekdayUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cWeekdayPersistence := NewPersistence(tr)
    ctx := context.Background()
    cWeekdayPersistence.CreateCWeekday(&ctx, &model.CWeekday{
                                                        WeekdayCd: 0,
                                                        LanguageCd: 1,
                                                        WeekdayName: "dummy-WeekdayName",
                                                        WeekdayAbbreviation: "dummy-WeekdayAbbreviation",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cWeekdayPersistence.UpdateCWeekday(&ctx, &model.CWeekday{
                                                        WeekdayCd: 0,
                                                        LanguageCd: 11,
                                                        WeekdayName: "dummy-WeekdayName2",
                                                        WeekdayAbbreviation: "dummy-WeekdayAbbreviation2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cWeekdays, _ := cWeekdayPersistence.GetCWeekdayWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cWeekdays[0].LanguageCd, 11)
    assert.Equal(t, cWeekdays[0].WeekdayName, "dummy-WeekdayName2")
    assert.Equal(t, cWeekdays[0].WeekdayAbbreviation, "dummy-WeekdayAbbreviation2")
    assert.Equal(t, cWeekdays[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cWeekdays[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cWeekdays[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cWeekdays[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
