package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CTypeOfContactApplication interface {
    GetCTypeOfContactWithKey(*context.Context, int,int) ([]*model.CTypeOfContact, *errordef.LogicError)
    GetCTypeOfContactWithLanguageCd(*context.Context, int) ([]*model.CTypeOfContact, *errordef.LogicError)
}

type cTypeOfContactApplication struct {
    cTypeOfContactService service.CTypeOfContactService
    transaction repository.Transaction
}

func NewCTypeOfContactApplication(ls service.CTypeOfContactService, tr repository.Transaction) CTypeOfContactApplication {
    return &cTypeOfContactApplication{
        cTypeOfContactService :ls,
        transaction :tr,
    }
}

func (sa cTypeOfContactApplication) GetCTypeOfContactWithKey(ctx *context.Context, typeOfContactCd int,languageCd int) ([]*model.CTypeOfContact, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cTypeOfContactService.GetCTypeOfContactWithKey(ctx, typeOfContactCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cTypeOfContactApplication) GetCTypeOfContactWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTypeOfContact, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cTypeOfContactService.GetCTypeOfContactWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
