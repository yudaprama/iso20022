package model

// Parameters applied to the settlement of a security transfer.
type ReceivingPartiesAndAccount3 struct {

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount3 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the Receiver's custodian uses to effect the receipt of a security, when the Receiver's custodian does not have a direct relationship with the Receiver agent.
	ReceiversIntermediaryDetails *PartyIdentificationAndAccount3 `xml:"RcvrsIntrmyDtls,omitempty"`

	// Party that receives securities from the delivering agent at the place of settlement, eg, central securities depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount3 `xml:"RcvgAgtDtls"`
}

func (r *ReceivingPartiesAndAccount3) AddReceiversCustodianDetails() *PartyIdentificationAndAccount3 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount3)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount3) AddReceiversIntermediaryDetails() *PartyIdentificationAndAccount3 {
	r.ReceiversIntermediaryDetails = new(PartyIdentificationAndAccount3)
	return r.ReceiversIntermediaryDetails
}

func (r *ReceivingPartiesAndAccount3) AddReceivingAgentDetails() *PartyIdentificationAndAccount3 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount3)
	return r.ReceivingAgentDetails
}
