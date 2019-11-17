package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount1 struct {

	// Party that sells goods or services, or a financial instrument.
	DelivererDetails *InvestmentAccount11 `xml:"DlvrrDtls"`

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount2 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediaryDetails *PartyIdentificationAndAccount2 `xml:"DlvrrsIntrmyDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount2 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentificationAndAccount2 `xml:"PlcOfSttlmDtls"`
}

func (d *DeliveringPartiesAndAccount1) AddDelivererDetails() *InvestmentAccount11 {
	d.DelivererDetails = new(InvestmentAccount11)
	return d.DelivererDetails
}

func (d *DeliveringPartiesAndAccount1) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount2 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount2)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount1) AddDeliverersIntermediaryDetails() *PartyIdentificationAndAccount2 {
	d.DeliverersIntermediaryDetails = new(PartyIdentificationAndAccount2)
	return d.DeliverersIntermediaryDetails
}

func (d *DeliveringPartiesAndAccount1) AddDeliveringAgentDetails() *PartyIdentificationAndAccount2 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount2)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount1) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount1) AddPlaceOfSettlementDetails() *PartyIdentificationAndAccount2 {
	d.PlaceOfSettlementDetails = new(PartyIdentificationAndAccount2)
	return d.PlaceOfSettlementDetails
}
