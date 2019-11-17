package model

// Choice of reference.
type References59Choice struct {

	// Reference to a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *GenericDocumentIdentification5 `xml:"OthrTxId"`

	// Unambiguous identification of the underlying securities financing transaction (not the underlying securities financing trade) as assigned by the instructing party.
	SecuritiesFinancingTransactionIdentification *SettlementTypeAndIdentification22 `xml:"SctiesFincgTxId"`

	// Unambiguous identification of the securities settlement transaction.
	SecuritiesSettlementTransactionIdentification *SettlementTypeAndIdentification22 `xml:"SctiesSttlmTxId"`

	// Reference to the intra-position movement transaction requested to be cancelled as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *RestrictedFINXMax16Text `xml:"IntraPosMvmntId"`
}

func (r *References59Choice) AddOtherTransactionIdentification() *GenericDocumentIdentification5 {
	r.OtherTransactionIdentification = new(GenericDocumentIdentification5)
	return r.OtherTransactionIdentification
}

func (r *References59Choice) AddSecuritiesFinancingTransactionIdentification() *SettlementTypeAndIdentification22 {
	r.SecuritiesFinancingTransactionIdentification = new(SettlementTypeAndIdentification22)
	return r.SecuritiesFinancingTransactionIdentification
}

func (r *References59Choice) AddSecuritiesSettlementTransactionIdentification() *SettlementTypeAndIdentification22 {
	r.SecuritiesSettlementTransactionIdentification = new(SettlementTypeAndIdentification22)
	return r.SecuritiesSettlementTransactionIdentification
}

func (r *References59Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*RestrictedFINXMax16Text)(&value)
}
