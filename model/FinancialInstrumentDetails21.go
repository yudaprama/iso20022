package model

// Reporting per financial instrument.
type FinancialInstrumentDetails21 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes63 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails32 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails21) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails21) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes63 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes63)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails21) AddSubBalance() *IntraPositionDetails32 {
	newValue := new(IntraPositionDetails32)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}
