package model

// Reporting per financial instrument.
type FinancialInstrumentDetails24 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes63 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails40 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails24) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails24) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes63 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes63)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails24) AddSubBalance() *IntraPositionDetails40 {
	newValue := new(IntraPositionDetails40)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}
