package model

// Period of time details related to the tax payment.
type TaxPeriod1 struct {

	// Year related to the tax payment.
	Year *ISODate `xml:"Yr,omitempty"`

	// Identification of the period related to the tax payment.
	Type *TaxRecordPeriod1Code `xml:"Tp,omitempty"`

	// Range of time between a start date and an end date for which the tax report is provided.
	FromToDate *DatePeriodDetails `xml:"FrToDt,omitempty"`
}

func (t *TaxPeriod1) SetYear(value string) {
	t.Year = (*ISODate)(&value)
}

func (t *TaxPeriod1) SetType(value string) {
	t.Type = (*TaxRecordPeriod1Code)(&value)
}

func (t *TaxPeriod1) AddFromToDate() *DatePeriodDetails {
	t.FromToDate = new(DatePeriodDetails)
	return t.FromToDate
}
