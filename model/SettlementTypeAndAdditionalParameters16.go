package model

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters16 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the client's point of view.
	ClientCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntCollTxId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"TrptyCollInstrId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters16) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetClientCollateralInstructionIdentification(value string) {
	s.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetClientCollateralTransactionIdentification(value string) {
	s.ClientCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters16) SetTripartyCollateralInstructionIdentification(value string) {
	s.TripartyCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}
