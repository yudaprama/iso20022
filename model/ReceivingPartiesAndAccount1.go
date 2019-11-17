package model

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount1 struct {

	// Party that buys goods or services, or a financial instrument.
	ReceiverDetails *InvestmentAccount11 `xml:"RcvrDtls"`

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount2 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the Receiver's custodian uses to effect the receipt of a security, when the Receiver's custodian does not have a direct relationship with the Receiver agent.
	ReceiversIntermediaryDetails *PartyIdentificationAndAccount2 `xml:"RcvrsIntrmyDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, eg, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount2 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentificationAndAccount2 `xml:"PlcOfSttlmDtls"`
}

func (r *ReceivingPartiesAndAccount1) AddReceiverDetails() *InvestmentAccount11 {
	r.ReceiverDetails = new(InvestmentAccount11)
	return r.ReceiverDetails
}

func (r *ReceivingPartiesAndAccount1) AddReceiversCustodianDetails() *PartyIdentificationAndAccount2 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount2)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount1) AddReceiversIntermediaryDetails() *PartyIdentificationAndAccount2 {
	r.ReceiversIntermediaryDetails = new(PartyIdentificationAndAccount2)
	return r.ReceiversIntermediaryDetails
}

func (r *ReceivingPartiesAndAccount1) AddReceivingAgentDetails() *PartyIdentificationAndAccount2 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount2)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount1) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount1) AddPlaceOfSettlementDetails() *PartyIdentificationAndAccount2 {
	r.PlaceOfSettlementDetails = new(PartyIdentificationAndAccount2)
	return r.PlaceOfSettlementDetails
}
