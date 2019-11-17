package model

// Choice of reference.
type References58Choice struct {

	// Unambiguous identification of a securities settlement transaction as known by the account owner (or instructing party acting on its behalf).
	SecuritiesSettlementTransactionIdentification *RestrictedFINXMax16Text `xml:"SctiesSttlmTxId"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId"`

	// Unambiguous identification of the intra-position movement transaction as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *RestrictedFINXMax16Text `xml:"IntraPosMvmntId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *RestrictedFINXMax16Text `xml:"OthrTxId"`
}

func (r *References58Choice) SetSecuritiesSettlementTransactionIdentification(value string) {
	r.SecuritiesSettlementTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References58Choice) SetPoolIdentification(value string) {
	r.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References58Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References58Choice) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References58Choice) SetOtherTransactionIdentification(value string) {
	r.OtherTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}
