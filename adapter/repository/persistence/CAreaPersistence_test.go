package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCAreaCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cAreaPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cAreaPersistence.CreateCArea(&ctx, &model.CArea{
                                                        LanguageCd: 0,
                                                        AreaCd: 1,
                                                        StateCd: 2,
                                                        AreaName: "dummy-AreaName",
                                                        SearchAreaNameSeo: "dummy-SearchAreaNameSeo",
                                                        WestLongitude: 5.2,
                                                        EastLongitude: 6.2,
                                                        NorthLatitude: 7.2,
                                                        SouthLatitude: 8.2,
                                                        SummaryAreaFlg: true,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cAreas, _ := cAreaPersistence.GetCAreaWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cAreas[0].LanguageCd, 0)
    assert.Equal(t, cAreas[0].AreaCd, 1)
    assert.Equal(t, cAreas[0].StateCd, 2)
    assert.Equal(t, cAreas[0].AreaName, "dummy-AreaName")
    assert.Equal(t, cAreas[0].SearchAreaNameSeo, "dummy-SearchAreaNameSeo")
    assert.Equal(t, cAreas[0].WestLongitude, 5.2)
    assert.Equal(t, cAreas[0].EastLongitude, 6.2)
    assert.Equal(t, cAreas[0].NorthLatitude, 7.2)
    assert.Equal(t, cAreas[0].SouthLatitude, 8.2)
    assert.Equal(t, cAreas[0].SummaryAreaFlg, true)
    assert.Equal(t, cAreas[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cAreas[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cAreas[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cAreas[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCAreaUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cAreaPersistence := NewPersistence(tr)
    ctx := context.Background()
    cAreaPersistence.CreateCArea(&ctx, &model.CArea{
                                                        LanguageCd: 0,
                                                        AreaCd: 1,
                                                        StateCd: 2,
                                                        AreaName: "dummy-AreaName",
                                                        SearchAreaNameSeo: "dummy-SearchAreaNameSeo",
                                                        WestLongitude: 5.2,
                                                        EastLongitude: 6.2,
                                                        NorthLatitude: 7.2,
                                                        SouthLatitude: 8.2,
                                                        SummaryAreaFlg: true,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cAreaPersistence.UpdateCArea(&ctx, &model.CArea{
                                                        LanguageCd: 0,
                                                        AreaCd: 1,
                                                        StateCd: 12,
                                                        AreaName: "dummy-AreaName2",
                                                        SearchAreaNameSeo: "dummy-SearchAreaNameSeo2",
                                                        WestLongitude: 15.2,
                                                        EastLongitude: 16.2,
                                                        NorthLatitude: 17.2,
                                                        SouthLatitude: 18.2,
                                                        SummaryAreaFlg: false,
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cAreas, _ := cAreaPersistence.GetCAreaWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cAreas[0].StateCd, 12)
    assert.Equal(t, cAreas[0].AreaName, "dummy-AreaName2")
    assert.Equal(t, cAreas[0].SearchAreaNameSeo, "dummy-SearchAreaNameSeo2")
    assert.Equal(t, cAreas[0].WestLongitude, 15.2)
    assert.Equal(t, cAreas[0].EastLongitude, 16.2)
    assert.Equal(t, cAreas[0].NorthLatitude, 17.2)
    assert.Equal(t, cAreas[0].SouthLatitude, 18.2)
    assert.Equal(t, cAreas[0].SummaryAreaFlg, false)
    assert.Equal(t, cAreas[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cAreas[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cAreas[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cAreas[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
