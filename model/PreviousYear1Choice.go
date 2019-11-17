package model

// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
type PreviousYear1Choice struct {

	// Selection ot the entirety of the investment plans.
	AllPreviousYears *PreviousAll `xml:"AllPrvsYrs"`

	// Selection of investment plans issued during previous years.
	SpecificPreviousYears []*ISOYear `xml:"SpcfcPrvsYrs"`
}

func (p *PreviousYear1Choice) SetAllPreviousYears(value string) {
	p.AllPreviousYears = (*PreviousAll)(&value)
}

func (p *PreviousYear1Choice) AddSpecificPreviousYears(value string) {
	p.SpecificPreviousYears = append(p.SpecificPreviousYears, (*ISOYear)(&value))
}
