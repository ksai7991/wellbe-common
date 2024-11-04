package repository

import (
	"database/sql"
	"time"

	settings "wellbe-common/settings"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"
)

func DbOpen() (*sql.DB, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	dbDriver, dsn := settings.GetDbSettings()
	db,err := sql.Open(dbDriver,dsn)
	openConnections := settings.GetDbMaxOpenConections()
	idleConnections := settings.GetDbMaxIdleConections()
	lifeTimeMinutes := settings.GetDbMaxLifeTimeMinutes()
	db.SetMaxOpenConns(openConnections)
	db.SetMaxIdleConns(idleConnections)
	db.SetConnMaxLifetime(lifeTimeMinutes*time.Minute)
	if err != nil {
		logger.Error(err.Error())
		return db, &errordef.LogicError{Msg: err.Error(), Code: constants.LOGIC_ERROR_CODE_DBERROR}
	}
	return db, nil
}