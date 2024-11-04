package datetime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentDateTime(t *testing.T) {
    datetime := GetCurrentDateTime()
    assert.NotNil(t, datetime)
}

func TestFormatDateTime2DateString(t *testing.T) {
    ti := time.Date(2018, 12, 20, 12, 34, 56, 123456, time.UTC)
    datetime := FormatDateTime2DateString(ti)
    expect := "20181220"
    assert.Equal(t, expect, datetime)
}

func TestFormatDateTime2DateTimeString(t *testing.T) {
    ti := time.Date(2018, 12, 20, 12, 34, 56, 123456, time.UTC)
    datetime := FormatDateTime2DateTimeString(ti)
    expect := "20181220123456"
    assert.Equal(t, expect, datetime)
}
func TestGetCurrentWeekDayCd(t *testing.T) {
    ti := time.Date(2023, 6, 25, 0, 0, 0, 0, time.UTC)
    i, err := GetWeekDayCd(ti, "Asia/Bangkok")
    assert.Nil(t, err)
    assert.Equal(t, 1, i)
}
func TestCalcDateTimeStringWithInterval(t *testing.T) {
    d, tt, err := CalcDateTimeStringWithInterval("20230801", "23:30", 60)
    assert.Nil(t, err)
    assert.Equal(t, "20230802", d)
    assert.Equal(t, "00:30", tt)
}

func TestIsCurrentDatetimeIsInTwoTime(t *testing.T) {
    ti := time.Date(2023, 6, 25, 2, 0, 0, 0, time.UTC)
    b, err := IsDatetimeIsInTwoTime(ti, "09:00", "21:00", "Asia/Bangkok")
    assert.Nil(t, err)
    assert.True(t, b)
    ti = time.Date(2023, 6, 25, 2, 1, 0, 0, time.UTC)
    b, err = IsDatetimeIsInTwoTime(ti, "09:00", "21:00", "Asia/Bangkok")
    assert.Nil(t, err)
    assert.True(t, b)
    ti = time.Date(2023, 6, 25, 14, 0, 0, 0, time.UTC)
    b, err = IsDatetimeIsInTwoTime(ti, "09:00", "21:00", "Asia/Bangkok")
    assert.Nil(t, err)
    assert.True(t, b)
    ti = time.Date(2023, 6, 25, 13, 59, 0, 0, time.UTC)
    b, err = IsDatetimeIsInTwoTime(ti, "09:00", "21:00", "Asia/Bangkok")
    assert.Nil(t, err)
    assert.True(t, b)
    ti = time.Date(2023, 6, 25, 1, 59, 0, 0, time.UTC)
    b, err = IsDatetimeIsInTwoTime(ti, "09:00", "21:00", "Asia/Bangkok")
    assert.Nil(t, err)
    assert.False(t, b)
    ti = time.Date(2023, 6, 25, 14, 1, 0, 0, time.UTC)
    b, err = IsDatetimeIsInTwoTime(ti, "09:00", "21:00", "Asia/Bangkok")
    assert.Nil(t, err)
    assert.False(t, b)

}
func TestFormatDateTimeAndString2DateTime(t *testing.T) {
    tt, err := FormatDateTimeAndString2DateTime("20230801", "24:05", "Asia/Tokyo")
    actt := FormatDateTime2DateTimeStringYYYYMMDDHHMMHyphen(*tt)
    assert.Nil(t, err)
    assert.Equal(t, "2023-08-02 00:05", actt)
}
func TestFormatDateTimeStringYYYYMMDDHHMMHyphen2DateTimeWithLocale1(t *testing.T) {
    tt, err := FormatDateTimeStringYYYYMMDDHHMMHyphen2DateTimeWithLocale("2023-08-01 15:00", "Asia/Tokyo")
    actt := FormatDateTime2DateTimeStringYYYYMMDDHHMMHyphen(tt.UTC())
    assert.Nil(t, err)
    assert.Equal(t, "2023-08-01 06:00", actt)
}
func TestFormatDateTimeStringYYYYMMDDHHMMHyphen2DateTimeWithLocale2(t *testing.T) {
    tt, err := FormatDateTimeStringYYYYMMDDHHMMHyphen2DateTimeWithLocale("2023-08-01 15:00", time.UTC.String())
    actt := FormatDateTime2DateTimeStringYYYYMMDDHHMMHyphen(tt.UTC())
    assert.Nil(t, err)
    assert.Equal(t, "2023-08-01 15:00", actt)
}