package model

// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
type PreviousYear2 struct {

	// Selection of investment plans issued during previous years.
	PreviousYears *PreviousYear1Choice `xml:"PrvsYrs"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`
}

func (p *PreviousYear2) AddPreviousYears() *PreviousYear1Choice {
	p.PreviousYears = new(PreviousYear1Choice)
	return p.PreviousYears
}

func (p *PreviousYear2) SetCashComponentIndicator(value string) {
	p.CashComponentIndicator = (*YesNoIndicator)(&value)
}
