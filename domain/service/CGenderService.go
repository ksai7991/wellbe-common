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

type CGenderService interface {
    CreateCGender(*context.Context, *model.CGender) (*model.CGender, *errordef.LogicError)
    UpdateCGender(*context.Context, *model.CGender) (*model.CGender, *errordef.LogicError)
    GetCGenderWithKey(*context.Context, int,int) ([]*model.CGender, *errordef.LogicError)
    GetCGenderWithLanguageCd(*context.Context, int) ([]*model.CGender, *errordef.LogicError)
    DeleteCGender(*context.Context, int, int) *errordef.LogicError
}

type cGenderService struct {
    cGenderRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCGenderService(pr repository.Repository, nu number.NumberUtil) CGenderService {
    return &cGenderService{
        cGenderRepository :pr,
        numberUtil :nu,
    }
}

func (ss cGenderService) CreateCGender(ctx *context.Context, cGender *model.CGender) (*model.CGender, *errordef.LogicError) {
    cGender.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cGender.CreateFunction = "CreateCGender"
    cGender.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cGender.UpdateFunction = "CreateCGender"
    cGenderRepository := ss.cGenderRepository
    created, err := cGenderRepository.CreateCGender(ctx, cGender)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cGenderService) UpdateCGender(ctx *context.Context, cGender *model.CGender) (*model.CGender, *errordef.LogicError) {
    cGender.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cGender.UpdateFunction = "UpdateCGender"
    cGenderRepository := ss.cGenderRepository
    results, err := cGenderRepository.GetCGenderWithKey(ctx, cGender.GenderCd, cGender.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cGenderRepository.UpdateCGender(ctx, cGender)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cGenderService) DeleteCGender(ctx *context.Context, genderCd int, languageCd int) *errordef.LogicError {
    cGenderRepository := ss.cGenderRepository
    results, err := cGenderRepository.GetCGenderWithKey(ctx, genderCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cGenderRepository.DeleteCGender(ctx, genderCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cGenderService) GetCGenderWithKey(ctx *context.Context, genderCd int,languageCd int) ([]*model.CGender, *errordef.LogicError) {
    cGenderRepository := ss.cGenderRepository
    cGender, err := cGenderRepository.GetCGenderWithKey(ctx, genderCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cGender, nil
}

func (ss cGenderService) GetCGenderWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CGender, *errordef.LogicError) {
    cGenderRepository := ss.cGenderRepository
    cGender, err := cGenderRepository.GetCGenderWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cGender, nil
}

