package model

// Identifies the details of the transaction.
type TransactionDetails64 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *SettlementTypeAndIdentification3 `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *Max35Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *Max35Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *Max35Text `xml:"PrcrTxId,omitempty"`

	// Identification of a transaction that cannot be identified using a standard reference element present in the message.
	OtherTransactionIdentification *Max35Text `xml:"OthrTxId,omitempty"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Provides additional details on the transaction which can not be included within the structured fields of the message.
	TransactionDetails *TransactionDetails28 `xml:"TxDtls,omitempty"`
}

func (t *TransactionDetails64) AddAccountOwnerTransactionIdentification() *SettlementTypeAndIdentification3 {
	t.AccountOwnerTransactionIdentification = new(SettlementTypeAndIdentification3)
	return t.AccountOwnerTransactionIdentification
}

func (t *TransactionDetails64) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails64) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails64) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails64) SetOtherTransactionIdentification(value string) {
	t.OtherTransactionIdentification = (*Max35Text)(&value)
}

func (t *TransactionDetails64) AddAccountOwner() *PartyIdentification36Choice {
	t.AccountOwner = new(PartyIdentification36Choice)
	return t.AccountOwner
}

func (t *TransactionDetails64) AddSafekeepingAccount() *SecuritiesAccount13 {
	t.SafekeepingAccount = new(SecuritiesAccount13)
	return t.SafekeepingAccount
}

func (t *TransactionDetails64) AddTransactionDetails() *TransactionDetails28 {
	t.TransactionDetails = new(TransactionDetails28)
	return t.TransactionDetails
}
