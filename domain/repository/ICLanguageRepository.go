package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CLanguageRepository interface {
    CreateCLanguage(*context.Context, *model.CLanguage) (*model.CLanguage, *errordef.LogicError)
    UpdateCLanguage(*context.Context, *model.CLanguage) (*model.CLanguage, *errordef.LogicError)
    DeleteCLanguage(*context.Context, int) *errordef.LogicError
    GetCLanguageWithKey(*context.Context, int) ([]*model.CLanguage, *errordef.LogicError)
    GetCLanguageWithLanguageCharCd(*context.Context, string) ([]*model.CLanguage, *errordef.LogicError)
    GetCLanguageWithFilterCol(*context.Context, string,string) ([]*model.CLanguage, *errordef.LogicError)
}
