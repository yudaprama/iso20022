package model

// Range of time defined by a start date and an end date.
type DatePeriod1 struct {

	// Start date of the range.
	FromDate *ISODate `xml:"FrDt,omitempty"`

	// End date of the range.
	ToDate *ISODate `xml:"ToDt"`
}

func (d *DatePeriod1) SetFromDate(value string) {
	d.FromDate = (*ISODate)(&value)
}

func (d *DatePeriod1) SetToDate(value string) {
	d.ToDate = (*ISODate)(&value)
}
