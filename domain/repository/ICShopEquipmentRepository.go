package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CShopEquipmentRepository interface {
    CreateCShopEquipment(*context.Context, *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError)
    UpdateCShopEquipment(*context.Context, *model.CShopEquipment) (*model.CShopEquipment, *errordef.LogicError)
    DeleteCShopEquipment(*context.Context, int, int) *errordef.LogicError
    GetCShopEquipmentWithKey(*context.Context, int,int) ([]*model.CShopEquipment, *errordef.LogicError)
    GetCShopEquipmentWithLanguageCd(*context.Context, int) ([]*model.CShopEquipment, *errordef.LogicError)
}
