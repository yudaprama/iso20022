package model

// Parameters applied to the settlement of a security transfer.
type DeliveringPartiesAndAccount16 struct {

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount147 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediary1Details *PartyIdentificationAndAccount147 `xml:"DlvrrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the deliverer's intermediary 1.
	DeliverersIntermediary2Details *PartyIdentificationAndAccount147 `xml:"DlvrrsIntrmy2Dtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount147 `xml:"DlvrgAgtDtls"`
}

func (d *DeliveringPartiesAndAccount16) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount147 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount147)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount16) AddDeliverersIntermediary1Details() *PartyIdentificationAndAccount147 {
	d.DeliverersIntermediary1Details = new(PartyIdentificationAndAccount147)
	return d.DeliverersIntermediary1Details
}

func (d *DeliveringPartiesAndAccount16) AddDeliverersIntermediary2Details() *PartyIdentificationAndAccount147 {
	d.DeliverersIntermediary2Details = new(PartyIdentificationAndAccount147)
	return d.DeliverersIntermediary2Details
}

func (d *DeliveringPartiesAndAccount16) AddDeliveringAgentDetails() *PartyIdentificationAndAccount147 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount147)
	return d.DeliveringAgentDetails
}
