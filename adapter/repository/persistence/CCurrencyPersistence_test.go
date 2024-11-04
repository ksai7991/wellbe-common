package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCCurrencyCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCurrencyPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cCurrencyPersistence.CreateCCurrency(&ctx, &model.CCurrency{
                                                        CurrencyCd: 0,
                                                        LanguageCd: 1,
                                                        CurrencyName: "dummy-CurrencyName",
                                                        CurrencyCdIso: "XXX",
                                                        SignificantDigit: 4,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cCurrencys, _ := cCurrencyPersistence.GetCCurrencyWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCurrencys[0].CurrencyCd, 0)
    assert.Equal(t, cCurrencys[0].LanguageCd, 1)
    assert.Equal(t, cCurrencys[0].CurrencyName, "dummy-CurrencyName")
    assert.Equal(t, cCurrencys[0].CurrencyCdIso, "XXX")
    assert.Equal(t, cCurrencys[0].SignificantDigit, 4)
    assert.Equal(t, cCurrencys[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cCurrencys[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cCurrencys[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cCurrencys[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCCurrencyUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCurrencyPersistence := NewPersistence(tr)
    ctx := context.Background()
    cCurrencyPersistence.CreateCCurrency(&ctx, &model.CCurrency{
                                                        CurrencyCd: 0,
                                                        LanguageCd: 1,
                                                        CurrencyName: "dummy-CurrencyName",
                                                        CurrencyCdIso: "XXX",
                                                        SignificantDigit: 4,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cCurrencyPersistence.UpdateCCurrency(&ctx, &model.CCurrency{
                                                        CurrencyCd: 0,
                                                        LanguageCd: 1,
                                                        CurrencyName: "dummy-CurrencyName2",
                                                        CurrencyCdIso: "YYY",
                                                        SignificantDigit: 14,
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cCurrencys, _ := cCurrencyPersistence.GetCCurrencyWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCurrencys[0].CurrencyName, "dummy-CurrencyName2")
    assert.Equal(t, cCurrencys[0].CurrencyCdIso, "YYY")
    assert.Equal(t, cCurrencys[0].SignificantDigit, 14)
    assert.Equal(t, cCurrencys[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cCurrencys[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cCurrencys[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cCurrencys[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
