package model

// Choice between date and date-time for the specification of a period.
type Period4Choice struct {

	// Date period is limited to a single date.
	Date *ISODate `xml:"Dt"`

	// Date at which the date period range starts.
	FromDate *ISODate `xml:"FrDt"`

	// Date which the range date period ends.
	ToDate *ISODate `xml:"ToDt"`

	// Time span defined by a start date, and an end date.
	FromDateToDate *Period2 `xml:"FrDtToDt"`
}

func (p *Period4Choice) SetDate(value string) {
	p.Date = (*ISODate)(&value)
}

func (p *Period4Choice) SetFromDate(value string) {
	p.FromDate = (*ISODate)(&value)
}

func (p *Period4Choice) SetToDate(value string) {
	p.ToDate = (*ISODate)(&value)
}

func (p *Period4Choice) AddFromDateToDate() *Period2 {
	p.FromDateToDate = new(Period2)
	return p.FromDateToDate
}
