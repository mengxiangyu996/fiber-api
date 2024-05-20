package datetime

import (
	"database/sql/driver"
	"errors"
	"time"
)

// 日期时间
type DateTime struct {
	time.Time
}

// 编码为自定义的Json格式
func (t DateTime) MarshalJSON() ([]byte, error) {

	// 时间为零返回null
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte("\"" + t.Format("2006-01-02 15:04:05") + "\""), nil
}

// 将Json格式解码
func (t *DateTime) UnmarshalJSON(data []byte) error {

	var err error

	if len(data) == 2 || string(data) == "null" {
		return err
	}

	var now time.Time

	// 自定义格式解析
	if now, err = time.ParseInLocation("2006-01-02 15:04:05", string(data), time.Local); err == nil {
		*t = DateTime{now}
		return err
	}

	// 带引号的自定义格式解析
	if now, err = time.ParseInLocation("\"2006-01-02 15:04:05\"", string(data), time.Local); err == nil {
		*t = DateTime{now}
		return err
	}

	// 默认格式解析
	if now, err = time.ParseInLocation(time.RFC3339, string(data), time.Local); err == nil {
		*t = DateTime{now}
		return err
	}

	if now, err = time.ParseInLocation("\""+time.RFC3339+"\"", string(data), time.Local); err == nil {
		*t = DateTime{now}
		return err
	}

	return err
}

// 转换为数据库值
func (t DateTime) Value() (driver.Value, error) {

	if t.IsZero() {
		return nil, nil
	}

	return t.Time, nil
}

// 数据库值转换为Datetime
func (t *DateTime) Scan(value interface{}) error {

	if val, ok := value.(time.Time); ok {
		*t = DateTime{Time: val}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}

// 返回字符串
func (t DateTime) String() string {
	return t.Format("2006-01-02 15:04:05")
}
