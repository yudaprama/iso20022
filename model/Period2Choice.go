package model

// Choice between date and date-time for the specification of a period.
type Period2Choice struct {

	// Time span defined by a start date and time, and an end date and time.
	FromDateTimeToDateTime *DateTimePeriodDetails `xml:"FrDtTmToDtTm"`

	// Time span defined by a start date and time, and an end date and time.
	FromDateToDate *Period2 `xml:"FrDtToDt"`
}

func (p *Period2Choice) AddFromDateTimeToDateTime() *DateTimePeriodDetails {
	p.FromDateTimeToDateTime = new(DateTimePeriodDetails)
	return p.FromDateTimeToDateTime
}

func (p *Period2Choice) AddFromDateToDate() *Period2 {
	p.FromDateToDate = new(Period2)
	return p.FromDateToDate
}
