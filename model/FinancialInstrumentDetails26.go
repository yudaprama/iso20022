package model

// Reporting per financial instrument.
type FinancialInstrumentDetails26 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes75 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails44 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails26) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails26) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes75 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes75)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails26) AddSubBalance() *IntraPositionDetails44 {
	newValue := new(IntraPositionDetails44)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}
