package model

// Provides a set if identifications to reference to a securities settlement transaction.
type References1 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *Max35Text `xml:"CmonId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification *Max35Text `xml:"TradId,omitempty"`
}

func (r *References1) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References1) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References1) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (r *References1) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References1) SetCommonIdentification(value string) {
	r.CommonIdentification = (*Max35Text)(&value)
}

func (r *References1) SetTradeIdentification(value string) {
	r.TradeIdentification = (*Max35Text)(&value)
}
