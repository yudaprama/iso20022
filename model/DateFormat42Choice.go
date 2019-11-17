package model

// Choice of formats for a date.
type DateFormat42Choice struct {

	// Year and month.
	YearMonth *ISOYearMonth `xml:"YrMnth"`

	// Year, month and day.
	YearMonthDay *ISODate `xml:"YrMnthDay"`
}

func (d *DateFormat42Choice) SetYearMonth(value string) {
	d.YearMonth = (*ISOYearMonth)(&value)
}

func (d *DateFormat42Choice) SetYearMonthDay(value string) {
	d.YearMonthDay = (*ISODate)(&value)
}
