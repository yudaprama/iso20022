package model

// Set of data which contains the link to the opening of the non deliverable trade and supplementary information on its valuation.
type ValuationData2 struct {

	// Reference to the latest trade identification of the NDF opening trade.
	ValuationReference *Max35Text `xml:"ValtnRef"`

	// Specifies the currency in which the non deliverable trade has to be settled ie the deliverable currency.
	SettlementCurrency *ActiveOrHistoricCurrencyCode `xml:"SttlmCcy,omitempty"`

	// Free format text that may contain information on the valuation such as the currency, the place, the time or the source of the rate.
	AdditionalValuationInformation *Max140Text `xml:"AddtlValtnInf,omitempty"`

	// Party through which the settlement will take place. It may contain the BIC of a central settlement system eg CLSBUS33.
	SettlementParty *PartyIdentification8Choice `xml:"SttlmPty,omitempty"`
}

func (v *ValuationData2) SetValuationReference(value string) {
	v.ValuationReference = (*Max35Text)(&value)
}

func (v *ValuationData2) SetSettlementCurrency(value string) {
	v.SettlementCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (v *ValuationData2) SetAdditionalValuationInformation(value string) {
	v.AdditionalValuationInformation = (*Max140Text)(&value)
}

func (v *ValuationData2) AddSettlementParty() *PartyIdentification8Choice {
	v.SettlementParty = new(PartyIdentification8Choice)
	return v.SettlementParty
}
