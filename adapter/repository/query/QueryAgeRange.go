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

func (q Query) QueryAgeRange(ctx *context.Context, language_cd string, age int) (*model.CAgeRange, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
    row, err := q.transaction.DoInTx(ctx, func(tx *sql.Tx) (interface{}, *errordef.LogicError) {
        row := tx.QueryRow(querySql.GetAgeRange, language_cd, age)
        return row, nil
    })
    rowv, _ := row.(*sql.Row)
    if err != nil {
        return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    ageRange := &model.CAgeRange{}
    var ageRangeCd sql.NullInt32
    var languageCd sql.NullInt32
    var ageRangeGender sql.NullString
    errScan := rowv.Scan(&ageRangeCd, &languageCd, &ageRangeGender)
    if errScan != nil {
        logger.Error(errScan.Error())
        return nil, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
    }
    if ageRangeCd.Valid {
        ageRange.AgeRangeCd = int(ageRangeCd.Int32)
    }
    if languageCd.Valid {
        ageRange.LanguageCd = int(languageCd.Int32)
    }
    if ageRangeGender.Valid {
        ageRange.AgeRangeGender = ageRangeGender.String
    }

    return ageRange, nil
}