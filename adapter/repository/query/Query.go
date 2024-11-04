package query

import (
	repository "wellbe-common/domain/repository"

	_ "github.com/lib/pq"
)

type Query struct{
    transaction repository.Transaction
} 

func NewQuery(tr repository.Transaction) *Query {
    return &Query{transaction: tr}
}