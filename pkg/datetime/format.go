package datetime

import "time"

// 将字符串解析为时间
func ParseTimeInLocation(value string, location *time.Location) (Time, error) {

	time, err := time.ParseInLocation("2006-01-02 15:04:05", value, location)

	return Time{Time: time}, err
}

// 将字符串解析为日期
func ParseDateInLocation(value string, location *time.Location) (Date, error) {

	date, err := time.ParseInLocation("2006-01-02", value, location)

	return Date{Time: date}, err
}
