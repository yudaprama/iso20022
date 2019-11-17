package model

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters18 struct {

	// Specifies if the movement on a securities account results from a deliver or a receive instruction.
	SecuritiesMovementType *ReceiveDelivery1Code `xml:"SctiesMvmntTp,omitempty"`

	// Specifies how the transaction is to be settled, for example, against payment.
	Payment *DeliveryReceiptType2Code `xml:"Pmt,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Indicates whether the settlement transaction was already sent on the market and that it is only sent by an account owner to an account servicer for reconciliation purposes.
	ReconciliationIndicator *YesNoIndicator `xml:"RcncltnInd,omitempty"`
}

func (s *SettlementTypeAndAdditionalParameters18) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters18) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters18) SetCommonIdentification(value string) {
	s.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters18) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters18) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}
