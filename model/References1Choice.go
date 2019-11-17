package model

// Choice of reference.
type References1Choice struct {

	// Unambiguous identification of a securities settlement transaction as known by the account owner (or instructing party acting on its behalf).
	SecuritiesSettlementTransactionIdentification *Max35Text `xml:"SctiesSttlmTxId"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId"`

	// Unambiguous identification of the intra-position movement transaction as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *Max35Text `xml:"IntraPosMvmntId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *Max35Text `xml:"OthrTxId"`
}

func (r *References1Choice) SetSecuritiesSettlementTransactionIdentification(value string) {
	r.SecuritiesSettlementTransactionIdentification = (*Max35Text)(&value)
}

func (r *References1Choice) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References1Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*Max35Text)(&value)
}

func (r *References1Choice) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References1Choice) SetOtherTransactionIdentification(value string) {
	r.OtherTransactionIdentification = (*Max35Text)(&value)
}
