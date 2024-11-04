package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCTypeOfContactCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cTypeOfContactPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cTypeOfContactPersistence.CreateCTypeOfContact(&ctx, &model.CTypeOfContact{
                                                        TypeOfContactCd: 0,
                                                        LanguageCd: 1,
                                                        TypeOfContactName: "dummy-TypeOfContactName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cTypeOfContacts, _ := cTypeOfContactPersistence.GetCTypeOfContactWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cTypeOfContacts[0].TypeOfContactCd, 0)
    assert.Equal(t, cTypeOfContacts[0].LanguageCd, 1)
    assert.Equal(t, cTypeOfContacts[0].TypeOfContactName, "dummy-TypeOfContactName")
    assert.Equal(t, cTypeOfContacts[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cTypeOfContacts[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cTypeOfContacts[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cTypeOfContacts[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCTypeOfContactUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cTypeOfContactPersistence := NewPersistence(tr)
    ctx := context.Background()
    cTypeOfContactPersistence.CreateCTypeOfContact(&ctx, &model.CTypeOfContact{
                                                        TypeOfContactCd: 0,
                                                        LanguageCd: 1,
                                                        TypeOfContactName: "dummy-TypeOfContactName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cTypeOfContactPersistence.UpdateCTypeOfContact(&ctx, &model.CTypeOfContact{
                                                        TypeOfContactCd: 0,
                                                        LanguageCd: 1,
                                                        TypeOfContactName: "dummy-TypeOfContactName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cTypeOfContacts, _ := cTypeOfContactPersistence.GetCTypeOfContactWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cTypeOfContacts[0].TypeOfContactName, "dummy-TypeOfContactName2")
    assert.Equal(t, cTypeOfContacts[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cTypeOfContacts[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cTypeOfContacts[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cTypeOfContacts[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
