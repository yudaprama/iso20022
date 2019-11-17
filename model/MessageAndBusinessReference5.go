package model

// Information to identify the funds order confirmations.
type MessageAndBusinessReference5 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifiers of one or more individual order instructions.
	OrderReference []*InvestmentFundOrder3 `xml:"OrdrRef,omitempty"`
}

func (m *MessageAndBusinessReference5) AddOtherReference() *AdditionalReference3 {
	m.OtherReference = new(AdditionalReference3)
	return m.OtherReference
}

func (m *MessageAndBusinessReference5) AddPreviousReference() *AdditionalReference3 {
	m.PreviousReference = new(AdditionalReference3)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference5) AddRelatedReference() *AdditionalReference3 {
	m.RelatedReference = new(AdditionalReference3)
	return m.RelatedReference
}

func (m *MessageAndBusinessReference5) AddOrderReference() *InvestmentFundOrder3 {
	newValue := new(InvestmentFundOrder3)
	m.OrderReference = append(m.OrderReference, newValue)
	return newValue
}
