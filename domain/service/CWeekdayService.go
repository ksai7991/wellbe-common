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

type CWeekdayService interface {
    CreateCWeekday(*context.Context, *model.CWeekday) (*model.CWeekday, *errordef.LogicError)
    UpdateCWeekday(*context.Context, *model.CWeekday) (*model.CWeekday, *errordef.LogicError)
    GetCWeekdayWithKey(*context.Context, int) ([]*model.CWeekday, *errordef.LogicError)
    GetCWeekdayWithLanguageCd(*context.Context, int) ([]*model.CWeekday, *errordef.LogicError)
    DeleteCWeekday(*context.Context, int) *errordef.LogicError
}

type cWeekdayService struct {
    cWeekdayRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCWeekdayService(pr repository.Repository, nu number.NumberUtil) CWeekdayService {
    return &cWeekdayService{
        cWeekdayRepository :pr,
        numberUtil :nu,
    }
}

func (ss cWeekdayService) CreateCWeekday(ctx *context.Context, cWeekday *model.CWeekday) (*model.CWeekday, *errordef.LogicError) {
    cWeekday.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cWeekday.CreateFunction = "CreateCWeekday"
    cWeekday.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cWeekday.UpdateFunction = "CreateCWeekday"
    cWeekdayRepository := ss.cWeekdayRepository
    created, err := cWeekdayRepository.CreateCWeekday(ctx, cWeekday)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cWeekdayService) UpdateCWeekday(ctx *context.Context, cWeekday *model.CWeekday) (*model.CWeekday, *errordef.LogicError) {
    cWeekday.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cWeekday.UpdateFunction = "UpdateCWeekday"
    cWeekdayRepository := ss.cWeekdayRepository
    results, err := cWeekdayRepository.GetCWeekdayWithKey(ctx, cWeekday.WeekdayCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cWeekdayRepository.UpdateCWeekday(ctx, cWeekday)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cWeekdayService) DeleteCWeekday(ctx *context.Context, weekdayCd int) *errordef.LogicError {
    cWeekdayRepository := ss.cWeekdayRepository
    results, err := cWeekdayRepository.GetCWeekdayWithKey(ctx, weekdayCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cWeekdayRepository.DeleteCWeekday(ctx, weekdayCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cWeekdayService) GetCWeekdayWithKey(ctx *context.Context, weekdayCd int) ([]*model.CWeekday, *errordef.LogicError) {
    cWeekdayRepository := ss.cWeekdayRepository
    cWeekday, err := cWeekdayRepository.GetCWeekdayWithKey(ctx, weekdayCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cWeekday, nil
}

func (ss cWeekdayService) GetCWeekdayWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CWeekday, *errordef.LogicError) {
    cWeekdayRepository := ss.cWeekdayRepository
    cWeekday, err := cWeekdayRepository.GetCWeekdayWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cWeekday, nil
}

