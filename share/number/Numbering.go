package number

import (
	"context"
	"fmt"
	"hash/crc32"
	"math/rand"
	"strconv"
	commondb "wellbe-common/share/commondb"
	persistence "wellbe-common/share/commondb/persistence"
	constants "wellbe-common/share/commonsettings/constants"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"
	messages "wellbe-common/share/messages"

	"github.com/google/uuid"
)

type NumberUtil interface {
	GetNumber(string) (string, *errordef.LogicError)
}

type numberUtil struct {
}

func NewNumberUtil() NumberUtil {
	return &numberUtil{}
}

// Get Numberd key
func (nu numberUtil)GetNumber(numberingKey string) (string, *errordef.LogicError) {
	logger := log.GetLogger()
	defer logger.Sync()
	ctx := context.Background()
	db, _ := commondb.DbOpen()
	defer commondb.Rollback(&ctx)
	defer db.Close()
	commonPersistence := persistence.NewCommonPersistence(db)

	// Get current number
	numbering, err := commonPersistence.GetOneNumbering(&ctx, numberingKey)
	if err != nil {
		return "", err
	}
	if numbering == nil {
		numberingMaster, err := commonPersistence.GetOneNumberingMaster(&ctx, numberingKey)
		if err != nil {
			return "", err
		}
		if numberingMaster == nil {
			logger.Error(messages.MESSAGE_EN_NUMBERING_NOTEXISTS)
			return "", &errordef.LogicError{Msg: messages.MESSAGE_EN_NUMBERING_NOTEXISTS, Code: constants.LOGIC_ERROR_CODE_NUMBERING_NOTEXISTS}
		}

		numbering, err = commonPersistence.CreateNumbering(&ctx, numberingMaster)
	}

	// If incremented number is over than max value will error
	if numbering.CurrentValue > numbering.MaxValue {
		logger.Error(messages.MESSAGE_EN_NUMBERING_OVERFLOW)
		return "", &errordef.LogicError{Msg: messages.MESSAGE_EN_NUMBERING_OVERFLOW, Code: constants.LOGIC_ERROR_CODE_NUMBERING_OVERFLOW}
	}

	// If Fixlength is upper than 1, padding with 0 till fixlength
	var result string
	if numbering.FixLength < 1 {
		result = fmt.Sprintf("%01d", numbering.CurrentValue)
	} else {
		s := strconv.Itoa(numbering.FixLength)
		result = fmt.Sprintf("%0"+s+"d", numbering.CurrentValue)
	}

	// Update number
	numbering.CurrentValue = numbering.CurrentValue + 1
	_, err = commonPersistence.UpdateNumbering(&ctx, numbering)
	if err != nil {
		return "", err
	}

	random := GenerateRandomString(constants.RANDOM_NUMBER_DIGIT)

	commondb.Commit(&ctx)
	return GetPrefix(numberingKey) + random + "-" + result, nil
}

func GenerateRandomString(digit int) string {
	uuidId, _ := uuid.NewRandom()
	uuidBytes := []byte(uuidId.String())
	hashValue := crc32.ChecksumIEEE(uuidBytes)
	int64Value := int64(hashValue)
	rand.Seed(int64Value)

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	charsetLength := len(charset)

	randomString := make([]byte, digit)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(charsetLength)]
	}

	return string(randomString)
}

func GetPrefix(key string) string {
	switch key {
	case constants.NUMBERING_KEY_SHOP_NO:
		return constants.NUMBERING_KEY_SHOP_NO_PRE
	case constants.NUMBERING_KEY_BOOKING_NO:
		return constants.NUMBERING_KEY_BOOKING_NO_PRE
	case constants.NUMBERING_KEY_ACCOUNT_NO:
		return constants.NUMBERING_KEY_ACCOUNT_NO_PRE
	case constants.NUMBERING_KEY_CUSTOMER_NO:
		return constants.NUMBERING_KEY_CUSTOMER_NO_PRE
	case constants.NUMBERING_KEY_SHOP_CONTACT_NO:
		return constants.NUMBERING_KEY_SHOP_CONTACT_NO_PRE
	default:
		return ""
	}

}