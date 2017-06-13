package finance

import (
	"fmt"
	"log"
	"time"
)

// Datetime is a simple time construct.
type Datetime struct {
	Month  int `json:"m,string"`
	Day    int `json:"d,string"`
	Year   int `json:"y,string"`
	Hour   int `json:",omitempty"`
	Minute int `json:",omitempty"`
	Second int `json:",omitempty"`
	t      time.Time
}

// ParseDatetime creates a new instance of Datetime from a string.
func ParseDatetime(s string) Datetime {

	t, err := time.Parse("1/2/2006", s)
	if err != nil {
		t, err = time.Parse("3:04pm", s)
		if err != nil {
			t, err = parseDashedDate(s)
			if err != nil {
				log.Printf("[go-finance] error parsing time: %s", err.Error())
			}
		}
	}
	return NewDatetime(t)
}

// NewDatetime creates a new instance of Datetime.
func NewDatetime(t time.Time) Datetime {

	// Its just a time.
	if t.Year() == 0 {
		hour, min, sec := t.Clock()
		return Datetime{
			Hour:   hour,
			Minute: min,
			Second: sec,
			t:      t,
		}
	}

	// Its a day.
	year, month, day := t.Date()
	return Datetime{
		Month: int(month),
		Day:   day,
		Year:  year,
		t:     t,
	}
}

func (d Datetime) unixTime() string {
	return fmt.Sprintf("%v", d.t.Unix())
}
