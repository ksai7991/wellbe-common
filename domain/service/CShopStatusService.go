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

type CShopStatusService interface {
    CreateCShopStatus(*context.Context, *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError)
    UpdateCShopStatus(*context.Context, *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError)
    GetCShopStatusWithKey(*context.Context, int,int) ([]*model.CShopStatus, *errordef.LogicError)
    GetCShopStatusWithLanguageCd(*context.Context, int) ([]*model.CShopStatus, *errordef.LogicError)
    DeleteCShopStatus(*context.Context, int, int) *errordef.LogicError
}

type cShopStatusService struct {
    cShopStatusRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCShopStatusService(pr repository.Repository, nu number.NumberUtil) CShopStatusService {
    return &cShopStatusService{
        cShopStatusRepository :pr,
        numberUtil :nu,
    }
}

func (ss cShopStatusService) CreateCShopStatus(ctx *context.Context, cShopStatus *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError) {
    cShopStatus.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopStatus.CreateFunction = "CreateCShopStatus"
    cShopStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopStatus.UpdateFunction = "CreateCShopStatus"
    cShopStatusRepository := ss.cShopStatusRepository
    created, err := cShopStatusRepository.CreateCShopStatus(ctx, cShopStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cShopStatusService) UpdateCShopStatus(ctx *context.Context, cShopStatus *model.CShopStatus) (*model.CShopStatus, *errordef.LogicError) {
    cShopStatus.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cShopStatus.UpdateFunction = "UpdateCShopStatus"
    cShopStatusRepository := ss.cShopStatusRepository
    results, err := cShopStatusRepository.GetCShopStatusWithKey(ctx, cShopStatus.ShopStatusCd, cShopStatus.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cShopStatusRepository.UpdateCShopStatus(ctx, cShopStatus)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cShopStatusService) DeleteCShopStatus(ctx *context.Context, shopStatusCd int, languageCd int) *errordef.LogicError {
    cShopStatusRepository := ss.cShopStatusRepository
    results, err := cShopStatusRepository.GetCShopStatusWithKey(ctx, shopStatusCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cShopStatusRepository.DeleteCShopStatus(ctx, shopStatusCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cShopStatusService) GetCShopStatusWithKey(ctx *context.Context, shopStatusCd int,languageCd int) ([]*model.CShopStatus, *errordef.LogicError) {
    cShopStatusRepository := ss.cShopStatusRepository
    cShopStatus, err := cShopStatusRepository.GetCShopStatusWithKey(ctx, shopStatusCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopStatus, nil
}

func (ss cShopStatusService) GetCShopStatusWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CShopStatus, *errordef.LogicError) {
    cShopStatusRepository := ss.cShopStatusRepository
    cShopStatus, err := cShopStatusRepository.GetCShopStatusWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cShopStatus, nil
}

