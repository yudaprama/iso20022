package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type DeliveringPartiesAndAccount14 struct {

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount124 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediary1Details *PartyIdentificationAndAccount124 `xml:"DlvrrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the deliverer's intermediary.
	DeliverersIntermediary2Details *PartyIdentificationAndAccount124 `xml:"DlvrrsIntrmy2Dtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount124 `xml:"DlvrgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification97 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (d *DeliveringPartiesAndAccount14) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount124 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount124)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount14) AddDeliverersIntermediary1Details() *PartyIdentificationAndAccount124 {
	d.DeliverersIntermediary1Details = new(PartyIdentificationAndAccount124)
	return d.DeliverersIntermediary1Details
}

func (d *DeliveringPartiesAndAccount14) AddDeliverersIntermediary2Details() *PartyIdentificationAndAccount124 {
	d.DeliverersIntermediary2Details = new(PartyIdentificationAndAccount124)
	return d.DeliverersIntermediary2Details
}

func (d *DeliveringPartiesAndAccount14) AddDeliveringAgentDetails() *PartyIdentificationAndAccount124 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount124)
	return d.DeliveringAgentDetails
}

func (d *DeliveringPartiesAndAccount14) SetSecuritiesSettlementSystem(value string) {
	d.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (d *DeliveringPartiesAndAccount14) AddPlaceOfSettlementDetails() *PartyIdentification97 {
	d.PlaceOfSettlementDetails = new(PartyIdentification97)
	return d.PlaceOfSettlementDetails
}
