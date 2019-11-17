package model

// Parameters applied to the settlement of a security transfer.
type ReceivingPartiesAndAccount16 struct {

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount147 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the receiver's custodian uses to effect the receipt of a security, when the receiver's custodian does not have a direct relationship with the receiving agent.
	ReceiversIntermediary1Details *PartyIdentificationAndAccount147 `xml:"RcvrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the receiver’s intermediary 1.
	ReceiversIntermediary2Details *PartyIdentificationAndAccount147 `xml:"RcvrsIntrmy2Dtls,omitempty"`

	// Party that receives securities from the delivering agent at the place of settlement, for example, central securities depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount147 `xml:"RcvgAgtDtls"`
}

func (r *ReceivingPartiesAndAccount16) AddReceiversCustodianDetails() *PartyIdentificationAndAccount147 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount147)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount16) AddReceiversIntermediary1Details() *PartyIdentificationAndAccount147 {
	r.ReceiversIntermediary1Details = new(PartyIdentificationAndAccount147)
	return r.ReceiversIntermediary1Details
}

func (r *ReceivingPartiesAndAccount16) AddReceiversIntermediary2Details() *PartyIdentificationAndAccount147 {
	r.ReceiversIntermediary2Details = new(PartyIdentificationAndAccount147)
	return r.ReceiversIntermediary2Details
}

func (r *ReceivingPartiesAndAccount16) AddReceivingAgentDetails() *PartyIdentificationAndAccount147 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount147)
	return r.ReceivingAgentDetails
}
