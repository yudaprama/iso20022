package model

// Identifies the details of the transaction.
type TransactionDetails11 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *References2Choice `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails10 `xml:"TxDtls,omitempty"`

	// Specifies whether an associated FX should be cancelled.
	FXCancellation *FXCancellation1Choice `xml:"FxCxl,omitempty"`
}

func (t *TransactionDetails11) AddAccountOwnerTransactionIdentification() *References2Choice {
	t.AccountOwnerTransactionIdentification = new(References2Choice)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails11) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails11) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails11) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails11) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails11) AddTransactionDetails() *TransactionDetails10 {
	t.TransactionDetails = new(TransactionDetails10)
	return t.TransactionDetails
}

func (t *TransactionDetails11) AddFXCancellation() *FXCancellation1Choice {
	t.FXCancellation = new(FXCancellation1Choice)
	return t.FXCancellation
}
