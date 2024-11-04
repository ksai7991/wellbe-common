package query

import (
	entity "wellbe-common/adapter/repository/entity"
	querySql "wellbe-common/adapter/repository/sql"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"

	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func (q Query) QueryArea(ctx *context.Context, language_cd string, country_cd string, state_cd string, area_cd string) ([]*entity.Area, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var areas []*entity.Area
    rows, err := q.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(querySql.GetArea, language_cd, country_cd, state_cd, area_cd)
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        return rows, nil
    })
    rowsv, _ := rows.(*sql.Rows)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    defer rowsv.Close()
    for rowsv.Next() {
        area := &entity.Area{}
        var areaCd            sql.NullInt32
        var areaName          sql.NullString
        var searchAreaNameSeo sql.NullString
        var stateCd           sql.NullInt32   
        var stateName         sql.NullString
        var stateCdIso        sql.NullString
        var countryCd         sql.NullInt32   
        var countryName       sql.NullString
        var countryCdIso      sql.NullString
        err := rowsv.Scan(&areaCd, &areaName, &searchAreaNameSeo, &stateCd, &stateName, &stateCdIso, &countryCd, &countryName, &countryCdIso)
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        if areaCd.Valid {
            area.AreaCd = int(areaCd.Int32)
        }
        if areaName.Valid {
            area.AreaName = areaName.String
        }
        if searchAreaNameSeo.Valid {
            area.SearchAreaNameSeo = searchAreaNameSeo.String
        }
        if stateCd.Valid {
            area.StateCd = int(stateCd.Int32)
        }
        if stateName.Valid {
            area.StateName = stateName.String
        }
        if stateCdIso.Valid {
            area.StateCdIso = stateCdIso.String
        }
        if countryCd.Valid {
            area.CountryCd = int(countryCd.Int32)
        }
        if countryName.Valid {
            area.CountryName = countryName.String
        }
        if countryCdIso.Valid {
            area.CountryCdIso = countryCdIso.String
        }
        areas = append(areas, area)
    }

    return areas, nil
}