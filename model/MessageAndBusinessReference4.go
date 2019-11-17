package model

// Information to identify funds order(s).
type MessageAndBusinessReference4 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifiers of one or more indiviudal order instructions or individual order cancellation requests.
	OrderReference []*InvestmentFundOrder2 `xml:"OrdrRef,omitempty"`
}

func (m *MessageAndBusinessReference4) AddOtherReference() *AdditionalReference3 {
	m.OtherReference = new(AdditionalReference3)
	return m.OtherReference
}

func (m *MessageAndBusinessReference4) AddPreviousReference() *AdditionalReference3 {
	m.PreviousReference = new(AdditionalReference3)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference4) AddRelatedReference() *AdditionalReference3 {
	m.RelatedReference = new(AdditionalReference3)
	return m.RelatedReference
}

func (m *MessageAndBusinessReference4) AddOrderReference() *InvestmentFundOrder2 {
	newValue := new(InvestmentFundOrder2)
	m.OrderReference = append(m.OrderReference, newValue)
	return newValue
}
