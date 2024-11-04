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

func (lp persistence) CreateCCountry(ctx *context.Context, cCountry *model.CCountry) (*model.CCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_country(country_cd, language_cd, country_name, country_cd_iso, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCountry.CountryCd,
                            cCountry.LanguageCd,
                            cCountry.CountryName,
                            cCountry.CountryCdIso,
                            cCountry.CreateDatetime,
                            cCountry.CreateFunction,
                            cCountry.UpdateDatetime,
                            cCountry.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCountry, nil
}

func (lp persistence) UpdateCCountry(ctx *context.Context, cCountry *model.CCountry) (*model.CCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_country "
    script = script + "SET country_name = $3, country_cd_iso = $4, create_datetime = $5, create_function = $6, update_datetime = $7, update_function = $8 "
    script = script + "WHERE country_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cCountry.CountryCd,
                            cCountry.LanguageCd,
                            cCountry.CountryName,
                            cCountry.CountryCdIso,
                            cCountry.CreateDatetime,
                            cCountry.CreateFunction,
                            cCountry.UpdateDatetime,
                            cCountry.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cCountry, nil
}

func (lp persistence) DeleteCCountry(ctx *context.Context, countryCd int, languageCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_country WHERE country_cd = $1 and language_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, countryCd, languageCd); err != nil {
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

func (lp persistence) GetCCountryWithKey(ctx *context.Context, countryCd int,languageCd int) ([]*model.CCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCountrys []*model.CCountry
    script := "SELECT country_cd, language_cd, country_name, country_cd_iso, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_country WHERE country_cd = $1 and language_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, countryCd,languageCd)
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
        cCountry := &model.CCountry{}
        var countryName sql.NullString
        var countryCdIso sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCountry.CountryCd, &cCountry.LanguageCd, &countryName, &countryCdIso, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if countryName.Valid {
            cCountry.CountryName = countryName.String
        }
        if countryCdIso.Valid {
            cCountry.CountryCdIso = countryCdIso.String
        }
        if createDatetime.Valid {
            cCountry.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCountry.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCountry.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCountry.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCountrys = append(cCountrys, cCountry)
    }

    return cCountrys, nil
}

func (lp persistence) GetCCountryWithCountryCdIso(ctx *context.Context, countryCdIso string) ([]*model.CCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCountrys []*model.CCountry
    script := "SELECT country_cd, language_cd, country_name, country_cd_iso, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_country WHERE country_cd_iso = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, countryCdIso)
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
        cCountry := &model.CCountry{}
        var countryName sql.NullString
        var countryCdIso sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCountry.CountryCd, &cCountry.LanguageCd, &countryName, &countryCdIso, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if countryName.Valid {
            cCountry.CountryName = countryName.String
        }
        if countryCdIso.Valid {
            cCountry.CountryCdIso = countryCdIso.String
        }
        if createDatetime.Valid {
            cCountry.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCountry.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCountry.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCountry.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCountrys = append(cCountrys, cCountry)
    }

    return cCountrys, nil
}

func (lp persistence) GetCCountryWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CCountry, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cCountrys []*model.CCountry
    script := "SELECT country_cd, language_cd, country_name, country_cd_iso, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_country WHERE language_cd = $1 ORDER BY create_datetime"
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
        cCountry := &model.CCountry{}
        var countryName sql.NullString
        var countryCdIso sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cCountry.CountryCd, &cCountry.LanguageCd, &countryName, &countryCdIso, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if countryName.Valid {
            cCountry.CountryName = countryName.String
        }
        if countryCdIso.Valid {
            cCountry.CountryCdIso = countryCdIso.String
        }
        if createDatetime.Valid {
            cCountry.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cCountry.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cCountry.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cCountry.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cCountrys = append(cCountrys, cCountry)
    }

    return cCountrys, nil
}
