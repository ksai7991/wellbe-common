package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCLanguageCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cLanguagePersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cLanguagePersistence.CreateCLanguage(&ctx, &model.CLanguage{
                                                        LanguageCd: 0,
                                                        LanguageCharCd: "XX",
                                                        LanguageName: "dummy-LanguageName",
                                                        SortNumber: 3,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cLanguages, _ := cLanguagePersistence.GetCLanguageWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cLanguages[0].LanguageCd, 0)
    assert.Equal(t, cLanguages[0].LanguageCharCd, "XX")
    assert.Equal(t, cLanguages[0].LanguageName, "dummy-LanguageName")
    assert.Equal(t, cLanguages[0].SortNumber, 3)
    assert.Equal(t, cLanguages[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cLanguages[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cLanguages[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cLanguages[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCLanguageUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cLanguagePersistence := NewPersistence(tr)
    ctx := context.Background()
    cLanguagePersistence.CreateCLanguage(&ctx, &model.CLanguage{
                                                        LanguageCd: 0,
                                                        LanguageCharCd: "XX",
                                                        LanguageName: "dummy-LanguageName",
                                                        SortNumber: 3,
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cLanguagePersistence.UpdateCLanguage(&ctx, &model.CLanguage{
                                                        LanguageCd: 0,
                                                        LanguageCharCd: "YY",
                                                        LanguageName: "dummy-LanguageName2",
                                                        SortNumber: 13,
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cLanguages, _ := cLanguagePersistence.GetCLanguageWithKey(&ctx, 0)
    tr.Rollback(&ctx)
    assert.Equal(t, cLanguages[0].LanguageCharCd, "YY")
    assert.Equal(t, cLanguages[0].LanguageName, "dummy-LanguageName2")
    assert.Equal(t, cLanguages[0].SortNumber, 13)
    assert.Equal(t, cLanguages[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cLanguages[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cLanguages[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cLanguages[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
