package migrate

import (
	"database/sql"
	"io/ioutil"
	log "wellbe-common/share/log"

	"github.com/tanimutomo/sqlfile"
)

func Migrate(db *sql.DB, dev string){
    logger := log.GetLogger()
    defer logger.Sync()
	s := sqlfile.New()
	errfile := s.Directory("./dll")
	if errfile != nil {
		logger.Error(errfile.Error())
		return
	}
	_, errexec := s.Exec(db)
	if errexec != nil {
		logger.Error(errexec.Error())
		return
	}
	
	if dev == "dev" {
		all, _ := ioutil.ReadDir("./dll_alter_dev")
		for _, f := range all {
			logger.Info(f.Name())
			dllAlter := sqlfile.New()
			dllAlter.File("./dll_alter_dev/" + f.Name())
			_, err := dllAlter.Exec(db)
			if err != nil {
				// ignore error
				logger.Warn(err.Error())
			}
		}
	}
	errfile = s.Directory("./dml")
	if errfile != nil {
		logger.Error(errfile.Error())
		return
	}
	_, errexec = s.Exec(db)
	if errexec != nil {
		logger.Error(errexec.Error())
		return
	}
	if dev == "dev" {
		errfile = s.Directory("./dml_test")
		if errfile != nil {
			logger.Error(errfile.Error())
			return
		}
		_, errexec = s.Exec(db)
		if errexec != nil {
			logger.Error(errexec.Error())
			return
		}
	}
}