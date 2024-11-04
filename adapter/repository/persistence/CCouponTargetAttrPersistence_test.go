package persistence

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "context"
    env "wellbe-common/settings/env"
    model "wellbe-common/domain/model"
    repository "wellbe-common/adapter/repository"
)

func TestCCouponTargetAttrCreate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCouponTargetAttrPersistence := NewPersistence(tr)
    ctx := context.Background()
    _, err := cCouponTargetAttrPersistence.CreateCCouponTargetAttr(&ctx, &model.CCouponTargetAttr{
                                                        CouponTargetAttrCd: 0,
                                                        LanguageCd: 1,
                                                        CouponTargetAttrName: "dummy-CouponTargetAttrName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    cCouponTargetAttrs, _ := cCouponTargetAttrPersistence.GetCCouponTargetAttrWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCouponTargetAttrs[0].CouponTargetAttrCd, 0)
    assert.Equal(t, cCouponTargetAttrs[0].LanguageCd, 1)
    assert.Equal(t, cCouponTargetAttrs[0].CouponTargetAttrName, "dummy-CouponTargetAttrName")
    assert.Equal(t, cCouponTargetAttrs[0].CreateDatetime, "dummy-CreateDatetime")
    assert.Equal(t, cCouponTargetAttrs[0].CreateFunction, "dummy-CreateFunction")
    assert.Equal(t, cCouponTargetAttrs[0].UpdateDatetime, "dummy-UpdateDatetime")
    assert.Equal(t, cCouponTargetAttrs[0].UpdateFunction, "dummy-UpdateFunction")
    assert.Nil(t, err)
}

func TestCCouponTargetAttrUpdate(t *testing.T) {
    env.EnvLoad("./../../../")
    db, _ := repository.DbOpen()
    tr := repository.NewTransaction(db)
    defer db.Close()
    cCouponTargetAttrPersistence := NewPersistence(tr)
    ctx := context.Background()
    cCouponTargetAttrPersistence.CreateCCouponTargetAttr(&ctx, &model.CCouponTargetAttr{
                                                        CouponTargetAttrCd: 0,
                                                        LanguageCd: 1,
                                                        CouponTargetAttrName: "dummy-CouponTargetAttrName",
                                                        CreateDatetime: "dummy-CreateDatetime",
                                                        CreateFunction: "dummy-CreateFunction",
                                                        UpdateDatetime: "dummy-UpdateDatetime",
                                                        UpdateFunction: "dummy-UpdateFunction",
                                                        })
    _, err := cCouponTargetAttrPersistence.UpdateCCouponTargetAttr(&ctx, &model.CCouponTargetAttr{
                                                        CouponTargetAttrCd: 0,
                                                        LanguageCd: 1,
                                                        CouponTargetAttrName: "dummy-CouponTargetAttrName2",
                                                        CreateDatetime: "dummy-CreateDatetime2",
                                                        CreateFunction: "dummy-CreateFunction2",
                                                        UpdateDatetime: "dummy-UpdateDatetime2",
                                                        UpdateFunction: "dummy-UpdateFunction2",
                                                        })
    cCouponTargetAttrs, _ := cCouponTargetAttrPersistence.GetCCouponTargetAttrWithKey(&ctx, 0, 1)
    tr.Rollback(&ctx)
    assert.Equal(t, cCouponTargetAttrs[0].CouponTargetAttrName, "dummy-CouponTargetAttrName2")
    assert.Equal(t, cCouponTargetAttrs[0].CreateDatetime, "dummy-CreateDatetime2")
    assert.Equal(t, cCouponTargetAttrs[0].CreateFunction, "dummy-CreateFunction2")
    assert.Equal(t, cCouponTargetAttrs[0].UpdateDatetime, "dummy-UpdateDatetime2")
    assert.Equal(t, cCouponTargetAttrs[0].UpdateFunction, "dummy-UpdateFunction2")
    assert.Nil(t, err)
}
