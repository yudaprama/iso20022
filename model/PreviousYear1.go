package model

// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
type PreviousYear1 struct {

	// Selection ot the entirety of the investment plans.
	AllPreviousYears *PreviousAll `xml:"AllPrvsYrs"`

	// Selection of investment plans issued during previous years.
	SpecificPreviousYears []*ISOYear `xml:"SpcfcPrvsYrs"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`
}

func (p *PreviousYear1) SetAllPreviousYears(value string) {
	p.AllPreviousYears = (*PreviousAll)(&value)
}

func (p *PreviousYear1) AddSpecificPreviousYears(value string) {
	p.SpecificPreviousYears = append(p.SpecificPreviousYears, (*ISOYear)(&value))
}

func (p *PreviousYear1) SetCashComponentIndicator(value string) {
	p.CashComponentIndicator = (*YesNoIndicator)(&value)
}
