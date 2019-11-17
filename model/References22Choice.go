package model

// Choice of reference.
type References22Choice struct {

	// Reference to a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *GenericDocumentIdentification1 `xml:"OthrTxId"`

	// Unambiguous identification of the underlying securities financing transaction (not the underlying securities financing trade) as assigned by the instructing party.
	SecuritiesFinancingTransactionIdentification *SettlementTypeAndIdentification13 `xml:"SctiesFincgTxId"`

	// Unambiguous identification of the securities settlement transaction.
	SecuritiesSettlementTransactionIdentification *SettlementTypeAndIdentification13 `xml:"SctiesSttlmTxId"`

	// Reference to the intra-position movement transaction requested to be cancelled as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *Max35Text `xml:"IntraPosMvmntId"`
}

func (r *References22Choice) AddOtherTransactionIdentification() *GenericDocumentIdentification1 {
	r.OtherTransactionIdentification = new(GenericDocumentIdentification1)
	return r.OtherTransactionIdentification
}

func (r *References22Choice) AddSecuritiesFinancingTransactionIdentification() *SettlementTypeAndIdentification13 {
	r.SecuritiesFinancingTransactionIdentification = new(SettlementTypeAndIdentification13)
	return r.SecuritiesFinancingTransactionIdentification
}

func (r *References22Choice) AddSecuritiesSettlementTransactionIdentification() *SettlementTypeAndIdentification13 {
	r.SecuritiesSettlementTransactionIdentification = new(SettlementTypeAndIdentification13)
	return r.SecuritiesSettlementTransactionIdentification
}

func (r *References22Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*Max35Text)(&value)
}
