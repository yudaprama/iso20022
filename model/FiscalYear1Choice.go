package model

// Choice of start date and end date for the fiscal year.
type FiscalYear1Choice struct {

	// Start date of the fiscal year.
	StartDate *ISODate `xml:"StartDt"`

	// End date of the fiscal year.
	EndDate *ISODate `xml:"EndDt"`
}

func (f *FiscalYear1Choice) SetStartDate(value string) {
	f.StartDate = (*ISODate)(&value)
}

func (f *FiscalYear1Choice) SetEndDate(value string) {
	f.EndDate = (*ISODate)(&value)
}
