package model

// Choice of format for the auto extend period.
type AutoExtend1Choice struct {

	// Number of days.
	Days *Number `xml:"Days"`

	// Number of months
	Months *Number `xml:"Mnths"`

	// Number of years.
	Years *Number `xml:"Yrs"`

	// Auto extension end date.
	Date *ISODate `xml:"Dt"`
}

func (a *AutoExtend1Choice) SetDays(value string) {
	a.Days = (*Number)(&value)
}

func (a *AutoExtend1Choice) SetMonths(value string) {
	a.Months = (*Number)(&value)
}

func (a *AutoExtend1Choice) SetYears(value string) {
	a.Years = (*Number)(&value)
}

func (a *AutoExtend1Choice) SetDate(value string) {
	a.Date = (*ISODate)(&value)
}
