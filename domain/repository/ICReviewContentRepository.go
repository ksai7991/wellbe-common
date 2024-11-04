package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CReviewContentRepository interface {
    CreateCReviewContent(*context.Context, *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError)
    UpdateCReviewContent(*context.Context, *model.CReviewContent) (*model.CReviewContent, *errordef.LogicError)
    DeleteCReviewContent(*context.Context, int, int) *errordef.LogicError
    GetCReviewContentWithKey(*context.Context, int,int) ([]*model.CReviewContent, *errordef.LogicError)
    GetCReviewContentWithLanguageCd(*context.Context, int) ([]*model.CReviewContent, *errordef.LogicError)
}
