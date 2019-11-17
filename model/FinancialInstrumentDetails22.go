package model

// Reporting per financial instrument.
type FinancialInstrumentDetails22 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes75 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails37 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails22) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails22) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes75 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes75)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails22) AddSubBalance() *IntraPositionDetails37 {
	newValue := new(IntraPositionDetails37)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}
