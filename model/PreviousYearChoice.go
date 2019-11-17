package model

// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
type PreviousYearChoice struct {

	// Selection ot the entirety of the investment plans.
	AllPreviousYears *PreviousAll `xml:"AllPrvsYrs"`

	// Selection of investment plans issued during previous years.
	SpecificPreviousYears []*ISOYear `xml:"SpcfcPrvsYrs"`
}

func (p *PreviousYearChoice) SetAllPreviousYears(value string) {
	p.AllPreviousYears = (*PreviousAll)(&value)
}

func (p *PreviousYearChoice) AddSpecificPreviousYears(value string) {
	p.SpecificPreviousYears = append(p.SpecificPreviousYears, (*ISOYear)(&value))
}
