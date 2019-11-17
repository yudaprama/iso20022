package model

// Specifies  the subtotal calculated tax applicable for this settlement.
type SettlementSubTotalCalculatedTax1 struct {

	// Type of tax applied.
	TypeCode *Max4Text `xml:"TpCd,omitempty"`

	// Reference used to identify the nature of tax levied, such as Value Added Tax (VAT).
	CategoryCode *Max4Text `xml:"CtgyCd,omitempty"`

	// Monetary value resulting from the calculation of this tax, levy or duty.
	CalculatedAmount []*CurrencyAndAmount `xml:"ClctdAmt,omitempty"`

	// Monetary value used as the basis on which this tax, levy or duty is calculated.
	BasisAmount []*CurrencyAndAmount `xml:"BsisAmt,omitempty"`

	// Rate used to calculate the amount of this tax, levy or duty.
	CalculatedRate *PercentageRate `xml:"ClctdRate,omitempty"`

	// Reason for a tax exemption.
	ExemptionReason *TaxExemptionReason1 `xml:"XmptnRsn,omitempty"`
}

func (s *SettlementSubTotalCalculatedTax1) SetTypeCode(value string) {
	s.TypeCode = (*Max4Text)(&value)
}

func (s *SettlementSubTotalCalculatedTax1) SetCategoryCode(value string) {
	s.CategoryCode = (*Max4Text)(&value)
}

func (s *SettlementSubTotalCalculatedTax1) AddCalculatedAmount(value, currency string) {
	s.CalculatedAmount = append(s.CalculatedAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementSubTotalCalculatedTax1) AddBasisAmount(value, currency string) {
	s.BasisAmount = append(s.BasisAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementSubTotalCalculatedTax1) SetCalculatedRate(value string) {
	s.CalculatedRate = (*PercentageRate)(&value)
}

func (s *SettlementSubTotalCalculatedTax1) AddExemptionReason() *TaxExemptionReason1 {
	s.ExemptionReason = new(TaxExemptionReason1)
	return s.ExemptionReason
}
