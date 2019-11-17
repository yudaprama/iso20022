package model

// Specifies  the subtotal calculated tax applicable for this settlement.
type SettlementSubTotalCalculatedTax2 struct {

	// Type of tax applied.
	TypeCode *Max4Text `xml:"TpCd,omitempty"`

	// Rate used to calculate the amount of this tax, levy or duty.
	CalculatedRate *PercentageRate `xml:"ClctdRate,omitempty"`

	// Monetary value used as the basis on which this tax, levy or duty is calculated.
	BasisAmount []*CurrencyAndAmount `xml:"BsisAmt,omitempty"`

	// Monetary value resulting from the calculation of this tax, levy or duty.
	CalculatedAmount []*CurrencyAndAmount `xml:"ClctdAmt,omitempty"`

	// Reason for tax exemption expressed as a code,  if invoice or card transaction is out of tax processing.
	ExemptionReasonCode *Max4Text `xml:"XmptnRsnCd,omitempty"`

	// Reason for a tax exemption,  if invoice or card transaction is out of tax processing.
	ExemptionReasonText *Max500Text `xml:"XmptnRsnTxt,omitempty"`

	// If tax currency in tax calculation is different from invoice currency, then applied exchange rate is given in this message structure.
	TaxCurrencyExchange *CurrencyReference3 `xml:"TaxCcyXchg,omitempty"`
}

func (s *SettlementSubTotalCalculatedTax2) SetTypeCode(value string) {
	s.TypeCode = (*Max4Text)(&value)
}

func (s *SettlementSubTotalCalculatedTax2) SetCalculatedRate(value string) {
	s.CalculatedRate = (*PercentageRate)(&value)
}

func (s *SettlementSubTotalCalculatedTax2) AddBasisAmount(value, currency string) {
	s.BasisAmount = append(s.BasisAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementSubTotalCalculatedTax2) AddCalculatedAmount(value, currency string) {
	s.CalculatedAmount = append(s.CalculatedAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementSubTotalCalculatedTax2) SetExemptionReasonCode(value string) {
	s.ExemptionReasonCode = (*Max4Text)(&value)
}

func (s *SettlementSubTotalCalculatedTax2) SetExemptionReasonText(value string) {
	s.ExemptionReasonText = (*Max500Text)(&value)
}

func (s *SettlementSubTotalCalculatedTax2) AddTaxCurrencyExchange() *CurrencyReference3 {
	s.TaxCurrencyExchange = new(CurrencyReference3)
	return s.TaxCurrencyExchange
}
