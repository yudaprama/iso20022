package model

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters9 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *Max35Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *Max35Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the client's point of view.
	ClientCollateralTransactionIdentification *Max35Text `xml:"ClntCollTxId,omitempty"`

	// Unique identification identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *Max35Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *Max35Text `xml:"TrptyCollInstrId,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters9) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetClientCollateralInstructionIdentification(value string) {
	s.ClientCollateralInstructionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetClientCollateralTransactionIdentification(value string) {
	s.ClientCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetTripartyAgentCollateralTransactionIdentification(value string) {
	s.TripartyAgentCollateralTransactionIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters9) SetTripartyCollateralInstructionIdentification(value string) {
	s.TripartyCollateralInstructionIdentification = (*Max35Text)(&value)
}
