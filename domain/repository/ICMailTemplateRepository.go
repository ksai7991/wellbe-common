package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CMailTemplateRepository interface {
    CreateCMailTemplate(*context.Context, *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError)
    UpdateCMailTemplate(*context.Context, *model.CMailTemplate) (*model.CMailTemplate, *errordef.LogicError)
    DeleteCMailTemplate(*context.Context, int, int) *errordef.LogicError
    GetCMailTemplateWithKey(*context.Context, int,int) ([]*model.CMailTemplate, *errordef.LogicError)
    GetCMailTemplateWithLanguageCd(*context.Context, int) ([]*model.CMailTemplate, *errordef.LogicError)
}
