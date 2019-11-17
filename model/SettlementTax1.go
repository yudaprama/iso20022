package model

// Species the tax applicable for this settlement.
type SettlementTax1 struct {

	// Type of tax applied.
	TypeCode *TaxTypeFormat1Choice `xml:"TpCd,omitempty"`

	// Monetary value resulting from the calculation of this tax, levy or duty.
	CalculatedAmount []*CurrencyAndAmount `xml:"ClctdAmt,omitempty"`

	// Monetary value used as the basis on which this tax, levy or duty is calculated.
	BasisAmount []*CurrencyAndAmount `xml:"BsisAmt,omitempty"`

	// Date of the tax point when this tax, levy or duty becomes applicable.
	TaxPointDate *ISODate `xml:"TaxPtDt,omitempty"`
}

func (s *SettlementTax1) AddTypeCode() *TaxTypeFormat1Choice {
	s.TypeCode = new(TaxTypeFormat1Choice)
	return s.TypeCode
}

func (s *SettlementTax1) AddCalculatedAmount(value, currency string) {
	s.CalculatedAmount = append(s.CalculatedAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementTax1) AddBasisAmount(value, currency string) {
	s.BasisAmount = append(s.BasisAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementTax1) SetTaxPointDate(value string) {
	s.TaxPointDate = (*ISODate)(&value)
}
