package ewsxml

import (
	"encoding/xml"
	"fmt"
	"time"
)

type TimeWindow struct {
	StartTime time.Time `xml:"StartTime"`
	EndTime   time.Time `xml:"EndTime"`
}

type TimeZone struct {
	Bias         int          `xml:"Bias"`
	StandardTime TimeZoneTime `xml:"StandardTime"`
	DaylightTime TimeZoneTime `xml:"DaylightTime"`
}

type TimeZoneTime struct {
	Bias      int    `xml:"Bias"`
	Time      string `xml:"Time"`
	DayOrder  int16  `xml:"DayOrder"`
	Month     int16  `xml:"Month"`
	DayOfWeek string `xml:"DayOfWeek"`
	Year      string `xml:",omitempty"`
}

type TimeZoneDefinition struct {
	Id   string `xml:",attr"`
	Name string `xml:",attr"`
}

type Time string

func (t Time) ToTime() (time.Time, error) {
	offset, err := getRFC3339Offset(time.Now())
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(time.RFC3339, string(t)+offset)
}

// return RFC3339 formatted offset, ex: +03:00 -03:30
func getRFC3339Offset(t time.Time) (string, error) {
	_, offset := t.Zone()
	i := int(float32(offset) / 36)

	sign := "+"
	if i < 0 {
		i = -i
		sign = "-"
	}
	hour := i / 100
	min := i % 100
	min = (60 * min) / 100

	return fmt.Sprintf("%s%02d:%02d", sign, hour, min), nil
}

type Minutes time.Duration

func (m Minutes) Duration() time.Duration { return time.Duration(m) }

func (m Minutes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.Duration()/time.Minute, start)
}

// func (m *Minutes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
// 	var v interface{}
// 	if err := d.DecodeElement(&v, &start); err != nil {
// 		return err
// 	}
//
// 	switch v := v.(type) {
// 	case string:
// 		dur, err := time.ParseDuration(v)
// 		if err == nil {
// 			*m = Minutes(dur)
// 			return nil
// 		}
// 		if int, err2 := strconv.Atoi(v); err2 == nil {
// 			*m = Minutes(time.Duration(int) * time.Minute)
// 		}
// 			return err
// 		}
//
// 		m.Duration = dur
// 		return nil
// 	case int64:
// 		m.Duration = time.Duration(v) * time.Minute
// 	}
// 	return nil
// }
