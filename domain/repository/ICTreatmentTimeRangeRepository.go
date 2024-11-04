package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CTreatmentTimeRangeRepository interface {
    CreateCTreatmentTimeRange(*context.Context, *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError)
    UpdateCTreatmentTimeRange(*context.Context, *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError)
    DeleteCTreatmentTimeRange(*context.Context, int, int) *errordef.LogicError
    GetCTreatmentTimeRangeWithKey(*context.Context, int,int) ([]*model.CTreatmentTimeRange, *errordef.LogicError)
    GetCTreatmentTimeRangeWithLanguageCd(*context.Context, int) ([]*model.CTreatmentTimeRange, *errordef.LogicError)
}
