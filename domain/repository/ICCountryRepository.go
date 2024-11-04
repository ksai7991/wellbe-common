package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CCountryRepository interface {
    CreateCCountry(*context.Context, *model.CCountry) (*model.CCountry, *errordef.LogicError)
    UpdateCCountry(*context.Context, *model.CCountry) (*model.CCountry, *errordef.LogicError)
    DeleteCCountry(*context.Context, int, int) *errordef.LogicError
    GetCCountryWithKey(*context.Context, int,int) ([]*model.CCountry, *errordef.LogicError)
    GetCCountryWithCountryCdIso(*context.Context, string) ([]*model.CCountry, *errordef.LogicError)
    GetCCountryWithLanguageCd(*context.Context, int) ([]*model.CCountry, *errordef.LogicError)
}
