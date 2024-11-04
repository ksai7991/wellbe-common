package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCAgeRangeCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cAgeRangePersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cAgeRangePersistence.CreateCAgeRange(&ctx, &model.CAgeRange{
                                                        AgeRangeCd: 0,
                                                        LanguageCd: 1,
                                                        AgeRangeGender: "dummy-AgeRangeGender",
                                                        AgeRangeFrom: 3,
                                                        AgeRangeTo: 4,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cAgeRanges, _ := cAgeRangePersistence.GetCAgeRangeWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cAgeRanges[0].AgeRangeCd, 0)
    assert.Equal(t, cAgeRanges[0].LanguageCd, 1)
    assert.Equal(t, cAgeRanges[0].AgeRangeGender, "dummy-AgeRangeGender")
    assert.Equal(t, cAgeRanges[0].AgeRangeFrom, 3)
    assert.Equal(t, cAgeRanges[0].AgeRangeTo, 4)
    assert.Equal(t, cAgeRanges[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cAgeRanges[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cAgeRanges[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cAgeRanges[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCAgeRangeUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cAgeRangePersistence := NewPersistence(tr)
    ctx := context.Background()
    cAgeRangePersistence.CreateCAgeRange(&ctx, &model.CAgeRange{
                                                        AgeRangeCd: 0,
                                                        LanguageCd: 1,
                                                        AgeRangeGender: "dummy-AgeRangeGender",
                                                        AgeRangeFrom: 3,
                                                        AgeRangeTo: 4,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cAgeRangePersistence.UpdateCAgeRange(&ctx, &model.CAgeRange{
                                                        AgeRangeCd: 0,
                                                        LanguageCd: 1,
                                                        AgeRangeGender: "dummy-AgeRangeGender2",
                                                        AgeRangeFrom: 13,
                                                        AgeRangeTo: 14,
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cAgeRanges, _ := cAgeRangePersistence.GetCAgeRangeWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cAgeRanges[0].AgeRangeGender, "dummy-AgeRangeGender2")
    assert.Equal(t, cAgeRanges[0].AgeRangeFrom, 13)
    assert.Equal(t, cAgeRanges[0].AgeRangeTo, 14)
    assert.Equal(t, cAgeRanges[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cAgeRanges[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cAgeRanges[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cAgeRanges[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
