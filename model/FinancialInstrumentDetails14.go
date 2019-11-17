package model

// Reporting per financial instrument.
type FinancialInstrumentDetails14 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes36 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails28 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails14) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails14) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes36 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes36)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails14) AddSubBalance() *IntraPositionDetails28 {
	newValue := new(IntraPositionDetails28)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}
