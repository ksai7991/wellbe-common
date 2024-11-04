package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CShopMaintenanceLabelRepository interface {
    CreateCShopMaintenanceLabel(*context.Context, *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError)
    UpdateCShopMaintenanceLabel(*context.Context, *model.CShopMaintenanceLabel) (*model.CShopMaintenanceLabel, *errordef.LogicError)
    DeleteCShopMaintenanceLabel(*context.Context, int, int) *errordef.LogicError
    GetCShopMaintenanceLabelWithKey(*context.Context, int,int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError)
    GetCShopMaintenanceLabelWithLanguageCd(*context.Context, int) ([]*model.CShopMaintenanceLabel, *errordef.LogicError)
}
