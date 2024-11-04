package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCurrencyExchangeRateCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    currencyExchangeRatePersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := currencyExchangeRatePersistence.CreateCurrencyExchangeRate(&ctx, &model.CurrencyExchangeRate{
                                                        BaseCurrencyCd: 0,
                                                        TargetCurrencyCd: 1,
                                                        PaireName: "XXXXXX",
                                                        Rate: 3.2,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    currencyExchangeRates, _ := currencyExchangeRatePersistence.GetCurrencyExchangeRateWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, currencyExchangeRates[0].BaseCurrencyCd, 0)
    assert.Equal(t, currencyExchangeRates[0].TargetCurrencyCd, 1)
    assert.Equal(t, currencyExchangeRates[0].PaireName, "XXXXXX")
    assert.Equal(t, currencyExchangeRates[0].Rate, 3.2)
    assert.Equal(t, currencyExchangeRates[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, currencyExchangeRates[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, currencyExchangeRates[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, currencyExchangeRates[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCurrencyExchangeRateUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    currencyExchangeRatePersistence := NewPersistence(tr)
    ctx := context.Background()
    currencyExchangeRatePersistence.CreateCurrencyExchangeRate(&ctx, &model.CurrencyExchangeRate{
                                                        BaseCurrencyCd: 0,
                                                        TargetCurrencyCd: 1,
                                                        PaireName: "XXXXXX",
                                                        Rate: 3.2,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := currencyExchangeRatePersistence.UpdateCurrencyExchangeRate(&ctx, &model.CurrencyExchangeRate{
                                                        BaseCurrencyCd: 0,
                                                        TargetCurrencyCd: 1,
                                                        PaireName: "YYYYYY",
                                                        Rate: 13.2,
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    currencyExchangeRates, _ := currencyExchangeRatePersistence.GetCurrencyExchangeRateWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, currencyExchangeRates[0].PaireName, "YYYYYY")
    assert.Equal(t, currencyExchangeRates[0].Rate, 13.2)
    assert.Equal(t, currencyExchangeRates[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, currencyExchangeRates[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, currencyExchangeRates[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, currencyExchangeRates[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
