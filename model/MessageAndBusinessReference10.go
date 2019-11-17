package model

// Information to identify funds order(s).
type MessageAndBusinessReference10 struct {

	// Reference to a linked message that was previously sent.
	Reference *References62Choice `xml:"Ref,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference *AdditionalReference8 `xml:"RltdRef,omitempty"`

	// Unique and unambiguous identifiers of one or more individual order instructions or individual order cancellation requests.
	OrderReference []*InvestmentFundOrder8 `xml:"OrdrRef,omitempty"`
}

func (m *MessageAndBusinessReference10) AddReference() *References62Choice {
	m.Reference = new(References62Choice)
	return m.Reference
}

func (m *MessageAndBusinessReference10) AddRelatedReference() *AdditionalReference8 {
	m.RelatedReference = new(AdditionalReference8)
	return m.RelatedReference
}

func (m *MessageAndBusinessReference10) AddOrderReference() *InvestmentFundOrder8 {
	newValue := new(InvestmentFundOrder8)
	m.OrderReference = append(m.OrderReference, newValue)
	return newValue
}
