package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCGenderCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cGenderPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cGenderPersistence.CreateCGender(&ctx, &model.CGender{
                                                        GenderCd: 0,
                                                        LanguageCd: 1,
                                                        GenderName: "dummy-GenderName",
                                                        GenderAbbreviation: "dummy-GenderAbbreviation",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cGenders, _ := cGenderPersistence.GetCGenderWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cGenders[0].GenderCd, 0)
    assert.Equal(t, cGenders[0].LanguageCd, 1)
    assert.Equal(t, cGenders[0].GenderName, "dummy-GenderName")
    assert.Equal(t, cGenders[0].GenderAbbreviation, "dummy-GenderAbbreviation")
    assert.Equal(t, cGenders[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cGenders[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cGenders[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cGenders[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCGenderUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cGenderPersistence := NewPersistence(tr)
    ctx := context.Background()
    cGenderPersistence.CreateCGender(&ctx, &model.CGender{
                                                        GenderCd: 0,
                                                        LanguageCd: 1,
                                                        GenderName: "dummy-GenderName",
                                                        GenderAbbreviation: "dummy-GenderAbbreviation",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cGenderPersistence.UpdateCGender(&ctx, &model.CGender{
                                                        GenderCd: 0,
                                                        LanguageCd: 1,
                                                        GenderName: "dummy-GenderName2",
                                                        GenderAbbreviation: "dummy-GenderAbbreviation2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cGenders, _ := cGenderPersistence.GetCGenderWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cGenders[0].GenderName, "dummy-GenderName2")
    assert.Equal(t, cGenders[0].GenderAbbreviation, "dummy-GenderAbbreviation2")
    assert.Equal(t, cGenders[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cGenders[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cGenders[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cGenders[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
