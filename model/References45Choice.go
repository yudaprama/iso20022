package model

// Choice of reference.
type References45Choice struct {

	// Unambiguous identification of the securities settlement transaction.
	SecuritiesSettlementTransactionIdentification *SettlementTypeAndIdentification18 `xml:"SctiesSttlmTxId"`

	// Unambiguous identification of the underlying securities financing transaction (not the underlying securities financing trade) as assigned by the instructing party.
	SecuritiesFinancingTransactionIdentification *SettlementTypeAndIdentification18 `xml:"SctiesFincgTxId"`

	// Reference to the intra-position movement transaction requested to be cancelled as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *Max35Text `xml:"IntraPosMvmntId"`

	// Reference to a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *GenericDocumentIdentification4 `xml:"OthrTxId"`
}

func (r *References45Choice) AddSecuritiesSettlementTransactionIdentification() *SettlementTypeAndIdentification18 {
	r.SecuritiesSettlementTransactionIdentification = new(SettlementTypeAndIdentification18)
	return r.SecuritiesSettlementTransactionIdentification
}

func (r *References45Choice) AddSecuritiesFinancingTransactionIdentification() *SettlementTypeAndIdentification18 {
	r.SecuritiesFinancingTransactionIdentification = new(SettlementTypeAndIdentification18)
	return r.SecuritiesFinancingTransactionIdentification
}

func (r *References45Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*Max35Text)(&value)
}

func (r *References45Choice) AddOtherTransactionIdentification() *GenericDocumentIdentification4 {
	r.OtherTransactionIdentification = new(GenericDocumentIdentification4)
	return r.OtherTransactionIdentification
}
