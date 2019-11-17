package model

// Choice of reference.
type References13Choice struct {

	// Unambiguous identification of a securities settlement transaction as known by the account owner (or instructing party acting on its behalf).
	SecuritiesSettlementTransactionIdentification *Max35Text `xml:"SctiesSttlmTxId,omitempty"`

	// Unambiguous identification of the intra-position movement transaction as known by the account owner (or instructing party acting on its behalf).
	IntraPositionMovementIdentification *Max35Text `xml:"IntraPosMvmntId,omitempty"`

	// Unambiguous identification of the intra balance movement transaction as known by the account owner.
	IntraBalanceMovementIdentification *Max35Text `xml:"IntraBalMvmntId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *Max35Text `xml:"TradId"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *Max35Text `xml:"OthrTxId,omitempty"`
}

func (r *References13Choice) SetSecuritiesSettlementTransactionIdentification(value string) {
	r.SecuritiesSettlementTransactionIdentification = (*Max35Text)(&value)
}

func (r *References13Choice) SetIntraPositionMovementIdentification(value string) {
	r.IntraPositionMovementIdentification = (*Max35Text)(&value)
}

func (r *References13Choice) SetIntraBalanceMovementIdentification(value string) {
	r.IntraBalanceMovementIdentification = (*Max35Text)(&value)
}

func (r *References13Choice) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References13Choice) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (r *References13Choice) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References13Choice) SetCommonIdentification(value string) {
	r.CommonIdentification = (*Max35Text)(&value)
}

func (r *References13Choice) SetTradeIdentification(value string) {
	r.TradeIdentification = (*Max35Text)(&value)
}

func (r *References13Choice) SetOtherTransactionIdentification(value string) {
	r.OtherTransactionIdentification = (*Max35Text)(&value)
}
