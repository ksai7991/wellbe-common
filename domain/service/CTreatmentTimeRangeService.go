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

type CTreatmentTimeRangeService interface {
    CreateCTreatmentTimeRange(*context.Context, *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError)
    UpdateCTreatmentTimeRange(*context.Context, *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError)
    GetCTreatmentTimeRangeWithKey(*context.Context, int,int) ([]*model.CTreatmentTimeRange, *errordef.LogicError)
    GetCTreatmentTimeRangeWithLanguageCd(*context.Context, int) ([]*model.CTreatmentTimeRange, *errordef.LogicError)
    DeleteCTreatmentTimeRange(*context.Context, int, int) *errordef.LogicError
}

type cTreatmentTimeRangeService struct {
    cTreatmentTimeRangeRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCTreatmentTimeRangeService(pr repository.Repository, nu number.NumberUtil) CTreatmentTimeRangeService {
    return &cTreatmentTimeRangeService{
        cTreatmentTimeRangeRepository :pr,
        numberUtil :nu,
    }
}

func (ss cTreatmentTimeRangeService) CreateCTreatmentTimeRange(ctx *context.Context, cTreatmentTimeRange *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError) {
    cTreatmentTimeRange.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTreatmentTimeRange.CreateFunction = "CreateCTreatmentTimeRange"
    cTreatmentTimeRange.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTreatmentTimeRange.UpdateFunction = "CreateCTreatmentTimeRange"
    cTreatmentTimeRangeRepository := ss.cTreatmentTimeRangeRepository
    created, err := cTreatmentTimeRangeRepository.CreateCTreatmentTimeRange(ctx, cTreatmentTimeRange)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cTreatmentTimeRangeService) UpdateCTreatmentTimeRange(ctx *context.Context, cTreatmentTimeRange *model.CTreatmentTimeRange) (*model.CTreatmentTimeRange, *errordef.LogicError) {
    cTreatmentTimeRange.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTreatmentTimeRange.UpdateFunction = "UpdateCTreatmentTimeRange"
    cTreatmentTimeRangeRepository := ss.cTreatmentTimeRangeRepository
    results, err := cTreatmentTimeRangeRepository.GetCTreatmentTimeRangeWithKey(ctx, cTreatmentTimeRange.TreatmentTimeCd, cTreatmentTimeRange.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cTreatmentTimeRangeRepository.UpdateCTreatmentTimeRange(ctx, cTreatmentTimeRange)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cTreatmentTimeRangeService) DeleteCTreatmentTimeRange(ctx *context.Context, treatmentTimeCd int, languageCd int) *errordef.LogicError {
    cTreatmentTimeRangeRepository := ss.cTreatmentTimeRangeRepository
    results, err := cTreatmentTimeRangeRepository.GetCTreatmentTimeRangeWithKey(ctx, treatmentTimeCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cTreatmentTimeRangeRepository.DeleteCTreatmentTimeRange(ctx, treatmentTimeCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cTreatmentTimeRangeService) GetCTreatmentTimeRangeWithKey(ctx *context.Context, treatmentTimeCd int,languageCd int) ([]*model.CTreatmentTimeRange, *errordef.LogicError) {
    cTreatmentTimeRangeRepository := ss.cTreatmentTimeRangeRepository
    cTreatmentTimeRange, err := cTreatmentTimeRangeRepository.GetCTreatmentTimeRangeWithKey(ctx, treatmentTimeCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cTreatmentTimeRange, nil
}

func (ss cTreatmentTimeRangeService) GetCTreatmentTimeRangeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTreatmentTimeRange, *errordef.LogicError) {
    cTreatmentTimeRangeRepository := ss.cTreatmentTimeRangeRepository
    cTreatmentTimeRange, err := cTreatmentTimeRangeRepository.GetCTreatmentTimeRangeWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cTreatmentTimeRange, nil
}

