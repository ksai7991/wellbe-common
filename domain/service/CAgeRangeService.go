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

type CAgeRangeService interface {
    CreateCAgeRange(*context.Context, *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError)
    UpdateCAgeRange(*context.Context, *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError)
    GetCAgeRangeWithKey(*context.Context, int,int) ([]*model.CAgeRange, *errordef.LogicError)
    GetCAgeRangeWithLanguageCd(*context.Context, int) ([]*model.CAgeRange, *errordef.LogicError)
    DeleteCAgeRange(*context.Context, int, int) *errordef.LogicError
}

type cAgeRangeService struct {
    cAgeRangeRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCAgeRangeService(pr repository.Repository, nu number.NumberUtil) CAgeRangeService {
    return &cAgeRangeService{
        cAgeRangeRepository :pr,
        numberUtil :nu,
    }
}

func (ss cAgeRangeService) CreateCAgeRange(ctx *context.Context, cAgeRange *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError) {
    cAgeRange.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cAgeRange.CreateFunction = "CreateCAgeRange"
    cAgeRange.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cAgeRange.UpdateFunction = "CreateCAgeRange"
    cAgeRangeRepository := ss.cAgeRangeRepository
    created, err := cAgeRangeRepository.CreateCAgeRange(ctx, cAgeRange)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cAgeRangeService) UpdateCAgeRange(ctx *context.Context, cAgeRange *model.CAgeRange) (*model.CAgeRange, *errordef.LogicError) {
    cAgeRange.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cAgeRange.UpdateFunction = "UpdateCAgeRange"
    cAgeRangeRepository := ss.cAgeRangeRepository
    results, err := cAgeRangeRepository.GetCAgeRangeWithKey(ctx, cAgeRange.AgeRangeCd, cAgeRange.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cAgeRangeRepository.UpdateCAgeRange(ctx, cAgeRange)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cAgeRangeService) DeleteCAgeRange(ctx *context.Context, ageRangeCd int, languageCd int) *errordef.LogicError {
    cAgeRangeRepository := ss.cAgeRangeRepository
    results, err := cAgeRangeRepository.GetCAgeRangeWithKey(ctx, ageRangeCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cAgeRangeRepository.DeleteCAgeRange(ctx, ageRangeCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cAgeRangeService) GetCAgeRangeWithKey(ctx *context.Context, ageRangeCd int,languageCd int) ([]*model.CAgeRange, *errordef.LogicError) {
    cAgeRangeRepository := ss.cAgeRangeRepository
    cAgeRange, err := cAgeRangeRepository.GetCAgeRangeWithKey(ctx, ageRangeCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cAgeRange, nil
}

func (ss cAgeRangeService) GetCAgeRangeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CAgeRange, *errordef.LogicError) {
    cAgeRangeRepository := ss.cAgeRangeRepository
    cAgeRange, err := cAgeRangeRepository.GetCAgeRangeWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cAgeRange, nil
}

