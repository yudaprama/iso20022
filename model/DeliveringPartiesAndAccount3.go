package model

// Parameters applied to the settlement of a security transfer.
type DeliveringPartiesAndAccount3 struct {

	// Party that acts on behalf of the seller of securities when the seller does not have a direct relationship with the delivering agent.
	DeliverersCustodianDetails *PartyIdentificationAndAccount3 `xml:"DlvrrsCtdnDtls,omitempty"`

	// Party that the deliverer's custodian uses to effect the delivery of a security, when the deliverer's custodian does not have a direct relationship with the delivering agent.
	DeliverersIntermediaryDetails *PartyIdentificationAndAccount3 `xml:"DlvrrsIntrmyDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, eg, central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount3 `xml:"DlvrgAgtDtls"`
}

func (d *DeliveringPartiesAndAccount3) AddDeliverersCustodianDetails() *PartyIdentificationAndAccount3 {
	d.DeliverersCustodianDetails = new(PartyIdentificationAndAccount3)
	return d.DeliverersCustodianDetails
}

func (d *DeliveringPartiesAndAccount3) AddDeliverersIntermediaryDetails() *PartyIdentificationAndAccount3 {
	d.DeliverersIntermediaryDetails = new(PartyIdentificationAndAccount3)
	return d.DeliverersIntermediaryDetails
}

func (d *DeliveringPartiesAndAccount3) AddDeliveringAgentDetails() *PartyIdentificationAndAccount3 {
	d.DeliveringAgentDetails = new(PartyIdentificationAndAccount3)
	return d.DeliveringAgentDetails
}
