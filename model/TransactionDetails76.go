package model

// Identifies the details of the transaction.
type TransactionDetails76 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *SettlementTypeAndIdentification18 `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *Max35Text `xml:"OthrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification98 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionDetails *TransactionDetails74 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails76) AddAccountOwnerTransactionIdentification() *SettlementTypeAndIdentification18 {
	t.AccountOwnerTransactionIdentification = new(SettlementTypeAndIdentification18)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails76) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails76) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails76) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails76) SetOtherTransactionIdentification(value string) {
	t.OtherTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails76) AddAccountOwner() *PartyIdentification98 {
	t.AccountOwner = new(PartyIdentification98)
	return t.AccountOwner
}

func (t *TransactionDetails76) AddSafekeepingAccount() *SecuritiesAccount19 {
	t.SafekeepingAccount = new(SecuritiesAccount19)
	return t.SafekeepingAccount
}

func (t *TransactionDetails76) AddTransactionDetails() *TransactionDetails74 {
	t.TransactionDetails = new(TransactionDetails74)
	return t.TransactionDetails
}
