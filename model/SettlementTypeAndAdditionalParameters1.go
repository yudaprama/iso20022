package model

// Provides transaction type and identification information.
type SettlementTypeAndAdditionalParameters1 struct {

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
}

func (s *SettlementTypeAndAdditionalParameters1) SetSecuritiesMovementType(value string) {
	s.SecuritiesMovementType = (*ReceiveDelivery1Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters1) SetPayment(value string) {
	s.Payment = (*DeliveryReceiptType2Code)(&value)
}

func (s *SettlementTypeAndAdditionalParameters1) SetCommonIdentification(value string) {
	s.CommonIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters1) SetCorporateActionEventIdentification(value string) {
	s.CorporateActionEventIdentification = (*Max35Text)(&value)
}

func (s *SettlementTypeAndAdditionalParameters1) SetReconciliationIndicator(value string) {
	s.ReconciliationIndicator = (*YesNoIndicator)(&value)
}
