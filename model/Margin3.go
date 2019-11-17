package model

// Provides details on the calculation of the margin.
type Margin3 struct {

	// Margin required for absorbing future market price fluctuations (market risks) occurring between the default of a member and close-out of unsettled securities positions by the central counterparty.
	InitialMargin *Amount2 `xml:"InitlMrgn,omitempty"`

	// Provides details on the calculation of the variation margin.
	VariationMargin []*VariationMargin3 `xml:"VartnMrgn,omitempty"`

	// Provides details on the margin type and amount.
	OtherMargin []*Margin4 `xml:"OthrMrgn,omitempty"`
}

func (m *Margin3) AddInitialMargin() *Amount2 {
	m.InitialMargin = new(Amount2)
	return m.InitialMargin
}

func (m *Margin3) AddVariationMargin() *VariationMargin3 {
	newValue := new(VariationMargin3)
	m.VariationMargin = append(m.VariationMargin, newValue)
	return newValue
}

func (m *Margin3) AddOtherMargin() *Margin4 {
	newValue := new(Margin4)
	m.OtherMargin = append(m.OtherMargin, newValue)
	return newValue
}
