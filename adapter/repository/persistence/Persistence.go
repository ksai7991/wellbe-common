package persistence

import (
    repository "wellbe-common/domain/repository"

    _ "github.com/lib/pq"
)

type persistence struct{
    transaction repository.Transaction
} 

func NewPersistence(tr repository.Transaction) repository.Repository {
    return &persistence{transaction: tr}
}