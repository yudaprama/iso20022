package model

// Amount of expected margin required by any of the parties of the margin calculation.
type MarginRequirement1 struct {

	// Amount of new margin that will be delivered to one party by the other party after rounding, threshold and minimum transfer amount are taken into account.
	DeliverMarginAmount *ActiveCurrencyAndAmount `xml:"DlvrMrgnAmt,omitempty"`

	// Amount of new margin that will be recalled to one party from the other party after rounding, threshold and minimum transfer amount are taken into account.
	ReturnMarginAmount *ActiveCurrencyAndAmount `xml:"RtrMrgnAmt,omitempty"`
}

func (m *MarginRequirement1) SetDeliverMarginAmount(value, currency string) {
	m.DeliverMarginAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginRequirement1) SetReturnMarginAmount(value, currency string) {
	m.ReturnMarginAmount = NewActiveCurrencyAndAmount(value, currency)
}
