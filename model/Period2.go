package model

// Time span defined by a start date and time, and an end date and time.
type Period2 struct {

	// Date and time at which the range starts.
	FromDate *ISODate `xml:"FrDt"`

	// Date and time at which the range ends.
	ToDate *ISODate `xml:"ToDt"`
}

func (p *Period2) SetFromDate(value string) {
	p.FromDate = (*ISODate)(&value)
}

func (p *Period2) SetToDate(value string) {
	p.ToDate = (*ISODate)(&value)
}
