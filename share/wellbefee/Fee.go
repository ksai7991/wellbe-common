package wellbefee

import (
	"context"
	"fmt"
	"wellbe-common/share/commondb"
	"wellbe-common/share/commondb/persistence"
	"wellbe-common/share/commonmodel"
	"wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	"wellbe-common/share/log"
	"wellbe-common/share/messages"
)

func GetDefaultFeeMasterWithKey(id string) (*commonmodel.DefaultFeeMaster, *errordef.LogicError) {
    logger := log.GetLogger()
    defer logger.Sync()
	db, _ := commondb.DbOpen()
	ctx := context.Background()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)
	defaultFees, err := commonPersistence.GetDefaultFeeMasterWithKey(&ctx, id)
	if err != nil {
		return nil, err
	}
	if len(defaultFees) == 0 {
		errmessage := fmt.Sprintf(messages.MESSAGE_EN_MASTER_DATA_IS_UNSETUP, "default_fee_master", "Can not get default_fee_master")
		logger.Error(errmessage)
		return nil, &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_MASTER_DATA_IS_UNSETUP, Msg: errmessage}
	}

	return defaultFees[0], nil
}