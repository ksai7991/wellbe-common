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

func (lp persistence) CreateCArea(ctx *context.Context, cArea *model.CArea) (*model.CArea, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "INSERT INTO wellbe_common.c_area(language_cd, area_cd, state_cd, area_name, search_area_name_seo, west_longitude, east_longitude, north_latitude, south_latitude, summary_area_flg, create_datetime, create_function, update_datetime, update_function)"
    script = script + " VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cArea.LanguageCd,
                            cArea.AreaCd,
                            cArea.StateCd,
                            cArea.AreaName,
                            cArea.SearchAreaNameSeo,
                            cArea.WestLongitude,
                            cArea.EastLongitude,
                            cArea.NorthLatitude,
                            cArea.SouthLatitude,
                            cArea.SummaryAreaFlg,
                            cArea.CreateDatetime,
                            cArea.CreateFunction,
                            cArea.UpdateDatetime,
                            cArea.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cArea, nil
}

func (lp persistence) UpdateCArea(ctx *context.Context, cArea *model.CArea) (*model.CArea, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "UPDATE wellbe_common.c_area "
    script = script + "SET state_cd = $3, area_name = $4, search_area_name_seo = $5, west_longitude = $6, east_longitude = $7, north_latitude = $8, south_latitude = $9, summary_area_flg = $10, create_datetime = $11, create_function = $12, update_datetime = $13, update_function = $14 "
    script = script + "WHERE language_cd = $1 and area_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script,
                            cArea.LanguageCd,
                            cArea.AreaCd,
                            cArea.StateCd,
                            cArea.AreaName,
                            cArea.SearchAreaNameSeo,
                            cArea.WestLongitude,
                            cArea.EastLongitude,
                            cArea.NorthLatitude,
                            cArea.SouthLatitude,
                            cArea.SummaryAreaFlg,
                            cArea.CreateDatetime,
                            cArea.CreateFunction,
                            cArea.UpdateDatetime,
                            cArea.UpdateFunction,
                            ); err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return nil, nil
    })
    if err != nil {
        return nil, err
    }
    return cArea, nil
}

func (lp persistence) DeleteCArea(ctx *context.Context, languageCd int, areaCd int) *errordef.LogicError {
    logger := log.GetLogger()
    defer logger.Sync()
    script := "DELETE FROM wellbe_common.c_area WHERE language_cd = $1 and area_cd = $2"
    _, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        if _, err := tx.Exec(script, languageCd, areaCd); err != nil {
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

func (lp persistence) GetCAreaWithKey(ctx *context.Context, languageCd int,areaCd int) ([]*model.CArea, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cAreas []*model.CArea
    script := "SELECT language_cd, area_cd, state_cd, area_name, search_area_name_seo, west_longitude, east_longitude, north_latitude, south_latitude, summary_area_flg, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_area WHERE language_cd = $1 and area_cd = $2 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, languageCd,areaCd)
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
        cArea := &model.CArea{}
        var areaName sql.NullString
        var searchAreaNameSeo sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cArea.LanguageCd, &cArea.AreaCd, &cArea.StateCd, &areaName, &searchAreaNameSeo, &cArea.WestLongitude, &cArea.EastLongitude, &cArea.NorthLatitude, &cArea.SouthLatitude, &cArea.SummaryAreaFlg, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if areaName.Valid {
            cArea.AreaName = areaName.String
        }
        if searchAreaNameSeo.Valid {
            cArea.SearchAreaNameSeo = searchAreaNameSeo.String
        }
        if createDatetime.Valid {
            cArea.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cArea.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cArea.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cArea.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cAreas = append(cAreas, cArea)
    }

    return cAreas, nil
}

func (lp persistence) GetCAreaWithLanguageCd(ctx *context.Context, languageCd int) ([]*model.CArea, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cAreas []*model.CArea
    script := "SELECT language_cd, area_cd, state_cd, area_name, search_area_name_seo, west_longitude, east_longitude, north_latitude, south_latitude, summary_area_flg, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_area WHERE language_cd = $1 ORDER BY create_datetime"
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
        cArea := &model.CArea{}
        var areaName sql.NullString
        var searchAreaNameSeo sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cArea.LanguageCd, &cArea.AreaCd, &cArea.StateCd, &areaName, &searchAreaNameSeo, &cArea.WestLongitude, &cArea.EastLongitude, &cArea.NorthLatitude, &cArea.SouthLatitude, &cArea.SummaryAreaFlg, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if areaName.Valid {
            cArea.AreaName = areaName.String
        }
        if searchAreaNameSeo.Valid {
            cArea.SearchAreaNameSeo = searchAreaNameSeo.String
        }
        if createDatetime.Valid {
            cArea.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cArea.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cArea.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cArea.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cAreas = append(cAreas, cArea)
    }

    return cAreas, nil
}

func (lp persistence) GetCAreaWithStateCd(ctx *context.Context, stateCd int) ([]*model.CArea, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var cAreas []*model.CArea
    script := "SELECT language_cd, area_cd, state_cd, area_name, search_area_name_seo, west_longitude, east_longitude, north_latitude, south_latitude, summary_area_flg, create_datetime, create_function, update_datetime, update_function FROM wellbe_common.c_area WHERE state_cd = $1 ORDER BY create_datetime"
    rows, err := lp.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(script, stateCd)
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
        cArea := &model.CArea{}
        var areaName sql.NullString
        var searchAreaNameSeo sql.NullString
        var createDatetime sql.NullString
        var createFunction sql.NullString
        var updateDatetime sql.NullString
        var updateFunction sql.NullString
        err := rowsv.Scan(&cArea.LanguageCd, &cArea.AreaCd, &cArea.StateCd, &areaName, &searchAreaNameSeo, &cArea.WestLongitude, &cArea.EastLongitude, &cArea.NorthLatitude, &cArea.SouthLatitude, &cArea.SummaryAreaFlg, &createDatetime, &createFunction, &updateDatetime, &updateFunction)
        if areaName.Valid {
            cArea.AreaName = areaName.String
        }
        if searchAreaNameSeo.Valid {
            cArea.SearchAreaNameSeo = searchAreaNameSeo.String
        }
        if createDatetime.Valid {
            cArea.CreateDatetime = createDatetime.String
        }
        if createFunction.Valid {
            cArea.CreateFunction = createFunction.String
        }
        if updateDatetime.Valid {
            cArea.UpdateDatetime = updateDatetime.String
        }
        if updateFunction.Valid {
            cArea.UpdateFunction = updateFunction.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        cAreas = append(cAreas, cArea)
    }

    return cAreas, nil
}
