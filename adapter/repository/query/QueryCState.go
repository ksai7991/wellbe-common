package query

import (
	querySql "wellbe-common/adapter/repository/sql"
	"wellbe-common/domain/model"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"

	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func (q Query) QueryCState(ctx *context.Context, state_name string) ([]*model.CState, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    var states []*model.CState
    rows, err := q.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        rows, err := tx.Query(querySql.GetCState, state_name)
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
        cState := &model.CState{}
        var stateName sql.NullString
        var stateCdIso sql.NullString
        var timezoneIana sql.NullString
        err := rowsv.Scan(&cState.StateCd, &cState.LanguageCd, &cState.CountryCd, &stateName, &stateCdIso, &timezoneIana)
        if stateName.Valid {
            cState.StateName = stateName.String
        }
        if stateCdIso.Valid {
            cState.StateCdIso = stateCdIso.String
        }
        if timezoneIana.Valid {
            cState.TimezoneIana = timezoneIana.String
        }
        if err != nil {
            logger.Error(err.Error())
            return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
        }
        states = append(states, cState)
    }

    return states, nil
}