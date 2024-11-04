package service

import (
    "testing"
    "context"
    "github.com/stretchr/testify/assert"
    repository "wellbe-common/domain/repository"
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    number "wellbe-common/share/number"
)

type cMailTemplateMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError)
    FakeUpdate func(*model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CMailTemplate, *errordef.LogicError)
}

func (lr cMailTemplateMockRepository) CreateCMailTemplate(ctx *context.Context, cMailTemplate *model.CMailTemplate)  (*model.CMailTemplate, *errordef.LogicError) {
    return lr.FakeCreate(cMailTemplate)
}

func (lr cMailTemplateMockRepository) UpdateCMailTemplate(ctx *context.Context, cMailTemplate *model.CMailTemplate)  (*model.CMailTemplate, *errordef.LogicError) {
    return lr.FakeUpdate(cMailTemplate)
}

func (lr cMailTemplateMockRepository) DeleteCMailTemplate(ctx *context.Context, mailTemplateCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(mailTemplateCd, languageCd)
}

func (lr cMailTemplateMockRepository)GetCMailTemplateWithKey(ctx *context.Context, mailTemplateCd int, languageCd int)  ([]*model.CMailTemplate, *errordef.LogicError) {
    return lr.FakeGet(mailTemplateCd, languageCd)
}


type cMailTemplateMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cMailTemplateMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCMailTemplate(t *testing.T) {
    ctx := context.Background()
    repository := &cMailTemplateMockRepository{
        FakeCreate: func(cMailTemplate *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError) {
            return cMailTemplate, nil
        },
    }
    numberUtil := &cMailTemplateMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cMailTemplateService := NewCMailTemplateService(repository, numberUtil)
    in_cMailTemplate := new(model.CMailTemplate)
    in_cMailTemplate.MailTemplateCd = 0
    in_cMailTemplate.LanguageCd = 1
    in_cMailTemplate.Subject = "dummy-Subject"
    in_cMailTemplate.Body = "dummy-Body"
    out_cMailTemplate, err := cMailTemplateService.CreateCMailTemplate(&ctx, in_cMailTemplate)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cMailTemplate.MailTemplateCd)
    assert.Equal(t, 1, out_cMailTemplate.LanguageCd)
    assert.Equal(t, "dummy-Subject", out_cMailTemplate.Subject)
    assert.Equal(t, "dummy-Body", out_cMailTemplate.Body)
    assert.NotNil(t, out_cMailTemplate.CreateDatetime)
    assert.NotEqual(t, "", out_cMailTemplate.CreateDatetime)
    assert.Equal(t, "CreateCMailTemplate", out_cMailTemplate.CreateFunction)
    assert.NotNil(t, out_cMailTemplate.UpdateDatetime)
    assert.NotEqual(t, "", out_cMailTemplate.UpdateDatetime)
    assert.Equal(t, "CreateCMailTemplate", out_cMailTemplate.UpdateFunction)
}

func TestUpdateCMailTemplate(t *testing.T) {
    ctx := context.Background()
    repository := &cMailTemplateMockRepository{
        FakeUpdate: func(cMailTemplate *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError) {
            return cMailTemplate, nil
        },
        FakeGet: func(mailTemplateCd int, languageCd int) ([]*model.CMailTemplate, *errordef.LogicError) {
            return []*model.CMailTemplate{&model.CMailTemplate{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cMailTemplateMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cMailTemplateService := NewCMailTemplateService(repository, numberUtil)
    in_cMailTemplate := new(model.CMailTemplate)
    in_cMailTemplate.MailTemplateCd = 0
    in_cMailTemplate.LanguageCd = 1
    in_cMailTemplate.Subject = "dummy-Subject"
    in_cMailTemplate.Body = "dummy-Body"
    out_cMailTemplate, err := cMailTemplateService.UpdateCMailTemplate(&ctx, in_cMailTemplate)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cMailTemplate.MailTemplateCd)
    assert.Equal(t, 1, out_cMailTemplate.LanguageCd)
    assert.Equal(t, "dummy-Subject", out_cMailTemplate.Subject)
    assert.Equal(t, "dummy-Body", out_cMailTemplate.Body)
    assert.NotNil(t, out_cMailTemplate.CreateDatetime)
    assert.Equal(t, "", out_cMailTemplate.CreateDatetime)
    assert.Equal(t, "", out_cMailTemplate.CreateFunction)
    assert.NotNil(t, out_cMailTemplate.UpdateDatetime)
    assert.NotEqual(t, "", out_cMailTemplate.UpdateDatetime)
    assert.Equal(t, "UpdateCMailTemplate", out_cMailTemplate.UpdateFunction)
}
