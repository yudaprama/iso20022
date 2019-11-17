package model

// Choice of reference.
type References60Choice struct {

	// Unambiguous identification of the securities settlement transaction.
	SecuritiesSettlementTransactionIdentification *SettlementTypeAndIdentification22 `xml:"SctiesSttlmTxId"`

	// Unambiguous identification of the underlying securities financing transaction (not the underlying securities financing trade) as assigned by the instructing party.
	SecuritiesFinancingTransactionIdentification *SettlementTypeAndIdentification22 `xml:"SctiesFincgTxId"`

	// Reference to the intra-position movement transaction requested to be cancelled as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *RestrictedFINXMax16Text `xml:"IntraPosMvmntId"`

	// Reference to a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *GenericDocumentIdentification6 `xml:"OthrTxId"`
}

func (r *References60Choice) AddSecuritiesSettlementTransactionIdentification() *SettlementTypeAndIdentification22 {
	r.SecuritiesSettlementTransactionIdentification = new(SettlementTypeAndIdentification22)
	return r.SecuritiesSettlementTransactionIdentification
}

func (r *References60Choice) AddSecuritiesFinancingTransactionIdentification() *SettlementTypeAndIdentification22 {
	r.SecuritiesFinancingTransactionIdentification = new(SettlementTypeAndIdentification22)
	return r.SecuritiesFinancingTransactionIdentification
}

func (r *References60Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References60Choice) AddOtherTransactionIdentification() *GenericDocumentIdentification6 {
	r.OtherTransactionIdentification = new(GenericDocumentIdentification6)
	return r.OtherTransactionIdentification
}
