package application

import (
    model "wellbe-common/domain/model"
    service "wellbe-common/domain/service"
    repository "wellbe-common/domain/repository"
    
    errordef "wellbe-common/share/errordef"

    "context"
)

type CTreatmentTimeRangeApplication interface {
    GetCTreatmentTimeRangeWithKey(*context.Context, int,int) ([]*model.CTreatmentTimeRange, *errordef.LogicError)
    GetCTreatmentTimeRangeWithLanguageCd(*context.Context, int) ([]*model.CTreatmentTimeRange, *errordef.LogicError)
}

type cTreatmentTimeRangeApplication struct {
    cTreatmentTimeRangeService service.CTreatmentTimeRangeService
    transaction repository.Transaction
}

func NewCTreatmentTimeRangeApplication(ls service.CTreatmentTimeRangeService, tr repository.Transaction) CTreatmentTimeRangeApplication {
    return &cTreatmentTimeRangeApplication{
        cTreatmentTimeRangeService :ls,
        transaction :tr,
    }
}

func (sa cTreatmentTimeRangeApplication) GetCTreatmentTimeRangeWithKey(ctx *context.Context, treatmentTimeCd int,languageCd int) ([]*model.CTreatmentTimeRange, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cTreatmentTimeRangeService.GetCTreatmentTimeRangeWithKey(ctx, treatmentTimeCd,languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}

func (sa cTreatmentTimeRangeApplication) GetCTreatmentTimeRangeWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTreatmentTimeRange, *errordef.LogicError) {
    defer sa.transaction.Rollback(ctx)
    result, err := sa.cTreatmentTimeRangeService.GetCTreatmentTimeRangeWithLanguageCd(ctx, languageCd)
    sa.transaction.Commit(ctx)
    return result, err
}
