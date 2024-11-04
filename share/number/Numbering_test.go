package number

import (
	"context"
	"testing"
	env "wellbe-common/settings/env"
	commondb "wellbe-common/share/commondb"
	persistence "wellbe-common/share/commondb/persistence"
	model "wellbe-common/share/commonmodel"
	constants "wellbe-common/share/commonsettings/constants"
	log "wellbe-common/share/log"
	messages "wellbe-common/share/messages"

	"github.com/stretchr/testify/assert"
)

func TestGetNumber(t *testing.T) {
	env.EnvLoad("./../../")
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	commonPersistence := persistence.NewCommonPersistence(db)
	ctx := context.Background()
	commonPersistence.CreateNumbering(&ctx, &model.Numbering{NumberingKey: "dummykey", 
																InitialValue: 10,
																CurrentValue: 12,
																MaxValue: 12,
																FixLength: 5,
																CreateDatetime: "dummy-CreateDatetime",
																CreateFunc: "dummy-CreateFunc",
																UpdateDatetime: "dummy-UpdateDatetime",
																UpdateFunc: "dummy-UpdateFunc"})
	commondb.Commit(&ctx)
	db.Close()
	nu := NewNumberUtil()
	numberedkey, err := nu.GetNumber("dummykey")
	ctx = context.Background()
	db, _ = commondb.DbOpen()
	commonPersistence = persistence.NewCommonPersistence(db)
	commonPersistence.DeleteNumbering(&ctx, "dummykey")
	commondb.Commit(&ctx)
	db.Close()
	assert.Nil(t, err)
	assert.Equal(t, numberedkey[4:], "00012")
}

func TestGetNumberNotExists(t *testing.T) {
	env.EnvLoad("./../../")
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	commonPersistence := persistence.NewCommonPersistence(db)
	ctx := context.Background()
	commonPersistence.CreateNumbering(&ctx, &model.Numbering{NumberingKey: "dummykey", 
																InitialValue: 10,
																CurrentValue: 12,
																MaxValue: 12,
																FixLength: 5,
																CreateDatetime: "dummy-CreateDatetime",
																CreateFunc: "dummy-CreateFunc",
																UpdateDatetime: "dummy-UpdateDatetime",
																UpdateFunc: "dummy-UpdateFunc"})
	commondb.Commit(&ctx)
	db.Close()
	nu := NewNumberUtil()
	numberedkey, err := nu.GetNumber("dummykey2")
	ctx = context.Background()
	db, _ = commondb.DbOpen()
	commonPersistence = persistence.NewCommonPersistence(db)
	commonPersistence.DeleteNumbering(&ctx, "dummykey")
	commondb.Commit(&ctx)
	db.Close()
	assert.NotNil(t, err)
	assert.Equal(t, err.Msg, messages.MESSAGE_EN_NUMBERING_NOTEXISTS)
	assert.Equal(t, err.Code, constants.LOGIC_ERROR_CODE_NUMBERING_NOTEXISTS)
	assert.Equal(t, numberedkey, "")
}

func TestGetNumberFromMaster(t *testing.T) {
	env.EnvLoad("./../../")
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	commonPersistence := persistence.NewCommonPersistence(db)
	ctx := context.Background()
	commonPersistence.CreateNumbering(&ctx, &model.Numbering{NumberingKey: "dummykey", 
																InitialValue: 10,
																CurrentValue: 12,
																MaxValue: 12,
																FixLength: 5,
																CreateDatetime: "dummy-CreateDatetime",
																CreateFunc: "dummy-CreateFunc",
																UpdateDatetime: "dummy-UpdateDatetime",
																UpdateFunc: "dummy-UpdateFunc"})
	commonPersistence.CreateNumberingMaster(&ctx, &model.Numbering{NumberingKey: "dummykey2", 
																InitialValue: 12,
																CurrentValue: 14,
																MaxValue: 15,
																FixLength: 6,
																CreateDatetime: "dummy-CreateDatetime",
																CreateFunc: "dummy-CreateFunc",
																UpdateDatetime: "dummy-UpdateDatetime",
																UpdateFunc: "dummy-UpdateFunc"})
	commondb.Commit(&ctx)
	db.Close()
	nu := NewNumberUtil()
	numberedkey, err := nu.GetNumber("dummykey2")
	ctx = context.Background()
	db, _ = commondb.DbOpen()
	commonPersistence = persistence.NewCommonPersistence(db)
	commonPersistence.DeleteNumbering(&ctx, "dummykey")
	commonPersistence.DeleteNumbering(&ctx, "dummykey2")
	commonPersistence.DeleteNumberingMaster(&ctx, "dummykey2")
	commondb.Commit(&ctx)
	db.Close()
	assert.Nil(t, err)
	assert.Equal(t, numberedkey[4:], "000014")
}

func TestGetNumberOverflow(t *testing.T) {
	env.EnvLoad("./../../")
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	commonPersistence := persistence.NewCommonPersistence(db)
	ctx := context.Background()
	commonPersistence.CreateNumbering(&ctx, &model.Numbering{NumberingKey: "dummykey", 
																InitialValue: 10,
																CurrentValue: 13,
																MaxValue: 12,
																FixLength: 5,
																CreateDatetime: "dummy-CreateDatetime",
																CreateFunc: "dummy-CreateFunc",
																UpdateDatetime: "dummy-UpdateDatetime",
																UpdateFunc: "dummy-UpdateFunc"})
	commondb.Commit(&ctx)
	db.Close()
	nu := NewNumberUtil()
	numberedkey, err := nu.GetNumber("dummykey")
	ctx = context.Background()
	db, _ = commondb.DbOpen()
	commonPersistence = persistence.NewCommonPersistence(db)
	commonPersistence.DeleteNumbering(&ctx, "dummykey")
	commondb.Commit(&ctx)
	db.Close()
	assert.NotNil(t, err)
	assert.Equal(t, err.Msg, messages.MESSAGE_EN_NUMBERING_OVERFLOW)
	assert.Equal(t, err.Code, constants.LOGIC_ERROR_CODE_NUMBERING_OVERFLOW)
	assert.Equal(t, numberedkey, "")
}

func TestGetNumberNotFixedlength(t *testing.T) {
	env.EnvLoad("./../../")
	logger := log.GetLogger()
	defer logger.Sync()
	db, _ := commondb.DbOpen()
	commonPersistence := persistence.NewCommonPersistence(db)
	ctx := context.Background()
	commonPersistence.CreateNumbering(&ctx, &model.Numbering{NumberingKey: "dummykey", 
																InitialValue: 10,
																CurrentValue: 12,
																MaxValue: 12,
																FixLength: 0,
																CreateDatetime: "dummy-CreateDatetime",
																CreateFunc: "dummy-CreateFunc",
																UpdateDatetime: "dummy-UpdateDatetime",
																UpdateFunc: "dummy-UpdateFunc"})
	commondb.Commit(&ctx)
	db.Close()
	nu := NewNumberUtil()
	numberedkey, err := nu.GetNumber("dummykey")
	ctx = context.Background()
	db, _ = commondb.DbOpen()
	commonPersistence = persistence.NewCommonPersistence(db)
	commonPersistence.DeleteNumbering(&ctx, "dummykey")
	commondb.Commit(&ctx)
	db.Close()
	assert.Nil(t, err)
	assert.Equal(t, numberedkey[4:], "12")
}

func TestGenerateRandomString(t *testing.T) {
	env.EnvLoad("./../../")
	logger := log.GetLogger()
	defer logger.Sync()
	random1 := GenerateRandomString(constants.RANDOM_NUMBER_DIGIT)
	random2 := GenerateRandomString(constants.RANDOM_NUMBER_DIGIT)
	assert.NotNil(t, random1)
	assert.NotNil(t, random2)
	assert.NotEqual(t, random1, random2)
	assert.Equal(t, constants.RANDOM_NUMBER_DIGIT, len(random1))
}

func TestGetPrefix(t *testing.T) {
	env.EnvLoad("./../../")
	logger := log.GetLogger()
	defer logger.Sync()
	pre := GetPrefix(constants.NUMBERING_KEY_SHOP_NO)
	assert.Equal(t, constants.NUMBERING_KEY_SHOP_NO_PRE, pre)
	pre = GetPrefix(constants.NUMBERING_KEY_BOOKING_NO)
	assert.Equal(t, constants.NUMBERING_KEY_BOOKING_NO_PRE, pre)
	pre = GetPrefix(constants.NUMBERING_KEY_ACCOUNT_NO)
	assert.Equal(t, constants.NUMBERING_KEY_ACCOUNT_NO_PRE, pre)
	pre = GetPrefix(constants.NUMBERING_KEY_CUSTOMER_NO)
	assert.Equal(t, constants.NUMBERING_KEY_CUSTOMER_NO_PRE, pre)
	pre = GetPrefix(constants.NUMBERING_KEY_SHOP_CONTACT_NO)
	assert.Equal(t, constants.NUMBERING_KEY_SHOP_CONTACT_NO_PRE, pre)
}