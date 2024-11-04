package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCInvoiceStatusCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cInvoiceStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cInvoiceStatusPersistence.CreateCInvoiceStatus(&ctx, &model.CInvoiceStatus{
                                                        InvoiceStatusCd: 0,
                                                        LanguageCd: 1,
                                                        InvoiceStatusName: "dummy-InvoiceStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cInvoiceStatuss, _ := cInvoiceStatusPersistence.GetCInvoiceStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cInvoiceStatuss[0].InvoiceStatusCd, 0)
    assert.Equal(t, cInvoiceStatuss[0].LanguageCd, 1)
    assert.Equal(t, cInvoiceStatuss[0].InvoiceStatusName, "dummy-InvoiceStatusName")
    assert.Equal(t, cInvoiceStatuss[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cInvoiceStatuss[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cInvoiceStatuss[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cInvoiceStatuss[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCInvoiceStatusUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cInvoiceStatusPersistence := NewPersistence(tr)
    ctx := context.Background()
    cInvoiceStatusPersistence.CreateCInvoiceStatus(&ctx, &model.CInvoiceStatus{
                                                        InvoiceStatusCd: 0,
                                                        LanguageCd: 1,
                                                        InvoiceStatusName: "dummy-InvoiceStatusName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cInvoiceStatusPersistence.UpdateCInvoiceStatus(&ctx, &model.CInvoiceStatus{
                                                        InvoiceStatusCd: 0,
                                                        LanguageCd: 1,
                                                        InvoiceStatusName: "dummy-InvoiceStatusName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cInvoiceStatuss, _ := cInvoiceStatusPersistence.GetCInvoiceStatusWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cInvoiceStatuss[0].InvoiceStatusName, "dummy-InvoiceStatusName2")
    assert.Equal(t, cInvoiceStatuss[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cInvoiceStatuss[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cInvoiceStatuss[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cInvoiceStatuss[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
