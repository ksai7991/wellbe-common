package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCCountryCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCountryPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cCountryPersistence.CreateCCountry(&ctx, &model.CCountry{
                                                        CountryCd: 0,
                                                        LanguageCd: 1,
                                                        CountryName: "dummy-CountryName",
                                                        CountryCdIso: "XX",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cCountrys, _ := cCountryPersistence.GetCCountryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCountrys[0].CountryCd, 0)
    assert.Equal(t, cCountrys[0].LanguageCd, 1)
    assert.Equal(t, cCountrys[0].CountryName, "dummy-CountryName")
    assert.Equal(t, cCountrys[0].CountryCdIso, "XX")
    assert.Equal(t, cCountrys[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cCountrys[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cCountrys[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cCountrys[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCCountryUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCountryPersistence := NewPersistence(tr)
    ctx := context.Background()
    cCountryPersistence.CreateCCountry(&ctx, &model.CCountry{
                                                        CountryCd: 0,
                                                        LanguageCd: 1,
                                                        CountryName: "dummy-CountryName",
                                                        CountryCdIso: "XX",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cCountryPersistence.UpdateCCountry(&ctx, &model.CCountry{
                                                        CountryCd: 0,
                                                        LanguageCd: 1,
                                                        CountryName: "dummy-CountryName2",
                                                        CountryCdIso: "YY",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cCountrys, _ := cCountryPersistence.GetCCountryWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCountrys[0].CountryName, "dummy-CountryName2")
    assert.Equal(t, cCountrys[0].CountryCdIso, "YY")
    assert.Equal(t, cCountrys[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cCountrys[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cCountrys[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cCountrys[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
