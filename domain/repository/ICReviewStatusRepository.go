package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CReviewStatusRepository interface {
    CreateCReviewStatus(*context.Context, *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError)
    UpdateCReviewStatus(*context.Context, *model.CReviewStatus) (*model.CReviewStatus, *errordef.LogicError)
    DeleteCReviewStatus(*context.Context, int, int) *errordef.LogicError
    GetCReviewStatusWithKey(*context.Context, int,int) ([]*model.CReviewStatus, *errordef.LogicError)
    GetCReviewStatusWithLanguageCd(*context.Context, int) ([]*model.CReviewStatus, *errordef.LogicError)
}
