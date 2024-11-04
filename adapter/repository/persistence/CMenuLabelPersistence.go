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

func (lp persistence) CreateCMenuLabel(ctx *context.Context, cMenuLabel *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_menu_label(menu_label_cd, language_cd, menu_label_name, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cMenuLabel.MenuLabelCd,
                            cMenuLabel.LanguageCd,
                            cMenuLabel.MenuLabelName,
                            cMenuLabel.CreateDatetime,
                            cMenuLabel.CreateFunction,
                            cMenuLabel.UpdateDatetime,
                            cMenuLabel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cMenuLabel, nil
}

func (lp persistence) UpdateCMenuLabel(ctx *context.Context, cMenuLabel *model.CMenuLabel) (*model.CMenuLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_menu_label "
    script = script + "SET language_cd = $2, menu_label_name = $3, create_datetime = $4, create_function = $5, update_datetime = $6, update_function = $7 "
    script = script + "WHERE menu_label_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cMenuLabel.MenuLabelCd,
                            cMenuLabel.LanguageCd,
                            cMenuLabel.MenuLabelName,
                            cMenuLabel.CreateDatetime,
                            cMenuLabel.CreateFunction,
                            cMenuLabel.UpdateDatetime,
                            cMenuLabel.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cMenuLabel, nil
}

func (lp persistence) DeleteCMenuLabel(ctx *context.Context, menuLabelCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_menu_label WHERE menu_label_cd = $1"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, menuLabelCd); err != nil {
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

func (lp persistence) GetCMenuLabelWithKey(ctx *context.Context, menuLabelCd int) ([]*model.CMenuLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cMenuLabels []*model.CMenuLabel
    script := "SELECT menu_label_cd, language_cd, menu_label_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_menu_label WHERE menu_label_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, menuLabelCd)
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
        cMenuLabel := &model.CMenuLabel{}
        var menuLabelName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cMenuLabel.MenuLabelCd, &cMenuLabel.LanguageCd, &menuLabelName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if menuLabelName.Valid {
            cMenuLabel.MenuLabelName = menuLabelName.String
        }
        if createDatetime.Valid {
            cMenuLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cMenuLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cMenuLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cMenuLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cMenuLabels = append(cMenuLabels, cMenuLabel)
    }

    return cMenuLabels, nil
}

func (lp persistence) GetCMenuLabelWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CMenuLabel, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cMenuLabels []*model.CMenuLabel
    script := "SELECT menu_label_cd, language_cd, menu_label_name, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_menu_label WHERE language_cd = $1 ORDER BY create_datetime"
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
        cMenuLabel := &model.CMenuLabel{}
        var menuLabelName sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cMenuLabel.MenuLabelCd, &cMenuLabel.LanguageCd, &menuLabelName, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if menuLabelName.Valid {
            cMenuLabel.MenuLabelName = menuLabelName.String
        }
        if createDatetime.Valid {
            cMenuLabel.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cMenuLabel.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cMenuLabel.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cMenuLabel.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cMenuLabels = append(cMenuLabels, cMenuLabel)
    }

    return cMenuLabels, nil
}
