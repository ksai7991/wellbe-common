package persistence

import (
    model "wellbe-common/domain/model"
    errordef "wellbe-common/share/errordef"
    log "wellbe-common/share/log"
    constants "wellbe-common/share/commonsettings/constants"

    _ "github.com/lib/pq"
    "database/sql"
    "context"
)

func (lp persistence) CreateCCouponTargetAttr(ctx *context.Context, cCouponTargetAttr *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_coupon_target_attr(coupon_target_attr_cd, language_cd, coupon_target_attr_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCouponTargetAttr.CouponTargetAttrCd,
                            cCouponTargetAttr.LanguageCd,
                            cCouponTargetAttr.CouponTargetAttrName,
                            cCouponTargetAttr.CreateDatetime,
                            cCouponTargetAttr.CreateFunction,
                            cCouponTargetAttr.UpdateDatetime,
                            cCouponTargetAttr.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCouponTargetAttr, nil
}

func (lp persistence) UpdateCCouponTargetAttr(ctx *context.Context, cCouponTargetAttr *model.CCouponTargetAttr) (*model.CCouponTargetAttr, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_coupon_target_attr "
    script = script + "SET coupon_target_attr_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE coupon_target_attr_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCouponTargetAttr.CouponTargetAttrCd,
                            cCouponTargetAttr.LanguageCd,
                            cCouponTargetAttr.CouponTargetAttrName,
                            cCouponTargetAttr.CreateDatetime,
                            cCouponTargetAttr.CreateFunction,
                            cCouponTargetAttr.UpdateDatetime,
                            cCouponTargetAttr.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCouponTargetAttr, nil
}

func (lp persistence) DeleteCCouponTargetAttr(ctx *context.Context, couponTargetAttrCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_coupon_target_attr WHERE coupon_target_attr_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, couponTargetAttrCd, languageCd); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return err
    }
    return nil
}

func (lp persistence) GetCCouponTargetAttrWithKey(ctx *context.Context, couponTargetAttrCd int,languageCd int) ([]*model.CCouponTargetAttr, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCouponTargetAttrs []*model.CCouponTargetAttr
    script := "SELECT coupon_target_attr_cd, language_cd, coupon_target_attr_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_coupon_target_attr WHERE coupon_target_attr_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, couponTargetAttrCd,languageCd)
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return rows, nil
    })
    rowsv, _ := rows.(*sql.Rows)
    if err != nil {
        logger.Error(err.Error())
        return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    defer rowsv.Close()
    for rowsv.Next() {
        cCouponTargetAttr := &model.CCouponTargetAttr{}
        var couponTargetAttrName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCouponTargetAttr.CouponTargetAttrCd, &cCouponTargetAttr.LanguageCd, &couponTargetAttrName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if couponTargetAttrName.Valid {
            cCouponTargetAttr.CouponTargetAttrName = couponTargetAttrName.String
        }
        if createDatetime.Valid {
            cCouponTargetAttr.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCouponTargetAttr.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCouponTargetAttr.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCouponTargetAttr.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCouponTargetAttrs = append(cCouponTargetAttrs, cCouponTargetAttr)
    }

    return cCouponTargetAttrs, nil
}

func (lp persistence) GetCCouponTargetAttrWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCouponTargetAttr, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCouponTargetAttrs []*model.CCouponTargetAttr
    script := "SELECT coupon_target_attr_cd, language_cd, coupon_target_attr_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_coupon_target_attr WHERE language_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, languageCd)
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return rows, nil
    })
    rowsv, _ := rows.(*sql.Rows)
    if err != nil {
        logger.Error(err.Error())
        return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    defer rowsv.Close()
    for rowsv.Next() {
        cCouponTargetAttr := &model.CCouponTargetAttr{}
        var couponTargetAttrName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCouponTargetAttr.CouponTargetAttrCd, &cCouponTargetAttr.LanguageCd, &couponTargetAttrName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if couponTargetAttrName.Valid {
            cCouponTargetAttr.CouponTargetAttrName = couponTargetAttrName.String
        }
        if createDatetime.Valid {
            cCouponTargetAttr.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCouponTargetAttr.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCouponTargetAttr.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCouponTargetAttr.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCouponTargetAttrs = append(cCouponTargetAttrs, cCouponTargetAttr)
    }

    return cCouponTargetAttrs, nil
}
