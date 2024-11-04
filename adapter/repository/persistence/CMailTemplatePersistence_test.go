package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCMailTemplateCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cMailTemplatePersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cMailTemplatePersistence.CreateCMailTemplate(&ctx, &model.CMailTemplate{
                                                        MailTemplateCd: 0,
                                                        LanguageCd: 1,
                                                        Subject: "dummy-Subject",
                                                        Body: "dummy-Body",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cMailTemplates, _ := cMailTemplatePersistence.GetCMailTemplateWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cMailTemplates[0].MailTemplateCd, 0)
    assert.Equal(t, cMailTemplates[0].LanguageCd, 1)
    assert.Equal(t, cMailTemplates[0].Subject, "dummy-Subject")
    assert.Equal(t, cMailTemplates[0].Body, "dummy-Body")
    assert.Equal(t, cMailTemplates[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cMailTemplates[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cMailTemplates[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cMailTemplates[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCMailTemplateUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cMailTemplatePersistence := NewPersistence(tr)
    ctx := context.Background()
    cMailTemplatePersistence.CreateCMailTemplate(&ctx, &model.CMailTemplate{
                                                        MailTemplateCd: 0,
                                                        LanguageCd: 1,
                                                        Subject: "dummy-Subject",
                                                        Body: "dummy-Body",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cMailTemplatePersistence.UpdateCMailTemplate(&ctx, &model.CMailTemplate{
                                                        MailTemplateCd: 0,
                                                        LanguageCd: 1,
                                                        Subject: "dummy-Subject2",
                                                        Body: "dummy-Body2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cMailTemplates, _ := cMailTemplatePersistence.GetCMailTemplateWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cMailTemplates[0].Subject, "dummy-Subject2")
    assert.Equal(t, cMailTemplates[0].Body, "dummy-Body2")
    assert.Equal(t, cMailTemplates[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cMailTemplates[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cMailTemplates[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cMailTemplates[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
