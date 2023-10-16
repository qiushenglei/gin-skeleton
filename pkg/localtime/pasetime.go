package localtime

import "time"

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (l *LocalTime) UnmarshalJSON(data []byte) error {

	// 空值不进行解析
	if len(data) == 2 {
		*l = LocalTime(time.Time{})
		return nil
	}

	// 这里居然有""两个双引号
	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*l = LocalTime(now)
	return err
}

func (l *LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(*l).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}
