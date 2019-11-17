package model

// Specifies the details for the adjustment of the mandate.
type MandateAdjustment1 struct {

	// Specifies whether an adjustment is to be applied on pre-agreed collection date or not.
	DateAdjustmentRuleIndicator *TrueFalseIndicator `xml:"DtAdjstmntRuleInd"`

	// Defines the category of adjustment.
	Category *Frequency37Choice `xml:"Ctgy,omitempty"`

	// Pre-agreed amount to increase or decrease the mandate amount as justified per information in the category.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Pre-agreed increase or decrease rate that will be applied to the collection amount.
	Rate *PercentageRate `xml:"Rate,omitempty"`
}

func (m *MandateAdjustment1) SetDateAdjustmentRuleIndicator(value string) {
	m.DateAdjustmentRuleIndicator = (*TrueFalseIndicator)(&value)
}

func (m *MandateAdjustment1) AddCategory() *Frequency37Choice {
	m.Category = new(Frequency37Choice)
	return m.Category
}

func (m *MandateAdjustment1) SetAmount(value, currency string) {
	m.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MandateAdjustment1) SetRate(value string) {
	m.Rate = (*PercentageRate)(&value)
}
