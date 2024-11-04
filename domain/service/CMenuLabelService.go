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

type CMenuLabelService interface {
    CreateCMenuLabel(*context.Context, *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError)
    UpdateCMenuLabel(*context.Context, *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError)
    GetCMenuLabelWithKey(*context.Context, int) ([]*model.CMenuLabel, *errordef.LogicError)
    GetCMenuLabelWithLanguageCd(*context.Context, int) ([]*model.CMenuLabel, *errordef.LogicError)
    DeleteCMenuLabel(*context.Context, int) *errordef.LogicError
}

type cMenuLabelService struct {
    cMenuLabelRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCMenuLabelService(pr repository.Repository, nu number.NumberUtil) CMenuLabelService {
    return &cMenuLabelService{
        cMenuLabelRepository :pr,
        numberUtil :nu,
    }
}

func (ss cMenuLabelService) CreateCMenuLabel(ctx *context.Context, cMenuLabel *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError) {
    cMenuLabel.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cMenuLabel.CreateFunction = "CreateCMenuLabel"
    cMenuLabel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cMenuLabel.UpdateFunction = "CreateCMenuLabel"
    cMenuLabelRepository := ss.cMenuLabelRepository
    created, err := cMenuLabelRepository.CreateCMenuLabel(ctx, cMenuLabel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cMenuLabelService) UpdateCMenuLabel(ctx *context.Context, cMenuLabel *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError) {
    cMenuLabel.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cMenuLabel.UpdateFunction = "UpdateCMenuLabel"
    cMenuLabelRepository := ss.cMenuLabelRepository
    results, err := cMenuLabelRepository.GetCMenuLabelWithKey(ctx, cMenuLabel.MenuLabelCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cMenuLabelRepository.UpdateCMenuLabel(ctx, cMenuLabel)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cMenuLabelService) DeleteCMenuLabel(ctx *context.Context, menuLabelCd int) *errordef.LogicError {
    cMenuLabelRepository := ss.cMenuLabelRepository
    results, err := cMenuLabelRepository.GetCMenuLabelWithKey(ctx, menuLabelCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cMenuLabelRepository.DeleteCMenuLabel(ctx, menuLabelCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cMenuLabelService) GetCMenuLabelWithKey(ctx *context.Context, menuLabelCd int) ([]*model.CMenuLabel, *errordef.LogicError) {
    cMenuLabelRepository := ss.cMenuLabelRepository
    cMenuLabel, err := cMenuLabelRepository.GetCMenuLabelWithKey(ctx, menuLabelCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cMenuLabel, nil
}

func (ss cMenuLabelService) GetCMenuLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CMenuLabel, *errordef.LogicError) {
    cMenuLabelRepository := ss.cMenuLabelRepository
    cMenuLabel, err := cMenuLabelRepository.GetCMenuLabelWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cMenuLabel, nil
}

