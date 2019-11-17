package model

// Set of parameters used to calculate the fixing rate to be applied to a non-deliverable agreement.
type NonDeliverableForwardValuationConditions2 struct {

	// Specifies the currency in which the non deliverable trade has to be settled ie the deliverable currency.
	SettlementCurrency *ActiveOrHistoricCurrencyCode `xml:"SttlmCcy"`

	// Date at which the rate used for calculating the net amount to be settled is observed.
	ValuationDate *ISODate `xml:"ValtnDt"`

	// Free format text that may contain valuation information such as the place, the time or the source of the rate.
	AdditionalValuationInformation *Max140Text `xml:"AddtlValtnInf,omitempty"`

	// Party through which the settlement will take place. It may contain the BIC of a central settlement system eg CLSBUS33.
	SettlementParty *PartyIdentification8Choice `xml:"SttlmPty,omitempty"`
}

func (n *NonDeliverableForwardValuationConditions2) SetSettlementCurrency(value string) {
	n.SettlementCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (n *NonDeliverableForwardValuationConditions2) SetValuationDate(value string) {
	n.ValuationDate = (*ISODate)(&value)
}

func (n *NonDeliverableForwardValuationConditions2) SetAdditionalValuationInformation(value string) {
	n.AdditionalValuationInformation = (*Max140Text)(&value)
}

func (n *NonDeliverableForwardValuationConditions2) AddSettlementParty() *PartyIdentification8Choice {
	n.SettlementParty = new(PartyIdentification8Choice)
	return n.SettlementParty
}
