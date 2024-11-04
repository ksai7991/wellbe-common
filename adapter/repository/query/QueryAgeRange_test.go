package query

import (
	"context"
	"testing"
	repository "wellbe-common/adapter/repository"
	env "wellbe-common/settings/env"

	"github.com/stretchr/testify/assert"
)

func TestQueryAgeRange(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    query := NewQuery(tr)
    ctx := context.Background()
    age, err := query.QueryAgeRange(&ctx, "2", 30)
    tr.Rollback(&ctx)
    assert.Nil(t, err)
    assert.Equal(t, 4, age.AgeRangeCd)
    assert.Equal(t, 2, age.LanguageCd)
    assert.Equal(t, "30代前半", age.AgeRangeGender)
}