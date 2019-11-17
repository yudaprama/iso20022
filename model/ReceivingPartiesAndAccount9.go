package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount9 struct {

	// Party that buys goods or services, or a financial instrument.
	ReceiverDetails *InvestmentAccount41 `xml:"RcvrDtls,omitempty"`

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount5 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the Receiver's custodian uses to effect the receipt of a security, when the Receiver's custodian does not have a direct relationship with the Receiver agent.
	ReceiversIntermediaryDetails *PartyIdentificationAndAccount5 `xml:"RcvrsIntrmyDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, eg, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount4 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification21 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (r *ReceivingPartiesAndAccount9) AddReceiverDetails() *InvestmentAccount41 {
	r.ReceiverDetails = new(InvestmentAccount41)
	return r.ReceiverDetails
}

func (r *ReceivingPartiesAndAccount9) AddReceiversCustodianDetails() *PartyIdentificationAndAccount5 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount5)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount9) AddReceiversIntermediaryDetails() *PartyIdentificationAndAccount5 {
	r.ReceiversIntermediaryDetails = new(PartyIdentificationAndAccount5)
	return r.ReceiversIntermediaryDetails
}

func (r *ReceivingPartiesAndAccount9) AddReceivingAgentDetails() *PartyIdentificationAndAccount4 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount4)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount9) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount9) AddPlaceOfSettlementDetails() *PartyIdentification21 {
	r.PlaceOfSettlementDetails = new(PartyIdentification21)
	return r.PlaceOfSettlementDetails
}
