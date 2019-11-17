package model

// Reporting per financial instrument.
type FinancialInstrumentDetails6 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes21 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails17 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails6) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails6) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes21 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes21)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails6) AddSubBalance() *IntraPositionDetails17 {
	newValue := new(IntraPositionDetails17)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}
