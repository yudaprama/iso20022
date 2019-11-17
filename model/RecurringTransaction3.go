package model

// Financial loan (instalment) or a recurring transaction.
type RecurringTransaction3 struct {

	// Date of first transfer.
	StartDate *ISODate `xml:"StartDt"`

	// Number of transfers to perform.
	NumberOfOccurrences *Number `xml:"NbOfOcrncs,omitempty"`

	// Date of last transfer.
	EndDate *ISODate `xml:"EndDt"`

	// Period of the recurring transfer.
	PeriodUnit *Frequency3Code `xml:"PrdUnit,omitempty"`

	// Day of the period when the transfer will be performed.
	IntervalDay *Number `xml:"IntrvlDay,omitempty"`
}

func (r *RecurringTransaction3) SetStartDate(value string) {
	r.StartDate = (*ISODate)(&value)
}

func (r *RecurringTransaction3) SetNumberOfOccurrences(value string) {
	r.NumberOfOccurrences = (*Number)(&value)
}

func (r *RecurringTransaction3) SetEndDate(value string) {
	r.EndDate = (*ISODate)(&value)
}

func (r *RecurringTransaction3) SetPeriodUnit(value string) {
	r.PeriodUnit = (*Frequency3Code)(&value)
}

func (r *RecurringTransaction3) SetIntervalDay(value string) {
	r.IntervalDay = (*Number)(&value)
}
