package datetime

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	constants "wellbe-common/share/commonsettings/constants"
	"wellbe-common/share/commonsettings/constants/code/cWeekDay"
	errordef "wellbe-common/share/errordef"
	log "wellbe-common/share/log"
	"wellbe-common/share/messages"
)

func GetCurrentDateTime() time.Time {
	logger := log.GetLogger()
	defer logger.Sync()
	t := time.Now().UTC()
	tokyo, _ := time.LoadLocation(constants.TIME_LOCATION_ASIA_TOKYO)
	timeTokyo := t.In(tokyo)
	return timeTokyo
}

func FormatDateTime2DBDateTimeMillSecString(t time.Time) string {
	return t.Format(constants.DATETIME_DB_MILLSEC_FORMAT)
}

func FormatDBDateTimeMillSecString2DateTime(s string) (time.Time, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	t, err := time.Parse(constants.DATETIME_DB_MILLSEC_FORMAT, s)
	if err != nil {
		logger.Error(err.Error())
	}
	return t, err
}

func ConvertTimezone(t time.Time, timeZone string) (time.Time, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	loadedTimezone, err := time.LoadLocation(timeZone)
	if err != nil {
		logger.Error(err.Error())
		return t, err
	}
	convertedT := t.In(loadedTimezone)
	return convertedT, nil
}

func FormatDateTime2DateString(t time.Time) string {
	return t.Format(constants.DATE_FORMAT)
}

func FormatDateTime2DateTimeString(t time.Time) string {
	return t.Format(constants.DATETIME_FORMAT)
}

func FormatDateTime2DateStringYYYYMMDDHyphen(t time.Time) string {
	return t.Format(constants.DATE_FORMAT_HYPHEN)
}

func FormatDateTime2DateTimeStringYYYYMMDDHHMMHyphen(t time.Time) string {
	return t.Format(constants.DATETIME_YYYYMMDDHHMM_FORMAT)
}

func FormatDateTimeStringYYYYMMDDHHMMHyphen2DateTimeWithLocale(s string, timeZone string) (*time.Time, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	ss := strings.Split(s, " ")
	if len(ss) < 2 {
		return nil, fmt.Errorf(fmt.Sprintf(messages.MESSAGE_EN_DATETIME_FORMAT_IS_INVALID, s))
	}
	return FormatDateTimeAndString2DateTime(FormatDateStringYYYYMMDDHyphenDateString(ss[0]), ss[1], timeZone)
}

func FormatDateString2DateTime(s string) (time.Time, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	t, err := time.Parse(constants.DATE_FORMAT, s)
	if err != nil {
		logger.Error(err.Error())
	}
	return t, err
}

func FormatDateString2DateStringYYYYMMDDHyphen(s string) (string, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	t, err := FormatDateString2DateTime(s)
	if err != nil {
		logger.Error(err.Error())
	}
	return t.Format(constants.DATE_FORMAT_HYPHEN), nil
}

func FormatDateStringYYYYMMDDHyphenDateString(s string) (string) {
	logger := log.GetLogger()
	defer logger.Sync()
	return strings.ReplaceAll(s, "-", "")
}

func FormatDateTimeString2DateTime(s string) (time.Time, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	t, err := time.Parse(constants.DATETIME_FORMAT, s)
	if err != nil {
		logger.Error(err.Error())
	}
	return t, err
}

func FormatDateTimeAndString2DateTime(date string, hhcolonmm string, timeZone string) (*time.Time, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	loadedTimezone, err := time.LoadLocation(timeZone)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if hhcolonmm >= "24:00" {
		t, err := FormatDateString2DateTime(date)
		if err != nil {
			return nil, err
		}

		newT := t.AddDate(0,0,1)
		date = FormatDateTime2DateString(newT)

		hhcolonmm, err = subtract24FromTimeString(hhcolonmm)
		if err != nil {
			return nil, err
		}
	}
	t, err := time.ParseInLocation(constants.DATETIME_YYYYMMDDHHCLOLONMM_FORMAT, date + hhcolonmm, loadedTimezone)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &t, err
}

func subtract24FromTimeString(timeString string) (string, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	parts := strings.Split(timeString, ":")
	if len(parts) != 2 {
		logger.Error("Invalid time format")
		return "", errors.New("invalid time format")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		logger.Error("Invalid hour format")
		return "", error(fmt.Errorf("invalid hour format"))
	}

	hour = (hour - 24 + 24) % 24 // -24 をした後、正の値に戻すために +24 し、24で割る
	newTimeString := fmt.Sprintf("%02d:%s", hour, parts[1])
	return newTimeString, nil
}

func FormatDateTimeAndString2DateTimeNotLocale(date string, hhcolonmm string) (*time.Time, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	t, err := time.Parse(constants.DATETIME_YYYYMMDDHHCLOLONMM_FORMAT, date + hhcolonmm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return &t, err
}

func FormatDateTimeAndString2DateTimeString(date string, hhcolonmm string, timeZone string) (string, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	t, err := FormatDateTimeAndString2DateTime(date, hhcolonmm, timeZone)
	if err != nil {
		logger.Error(err.Error())
	}
	return t.Format(constants.DATETIME_YYYYMMDDHHMM_FORMAT), err
}

func FormatDateTimeAndString2DateTimeStringNotLocale(date string, hhcolonmm string) (string, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	t, err := FormatDateTimeAndString2DateTimeNotLocale(date, hhcolonmm)
	if err != nil {
		logger.Error(err.Error())
	}
	return t.Format(constants.DATETIME_YYYYMMDDHHMM_FORMAT), err
}

func FormatHHMMString2Time(s string) (time.Time, error) {
	logger := log.GetLogger()
	defer logger.Sync()
	t, err := time.Parse(constants.TIME_HHMM_FORMAT, s)
	if err != nil {
		logger.Error(err.Error())
	}
	return t, err
}

func FormatTime2HHMMString(t time.Time) (s string) {
	return t.Format(constants.TIME_HHMM_FORMAT)
}

func CalcTimeStringWithInterval(startTime string, interValMinutes int) (endTime string, err error) {
	t, err := FormatHHMMString2Time(startTime)
	if err != nil {
		return "", err
	}

	return FormatTime2HHMMString(CalcMinutesWithInterval(t, interValMinutes)), nil
}

func CalcDateTimeStringWithInterval(startDate string, startTime string, interValMinutes int) (string, string, error) {
	t, err := FormatDateTimeAndString2DateTimeNotLocale(startDate, startTime)
	if err != nil {
		return "", "", err
	}

	endDatetime := CalcMinutesWithInterval(*t, interValMinutes)
	endDate := FormatDateTime2DateString(endDatetime)
	endTime := FormatTime2HHMMString(endDatetime)


	return endDate, endTime, nil
}

func CalcMinutesWithInterval(startTime time.Time, interValMinutes int) (endTime time.Time) {
	return startTime.Add(time.Minute * time.Duration(interValMinutes))
}

func CompareTwoDatetimeWithComparestring(t1 time.Time, t2 time.Time, comparestring string) (bool, *errordef.LogicError) {
	if comparestring == constants.COMPARE_BEFORE {
		return t1.Before(t2), nil
	} else if comparestring == constants.COMPARE_AFTER {
		return t1.After(t2), nil
	} else if comparestring == constants.COMPARE_EQUAL {
		return t1.Equal(t2), nil
	} else {
		return false, &errordef.LogicError{Code: constants.LOGIC_ERROR_CODE_COMPARE_STRING_INVALID, Msg: messages.MESSAGE_EN_COMPARE_STRING_IS_INVALID}
	}
}

func GetWeekDayCd(utcDate time.Time, timeZone string) (int, error) {
	localDate, err := ConvertTimezone(utcDate, timeZone)
	if err != nil {
		return 0, err
	}
	daysOfWeek := [...]int{cWeekDay.SUNDAY, cWeekDay.MONDAY, cWeekDay.TUESDAY, cWeekDay.WEDNESDAY, cWeekDay.THURSDAY, cWeekDay.FRIDAY, cWeekDay.SATURDAY}
	return daysOfWeek[localDate.Weekday()], nil
}

func IsDatetimeIsInTwoTime(utcDate time.Time, fromHHMM_hyphen string, toHHMM_hyphen, timeZone string) (bool, error) {
	localDate, err := ConvertTimezone(utcDate, timeZone)
	if err != nil {
		return false, err
	}
	yyyymmdd := FormatDateTime2DateString(localDate)
	from, err := FormatDateTimeAndString2DateTime(yyyymmdd, fromHHMM_hyphen, timeZone)
	if err != nil {
		return false, err
	}
	to, err := FormatDateTimeAndString2DateTime(yyyymmdd, toHHMM_hyphen, timeZone)
	if err != nil {
		return false, err
	}
	if (localDate.Equal(*from) || localDate.Equal(*to) || localDate.After(*from) && localDate.Before(*to)) {
		return true, nil
	}
	return false, nil
}

func CalculateAge(birthday string) (int, error) {
    // 誕生日をyyyymmdd形式からパース
    birthDate, err := FormatDateString2DateTime(birthday)
    if err != nil {
        return 0, err
    }

    // 現在の日付を取得
    currentDate := GetCurrentDateTime()

    // 年齢を計算
    age := currentDate.Year() - birthDate.Year()

    // 誕生日がまだ来ていない場合、年齢を1つ減らす
    if currentDate.YearDay() < birthDate.YearDay() {
        age--
    }

    return age, nil
}