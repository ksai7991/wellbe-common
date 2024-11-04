package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CRecommendLabelRepository interface {
    CreateCRecommendLabel(*context.Context, *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError)
    UpdateCRecommendLabel(*context.Context, *model.CRecommendLabel) (*model.CRecommendLabel, *errordef.LogicError)
    DeleteCRecommendLabel(*context.Context, int, int) *errordef.LogicError
    GetCRecommendLabelWithKey(*context.Context, int,int) ([]*model.CRecommendLabel, *errordef.LogicError)
    GetCRecommendLabelWithLanguageCd(*context.Context, int) ([]*model.CRecommendLabel, *errordef.LogicError)
}
