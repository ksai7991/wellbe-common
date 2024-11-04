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

type CTypeOfContactService interface {
    CreateCTypeOfContact(*context.Context, *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError)
    UpdateCTypeOfContact(*context.Context, *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError)
    GetCTypeOfContactWithKey(*context.Context, int,int) ([]*model.CTypeOfContact, *errordef.LogicError)
    GetCTypeOfContactWithLanguageCd(*context.Context, int) ([]*model.CTypeOfContact, *errordef.LogicError)
    DeleteCTypeOfContact(*context.Context, int, int) *errordef.LogicError
}

type cTypeOfContactService struct {
    cTypeOfContactRepository repository.Repository
    numberUtil number.NumberUtil
}

func NewCTypeOfContactService(pr repository.Repository, nu number.NumberUtil) CTypeOfContactService {
    return &cTypeOfContactService{
        cTypeOfContactRepository :pr,
        numberUtil :nu,
    }
}

func (ss cTypeOfContactService) CreateCTypeOfContact(ctx *context.Context, cTypeOfContact *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError) {
    cTypeOfContact.CreateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTypeOfContact.CreateFunction = "CreateCTypeOfContact"
    cTypeOfContact.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTypeOfContact.UpdateFunction = "CreateCTypeOfContact"
    cTypeOfContactRepository := ss.cTypeOfContactRepository
    created, err := cTypeOfContactRepository.CreateCTypeOfContact(ctx, cTypeOfContact)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return created, nil
}

func (ss cTypeOfContactService) UpdateCTypeOfContact(ctx *context.Context, cTypeOfContact *model.CTypeOfContact) (*model.CTypeOfContact, *errordef.LogicError) {
    cTypeOfContact.UpdateDatetime = datetime.FormatDateTime2DBDateTimeMillSecString(datetime.GetCurrentDateTime())
    cTypeOfContact.UpdateFunction = "UpdateCTypeOfContact"
    cTypeOfContactRepository := ss.cTypeOfContactRepository
    results, err := cTypeOfContactRepository.GetCTypeOfContactWithKey(ctx, cTypeOfContact.TypeOfContactCd, cTypeOfContact.LanguageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return nil, &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_UPDATE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_UPDATE}
    }
    updated, err := cTypeOfContactRepository.UpdateCTypeOfContact(ctx, cTypeOfContact)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return updated, nil
}

func (ss cTypeOfContactService) DeleteCTypeOfContact(ctx *context.Context, typeOfContactCd int, languageCd int) *errordef.LogicError {
    cTypeOfContactRepository := ss.cTypeOfContactRepository
    results, err := cTypeOfContactRepository.GetCTypeOfContactWithKey(ctx, typeOfContactCd, languageCd)
    if err != nil {
        return  &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    if len(results) == 0 {
        return &errordef.LogicError{Msg: messages.MESSAGE_EN_NOTEXISTS_DELETE, Code:commonconstants.LOGIC_ERROR_CODE_NOTEXISTS_DELETE}
    }
    err = cTypeOfContactRepository.DeleteCTypeOfContact(ctx, typeOfContactCd, languageCd)
    if err != nil {
        return &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return nil
}

func (ss cTypeOfContactService) GetCTypeOfContactWithKey(ctx *context.Context, typeOfContactCd int,languageCd int) ([]*model.CTypeOfContact, *errordef.LogicError) {
    cTypeOfContactRepository := ss.cTypeOfContactRepository
    cTypeOfContact, err := cTypeOfContactRepository.GetCTypeOfContactWithKey(ctx, typeOfContactCd,languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cTypeOfContact, nil
}

func (ss cTypeOfContactService) GetCTypeOfContactWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTypeOfContact, *errordef.LogicError) {
    cTypeOfContactRepository := ss.cTypeOfContactRepository
    cTypeOfContact, err := cTypeOfContactRepository.GetCTypeOfContactWithLanguageCd(ctx, languageCd)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code:commonconstants.LOGIC_ERROR_CODE_DBERROR}
    }
    return cTypeOfContact, nil
}

