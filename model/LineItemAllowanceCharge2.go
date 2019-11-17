package model

// Pricing component, such as a service, promotion, allowance or charge, for this line item.
type LineItemAllowanceCharge2 struct {

	// Indication of whether or not this allowance charge is a charge.
	ChargeIndicator *YesNoIndicator `xml:"ChrgInd,omitempty"`

	// Actual monetary value of this allowance charge.
	ActualAmount []*CurrencyAndAmount `xml:"ActlAmt,omitempty"`

	// Quantity on which this allowance charge is based.
	BasisQuantity *Quantity10 `xml:"BsisQty,omitempty"`

	// Percentage applied to calculate this allowance charge.
	CalculationPercent *PercentageRate `xml:"ClctnPct,omitempty"`

	// Specifies the order in which the allowance or charge is applied.
	SequenceNumber *Number `xml:"SeqNb,omitempty"`

	// Reason, expressed as text, for this allowance charge.
	Reason *DiscountOrChargeType1Choice `xml:"Rsn,omitempty"`
}

func (l *LineItemAllowanceCharge2) SetChargeIndicator(value string) {
	l.ChargeIndicator = (*YesNoIndicator)(&value)
}

func (l *LineItemAllowanceCharge2) AddActualAmount(value, currency string) {
	l.ActualAmount = append(l.ActualAmount, NewCurrencyAndAmount(value, currency))
}

func (l *LineItemAllowanceCharge2) AddBasisQuantity() *Quantity10 {
	l.BasisQuantity = new(Quantity10)
	return l.BasisQuantity
}

func (l *LineItemAllowanceCharge2) SetCalculationPercent(value string) {
	l.CalculationPercent = (*PercentageRate)(&value)
}

func (l *LineItemAllowanceCharge2) SetSequenceNumber(value string) {
	l.SequenceNumber = (*Number)(&value)
}

func (l *LineItemAllowanceCharge2) AddReason() *DiscountOrChargeType1Choice {
	l.Reason = new(DiscountOrChargeType1Choice)
	return l.Reason
}
