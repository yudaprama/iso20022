package model

// Information about the message reference of the message for which the status is requested and the business reference of the order.
type MessageAndBusinessReference2 struct {

	// Reference to a linked message sent in a proprietary way or reference of a system.
	OtherReference *AdditionalReference3 `xml:"OthrRef"`

	// Reference to a linked message that was previously sent.
	PreviousReference *AdditionalReference3 `xml:"PrvsRef"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	IndividualOrderReference []*Max35Text `xml:"IndvOrdrRef,omitempty"`

	// Account information of the order message for which the status is requested.
	InvestmentAccount *InvestmentAccount13 `xml:"InvstmtAcct,omitempty"`
}

func (m *MessageAndBusinessReference2) AddOtherReference() *AdditionalReference3 {
	m.OtherReference = new(AdditionalReference3)
	return m.OtherReference
}

func (m *MessageAndBusinessReference2) AddPreviousReference() *AdditionalReference3 {
	m.PreviousReference = new(AdditionalReference3)
	return m.PreviousReference
}

func (m *MessageAndBusinessReference2) AddIndividualOrderReference(value string) {
	m.IndividualOrderReference = append(m.IndividualOrderReference, (*Max35Text)(&value))
}

func (m *MessageAndBusinessReference2) AddInvestmentAccount() *InvestmentAccount13 {
	m.InvestmentAccount = new(InvestmentAccount13)
	return m.InvestmentAccount
}
