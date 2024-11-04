package query

import (
	"context"
	"testing"
	repository "wellbe-common/adapter/repository"
	env "wellbe-common/settings/env"

	"github.com/stretchr/testify/assert"
)

func TestQueryCurrencyForPayment(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    query := NewQuery(tr)
    ctx := context.Background()
    counts, err := query.QueryCurrencyForPayment(&ctx, "2")
    tr.Rollback(&ctx)
    assert.Nil(t, err)
    assert.Equal(t, 2, len(counts))
    assert.Equal(t, 1, counts[0].CurrencyCd)
    assert.Equal(t, 2, counts[0].LanguageCd)
    assert.Equal(t, "ドル", counts[0].CurrencyName)
    assert.Equal(t, "USD", counts[0].CurrencyCdIso)
}