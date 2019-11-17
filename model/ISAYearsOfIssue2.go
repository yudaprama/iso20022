package model

// Year in which the ISA plan is issued.
type ISAYearsOfIssue2 struct {

	// ISA that was issued during the current fiscal year.
	CurrentYearType *ISAType2Code `xml:"CurYrTp,omitempty"`

	// Current year ISA is an ISA that was issued during the current fiscal year.
	ExtendedCurrentYearType *Extended350Code `xml:"XtndedCurYrTp,omitempty"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYearChoice `xml:"PrvsYrs,omitempty"`
}

func (i *ISAYearsOfIssue2) SetCurrentYearType(value string) {
	i.CurrentYearType = (*ISAType2Code)(&value)
}

func (i *ISAYearsOfIssue2) SetExtendedCurrentYearType(value string) {
	i.ExtendedCurrentYearType = (*Extended350Code)(&value)
}

func (i *ISAYearsOfIssue2) AddPreviousYears() *PreviousYearChoice {
	i.PreviousYears = new(PreviousYearChoice)
	return i.PreviousYears
}
