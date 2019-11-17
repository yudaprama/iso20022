package model

// Unique identifier of a document, message or transaction.
type TransactionIdentifications4 struct {

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *References4Choice `xml:"AcctOwnrTxId"`
}

func (t *TransactionIdentifications4) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionIdentifications4) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionIdentifications4) AddAccountOwnerTransactionIdentification() *References4Choice {
	t.AccountOwnerTransactionIdentification = new(References4Choice)
	return t.AccountOwnerTransactionIdentification
}
