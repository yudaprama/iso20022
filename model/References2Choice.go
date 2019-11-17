package model

// Choice of reference.
type References2Choice struct {

	// Unambiguous identification of the securities settlement transaction.
	SecuritiesSettlementTransactionIdentification *SettlementTypeAndIdentification3 `xml:"SctiesSttlmTxId"`

	// Unambiguous identification of the underlying securities financing transaction (not the underlying securities financing trade) as assigned by the instructing party.
	SecuritiesFinancingTransactionIdentification *SettlementTypeAndIdentification3 `xml:"SctiesFincgTxId"`

	// Reference to the intra-position movement transaction requested to be cancelled as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *Max35Text `xml:"IntraPosMvmntId"`

	// Reference to a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *GenericDocumentIdentification1 `xml:"OthrTxId"`
}

func (r *References2Choice) AddSecuritiesSettlementTransactionIdentification() *SettlementTypeAndIdentification3 {
	r.SecuritiesSettlementTransactionIdentification = new(SettlementTypeAndIdentification3)
	return r.SecuritiesSettlementTransactionIdentification
}

func (r *References2Choice) AddSecuritiesFinancingTransactionIdentification() *SettlementTypeAndIdentification3 {
	r.SecuritiesFinancingTransactionIdentification = new(SettlementTypeAndIdentification3)
	return r.SecuritiesFinancingTransactionIdentification
}

func (r *References2Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*Max35Text)(&value)
}

func (r *References2Choice) AddOtherTransactionIdentification() *GenericDocumentIdentification1 {
	r.OtherTransactionIdentification = new(GenericDocumentIdentification1)
	return r.OtherTransactionIdentification
}
