package model

// Specifies a pricing component, such as a service, promotion, allowance or charge, for this trade settlement.
type SettlementAllowanceCharge1 struct {

	// Indication of whether or not this allowance charge is a charge.
	AllowanceChargeIndicator *YesNoIndicator `xml:"AllwncChrgInd,omitempty"`

	// Actual monetary amount specified for the allowance or charge.
	ActualAmount []*CurrencyAndAmount `xml:"ActlAmt,omitempty"`

	// Reason, expressed as text, for this allowance or charge.
	Reason *DiscountOrChargeType1Choice `xml:"Rsn,omitempty"`
}

func (s *SettlementAllowanceCharge1) SetAllowanceChargeIndicator(value string) {
	s.AllowanceChargeIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementAllowanceCharge1) AddActualAmount(value, currency string) {
	s.ActualAmount = append(s.ActualAmount, NewCurrencyAndAmount(value, currency))
}

func (s *SettlementAllowanceCharge1) AddReason() *DiscountOrChargeType1Choice {
	s.Reason = new(DiscountOrChargeType1Choice)
	return s.Reason
}
