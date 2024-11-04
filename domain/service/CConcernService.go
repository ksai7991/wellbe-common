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

type CConcernService interface {
    CreateCConcern(*context.Context, *model.CConcern) (*model.CConcern, *errordef.LogicError)
    UpdateCConcern(*context.Context, *model.CConcern) (*model.CConcern, *errordef.LogicError)
    GetCConcernWithKey(*context.Context, int,int) ([]*model.CConcern, *errordef.LogicError)
    GetCConcernWithLanguageCd(*context.Context, int) ([]*model.CConcern, *errordef.LogicError)
    DeleteCConcern(*context.Context, int, int) *errordef.LogicError
}

type cConcernService struct {
    cConcernRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCConcernService(pr repository.Repository, nu number.NumberUtil) CConcernService {
    return &cConcernService{
        cConcernRepository :pr,
        numberUtil :nu,
    }
}

func (ss cConcernService) CreateCConcern(ctx *context.Context, cConcern *model.CConcern) (*model.CConcern, *errordef.LogicError) {
    cConcern.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cConcern.CreateFunction = "CreateCConcern"
    cConcern.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cConcern.UpdateFunction = "CreateCConcern"
    cConcernRepository := ss.cConcernRepository
    created, err := cConcernRepository.CreateCConcern(ctx, cConcern)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cConcernService) UpdateCConcern(ctx *context.Context, cConcern *model.CConcern) (*model.CConcern, *errordef.LogicError) {
    cConcern.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cConcern.UpdateFunction = "UpdateCConcern"
    cConcernRepository := ss.cConcernRepository
    results, err := cConcernRepository.GetCConcernWithKey(ctx, cConcern.ConcernCd, cConcern.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cConcernRepository.UpdateCConcern(ctx, cConcern)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cConcernService) DeleteCConcern(ctx *context.Context, concernCd int, languageCd int) *errordef.LogicError {
    cConcernRepository := ss.cConcernRepository
    results, err := cConcernRepository.GetCConcernWithKey(ctx, concernCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cConcernRepository.DeleteCConcern(ctx, concernCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cConcernService) GetCConcernWithKey(ctx *context.Context, concernCd int,languageCd int) ([]*model.CConcern, *errordef.LogicError) {
    cConcernRepository := ss.cConcernRepository
    cConcern, err := cConcernRepository.GetCConcernWithKey(ctx, concernCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cConcern, nil
}

func (ss cConcernService) GetCConcernWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CConcern, *errordef.LogicError) {
    cConcernRepository := ss.cConcernRepository
    cConcern, err := cConcernRepository.GetCConcernWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cConcern, nil
}

