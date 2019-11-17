package model

// Period as a date, a month or a date time span for which the statement is provided.
type DatePeriod1Choice struct {

	// Time span defined by a specific date.
	Date *ISODate `xml:"Dt"`

	// Time span defined by a month and a year.
	DateMonth *ISOYearMonth `xml:"DtMnth"`

	// Time span defined by a start date and an end date.
	FromDateToDate *Period2 `xml:"FrDtToDt"`
}

func (d *DatePeriod1Choice) SetDate(value string) {
	d.Date = (*ISODate)(&value)
}

func (d *DatePeriod1Choice) SetDateMonth(value string) {
	d.DateMonth = (*ISOYearMonth)(&value)
}

func (d *DatePeriod1Choice) AddFromDateToDate() *Period2 {
	d.FromDateToDate = new(Period2)
	return d.FromDateToDate
}
