package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CPayoutItemCategoryRepository interface {
    CreateCPayoutItemCategory(*context.Context, *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError)
    UpdateCPayoutItemCategory(*context.Context, *model.CPayoutItemCategory) (*model.CPayoutItemCategory, *errordef.LogicError)
    DeleteCPayoutItemCategory(*context.Context, int) *errordef.LogicError
    GetCPayoutItemCategoryWithKey(*context.Context, int) ([]*model.CPayoutItemCategory, *errordef.LogicError)
    GetCPayoutItemCategoryWithLanguageCd(*context.Context, int) ([]*model.CPayoutItemCategory, *errordef.LogicError)
}
