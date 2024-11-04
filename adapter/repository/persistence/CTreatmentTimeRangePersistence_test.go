package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCTreatmentTimeRangeCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cTreatmentTimeRangePersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cTreatmentTimeRangePersistence.CreateCTreatmentTimeRange(&ctx, &model.CTreatmentTimeRange{
                                                        TreatmentTimeCd: 0,
                                                        LanguageCd: 1,
                                                        TreatmentTimeName: "dummy-TreatmentTimeName",
                                                        MinTime: 3,
                                                        MaxTime: 4,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cTreatmentTimeRanges, _ := cTreatmentTimeRangePersistence.GetCTreatmentTimeRangeWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cTreatmentTimeRanges[0].TreatmentTimeCd, 0)
    assert.Equal(t, cTreatmentTimeRanges[0].LanguageCd, 1)
    assert.Equal(t, cTreatmentTimeRanges[0].TreatmentTimeName, "dummy-TreatmentTimeName")
    assert.Equal(t, cTreatmentTimeRanges[0].MinTime, 3)
    assert.Equal(t, cTreatmentTimeRanges[0].MaxTime, 4)
    assert.Equal(t, cTreatmentTimeRanges[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cTreatmentTimeRanges[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cTreatmentTimeRanges[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cTreatmentTimeRanges[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCTreatmentTimeRangeUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cTreatmentTimeRangePersistence := NewPersistence(tr)
    ctx := context.Background()
    cTreatmentTimeRangePersistence.CreateCTreatmentTimeRange(&ctx, &model.CTreatmentTimeRange{
                                                        TreatmentTimeCd: 0,
                                                        LanguageCd: 1,
                                                        TreatmentTimeName: "dummy-TreatmentTimeName",
                                                        MinTime: 3,
                                                        MaxTime: 4,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cTreatmentTimeRangePersistence.UpdateCTreatmentTimeRange(&ctx, &model.CTreatmentTimeRange{
                                                        TreatmentTimeCd: 0,
                                                        LanguageCd: 1,
                                                        TreatmentTimeName: "dummy-TreatmentTimeName2",
                                                        MinTime: 13,
                                                        MaxTime: 14,
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cTreatmentTimeRanges, _ := cTreatmentTimeRangePersistence.GetCTreatmentTimeRangeWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cTreatmentTimeRanges[0].TreatmentTimeName, "dummy-TreatmentTimeName2")
    assert.Equal(t, cTreatmentTimeRanges[0].MinTime, 13)
    assert.Equal(t, cTreatmentTimeRanges[0].MaxTime, 14)
    assert.Equal(t, cTreatmentTimeRanges[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cTreatmentTimeRanges[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cTreatmentTimeRanges[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cTreatmentTimeRanges[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
