package service

import (
    model "wellbe-common/domain/model"
    repository "wellbe-common/domain/repository"
    errordef "wellbe-common/share/errordef"
    number "wellbe-common/share/number"
    commonconstants "wellbe-common/share/commonsettings/constants"
    datetime "wellbe-common/share/datetime"
    messages "wellbe-common/share/messages"

    "context"
)

type CInvoiceStatusService interface {
    CreateCInvoiceStatus(*context.Context, *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError)
    UpdateCInvoiceStatus(*context.Context, *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError)
    GetCInvoiceStatusWithKey(*context.Context, int,int) ([]*model.CInvoiceStatus, *errordef.LogicError)
    GetCInvoiceStatusWithLanguageCd(*context.Context, int) ([]*model.CInvoiceStatus, *errordef.LogicError)
    DeleteCInvoiceStatus(*context.Context, int, int) *errordef.LogicError
}

type cInvoiceStatusService struct {
    cInvoiceStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCInvoiceStatusService(pr repository.Repository, nu number.NumberUtil) CInvoiceStatusService {
    return &cInvoiceStatusService{
        cInvoiceStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cInvoiceStatusService) CreateCInvoiceStatus(ctx *context.Context, cInvoiceStatus *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError) {
    cInvoiceStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cInvoiceStatus.CreateFunction = "CreateCInvoiceStatus"
    cInvoiceStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cInvoiceStatus.UpdateFunction = "CreateCInvoiceStatus"
    cInvoiceStatusRepository := ss.cInvoiceStatusRepository
    created, err := cInvoiceStatusRepository.CreateCInvoiceStatus(ctx, cInvoiceStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cInvoiceStatusService) UpdateCInvoiceStatus(ctx *context.Context, cInvoiceStatus *model.CInvoiceStatus) (*model.CInvoiceStatus, *errordef.LogicError) {
    cInvoiceStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cInvoiceStatus.UpdateFunction = "UpdateCInvoiceStatus"
    cInvoiceStatusRepository := ss.cInvoiceStatusRepository
    results, err := cInvoiceStatusRepository.GetCInvoiceStatusWithKey(ctx, cInvoiceStatus.InvoiceStatusCd, cInvoiceStatus.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cInvoiceStatusRepository.UpdateCInvoiceStatus(ctx, cInvoiceStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cInvoiceStatusService) DeleteCInvoiceStatus(ctx *context.Context, invoiceStatusCd int, languageCd int) *errordef.LogicError {
    cInvoiceStatusRepository := ss.cInvoiceStatusRepository
    results, err := cInvoiceStatusRepository.GetCInvoiceStatusWithKey(ctx, invoiceStatusCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cInvoiceStatusRepository.DeleteCInvoiceStatus(ctx, invoiceStatusCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cInvoiceStatusService) GetCInvoiceStatusWithKey(ctx *context.Context, invoiceStatusCd int,languageCd int) ([]*model.CInvoiceStatus, *errordef.LogicError) {
    cInvoiceStatusRepository := ss.cInvoiceStatusRepository
    cInvoiceStatus, err := cInvoiceStatusRepository.GetCInvoiceStatusWithKey(ctx, invoiceStatusCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cInvoiceStatus, nil
}

func (ss cInvoiceStatusService) GetCInvoiceStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CInvoiceStatus, *errordef.LogicError) {
    cInvoiceStatusRepository := ss.cInvoiceStatusRepository
    cInvoiceStatus, err := cInvoiceStatusRepository.GetCInvoiceStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cInvoiceStatus, nil
}

