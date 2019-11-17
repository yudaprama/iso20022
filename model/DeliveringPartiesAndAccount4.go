package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount4 struct {

	// Party that sells goods or services, or a financial instrument.
	DelivererDetails *InvestmentAccount24 `xml:"DlvrrDtls,omitempty"`

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount5 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediaryDetails *PartyIdentificationAndAccount5 `xml:"DlvrrsIntrmyDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount4 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification21 `xml:"PlcOfSttlmDtls"`
}

func (d *DeliveringPartiesAndAccount4) AddDelivererDetails() *InvestmentAccount24 {
	d.DelivererDetails = new(InvestmentAccount24)
	return d.DelivererDetails
}

func (d *DeliveringPartiesAndAccount4) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount5 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount5)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount4) AddDeliverersIntermediaryDetails() *PartyIdentificationAndAccount5 {
	d.DeliverersIntermediaryDetails = new(PartyIdentificationAndAccount5)
	return d.DeliverersIntermediaryDetails
}

func (d *DeliveringPartiesAndAccount4) AddDeliveringAgentDetails() *PartyIdentificationAndAccount4 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount4)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount4) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount4) AddPlaceOfSettlementDetails() *PartyIdentification21 {
	d.PlaceOfSettlementDetails = new(PartyIdentification21)
	return d.PlaceOfSettlementDetails
}
