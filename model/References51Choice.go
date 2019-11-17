package model

// Choice of reference.
type References51Choice struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner, the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId"`
}

func (r *References51Choice) SetAccountOwnerTransactionIdentification(value string) {
	r.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References51Choice) SetAccountServicerTransactionIdentification(value string) {
	r.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References51Choice) SetPoolIdentification(value string) {
	r.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References51Choice) SetMarketInfrastructureTransactionIdentification(value string) {
	r.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (r *References51Choice) SetProcessorTransactionIdentification(value string) {
	r.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}
