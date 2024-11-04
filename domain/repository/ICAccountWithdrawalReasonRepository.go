package repository

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    "context"
)
type CAccountWithdrawalReasonRepository interface {
    CreateCAccountWithdrawalReason(*context.Context, *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError)
    UpdateCAccountWithdrawalReason(*context.Context, *model.CAccountWithdrawalReason) (*model.CAccountWithdrawalReason, *errordef.LogicError)
    DeleteCAccountWithdrawalReason(*context.Context, int, int) *errordef.LogicError
    GetCAccountWithdrawalReasonWithKey(*context.Context, int,int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError)
    GetCAccountWithdrawalReasonWithLanguageCd(*context.Context, int) ([]*model.CAccountWithdrawalReason, *errordef.LogicError)
}
