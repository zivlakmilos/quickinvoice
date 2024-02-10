package quickinvoice

import (
	"fmt"
	"strings"
	"time"
)

type Date time.Time

const layout = time.DateOnly

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		*d = Date(time.Time{})
		return nil
	}

	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	t := time.Time(*d)
	if t.IsZero() {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf(`"%s"`, t.Format(layout))), nil
}

func (d Date) String() string {
	t := time.Time(d)
	return t.Format(layout)
}
