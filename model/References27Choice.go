package model

// Choice of reference.
type References27Choice struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *Max35Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId"`

	// Collective reference identifying a set of messages.
	PoolIdentification *Max35Text `xml:"PoolId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner, the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId"`
}

func (r *References27Choice) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References27Choice) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (r *References27Choice) SetPoolIdentification(value string) {
	r.PoolIdentification = (*Max35Text)(&value)
}

func (r *References27Choice) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (r *References27Choice) SetProcessorTransactionIdentification(value string) {
	r.ProcessorTransactionIdentification = (*Max35Text)(&value)
}
