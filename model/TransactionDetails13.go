package model

// Identifies the details of the transaction.
type TransactionDetails13 struct {

	// Provides transaction type and identification information.
	AccountServicerTransactionIdentification *SettlementTypeAndIdentification3 `xml:"AcctSvcrTxId"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Identification1 `xml:"MktInfrstrctrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails10 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails13) AddAccountServicerTransactionIdentification() *SettlementTypeAndIdentification3 {
	t.AccountServicerTransactionIdentification = new(SettlementTypeAndIdentification3)
	return t.AccountServicerTransactionIdentification
}

func (t *TransactionDetails13) AddMarketInfrastructureTransactionIdentification() *Identification1 {
	t.MarketInfrastructureTransactionIdentification = new(Identification1)
	return t.MarketInfrastructureTransactionIdentification
}

func (t *TransactionDetails13) AddAccountOwner() *PartyIdentification13Choice {
	t.AccountOwner = new(PartyIdentification13Choice)
	return t.AccountOwner
}

func (t *TransactionDetails13) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails13) AddTransactionDetails() *TransactionDetails10 {
	t.TransactionDetails = new(TransactionDetails10)
	return t.TransactionDetails
}
