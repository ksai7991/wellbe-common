package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCTellCountryCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cTellCountryPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cTellCountryPersistence.CreateCTellCountry(&ctx, &model.CTellCountry{
                                                        LanguageCd: 0,
                                                        TellCountryCd: 1,
                                                        CountryName: "dummy-CountryName",
                                                        CountryNo: "dummy-CountryNo",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cTellCountrys, _ := cTellCountryPersistence.GetCTellCountryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cTellCountrys[0].LanguageCd, 0)
    assert.Equal(t, cTellCountrys[0].TellCountryCd, 1)
    assert.Equal(t, cTellCountrys[0].CountryName, "dummy-CountryName")
    assert.Equal(t, cTellCountrys[0].CountryNo, "dummy-CountryNo")
    assert.Equal(t, cTellCountrys[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cTellCountrys[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cTellCountrys[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cTellCountrys[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCTellCountryUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cTellCountryPersistence := NewPersistence(tr)
    ctx := context.Background()
    cTellCountryPersistence.CreateCTellCountry(&ctx, &model.CTellCountry{
                                                        LanguageCd: 0,
                                                        TellCountryCd: 1,
                                                        CountryName: "dummy-CountryName",
                                                        CountryNo: "dummy-CountryNo",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cTellCountryPersistence.UpdateCTellCountry(&ctx, &model.CTellCountry{
                                                        LanguageCd: 0,
                                                        TellCountryCd: 1,
                                                        CountryName: "dummy-CountryName2",
                                                        CountryNo: "dummy-CountryNo2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cTellCountrys, _ := cTellCountryPersistence.GetCTellCountryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cTellCountrys[0].CountryName, "dummy-CountryName2")
    assert.Equal(t, cTellCountrys[0].CountryNo, "dummy-CountryNo2")
    assert.Equal(t, cTellCountrys[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cTellCountrys[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cTellCountrys[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cTellCountrys[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
