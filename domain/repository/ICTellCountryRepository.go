package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CTellCountryRepository interface {
    CreateCTellCountry(*context.Context, *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError)
    UpdateCTellCountry(*context.Context, *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError)
    DeleteCTellCountry(*context.Context, int, int) *errordef.LogicError
    GetCTellCountryWithKey(*context.Context, int,int) ([]*model.CTellCountry, *errordef.LogicError)
    GetCTellCountryWithLanguageCd(*context.Context, int) ([]*model.CTellCountry, *errordef.LogicError)
}
