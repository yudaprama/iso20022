package model

// Reporting per financial instrument.
type FinancialInstrumentDetails1 struct {

	// Financial instruments representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification11 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes4 `xml:"FinInstrmAttrbts,omitempty"`

	// Identification of the sub-balance.
	SubBalance []*IntraPositionDetails3 `xml:"SubBal"`
}

func (f *FinancialInstrumentDetails1) AddFinancialInstrumentIdentification() *SecurityIdentification11 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification11)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentDetails1) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes4 {
	f.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes4)
	return f.FinancialInstrumentAttributes
}

func (f *FinancialInstrumentDetails1) AddSubBalance() *IntraPositionDetails3 {
	newValue := new(IntraPositionDetails3)
	f.SubBalance = append(f.SubBalance, newValue)
	return newValue
}
