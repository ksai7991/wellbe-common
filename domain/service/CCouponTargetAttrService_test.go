package service

import (
    "testing"
    "context"
    "github.com/stretchr/testify/assert"
    repository "wellbe-common/domain/repository"
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    number "wellbe-common/share/number"
)

type cCouponTargetAttrMockRepository struct{
    repository.Repository
    FakeCreate func(*model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError)
    FakeUpdate func(*model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError)
    FakeDelete func(int, int) *errordef.LogicError
    FakeGet func(int, int) ([]*model.CCouponTargetAttr, *errordef.LogicError)
}

func (lr cCouponTargetAttrMockRepository) CreateCCouponTargetAttr(ctx *context.Context, cCouponTargetAttr *model.CCouponTargetAttr)  (*model.CCouponTargetAttr, *errordef.LogicError) {
    return lr.FakeCreate(cCouponTargetAttr)
}

func (lr cCouponTargetAttrMockRepository) UpdateCCouponTargetAttr(ctx *context.Context, cCouponTargetAttr *model.CCouponTargetAttr)  (*model.CCouponTargetAttr, *errordef.LogicError) {
    return lr.FakeUpdate(cCouponTargetAttr)
}

func (lr cCouponTargetAttrMockRepository) DeleteCCouponTargetAttr(ctx *context.Context, couponTargetAttrCd int, languageCd int)  *errordef.LogicError {
    return lr.FakeDelete(couponTargetAttrCd, languageCd)
}

func (lr cCouponTargetAttrMockRepository)GetCCouponTargetAttrWithKey(ctx *context.Context, couponTargetAttrCd int, languageCd int)  ([]*model.CCouponTargetAttr, *errordef.LogicError) {
    return lr.FakeGet(couponTargetAttrCd, languageCd)
}


type cCouponTargetAttrMockNumberUtil struct{
    number.NumberUtil
    FakeGetNumber func(string) (string, *errordef.LogicError)
}

func (lr cCouponTargetAttrMockNumberUtil) GetNumber(s string)  (string, *errordef.LogicError) {
    return lr.FakeGetNumber(s)
}

func TestCreateCCouponTargetAttr(t *testing.T) {
    ctx := context.Background()
    repository := &cCouponTargetAttrMockRepository{
        FakeCreate: func(cCouponTargetAttr *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError) {
            return cCouponTargetAttr, nil
        },
    }
    numberUtil := &cCouponTargetAttrMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCouponTargetAttrService := NewCCouponTargetAttrService(repository, numberUtil)
    in_cCouponTargetAttr := new(model.CCouponTargetAttr)
    in_cCouponTargetAttr.CouponTargetAttrCd = 0
    in_cCouponTargetAttr.LanguageCd = 1
    in_cCouponTargetAttr.CouponTargetAttrName = "dummy-CouponTargetAttrName"
    out_cCouponTargetAttr, err := cCouponTargetAttrService.CreateCCouponTargetAttr(&ctx, in_cCouponTargetAttr)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCouponTargetAttr.CouponTargetAttrCd)
    assert.Equal(t, 1, out_cCouponTargetAttr.LanguageCd)
    assert.Equal(t, "dummy-CouponTargetAttrName", out_cCouponTargetAttr.CouponTargetAttrName)
    assert.NotNil(t, out_cCouponTargetAttr.CreateDatetime)
    assert.NotEqual(t, "", out_cCouponTargetAttr.CreateDatetime)
    assert.Equal(t, "CreateCCouponTargetAttr", out_cCouponTargetAttr.CreateFunction)
    assert.NotNil(t, out_cCouponTargetAttr.UpdateDatetime)
    assert.NotEqual(t, "", out_cCouponTargetAttr.UpdateDatetime)
    assert.Equal(t, "CreateCCouponTargetAttr", out_cCouponTargetAttr.UpdateFunction)
}

func TestUpdateCCouponTargetAttr(t *testing.T) {
    ctx := context.Background()
    repository := &cCouponTargetAttrMockRepository{
        FakeUpdate: func(cCouponTargetAttr *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError) {
            return cCouponTargetAttr, nil
        },
        FakeGet: func(couponTargetAttrCd int, languageCd int) ([]*model.CCouponTargetAttr, *errordef.LogicError) {
            return []*model.CCouponTargetAttr{&model.CCouponTargetAttr{CreateDatetime: "dummy"}}, nil
        },
    }
    numberUtil := &cCouponTargetAttrMockNumberUtil{
        FakeGetNumber: func(s string) (string, *errordef.LogicError) {
            return s, nil
        },
    }
    cCouponTargetAttrService := NewCCouponTargetAttrService(repository, numberUtil)
    in_cCouponTargetAttr := new(model.CCouponTargetAttr)
    in_cCouponTargetAttr.CouponTargetAttrCd = 0
    in_cCouponTargetAttr.LanguageCd = 1
    in_cCouponTargetAttr.CouponTargetAttrName = "dummy-CouponTargetAttrName"
    out_cCouponTargetAttr, err := cCouponTargetAttrService.UpdateCCouponTargetAttr(&ctx, in_cCouponTargetAttr)
    assert.Nil(t, err)
    assert.Equal(t, 0, out_cCouponTargetAttr.CouponTargetAttrCd)
    assert.Equal(t, 1, out_cCouponTargetAttr.LanguageCd)
    assert.Equal(t, "dummy-CouponTargetAttrName", out_cCouponTargetAttr.CouponTargetAttrName)
    assert.NotNil(t, out_cCouponTargetAttr.CreateDatetime)
    assert.Equal(t, "", out_cCouponTargetAttr.CreateDatetime)
    assert.Equal(t, "", out_cCouponTargetAttr.CreateFunction)
    assert.NotNil(t, out_cCouponTargetAttr.UpdateDatetime)
    assert.NotEqual(t, "", out_cCouponTargetAttr.UpdateDatetime)
    assert.Equal(t, "UpdateCCouponTargetAttr", out_cCouponTargetAttr.UpdateFunction)
}
