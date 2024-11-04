package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CReviewCategoryRepository interface {
    CreateCReviewCategory(*context.Context, *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError)
    UpdateCReviewCategory(*context.Context, *model.CReviewCategory) (*model.CReviewCategory, *errordef.LogicError)
    DeleteCReviewCategory(*context.Context, int, int) *errordef.LogicError
    GetCReviewCategoryWithKey(*context.Context, int,int) ([]*model.CReviewCategory, *errordef.LogicError)
    GetCReviewCategoryWithLanguageCd(*context.Context, int) ([]*model.CReviewCategory, *errordef.LogicError)
}
