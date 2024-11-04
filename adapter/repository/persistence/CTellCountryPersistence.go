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

func (lp persistence) CreateCTellCountry(ctx *context.Context, cTellCountry *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_tell_country(language_cd, tell_country_cd, country_name, country_no, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cTellCountry.LanguageCd,
                            cTellCountry.TellCountryCd,
                            cTellCountry.CountryName,
                            cTellCountry.CountryNo,
                            cTellCountry.CreateDatetime,
                            cTellCountry.CreateFunction,
                            cTellCountry.UpdateDatetime,
                            cTellCountry.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cTellCountry, nil
}

func (lp persistence) UpdateCTellCountry(ctx *context.Context, cTellCountry *model.CTellCountry) (*model.CTellCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_tell_country "
    script = script + "SET country_name = $3, country_no = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE language_cd = $1 and tell_country_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cTellCountry.LanguageCd,
                            cTellCountry.TellCountryCd,
                            cTellCountry.CountryName,
                            cTellCountry.CountryNo,
                            cTellCountry.CreateDatetime,
                            cTellCountry.CreateFunction,
                            cTellCountry.UpdateDatetime,
                            cTellCountry.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cTellCountry, nil
}

func (lp persistence) DeleteCTellCountry(ctx *context.Context, languageCd int, tellCountryCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_tell_country WHERE language_cd = $1 and tell_country_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, languageCd, tellCountryCd); err != nil {
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

func (lp persistence) GetCTellCountryWithKey(ctx *context.Context, languageCd int,tellCountryCd int) ([]*model.CTellCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cTellCountrys []*model.CTellCountry
    script := "SELECT language_cd, tell_country_cd, country_name, country_no, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_tell_country WHERE language_cd = $1 and tell_country_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, languageCd,tellCountryCd)
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
        cTellCountry := &model.CTellCountry{}
        var countryName sql.NullString
        var countryNo sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cTellCountry.LanguageCd, &cTellCountry.TellCountryCd, &countryName, &countryNo, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if countryName.Valid {
            cTellCountry.CountryName = countryName.String
        }
        if countryNo.Valid {
            cTellCountry.CountryNo = countryNo.String
        }
        if createDatetime.Valid {
            cTellCountry.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cTellCountry.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cTellCountry.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cTellCountry.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cTellCountrys = append(cTellCountrys, cTellCountry)
    }

    return cTellCountrys, nil
}

func (lp persistence) GetCTellCountryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CTellCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cTellCountrys []*model.CTellCountry
    script := "SELECT language_cd, tell_country_cd, country_name, country_no, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_tell_country WHERE language_cd = $1 ORDER BY create_datetime"
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
        cTellCountry := &model.CTellCountry{}
        var countryName sql.NullString
        var countryNo sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cTellCountry.LanguageCd, &cTellCountry.TellCountryCd, &countryName, &countryNo, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if countryName.Valid {
            cTellCountry.CountryName = countryName.String
        }
        if countryNo.Valid {
            cTellCountry.CountryNo = countryNo.String
        }
        if createDatetime.Valid {
            cTellCountry.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cTellCountry.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cTellCountry.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cTellCountry.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cTellCountrys = append(cTellCountrys, cTellCountry)
    }

    return cTellCountrys, nil
}
