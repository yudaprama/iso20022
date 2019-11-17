package model

// Year in which the ISA plan is issued.
type ISAYearsOfIssue6 struct {

	// ISA that was issued during the current fiscal year.
	CurrentYear *CurrentYearType2Choice `xml:"CurYr,omitempty"`

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYearChoice `xml:"PrvsYrs,omitempty"`
}

func (i *ISAYearsOfIssue6) AddCurrentYear() *CurrentYearType2Choice {
	i.CurrentYear = new(CurrentYearType2Choice)
	return i.CurrentYear
}

func (i *ISAYearsOfIssue6) AddPreviousYears() *PreviousYearChoice {
	i.PreviousYears = new(PreviousYearChoice)
	return i.PreviousYears
}
