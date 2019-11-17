package model

// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
type PreviousYear3 struct {

	// Choice between selected investment plans issued during previous years or the entirety of the investment plans.
	PreviousYear *PreviousYear1Choice `xml:"PrvsYr"`

	// Indicates whether the ISA contains a cash component asset for transfer.
	CashComponentIndicator *YesNoIndicator `xml:"CshCmpntInd"`
}

func (p *PreviousYear3) AddPreviousYear() *PreviousYear1Choice {
	p.PreviousYear = new(PreviousYear1Choice)
	return p.PreviousYear
}

func (p *PreviousYear3) SetCashComponentIndicator(value string) {
	p.CashComponentIndicator = (*YesNoIndicator)(&value)
}
